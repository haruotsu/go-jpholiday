package holiday

const (
	// File permissions for cache operations
	DirPermission  = 0755
	FilePermission = 0644
)

// Default file paths
const (
	DefaultCacheFile = "data/holidays.json"
)

// Year range configuration
const (
	DefaultYearRange = 5  // ±5 years from current year
	MaxYearRange     = 15 // Maximum allowed year range
)
