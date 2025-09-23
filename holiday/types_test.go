package holiday

import (
	"encoding/json"
	"testing"
	"time"
)

func TestHolidayStructCreation(t *testing.T) {
	// Holiday構造体が正しく生成できる
	date := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	h := Holiday{
		Date:        date,
		Name:        "元日",
		Description: "新年の始まり",
	}

	if h.Date != date {
		t.Errorf("expected date %v, got %v", date, h.Date)
	}

	if h.Name != "元日" {
		t.Errorf("expected name '元日', got '%s'", h.Name)
	}

	if h.Description != "新年の始まり" {
		t.Errorf("expected description '新年の始まり', got '%s'", h.Description)
	}
}

func TestHolidayJSONSerialization(t *testing.T) {
	// Holiday構造体をJSONにシリアライズできる
	date := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	h := Holiday{
		Date:        date,
		Name:        "元日",
		Description: "新年の始まり",
	}

	jsonData, err := json.Marshal(h)
	if err != nil {
		t.Fatalf("failed to marshal Holiday: %v", err)
	}

	// JSONが期待通りの構造を持つことを確認
	var result map[string]interface{}
	if err := json.Unmarshal(jsonData, &result); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	if _, ok := result["date"]; !ok {
		t.Error("JSON does not contain 'date' field")
	}

	if name, ok := result["name"].(string); !ok || name != "元日" {
		t.Errorf("expected name '元日', got '%v'", result["name"])
	}

	if desc, ok := result["description"].(string); !ok || desc != "新年の始まり" {
		t.Errorf("expected description '新年の始まり', got '%v'", result["description"])
	}
}

func TestHolidayJSONDeserialization(t *testing.T) {
	// JSONからHoliday構造体にデシリアライズできる
	jsonStr := `{
		"date": "2024-01-01T00:00:00Z",
		"name": "元日",
		"description": "新年の始まり"
	}`

	var h Holiday
	if err := json.Unmarshal([]byte(jsonStr), &h); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	expectedDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	if !h.Date.Equal(expectedDate) {
		t.Errorf("expected date %v, got %v", expectedDate, h.Date)
	}

	if h.Name != "元日" {
		t.Errorf("expected name '元日', got '%s'", h.Name)
	}

	if h.Description != "新年の始まり" {
		t.Errorf("expected description '新年の始まり', got '%s'", h.Description)
	}
}

func TestHolidayCacheStructCreation(t *testing.T) {
	// HolidayCache構造体が正しく生成できる
	now := time.Now()
	cache := HolidayCache{
		LastUpdated: now,
		Holidays:    make(map[string]Holiday),
	}

	if !cache.LastUpdated.Equal(now) {
		t.Errorf("expected LastUpdated %v, got %v", now, cache.LastUpdated)
	}

	if cache.Holidays == nil {
		t.Error("Holidays map should not be nil")
	}

	if len(cache.Holidays) != 0 {
		t.Errorf("expected empty Holidays map, got %d items", len(cache.Holidays))
	}
}

func TestHolidayCacheJSONSerializationAndDeserialization(t *testing.T) {
	// HolidayCache構造体のJSONシリアライズ/デシリアライズが動作する
	now := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	holiday := Holiday{
		Date:        now,
		Name:        "元日",
		Description: "新年の始まり",
	}

	cache := HolidayCache{
		LastUpdated: now,
		Holidays: map[string]Holiday{
			"2024-01-01": holiday,
		},
	}

	// シリアライズ
	jsonData, err := json.Marshal(cache)
	if err != nil {
		t.Fatalf("failed to marshal HolidayCache: %v", err)
	}

	// デシリアライズ
	var deserializedCache HolidayCache
	if err := json.Unmarshal(jsonData, &deserializedCache); err != nil {
		t.Fatalf("failed to unmarshal HolidayCache: %v", err)
	}

	// 値の検証
	if !deserializedCache.LastUpdated.Equal(cache.LastUpdated) {
		t.Errorf("expected LastUpdated %v, got %v", cache.LastUpdated, deserializedCache.LastUpdated)
	}

	if len(deserializedCache.Holidays) != 1 {
		t.Errorf("expected 1 holiday, got %d", len(deserializedCache.Holidays))
	}

	if h, ok := deserializedCache.Holidays["2024-01-01"]; !ok {
		t.Error("expected holiday for key '2024-01-01' not found")
	} else {
		if !h.Date.Equal(holiday.Date) {
			t.Errorf("expected date %v, got %v", holiday.Date, h.Date)
		}
		if h.Name != holiday.Name {
			t.Errorf("expected name '%s', got '%s'", holiday.Name, h.Name)
		}
		if h.Description != holiday.Description {
			t.Errorf("expected description '%s', got '%s'", holiday.Description, h.Description)
		}
	}
}
