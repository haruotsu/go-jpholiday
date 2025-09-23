package holiday

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

var (
	cache   *HolidayCache
	cacheMu sync.RWMutex
)

// init initializes the package with default holiday data
func init() {
	initializeCache()
}

// initializeCache initializes the holiday cache with the available data
func initializeCache() {
	// Don't use once.Do in tests where we need to reset cache
	cacheMu.Lock()
	if cache != nil {
		cacheMu.Unlock()
		return
	}
	cacheMu.Unlock()

	// Try to load from cache file (created by GitHub Actions)
	if loadedCache, err := LoadCache("data/holidays.json"); err == nil {
		SetCache(loadedCache)
		return
	}

	// If no cache file found, this means the package is used without
	// the Actions-generated data, so we need some basic data
	defaultCache := getDefaultHolidayData()
	SetCache(defaultCache)
}

// EnsureInitialized ensures the cache is initialized (useful for testing)
func EnsureInitialized() {
	initializeCache()
}

// SetCache sets the holiday cache
func SetCache(c *HolidayCache) {
	cacheMu.Lock()
	defer cacheMu.Unlock()
	cache = c
}

// GetCache returns the current holiday cache
func GetCache() *HolidayCache {
	cacheMu.RLock()
	defer cacheMu.RUnlock()
	return cache
}

// IsHoliday checks if the given date is a Japanese holiday
func IsHoliday(date time.Time) bool {
	ensureCacheLoaded()

	cacheMu.RLock()
	defer cacheMu.RUnlock()

	if cache == nil {
		return false
	}

	key := formatDateKey(date)
	_, exists := cache.Holidays[key]
	return exists
}

// ensureCacheLoaded ensures cache is initialized (fallback safety)
func ensureCacheLoaded() {
	cacheMu.RLock()
	if cache != nil {
		cacheMu.RUnlock()
		return
	}
	cacheMu.RUnlock()

	// Cache is nil, initialize it
	initializeCache()
}

// GetHolidayName returns the name of the holiday for the given date
// Returns empty string if the date is not a holiday
func GetHolidayName(date time.Time) string {
	ensureCacheLoaded()

	cacheMu.RLock()
	defer cacheMu.RUnlock()

	if cache == nil {
		return ""
	}

	key := formatDateKey(date)
	if holiday, exists := cache.Holidays[key]; exists {
		return holiday.Name
	}
	return ""
}

// GetHolidaysInYear returns all holidays in the specified year
func GetHolidaysInYear(year int) []Holiday {
	ensureCacheLoaded()

	cacheMu.RLock()
	defer cacheMu.RUnlock()

	if cache == nil {
		return []Holiday{}
	}

	var holidays []Holiday
	startOfYear := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	endOfYear := time.Date(year, 12, 31, 23, 59, 59, 999999999, time.UTC)

	for _, holiday := range cache.Holidays {
		if !holiday.Date.Before(startOfYear) && !holiday.Date.After(endOfYear) {
			holidays = append(holidays, holiday)
		}
	}

	// Sort holidays by date
	sort.Slice(holidays, func(i, j int) bool {
		return holidays[i].Date.Before(holidays[j].Date)
	})

	return holidays
}

// GetHolidaysInRange returns holidays within the specified date range
func GetHolidaysInRange(start, end time.Time) []Holiday {
	ensureCacheLoaded()

	cacheMu.RLock()
	defer cacheMu.RUnlock()

	if cache == nil {
		return []Holiday{}
	}

	var holidays []Holiday
	// Normalize times to beginning of day for comparison
	startDay := time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
	endDay := time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, 999999999, end.Location())

	for _, holiday := range cache.Holidays {
		if !holiday.Date.Before(startDay) && !holiday.Date.After(endDay) {
			holidays = append(holidays, holiday)
		}
	}

	// Sort holidays by date
	sort.Slice(holidays, func(i, j int) bool {
		return holidays[i].Date.Before(holidays[j].Date)
	})

	return holidays
}

// formatDateKey formats a date as "YYYY-MM-DD" for use as a map key
func formatDateKey(date time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02d", date.Year(), date.Month(), date.Day())
}