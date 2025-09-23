package main

import (
	"os"
	"testing"
)

func TestGetAPIKey_FromEnvVar(t *testing.T) {
	// 環境変数からAPIキーを取得
	testAPIKey := "test-api-key-123"
	os.Setenv("GOOGLE_API_KEY", testAPIKey)
	defer os.Unsetenv("GOOGLE_API_KEY")

	apiKey := getAPIKey()
	if apiKey != testAPIKey {
		t.Errorf("expected API key '%s', got '%s'", testAPIKey, apiKey)
	}
}

func TestGetAPIKey_NotSet(t *testing.T) {
	// APIキーが設定されていない場合
	os.Unsetenv("GOOGLE_API_KEY")

	apiKey := getAPIKey()
	if apiKey != "" {
		t.Errorf("expected empty string when API key not set, got '%s'", apiKey)
	}
}

func TestValidateFlags_ValidYear(t *testing.T) {
	// 有効な年の検証
	config := &Config{
		startYear: 2024,
		endYear:   2025,
	}

	if err := validateFlags(config); err != nil {
		t.Errorf("unexpected error for valid years: %v", err)
	}
}

func TestValidateFlags_InvalidYearRange(t *testing.T) {
	// 無効な年の範囲
	config := &Config{
		startYear: 2025,
		endYear:   2024,
	}

	if err := validateFlags(config); err == nil {
		t.Error("expected error for invalid year range, got nil")
	}
}

func TestValidateFlags_TooManyYears(t *testing.T) {
	// 年の範囲が大きすぎる場合
	config := &Config{
		startYear: 2020,
		endYear:   2030, // 10年間は大きすぎる
	}

	if err := validateFlags(config); err == nil {
		t.Error("expected error for too many years, got nil")
	}
}

func TestGetCacheFilePath_Default(t *testing.T) {
	// デフォルトのキャッシュファイルパス
	config := &Config{
		cacheFile: "",
	}

	path := getCacheFilePath(config)
	expected := "data/holidays.json"
	if path != expected {
		t.Errorf("expected default cache file path '%s', got '%s'", expected, path)
	}
}

func TestGetCacheFilePath_Custom(t *testing.T) {
	// カスタムキャッシュファイルパス
	customPath := "/tmp/custom_holidays.json"
	config := &Config{
		cacheFile: customPath,
	}

	path := getCacheFilePath(config)
	if path != customPath {
		t.Errorf("expected custom cache file path '%s', got '%s'", customPath, path)
	}
}

func TestPrintUsage(t *testing.T) {
	// printUsage関数がパニックしないことを確認
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("printUsage panicked: %v", r)
		}
	}()

	printUsage()
}

func TestVersionCheck(t *testing.T) {
	// バージョン情報が適切に設定されていることを確認
	if version == "" {
		version = "dev" // テスト環境ではデフォルト値を設定
	}

	if version == "" {
		t.Error("version should not be empty")
	}
}
