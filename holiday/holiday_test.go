package holiday

import (
	"testing"
	"time"
)

// テスト用のモックデータを作成
func setupTestHolidayCache() *HolidayCache {
	cache := &HolidayCache{
		LastUpdated: time.Now(),
		Holidays: map[string]Holiday{
			"2024-01-01": {
				Date:        time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				Name:        "元日",
				Description: "新年の始まり",
			},
			"2024-02-11": {
				Date:        time.Date(2024, 2, 11, 0, 0, 0, 0, time.UTC),
				Name:        "建国記念の日",
				Description: "建国記念の日",
			},
			"2024-02-12": {
				Date:        time.Date(2024, 2, 12, 0, 0, 0, 0, time.UTC),
				Name:        "振替休日",
				Description: "建国記念の日の振替休日",
			},
			"2024-05-03": {
				Date:        time.Date(2024, 5, 3, 0, 0, 0, 0, time.UTC),
				Name:        "憲法記念日",
				Description: "憲法記念日",
			},
			"2025-01-01": {
				Date:        time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				Name:        "元日",
				Description: "新年の始まり",
			},
		},
	}
	return cache
}

func TestIsHoliday_WithHolidayDate(t *testing.T) {
	// IsHoliday: 祝日の日付でtrueを返す
	cache := setupTestHolidayCache()
	SetCache(cache)

	testDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	if !IsHoliday(testDate) {
		t.Errorf("expected true for holiday date %v, got false", testDate)
	}
}

func TestIsHoliday_WithWeekday(t *testing.T) {
	// IsHoliday: 平日の日付でfalseを返す
	cache := setupTestHolidayCache()
	SetCache(cache)

	testDate := time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC) // 2024年1月2日（火）
	if IsHoliday(testDate) {
		t.Errorf("expected false for weekday %v, got true", testDate)
	}
}

func TestIsHoliday_WithWeekend(t *testing.T) {
	// IsHoliday: 土日でもfalseを返す（祝日ではないため）
	cache := setupTestHolidayCache()
	SetCache(cache)

	// 2024年1月6日（土）
	saturday := time.Date(2024, 1, 6, 0, 0, 0, 0, time.UTC)
	if IsHoliday(saturday) {
		t.Errorf("expected false for Saturday %v, got true", saturday)
	}

	// 2024年1月7日（日）
	sunday := time.Date(2024, 1, 7, 0, 0, 0, 0, time.UTC)
	if IsHoliday(sunday) {
		t.Errorf("expected false for Sunday %v, got true", sunday)
	}
}

func TestGetHolidayName_WithHolidayDate(t *testing.T) {
	// GetHolidayName: 祝日の名前を返す
	cache := setupTestHolidayCache()
	SetCache(cache)

	testDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	name := GetHolidayName(testDate)
	expected := "元日"
	if name != expected {
		t.Errorf("expected holiday name '%s', got '%s'", expected, name)
	}
}

func TestGetHolidayName_WithWeekday(t *testing.T) {
	// GetHolidayName: 平日で空文字を返す
	cache := setupTestHolidayCache()
	SetCache(cache)

	testDate := time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)
	name := GetHolidayName(testDate)
	if name != "" {
		t.Errorf("expected empty string for weekday, got '%s'", name)
	}
}

func TestGetHolidaysInYear_WithDataExists(t *testing.T) {
	// GetHolidaysInYear: 指定年の全祝日を返す
	cache := setupTestHolidayCache()
	SetCache(cache)

	holidays := GetHolidaysInYear(2024)
	if len(holidays) != 4 { // 2024年のテストデータには4つの祝日がある
		t.Errorf("expected 4 holidays in 2024, got %d", len(holidays))
	}

	// 最初の祝日が元日であることを確認
	if len(holidays) > 0 && holidays[0].Name != "元日" {
		t.Errorf("expected first holiday to be '元日', got '%s'", holidays[0].Name)
	}
}

func TestGetHolidaysInYear_WithNoData(t *testing.T) {
	// GetHolidaysInYear: データがない年で空配列を返す
	cache := setupTestHolidayCache()
	SetCache(cache)

	holidays := GetHolidaysInYear(2023) // 2023年のデータはない
	if len(holidays) != 0 {
		t.Errorf("expected empty array for year with no data, got %d holidays", len(holidays))
	}
}

func TestGetHolidaysInRange_WithHolidays(t *testing.T) {
	// GetHolidaysInRange: 指定期間の祝日を返す
	cache := setupTestHolidayCache()
	SetCache(cache)

	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)

	holidays := GetHolidaysInRange(start, end)
	if len(holidays) != 3 { // 1月1日、2月11日、2月12日の3つ
		t.Errorf("expected 3 holidays in range, got %d", len(holidays))
	}
}

func TestGetHolidaysInRange_WithNoHolidays(t *testing.T) {
	// GetHolidaysInRange: 期間内に祝日がない場合空配列を返す
	cache := setupTestHolidayCache()
	SetCache(cache)

	start := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 4, 30, 0, 0, 0, 0, time.UTC)

	holidays := GetHolidaysInRange(start, end)
	if len(holidays) != 0 {
		t.Errorf("expected empty array when no holidays in range, got %d holidays", len(holidays))
	}
}