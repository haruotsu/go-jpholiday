package holiday

import (
	"testing"
	"time"
)

func TestAutoInitialization(t *testing.T) {
	// パッケージの自動初期化をテスト
	// キャッシュがnilの状態から始めて、自動初期化されることを確認

	// キャッシュをリセット（テスト用）
	cacheMu.Lock()
	cache = nil
	cacheMu.Unlock()

	// IsHolidayが自動的にキャッシュを初期化することを確認
	date := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	isHoliday := IsHoliday(date)

	// 元日は祝日なので、true が返されるはず
	if !isHoliday {
		t.Error("Expected New Year's Day (2024-01-01) to be a holiday")
	}

	// キャッシュが初期化されていることを確認
	cacheMu.RLock()
	if cache == nil {
		t.Error("Cache should be initialized after calling IsHoliday")
	}
	cacheMu.RUnlock()
}

func TestGetHolidayNameAutoInit(t *testing.T) {
	// キャッシュをリセット
	cacheMu.Lock()
	cache = nil
	cacheMu.Unlock()

	// GetHolidayNameが自動的にキャッシュを初期化することを確認
	date := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	name := GetHolidayName(date)

	if name != "元日" {
		t.Errorf("Expected holiday name '元日', got '%s'", name)
	}
}

func TestGetHolidaysInYearAutoInit(t *testing.T) {
	// キャッシュをリセット
	cacheMu.Lock()
	cache = nil
	cacheMu.Unlock()

	// GetHolidaysInYearが自動的にキャッシュを初期化することを確認
	holidays := GetHolidaysInYear(2024)

	if len(holidays) == 0 {
		t.Error("Expected holidays to be found for 2024")
	}

	// 元日が含まれていることを確認
	foundNewYear := false
	for _, h := range holidays {
		if h.Date.Month() == 1 && h.Date.Day() == 1 && h.Name == "元日" {
			foundNewYear = true
			break
		}
	}

	if !foundNewYear {
		t.Error("Expected to find New Year's Day in 2024 holidays")
	}
}

func TestGetHolidaysInRangeAutoInit(t *testing.T) {
	// キャッシュをリセット
	cacheMu.Lock()
	cache = nil
	cacheMu.Unlock()

	// GetHolidaysInRangeが自動的にキャッシュを初期化することを確認
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC)
	holidays := GetHolidaysInRange(start, end)

	if len(holidays) == 0 {
		t.Error("Expected holidays to be found in January 2024")
	}

	// 元日が含まれていることを確認
	foundNewYear := false
	for _, h := range holidays {
		if h.Date.Month() == 1 && h.Date.Day() == 1 && h.Name == "元日" {
			foundNewYear = true
			break
		}
	}

	if !foundNewYear {
		t.Error("Expected to find New Year's Day in January 2024 holidays")
	}
}

func TestDefaultDataContent(t *testing.T) {
	// デフォルトデータが適切な祝日を含んでいることを確認
	defaultCache := getDefaultHolidayData()

	if defaultCache == nil {
		t.Fatal("Default cache should not be nil")
	}

	if len(defaultCache.Holidays) == 0 {
		t.Fatal("Default cache should contain holidays")
	}

	// 主要な祝日がデフォルトデータに含まれていることを確認
	expectedHolidays := map[string]string{
		"2024-01-01": "元日",
		"2024-05-03": "憲法記念日",
		"2024-05-05": "こどもの日",
		"2024-11-23": "勤労感謝の日",
	}

	for dateKey, expectedName := range expectedHolidays {
		if holiday, exists := defaultCache.Holidays[dateKey]; !exists {
			t.Errorf("Expected holiday for %s not found in default data", dateKey)
		} else if holiday.Name != expectedName {
			t.Errorf("Expected holiday name '%s' for %s, got '%s'", expectedName, dateKey, holiday.Name)
		}
	}
}

func TestEnsureInitialized(t *testing.T) {
	// キャッシュをリセット
	cacheMu.Lock()
	cache = nil
	cacheMu.Unlock()

	// EnsureInitializedを呼び出し
	EnsureInitialized()

	// キャッシュが初期化されていることを確認
	cacheMu.RLock()
	if cache == nil {
		t.Error("Cache should be initialized after calling EnsureInitialized")
	}
	cacheMu.RUnlock()
}