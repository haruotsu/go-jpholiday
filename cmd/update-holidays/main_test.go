package main

import (
	"os"
	"testing"

	"github.com/haruotsu/go-jpholiday/model"
)

func TestGetAPIKey_FromEnvVar(t *testing.T) {
	testAPIKey := "test-api-key-123"
	os.Setenv("GOOGLE_API_KEY", testAPIKey)
	defer os.Unsetenv("GOOGLE_API_KEY")

	apiKey := getAPIKey()
	if apiKey != testAPIKey {
		t.Errorf("expected API key '%s', got '%s'", testAPIKey, apiKey)
	}
}

func TestGetAPIKey_NotSet(t *testing.T) {
	os.Unsetenv("GOOGLE_API_KEY")

	apiKey := getAPIKey()
	if apiKey != "" {
		t.Errorf("expected empty string when API key not set, got '%s'", apiKey)
	}
}

func TestValidateFlags_ValidYear(t *testing.T) {
	config := &model.Config{
		StartYear: 2024,
		EndYear:   2025,
	}

	if err := validateFlags(config); err != nil {
		t.Errorf("unexpected error for valid years: %v", err)
	}
}

func TestValidateFlags_InvalidYearRange(t *testing.T) {
	config := &model.Config{
		StartYear: 2025,
		EndYear:   2024,
	}

	if err := validateFlags(config); err == nil {
		t.Error("expected error for invalid year range, got nil")
	}
}

func TestValidateFlags_TooManyYears(t *testing.T) {
	config := &model.Config{
		StartYear: 2020,
		EndYear:   2036,
	}

	if err := validateFlags(config); err == nil {
		t.Error("expected error for too many years, got nil")
	}
}

func TestGetCacheFilePath_Default(t *testing.T) {
	config := &model.Config{
		CacheFile: "",
	}

	path := getCacheFilePath(config)
	expected := "data/holidays.json"
	if path != expected {
		t.Errorf("expected default cache file path '%s', got '%s'", expected, path)
	}
}

func TestGetCacheFilePath_Custom(t *testing.T) {
	customPath := "/tmp/custom_holidays.json"
	config := &model.Config{
		CacheFile: customPath,
	}

	path := getCacheFilePath(config)
	if path != customPath {
		t.Errorf("expected custom cache file path '%s', got '%s'", customPath, path)
	}
}

func TestPrintUsage(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("printUsage panicked: %v", r)
		}
	}()

	printUsage()
}

func TestVersionCheck(t *testing.T) {
	// version変数が存在し、デフォルト値を持つことを確認
	if version == "" {
		t.Error("version should have default value 'dev'")
	}

	if version != "dev" {
		t.Errorf("expected version to be 'dev', got '%s'", version)
	}
}
