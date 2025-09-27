package holiday

import (
	"time"

	"github.com/haruotsu/go-jpholiday/model"
)

// NewTestHolidayCache creates a test holiday cache with sample data
func NewTestHolidayCache() *model.HolidayCache {
	return &model.HolidayCache{
		LastUpdated: time.Now(),
		Holidays: map[string]model.Holiday{
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
}

// NewTestHolidayCacheWithDate creates a test cache with a specific date
func NewTestHolidayCacheWithDate(date time.Time, name, description string) *model.HolidayCache {
	return &model.HolidayCache{
		LastUpdated: time.Now(),
		Holidays: map[string]model.Holiday{
			formatDateKey(date): {
				Date:        date,
				Name:        name,
				Description: description,
			},
		},
	}
}
