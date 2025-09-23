package fetcher

import (
	"errors"
	"testing"
	"time"

	"github.com/haruotsu/go-jpholiday/holiday"
)

func TestNewFetcher_WithAPIKey(t *testing.T) {
	// NewFetcher: APIキーでFetcherを生成できる
	apiKey := "test-api-key"
	f := NewFetcher(apiKey)

	if f == nil {
		t.Fatal("expected fetcher to be created, got nil")
		return
	}

	if f.APIKey != apiKey {
		t.Errorf("expected API key '%s', got '%s'", apiKey, f.APIKey)
	}
}

func TestNewFetcher_WithEmptyAPIKey(t *testing.T) {
	// NewFetcher: APIキーが空でエラーを返す
	f := NewFetcher("")

	if f != nil {
		t.Error("expected nil for empty API key, got fetcher")
	}
}

func TestMockFetcher_FetchHolidays(t *testing.T) {
	// モック実装が正しく動作する
	mockFetcher := &MockFetcher{
		Holidays: []holiday.Holiday{
			{
				Date:        time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				Name:        "元日",
				Description: "新年の始まり",
			},
			{
				Date:        time.Date(2024, 2, 11, 0, 0, 0, 0, time.UTC),
				Name:        "建国記念の日",
				Description: "建国記念の日",
			},
		},
	}

	holidays, err := mockFetcher.FetchHolidays(2024)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(holidays) != 2 {
		t.Errorf("expected 2 holidays, got %d", len(holidays))
	}

	if holidays[0].Name != "元日" {
		t.Errorf("expected first holiday to be '元日', got '%s'", holidays[0].Name)
	}

	if holidays[1].Name != "建国記念の日" {
		t.Errorf("expected second holiday to be '建国記念の日', got '%s'", holidays[1].Name)
	}
}

func TestMockFetcher_FetchHolidaysWithError(t *testing.T) {
	// モック実装でエラーが正しく返される
	mockFetcher := &MockFetcher{
		ShouldError: true,
		ErrorMsg:    "API error",
	}

	_, err := mockFetcher.FetchHolidays(2024)
	if err == nil {
		t.Error("expected error, got nil")
	}

	if err.Error() != "API error" {
		t.Errorf("expected error message 'API error', got '%s'", err.Error())
	}
}

func TestMockFetcher_FetchHolidaysRange(t *testing.T) {
	// FetchHolidaysRange: 複数年の祝日を取得できる
	mockFetcher := &MockFetcher{
		Holidays: []holiday.Holiday{
			{
				Date:        time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				Name:        "元日",
				Description: "新年の始まり",
			},
			{
				Date:        time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				Name:        "元日",
				Description: "新年の始まり",
			},
		},
	}

	holidays, err := mockFetcher.FetchHolidaysRange(2024, 2025)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(holidays) != 2 {
		t.Errorf("expected 2 holidays, got %d", len(holidays))
	}

	// Check that we got holidays from both years
	foundYears := make(map[int]bool)
	for _, h := range holidays {
		foundYears[h.Date.Year()] = true
	}

	if !foundYears[2024] {
		t.Error("expected holidays from 2024")
	}

	if !foundYears[2025] {
		t.Error("expected holidays from 2025")
	}
}

// テスト用のヘルパー関数群

func TestFetcherInterface(t *testing.T) {
	// Fetcherインターフェースが正しく定義されていることを確認
	// MockFetcherがインターフェースを実装していることを確認
	var fetcher Fetcher = &MockFetcher{}
	_ = fetcher
}

func TestRetryMechanism(t *testing.T) {
	// リトライ機構のテスト（実装後に有効化）
	t.Skip("Retry mechanism not implemented yet")

	mockFetcher := &MockFetcher{
		RetryCount:  2,
		ShouldError: true,
		ErrorMsg:    "temporary error",
	}

	_, err := mockFetcher.FetchHolidays(2024)
	if err == nil {
		t.Error("expected error after retries, got nil")
	}

	if mockFetcher.CallCount != 3 { // 1 initial + 2 retries
		t.Errorf("expected 3 calls (1 + 2 retries), got %d", mockFetcher.CallCount)
	}
}

func TestNetworkErrorHandling(t *testing.T) {
	// ネットワークエラー時の適切なエラー処理
	mockFetcher := &MockFetcher{
		ShouldError: true,
		ErrorMsg:    "network error",
	}

	_, err := mockFetcher.FetchHolidays(2024)
	if err == nil {
		t.Error("expected network error, got nil")
	}

	if !errors.Is(err, errors.New("network error")) {
		// For simple string comparison in this test
		if err.Error() != "network error" {
			t.Errorf("expected network error, got '%s'", err.Error())
		}
	}
}

func TestIsOfficialHoliday(t *testing.T) {
	// 公式祝日判定のテスト
	tests := []struct {
		description string
		expected    bool
		name        string
	}{
		{
			description: "祝日",
			expected:    true,
			name:        "Official holiday",
		},
		{
			description: "祭日\n祭日を非表示にするには、Google カレンダーの [設定] > [日本の祝日] に移動してください",
			expected:    false,
			name:        "Festival/cultural observance",
		},
		{
			description: "",
			expected:    false,
			name:        "Empty description",
		},
		{
			description: "some other description",
			expected:    false,
			name:        "Other description",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isOfficialHoliday(tt.description)
			if result != tt.expected {
				t.Errorf("isOfficialHoliday(%q) = %v, expected %v", tt.description, result, tt.expected)
			}
		})
	}
}
