package holiday

import (
	"fmt"
	"time"
)

// formatDateKey formats a date as "YYYY-MM-DD" for use as a map key
func formatDateKey(date time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02d", date.Year(), date.Month(), date.Day())
}
