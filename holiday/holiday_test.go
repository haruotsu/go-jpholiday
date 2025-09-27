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
