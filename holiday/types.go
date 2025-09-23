package holiday

import (
	"time"
)

// Holiday represents a Japanese holiday
type Holiday struct {
	Date        time.Time `json:"date"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
}

// HolidayCache represents cached holiday data
type HolidayCache struct {
	LastUpdated time.Time          `json:"last_updated"`
	Holidays    map[string]Holiday `json:"holidays"` // Key: "YYYY-MM-DD"
}
