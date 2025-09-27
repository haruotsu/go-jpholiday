package fetcher

import (
	"fmt"

	"github.com/haruotsu/go-jpholiday/model"
)

// MockFetcher is a mock implementation of Fetcher for testing
type MockFetcher struct {
	Holidays    []model.Holiday
	ShouldError bool
	ErrorMsg    string
	RetryCount  int
	CallCount   int
}

// FetchHolidays implements the Fetcher interface for testing
func (m *MockFetcher) FetchHolidays(year int) ([]model.Holiday, error) {
	m.CallCount++

	if m.ShouldError {
		if m.CallCount <= m.RetryCount+1 {
			return nil, fmt.Errorf("%s", m.ErrorMsg)
		}
	}

	// Filter holidays by year
	var yearHolidays []model.Holiday
	for _, h := range m.Holidays {
		if h.Date.Year() == year {
			yearHolidays = append(yearHolidays, h)
		}
	}

	return yearHolidays, nil
}

// FetchHolidaysRange implements the Fetcher interface for testing
func (m *MockFetcher) FetchHolidaysRange(startYear, endYear int) ([]model.Holiday, error) {
	m.CallCount++

	if m.ShouldError {
		return nil, fmt.Errorf("%s", m.ErrorMsg)
	}

	// Filter holidays by year range
	var rangeHolidays []model.Holiday
	for _, h := range m.Holidays {
		year := h.Date.Year()
		if year >= startYear && year <= endYear {
			rangeHolidays = append(rangeHolidays, h)
		}
	}

	return rangeHolidays, nil
}
