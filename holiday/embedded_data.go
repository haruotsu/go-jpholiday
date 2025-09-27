package holiday

import (
	"time"

	"github.com/haruotsu/go-jpholiday/model"
)

// getDefaultHolidayData returns a basic set of holiday data for fallback
func getDefaultHolidayData() *model.HolidayCache {
	now := time.Now()
	holidays := map[string]model.Holiday{
		"2024-01-01": {
			Date:        time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			Name:        "元日",
			Description: "新年の始まり",
		},
		"2024-01-08": {
			Date:        time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC),
			Name:        "成人の日",
			Description: "成人の日",
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
		"2024-02-23": {
			Date:        time.Date(2024, 2, 23, 0, 0, 0, 0, time.UTC),
			Name:        "天皇誕生日",
			Description: "天皇誕生日",
		},
		"2024-03-20": {
			Date:        time.Date(2024, 3, 20, 0, 0, 0, 0, time.UTC),
			Name:        "春分の日",
			Description: "春分の日",
		},
		"2024-04-29": {
			Date:        time.Date(2024, 4, 29, 0, 0, 0, 0, time.UTC),
			Name:        "昭和の日",
			Description: "昭和の日",
		},
		"2024-05-03": {
			Date:        time.Date(2024, 5, 3, 0, 0, 0, 0, time.UTC),
			Name:        "憲法記念日",
			Description: "憲法記念日",
		},
		"2024-05-04": {
			Date:        time.Date(2024, 5, 4, 0, 0, 0, 0, time.UTC),
			Name:        "みどりの日",
			Description: "みどりの日",
		},
		"2024-05-05": {
			Date:        time.Date(2024, 5, 5, 0, 0, 0, 0, time.UTC),
			Name:        "こどもの日",
			Description: "こどもの日",
		},
		"2024-05-06": {
			Date:        time.Date(2024, 5, 6, 0, 0, 0, 0, time.UTC),
			Name:        "振替休日",
			Description: "こどもの日の振替休日",
		},
		"2024-07-15": {
			Date:        time.Date(2024, 7, 15, 0, 0, 0, 0, time.UTC),
			Name:        "海の日",
			Description: "海の日",
		},
		"2024-08-11": {
			Date:        time.Date(2024, 8, 11, 0, 0, 0, 0, time.UTC),
			Name:        "山の日",
			Description: "山の日",
		},
		"2024-08-12": {
			Date:        time.Date(2024, 8, 12, 0, 0, 0, 0, time.UTC),
			Name:        "振替休日",
			Description: "山の日の振替休日",
		},
		"2024-09-16": {
			Date:        time.Date(2024, 9, 16, 0, 0, 0, 0, time.UTC),
			Name:        "敬老の日",
			Description: "敬老の日",
		},
		"2024-09-22": {
			Date:        time.Date(2024, 9, 22, 0, 0, 0, 0, time.UTC),
			Name:        "秋分の日",
			Description: "秋分の日",
		},
		"2024-09-23": {
			Date:        time.Date(2024, 9, 23, 0, 0, 0, 0, time.UTC),
			Name:        "振替休日",
			Description: "秋分の日の振替休日",
		},
		"2024-10-14": {
			Date:        time.Date(2024, 10, 14, 0, 0, 0, 0, time.UTC),
			Name:        "スポーツの日",
			Description: "スポーツの日",
		},
		"2024-11-03": {
			Date:        time.Date(2024, 11, 3, 0, 0, 0, 0, time.UTC),
			Name:        "文化の日",
			Description: "文化の日",
		},
		"2024-11-04": {
			Date:        time.Date(2024, 11, 4, 0, 0, 0, 0, time.UTC),
			Name:        "振替休日",
			Description: "文化の日の振替休日",
		},
		"2024-11-23": {
			Date:        time.Date(2024, 11, 23, 0, 0, 0, 0, time.UTC),
			Name:        "勤労感謝の日",
			Description: "勤労感謝の日",
		},
	}

	return &model.HolidayCache{
		LastUpdated: now,
		Holidays:    holidays,
	}
}
