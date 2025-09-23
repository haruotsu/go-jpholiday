package holiday

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestLoadCache_ValidFile(t *testing.T) {
	// LoadCache: JSONファイルから読み込める
	// Create a temporary cache file
	tmpDir, err := os.MkdirTemp("", "holiday_cache_test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	cacheFile := filepath.Join(tmpDir, "holidays.json")

	// Create test cache data
	testCache := HolidayCache{
		LastUpdated: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		Holidays: map[string]Holiday{
			"2024-01-01": {
				Date:        time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				Name:        "元日",
				Description: "新年の始まり",
			},
		},
	}

	// Write test data to file
	data, err := json.Marshal(testCache)
	if err != nil {
		t.Fatalf("failed to marshal test data: %v", err)
	}

	if err := os.WriteFile(cacheFile, data, 0644); err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	// Test LoadCache
	loadedCache, err := LoadCache(cacheFile)
	if err != nil {
		t.Fatalf("LoadCache failed: %v", err)
	}

	if !loadedCache.LastUpdated.Equal(testCache.LastUpdated) {
		t.Errorf("expected LastUpdated %v, got %v", testCache.LastUpdated, loadedCache.LastUpdated)
	}

	if len(loadedCache.Holidays) != 1 {
		t.Errorf("expected 1 holiday, got %d", len(loadedCache.Holidays))
	}

	if holiday, ok := loadedCache.Holidays["2024-01-01"]; !ok {
		t.Error("expected holiday for 2024-01-01 not found")
	} else {
		if holiday.Name != "元日" {
			t.Errorf("expected holiday name '元日', got '%s'", holiday.Name)
		}
	}
}

func TestLoadCache_FileNotExists(t *testing.T) {
	// LoadCache: ファイルが存在しない場合のエラー処理
	nonExistentFile := "/tmp/nonexistent_cache.json"

	_, err := LoadCache(nonExistentFile)
	if err == nil {
		t.Error("expected error for non-existent file, got nil")
	}

	if !os.IsNotExist(err) {
		t.Errorf("expected os.IsNotExist error, got %v", err)
	}
}

func TestLoadCache_CorruptedJSON(t *testing.T) {
	// LoadCache: 破損したJSONファイルのエラー処理
	tmpDir, err := os.MkdirTemp("", "holiday_cache_test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	cacheFile := filepath.Join(tmpDir, "corrupted.json")

	// Write invalid JSON
	if err := os.WriteFile(cacheFile, []byte("{invalid json"), 0644); err != nil {
		t.Fatalf("failed to write corrupted file: %v", err)
	}

	_, err = LoadCache(cacheFile)
	if err == nil {
		t.Error("expected error for corrupted JSON, got nil")
	}
}

func TestSaveCache_ValidData(t *testing.T) {
	// SaveCache: JSONファイルに保存できる
	tmpDir, err := os.MkdirTemp("", "holiday_cache_test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	cacheFile := filepath.Join(tmpDir, "save_test.json")

	testCache := &HolidayCache{
		LastUpdated: time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
		Holidays: map[string]Holiday{
			"2024-02-11": {
				Date:        time.Date(2024, 2, 11, 0, 0, 0, 0, time.UTC),
				Name:        "建国記念の日",
				Description: "建国記念の日",
			},
		},
	}

	// Test SaveCache
	if err := SaveCache(cacheFile, testCache); err != nil {
		t.Fatalf("SaveCache failed: %v", err)
	}

	// Verify file was created and has correct content
	data, err := os.ReadFile(cacheFile)
	if err != nil {
		t.Fatalf("failed to read saved file: %v", err)
	}

	var savedCache HolidayCache
	if err := json.Unmarshal(data, &savedCache); err != nil {
		t.Fatalf("failed to unmarshal saved data: %v", err)
	}

	if !savedCache.LastUpdated.Equal(testCache.LastUpdated) {
		t.Errorf("expected LastUpdated %v, got %v", testCache.LastUpdated, savedCache.LastUpdated)
	}

	if len(savedCache.Holidays) != 1 {
		t.Errorf("expected 1 holiday, got %d", len(savedCache.Holidays))
	}
}

func TestSaveCache_CreateDirectory(t *testing.T) {
	// SaveCache: ディレクトリが存在しない場合に作成する
	tmpDir, err := os.MkdirTemp("", "holiday_cache_test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create nested directory path that doesn't exist
	cacheFile := filepath.Join(tmpDir, "new", "directory", "cache.json")

	testCache := &HolidayCache{
		LastUpdated: time.Now(),
		Holidays:    make(map[string]Holiday),
	}

	// Test SaveCache with non-existent directory
	if err := SaveCache(cacheFile, testCache); err != nil {
		t.Fatalf("SaveCache failed to create directory: %v", err)
	}

	// Verify file was created
	if _, err := os.Stat(cacheFile); err != nil {
		t.Errorf("cache file was not created: %v", err)
	}
}

func TestUpdateCache_NewData(t *testing.T) {
	// UpdateCache: 新しいデータでキャッシュを更新できる
	// Setup initial cache
	cache := &HolidayCache{
		LastUpdated: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		Holidays: map[string]Holiday{
			"2024-01-01": {
				Date:        time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				Name:        "元日",
				Description: "新年の始まり",
			},
		},
	}

	// New holidays to add
	newHolidays := []Holiday{
		{
			Date:        time.Date(2024, 2, 11, 0, 0, 0, 0, time.UTC),
			Name:        "建国記念の日",
			Description: "建国記念の日",
		},
		{
			Date:        time.Date(2024, 5, 3, 0, 0, 0, 0, time.UTC),
			Name:        "憲法記念日",
			Description: "憲法記念日",
		},
	}

	updateTime := time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC)

	// Test UpdateCache
	UpdateCache(cache, newHolidays, updateTime)

	// Verify cache was updated
	if !cache.LastUpdated.Equal(updateTime) {
		t.Errorf("expected LastUpdated %v, got %v", updateTime, cache.LastUpdated)
	}

	expectedCount := 3 // 1 original + 2 new
	if len(cache.Holidays) != expectedCount {
		t.Errorf("expected %d holidays, got %d", expectedCount, len(cache.Holidays))
	}

	// Verify new holidays were added
	if _, ok := cache.Holidays["2024-02-11"]; !ok {
		t.Error("expected holiday for 2024-02-11 not found")
	}

	if _, ok := cache.Holidays["2024-05-03"]; !ok {
		t.Error("expected holiday for 2024-05-03 not found")
	}
}

func TestIsStale_OldCache(t *testing.T) {
	// IsStale: 古いキャッシュを検出できる
	oldTime := time.Now().AddDate(0, -2, 0) // 2 months ago
	cache := &HolidayCache{
		LastUpdated: oldTime,
		Holidays:    make(map[string]Holiday),
	}

	if !IsStale(cache, 30*24*time.Hour) { // 30 days threshold
		t.Error("expected cache to be stale, got false")
	}
}

func TestIsStale_FreshCache(t *testing.T) {
	// IsStale: 新しいキャッシュは古くないと判定
	recentTime := time.Now().Add(-1 * time.Hour) // 1 hour ago
	cache := &HolidayCache{
		LastUpdated: recentTime,
		Holidays:    make(map[string]Holiday),
	}

	if IsStale(cache, 30*24*time.Hour) { // 30 days threshold
		t.Error("expected cache to be fresh, got stale")
	}
}

func TestIsStale_NilCache(t *testing.T) {
	// IsStale: nilキャッシュは古いと判定
	if !IsStale(nil, 30*24*time.Hour) {
		t.Error("expected nil cache to be stale, got false")
	}
}