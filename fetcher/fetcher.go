package fetcher

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/haruotsu/go-jpholiday/holiday"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const (
	// Japanese Official Holiday Calendar ID (excludes cultural observances like 節分, 七夕, クリスマス)
	japaneseHolidayCalendarID = "ja.japanese.official#holiday@group.v.calendar.google.com"

	// Retry configuration
	maxRetries = 3
	retryDelay = 1 * time.Second
)

// isOfficialHoliday determines if an event is an official Japanese holiday
// by checking the description field. Official holidays have "祝日" in the description,
// while festivals and cultural observances have "祭日" in the description.
func isOfficialHoliday(description string) bool {
	// First, exclude events with "祭日" (festival/cultural observance) in description
	if strings.Contains(description, "祭日") {
		return false
	}

	// Then, include events with "祝日" (official holiday) in description
	if strings.Contains(description, "祝日") {
		return true
	}

	// For events without clear description, be conservative and exclude
	return false
}

// Fetcher interface defines methods for fetching holiday data
type Fetcher interface {
	FetchHolidays(year int) ([]holiday.Holiday, error)
	FetchHolidaysRange(startYear, endYear int) ([]holiday.Holiday, error)
}

// GoogleCalendarFetcher implements the Fetcher interface using Google Calendar API
type GoogleCalendarFetcher struct {
	APIKey  string
	service *calendar.Service
}

// NewFetcher creates a new GoogleCalendarFetcher
func NewFetcher(apiKey string) *GoogleCalendarFetcher {
	if apiKey == "" {
		return nil
	}

	return &GoogleCalendarFetcher{
		APIKey: apiKey,
	}
}

// initService initializes the Google Calendar service if not already initialized
func (f *GoogleCalendarFetcher) initService() error {
	if f.service == nil {
		ctx := context.Background()
		service, err := calendar.NewService(ctx, option.WithAPIKey(f.APIKey))
		if err != nil {
			return fmt.Errorf("failed to create calendar service: %w", err)
		}
		f.service = service
	}
	return nil
}

// FetchHolidays fetches holidays for a specific year
func (f *GoogleCalendarFetcher) FetchHolidays(year int) ([]holiday.Holiday, error) {
	if err := f.initService(); err != nil {
		return nil, err
	}

	// Define the time range for the year
	timeMin := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC).Format(time.RFC3339)
	timeMax := time.Date(year+1, 1, 1, 0, 0, 0, 0, time.UTC).Format(time.RFC3339)

	// Fetch events with retry mechanism
	var events *calendar.Events
	var err error

	for attempt := 0; attempt <= maxRetries; attempt++ {
		events, err = f.service.Events.List(japaneseHolidayCalendarID).
			TimeMin(timeMin).
			TimeMax(timeMax).
			SingleEvents(true).
			OrderBy("startTime").
			Do()

		if err == nil {
			break
		}

		if attempt < maxRetries {
			time.Sleep(retryDelay * time.Duration(attempt+1)) // Exponential backoff
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to fetch holidays after %d retries: %w", maxRetries, err)
	}

	// Convert calendar events to our Holiday struct
	var holidays []holiday.Holiday
	for _, event := range events.Items {
		if event.Start == nil || event.Start.Date == "" {
			continue
		}

		// Parse the date
		eventDate, err := time.Parse("2006-01-02", event.Start.Date)
		if err != nil {
			continue
		}

		// Filter out non-official holidays (festivals)
		description := strings.TrimSpace(event.Description)
		if isOfficialHoliday(description) {
			// Create Holiday struct
			h := holiday.Holiday{
				Date:        eventDate,
				Name:        event.Summary,
				Description: description,
			}

			holidays = append(holidays, h)
		}
	}

	return holidays, nil
}

// FetchHolidaysRange fetches holidays for a range of years
func (f *GoogleCalendarFetcher) FetchHolidaysRange(startYear, endYear int) ([]holiday.Holiday, error) {
	var allHolidays []holiday.Holiday

	for year := startYear; year <= endYear; year++ {
		holidays, err := f.FetchHolidays(year)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch holidays for year %d: %w", year, err)
		}
		allHolidays = append(allHolidays, holidays...)
	}

	return allHolidays, nil
}
