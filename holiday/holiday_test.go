package holiday

import (
	"testing"
	"time"
)

func TestIsHoliday_WithHolidayDate(t *testing.T) {
	cache := NewTestHolidayCache()
	SetCache(cache)

	testDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	if !IsHoliday(testDate) {
		t.Errorf("expected true for holiday date %v, got false", testDate)
	}
}

func TestIsHoliday_WithWeekday(t *testing.T) {
	cache := NewTestHolidayCache()
	SetCache(cache)

	testDate := time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)
	if IsHoliday(testDate) {
		t.Errorf("expected false for weekday %v, got true", testDate)
	}
}

func TestIsHoliday_WithWeekend(t *testing.T) {
	cache := NewTestHolidayCache()
	SetCache(cache)

	saturday := time.Date(2024, 1, 6, 0, 0, 0, 0, time.UTC)
	if IsHoliday(saturday) {
		t.Errorf("expected false for Saturday %v, got true", saturday)
	}

	sunday := time.Date(2024, 1, 7, 0, 0, 0, 0, time.UTC)
	if IsHoliday(sunday) {
		t.Errorf("expected false for Sunday %v, got true", sunday)
	}
}

func TestGetHolidayName_WithHolidayDate(t *testing.T) {
	cache := NewTestHolidayCache()
	SetCache(cache)

	testDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	name := GetHolidayName(testDate)
	expected := "元日"
	if name != expected {
		t.Errorf("expected holiday name '%s', got '%s'", expected, name)
	}
}

func TestGetHolidayName_WithWeekday(t *testing.T) {
	cache := NewTestHolidayCache()
	SetCache(cache)

	testDate := time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)
	name := GetHolidayName(testDate)
	if name != "" {
		t.Errorf("expected empty string for weekday, got '%s'", name)
	}
}

func TestGetHolidaysInYear_WithDataExists(t *testing.T) {
	cache := NewTestHolidayCache()
	SetCache(cache)

	holidays := GetHolidaysInYear(2024)
	if len(holidays) != 4 {
		t.Errorf("expected 4 holidays in 2024, got %d", len(holidays))
	}

	if len(holidays) > 0 && holidays[0].Name != "元日" {
		t.Errorf("expected first holiday to be '元日', got '%s'", holidays[0].Name)
	}
}

func TestGetHolidaysInYear_WithNoData(t *testing.T) {
	cache := NewTestHolidayCache()
	SetCache(cache)

	holidays := GetHolidaysInYear(2023)
	if len(holidays) != 0 {
		t.Errorf("expected empty array for year with no data, got %d holidays", len(holidays))
	}
}

func TestGetHolidaysInRange_WithHolidays(t *testing.T) {
	cache := NewTestHolidayCache()
	SetCache(cache)

	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)

	holidays := GetHolidaysInRange(start, end)
	if len(holidays) != 3 {
		t.Errorf("expected 3 holidays in range, got %d", len(holidays))
	}
}

func TestGetHolidaysInRange_WithNoHolidays(t *testing.T) {
	cache := NewTestHolidayCache()
	SetCache(cache)

	start := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 4, 30, 0, 0, 0, 0, time.UTC)

	holidays := GetHolidaysInRange(start, end)
	if len(holidays) != 0 {
		t.Errorf("expected empty array when no holidays in range, got %d holidays", len(holidays))
	}
}

func TestIsHoliday_SpecificHolidays(t *testing.T) {
	testCases := []struct {
		name        string
		date        time.Time
		holidayName string
		expected    bool
	}{
		{
			name:        "New Year's Day (1/1) returns true",
			date:        time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			holidayName: "元日",
			expected:    true,
		},
		{
			name:        "National Foundation Day (2/11) returns true",
			date:        time.Date(2024, 2, 11, 0, 0, 0, 0, time.UTC),
			holidayName: "建国記念の日",
			expected:    true,
		},
		{
			name:        "Constitution Memorial Day (5/3) returns true",
			date:        time.Date(2024, 5, 3, 0, 0, 0, 0, time.UTC),
			holidayName: "憲法記念日",
			expected:    true,
		},
		{
			name:        "Substitute holiday returns true",
			date:        time.Date(2024, 2, 12, 0, 0, 0, 0, time.UTC),
			holidayName: "振替休日",
			expected:    true,
		},
		{
			name:     "Weekday returns false",
			date:     time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "Saturday returns false (not a holiday)",
			date:     time.Date(2024, 1, 6, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "Sunday returns false (not a holiday)",
			date:     time.Date(2024, 1, 7, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
	}

	cache := NewTestHolidayCache()
	SetCache(cache)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsHoliday(tc.date)
			if result != tc.expected {
				t.Errorf("IsHoliday(%v) = %v, expected %v", tc.date, result, tc.expected)
			}
		})
	}
}

func TestGetHolidayName_SpecificHolidays(t *testing.T) {
	testCases := []struct {
		name         string
		date         time.Time
		expectedName string
	}{
		{
			name:         "New Year's Day returns '元日'",
			date:         time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			expectedName: "元日",
		},
		{
			name:         "National Foundation Day returns '建国記念の日'",
			date:         time.Date(2024, 2, 11, 0, 0, 0, 0, time.UTC),
			expectedName: "建国記念の日",
		},
		{
			name:         "Weekday returns empty string",
			date:         time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
			expectedName: "",
		},
		{
			name:         "Saturday returns empty string",
			date:         time.Date(2024, 1, 6, 0, 0, 0, 0, time.UTC),
			expectedName: "",
		},
		{
			name:         "Sunday returns empty string",
			date:         time.Date(2024, 1, 7, 0, 0, 0, 0, time.UTC),
			expectedName: "",
		},
		{
			name:         "Substitute holiday returns '振替休日'",
			date:         time.Date(2024, 2, 12, 0, 0, 0, 0, time.UTC),
			expectedName: "振替休日",
		},
	}

	cache := NewTestHolidayCache()
	SetCache(cache)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := GetHolidayName(tc.date)
			if result != tc.expectedName {
				t.Errorf("GetHolidayName(%v) = '%s', expected '%s'", tc.date, result, tc.expectedName)
			}
		})
	}
}

func TestGetHolidaysInYear_DetailedTests(t *testing.T) {
	cache := NewTestHolidayCache()
	SetCache(cache)

	t.Run("Returns all holidays for 2024", func(t *testing.T) {
		holidays := GetHolidaysInYear(2024)
		expectedCount := 4
		if len(holidays) != expectedCount {
			t.Errorf("Expected %d holidays in 2024, got %d", expectedCount, len(holidays))
		}

		for i := 1; i < len(holidays); i++ {
			if holidays[i-1].Date.After(holidays[i].Date) {
				t.Errorf("Holidays are not sorted by date: %v is after %v", holidays[i-1].Date, holidays[i].Date)
			}
		}
	})

	t.Run("Returns all holidays for 2025", func(t *testing.T) {
		holidays := GetHolidaysInYear(2025)
		expectedCount := 1
		if len(holidays) != expectedCount {
			t.Errorf("Expected %d holidays in 2025, got %d", expectedCount, len(holidays))
		}
	})

	t.Run("Returns empty array for year with no data", func(t *testing.T) {
		holidays := GetHolidaysInYear(2020)
		if len(holidays) != 0 {
			t.Errorf("Expected empty array for year with no data, got %d holidays", len(holidays))
		}
	})
}

func TestGetHolidaysInRange_DetailedTests(t *testing.T) {
	cache := NewTestHolidayCache()
	SetCache(cache)

	t.Run("Returns holidays for January 2024 (New Year's Day)", func(t *testing.T) {
		start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
		end := time.Date(2024, 1, 31, 23, 59, 59, 0, time.UTC)
		holidays := GetHolidaysInRange(start, end)

		expectedCount := 1
		if len(holidays) != expectedCount {
			t.Errorf("Expected %d holidays in January 2024, got %d", expectedCount, len(holidays))
		}

		if len(holidays) > 0 && holidays[0].Name != "元日" {
			t.Errorf("Expected first holiday to be '元日', got '%s'", holidays[0].Name)
		}
	})

	t.Run("Returns holidays for February 2024 (National Foundation Day, Substitute holiday)", func(t *testing.T) {
		start := time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)
		end := time.Date(2024, 2, 29, 23, 59, 59, 0, time.UTC)
		holidays := GetHolidaysInRange(start, end)

		expectedCount := 2
		if len(holidays) != expectedCount {
			t.Errorf("Expected %d holidays in February 2024, got %d", expectedCount, len(holidays))
		}
	})

	t.Run("Returns empty array when no holidays in range", func(t *testing.T) {
		start := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
		end := time.Date(2024, 4, 30, 23, 59, 59, 0, time.UTC)
		holidays := GetHolidaysInRange(start, end)

		if len(holidays) != 0 {
			t.Errorf("Expected empty array when no holidays in range, got %d holidays", len(holidays))
		}
	})

	t.Run("Returns holidays correctly for cross-year range", func(t *testing.T) {
		start := time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC)
		end := time.Date(2025, 1, 31, 23, 59, 59, 0, time.UTC)
		holidays := GetHolidaysInRange(start, end)

		expectedCount := 1
		if len(holidays) != expectedCount {
			t.Errorf("Expected %d holidays in cross-year range, got %d", expectedCount, len(holidays))
		}

		if len(holidays) > 0 && holidays[0].Name != "元日" {
			t.Errorf("Expected holiday to be '元日', got '%s'", holidays[0].Name)
		}
	})
}

func TestGetHolidaysInYear_ErrorCases(t *testing.T) {
	cache := NewTestHolidayCache()
	SetCache(cache)

	testCases := []struct {
		name string
		year int
	}{
		{
			name: "Negative year",
			year: -1,
		},
		{
			name: "Year zero",
			year: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			holidays := GetHolidaysInYear(tc.year)
			if len(holidays) != 0 {
				t.Errorf("Expected empty array for invalid year %d, got %d holidays", tc.year, len(holidays))
			}
		})
	}
}

func TestIsHoliday_WithDefaultCache(t *testing.T) {
	SetCache(nil)

	testCases := []struct {
		name     string
		date     time.Time
		expected bool
	}{
		{
			name:     "New Year's Day (1/1) returns true",
			date:     time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Coming of Age Day returns true",
			date:     time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "National Foundation Day (2/11) returns true",
			date:     time.Date(2024, 2, 11, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Emperor's Birthday (2/23) returns true",
			date:     time.Date(2024, 2, 23, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Vernal Equinox Day returns true",
			date:     time.Date(2024, 3, 20, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Showa Day (4/29) returns true",
			date:     time.Date(2024, 4, 29, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Constitution Memorial Day (5/3) returns true",
			date:     time.Date(2024, 5, 3, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Greenery Day (5/4) returns true",
			date:     time.Date(2024, 5, 4, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Children's Day (5/5) returns true",
			date:     time.Date(2024, 5, 5, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Substitute holiday returns true",
			date:     time.Date(2024, 2, 12, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "National holiday returns true",
			date:     time.Date(2024, 5, 6, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Weekday returns false",
			date:     time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "Saturday returns false (not a holiday)",
			date:     time.Date(2024, 1, 6, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "Sunday returns false (not a holiday)",
			date:     time.Date(2024, 1, 7, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsHoliday(tc.date)
			if result != tc.expected {
				t.Errorf("IsHoliday(%v) = %v, expected %v", tc.date, result, tc.expected)
			}
		})
	}
}

func TestHolidayFunctions_WithNilCache(t *testing.T) {
	SetCache(nil)

	t.Run("IsHoliday with nil cache - auto initialization", func(t *testing.T) {
		date := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
		result := IsHoliday(date)
		if result != true {
			t.Errorf("Expected true after auto-initialization for New Year's Day, got %v", result)
		}
	})

	t.Run("GetHolidayName with nil cache - auto initialization", func(t *testing.T) {
		date := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
		result := GetHolidayName(date)
		if result != "元日" {
			t.Errorf("Expected '元日' after auto-initialization, got '%s'", result)
		}
	})

	t.Run("GetHolidaysInYear with nil cache - auto initialization", func(t *testing.T) {
		holidays := GetHolidaysInYear(2024)
		if len(holidays) == 0 {
			t.Errorf("Expected holidays after auto-initialization, got empty array")
		}
	})

	t.Run("GetHolidaysInRange with nil cache - auto initialization", func(t *testing.T) {
		start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
		end := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)
		holidays := GetHolidaysInRange(start, end)
		if len(holidays) == 0 {
			t.Errorf("Expected holidays after auto-initialization, got empty array")
		}
	})
}
