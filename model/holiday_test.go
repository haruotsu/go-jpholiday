package model

import (
	"encoding/json"
	"testing"
	"time"
)

func TestHoliday_Creation(t *testing.T) {
	date := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	holiday := Holiday{
		Date: date,
		Name: "元日",
	}

	if holiday.Date != date {
		t.Errorf("Expected date %v, got %v", date, holiday.Date)
	}

	if holiday.Name != "元日" {
		t.Errorf("Expected name '元日', got '%s'", holiday.Name)
	}
}

func TestHoliday_DateField(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		expected time.Time
	}{
		{
			name:     "Valid date",
			date:     time.Date(2024, 5, 3, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2024, 5, 3, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Year-end date",
			date:     time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "New Year date",
			date:     time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			holiday := Holiday{
				Date: tc.date,
				Name: "Test Holiday",
			}

			if !holiday.Date.Equal(tc.expected) {
				t.Errorf("Expected date %v, got %v", tc.expected, holiday.Date)
			}
		})
	}
}

func TestHoliday_NameField(t *testing.T) {
	testCases := []struct {
		name         string
		holidayName  string
		expectedName string
	}{
		{
			name:         "New Year's Day",
			holidayName:  "元日",
			expectedName: "元日",
		},
		{
			name:         "Coming of Age Day",
			holidayName:  "成人の日",
			expectedName: "成人の日",
		},
		{
			name:         "Empty string",
			holidayName:  "",
			expectedName: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			holiday := Holiday{
				Date: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				Name: tc.holidayName,
			}

			if holiday.Name != tc.expectedName {
				t.Errorf("Expected name '%s', got '%s'", tc.expectedName, holiday.Name)
			}
		})
	}
}

func TestHoliday_JSONSerialization(t *testing.T) {
	date := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	holiday := Holiday{
		Date:        date,
		Name:        "元日",
		Description: "新年を祝う日",
	}

	jsonData, err := json.Marshal(holiday)
	if err != nil {
		t.Fatalf("Failed to marshal Holiday to JSON: %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if result["name"] != "元日" {
		t.Errorf("Expected name '元日' in JSON, got '%v'", result["name"])
	}

	if result["description"] != "新年を祝う日" {
		t.Errorf("Expected description '新年を祝う日' in JSON, got '%v'", result["description"])
	}
}

func TestHoliday_JSONDeserialization(t *testing.T) {
	jsonData := `{
		"date": "2024-01-01T00:00:00Z",
		"name": "元日",
		"description": "新年を祝う日"
	}`

	var holiday Holiday
	err := json.Unmarshal([]byte(jsonData), &holiday)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON to Holiday: %v", err)
	}

	expectedDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	if !holiday.Date.Equal(expectedDate) {
		t.Errorf("Expected date %v, got %v", expectedDate, holiday.Date)
	}

	if holiday.Name != "元日" {
		t.Errorf("Expected name '元日', got '%s'", holiday.Name)
	}

	if holiday.Description != "新年を祝う日" {
		t.Errorf("Expected description '新年を祝う日', got '%s'", holiday.Description)
	}
}

func TestHoliday_JSONDeserializationInvalid(t *testing.T) {
	testCases := []struct {
		name     string
		jsonData string
	}{
		{
			name:     "Invalid JSON syntax",
			jsonData: `{"date": "2024-01-01", "name": "元日"`,
		},
		{
			name:     "Invalid date format",
			jsonData: `{"date": "invalid-date", "name": "元日"}`,
		},
		{
			name:     "Completely invalid JSON",
			jsonData: `invalid json`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var holiday Holiday
			err := json.Unmarshal([]byte(tc.jsonData), &holiday)
			if err == nil {
				t.Errorf("Expected error for invalid JSON, but got none")
			}
		})
	}
}

func TestHolidayCache_Creation(t *testing.T) {
	now := time.Now()
	cache := HolidayCache{
		LastUpdated: now,
		Holidays:    make(map[string]Holiday),
	}

	if !cache.LastUpdated.Equal(now) {
		t.Errorf("Expected LastUpdated %v, got %v", now, cache.LastUpdated)
	}

	if cache.Holidays == nil {
		t.Error("Expected Holidays map to be initialized, got nil")
	}

	if len(cache.Holidays) != 0 {
		t.Errorf("Expected empty Holidays map, got %d items", len(cache.Holidays))
	}
}

func TestHolidayCache_LastUpdatedField(t *testing.T) {
	testTime := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	cache := HolidayCache{
		LastUpdated: testTime,
		Holidays:    make(map[string]Holiday),
	}

	if !cache.LastUpdated.Equal(testTime) {
		t.Errorf("Expected LastUpdated %v, got %v", testTime, cache.LastUpdated)
	}
}

func TestHolidayCache_HolidaysMapInitialization(t *testing.T) {
	cache := HolidayCache{
		LastUpdated: time.Now(),
		Holidays:    make(map[string]Holiday),
	}

	holiday := Holiday{
		Date: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		Name: "元日",
	}
	cache.Holidays["2024-01-01"] = holiday

	if len(cache.Holidays) != 1 {
		t.Errorf("Expected 1 holiday in map, got %d", len(cache.Holidays))
	}

	if retrievedHoliday, exists := cache.Holidays["2024-01-01"]; !exists {
		t.Error("Expected holiday to exist in map")
	} else if retrievedHoliday.Name != "元日" {
		t.Errorf("Expected holiday name '元日', got '%s'", retrievedHoliday.Name)
	}
}

func TestHolidayCache_JSONSerialization(t *testing.T) {
	now := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	holiday := Holiday{
		Date: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		Name: "元日",
	}

	cache := HolidayCache{
		LastUpdated: now,
		Holidays: map[string]Holiday{
			"2024-01-01": holiday,
		},
	}

	jsonData, err := json.Marshal(cache)
	if err != nil {
		t.Fatalf("Failed to marshal HolidayCache to JSON: %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if _, exists := result["last_updated"]; !exists {
		t.Error("Expected 'last_updated' field in JSON")
	}

	if _, exists := result["holidays"]; !exists {
		t.Error("Expected 'holidays' field in JSON")
	}
}

func TestHolidayCache_JSONDeserialization(t *testing.T) {
	jsonData := `{
		"last_updated": "2024-01-01T12:00:00Z",
		"holidays": {
			"2024-01-01": {
				"date": "2024-01-01T00:00:00Z",
				"name": "元日"
			}
		}
	}`

	var cache HolidayCache
	err := json.Unmarshal([]byte(jsonData), &cache)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON to HolidayCache: %v", err)
	}

	expectedTime := time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
	if !cache.LastUpdated.Equal(expectedTime) {
		t.Errorf("Expected LastUpdated %v, got %v", expectedTime, cache.LastUpdated)
	}

	if len(cache.Holidays) != 1 {
		t.Errorf("Expected 1 holiday, got %d", len(cache.Holidays))
	}

	if holiday, exists := cache.Holidays["2024-01-01"]; !exists {
		t.Error("Expected holiday '2024-01-01' to exist")
	} else if holiday.Name != "元日" {
		t.Errorf("Expected holiday name '元日', got '%s'", holiday.Name)
	}
}

func TestHolidayCache_EmptyCache(t *testing.T) {
	cache := HolidayCache{
		LastUpdated: time.Now(),
		Holidays:    make(map[string]Holiday),
	}

	jsonData, err := json.Marshal(cache)
	if err != nil {
		t.Fatalf("Failed to marshal empty cache: %v", err)
	}

	var deserializedCache HolidayCache
	err = json.Unmarshal(jsonData, &deserializedCache)
	if err != nil {
		t.Fatalf("Failed to unmarshal empty cache: %v", err)
	}

	if len(deserializedCache.Holidays) != 0 {
		t.Errorf("Expected empty holidays map, got %d items", len(deserializedCache.Holidays))
	}
}
