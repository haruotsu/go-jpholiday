# go-jpholiday 🎌🌸

A Go library for determining Japanese holidays using the Google Calendar API.

Holiday data is automatically updated every month via GitHub Actions, ensuring you always have the latest Japanese holiday information without any manual intervention!

[![Test](https://github.com/haruotsu/go-jpholiday/actions/workflows/test.yml/badge.svg)](https://github.com/haruotsu/go-jpholiday/actions/workflows/test.yml)
[![Update Holidays](https://github.com/haruotsu/go-jpholiday/actions/workflows/update-holidays.yml/badge.svg)](https://github.com/haruotsu/go-jpholiday/actions/workflows/update-holidays.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/haruotsu/go-jpholiday.svg)](https://pkg.go.dev/github.com/haruotsu/go-jpholiday)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/haruotsu/go-jpholiday)](https://goreportcard.com/report/github.com/haruotsu/go-jpholiday)

## Installation

```bash
go get github.com/haruotsu/go-jpholiday
```


## Quick Start

```go
package main

import (
    "fmt"
    "time"
    "github.com/haruotsu/go-jpholiday/holiday"
)

func main() {
    // Check if a date is a holiday
    date := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
    if holiday.IsHoliday(date) {
        fmt.Printf("%s is a holiday: %s\n",
            date.Format("2006-01-02"),
            holiday.GetHolidayName(date))
    }
    // Output: 2024-01-01 is a holiday: 元日

    // Get all holidays for a year
    holidays := holiday.GetHolidaysInYear(2024)
    fmt.Printf("Number of holidays in 2024: %d\n", len(holidays))

    // Get holidays in a date range (e.g., Golden Week)
    start := time.Date(2024, 4, 29, 0, 0, 0, 0, time.Local)
    end := time.Date(2024, 5, 5, 0, 0, 0, 0, time.Local)
    gwHolidays := holiday.GetHolidaysInRange(start, end)
    for _, h := range gwHolidays {
        fmt.Printf("  %s: %s\n",
            h.Date.Format("2006-01-02"),
            h.Name)
    }
}
```


## Data Management

### Holiday Data Source

This library retrieves Japanese national holiday data from the official Google Calendar API (Calendar ID: `ja.japanese.official#holiday@group.v.calendar.google.com`). This calendar is maintained by Google and contains all official Japanese public holidays, including:

- Regular national holidays (元日, 成人の日, etc.)
- Special holidays (天皇即位の日, etc.)
- Substitute holidays (振替休日)

### Data Storage & Caching
- **Data Retention**: By default, the library fetches and stores 2 years of holiday data (current year + next year)
- **Cache Format**: JSON format with date keys ("YYYY-MM-DD") for O(1) lookup performance

### Update Schedule

| Update Type | Frequency | Method | Description |
|------------|-----------|---------|-------------|
| **Automatic** | Monthly (1st day) | GitHub Actions | Ensures holiday data is always current |
| **On-demand** | Anytime | CLI tool | Update manually when needed |
| **Package Release** | With new versions | Embedded data | Pre-bundled data updated with each release |

### Data Freshness

- Holiday data is typically published by the Japanese government 1-2 years in advance
- Special holidays (like imperial ceremonies) may be announced with shorter notice
- The automatic monthly updates ensure any newly announced holidays are captured promptly

### Automatic Updates

Holiday data is automatically updated through multiple mechanisms:

1. GitHub Actions (Monthly): Runs on the 1st of every month at 00:00 UTC
2. Package Updates: New releases include the latest holiday data
3. Runtime Fallback: If no cached data exists, the library uses embedded defaults

### Installation

```bash
go install github.com/haruotsu/go-jpholiday/cmd/update-holidays@latest
```

### Usage

```bash
# Set your Google Calendar API key
export GOOGLE_API_KEY=your-google-calendar-api-key

# Fetch holidays for current and next year
update-holidays

# Specify a custom year range
update-holidays -start-year 2024 -end-year 2025

# Dry run (preview without updating)
update-holidays -dry-run

# Enable debug output
update-holidays -debug
```

### CLI Options

| Option | Description | Default |
|--------|-------------|---------|
| `-start-year` | Start year for fetching | Current year |
| `-end-year` | End year for fetching | Current year + 1 |
| `-cache-file` | Path to cache file | `data/holidays.json` |
| `-dry-run` | Preview changes without updating | `false` |
| `-debug` | Enable debug output | `false` |
| `-help, -h` | Show help message | - |
| `-version, -v` | Show version information | - |


### Setup

```bash
# Clone the repository
git clone https://github.com/haruotsu/go-jpholiday.git
cd go-jpholiday

# Install dependencies
make setup

# Run tests
make test

# Run with coverage
make test-coverage

# Run linters
make lint

# Format code
make fmt
```

### Updating Holiday Data

#### Manual Update

To update the holiday data manually:

```bash
export GOOGLE_API_KEY=your-api-key
make run
# Or directly:
go run cmd/update-holidays/main.go
```


#### Custom Year Range

You can fetch holiday data for specific years:

```bash
# Fetch 5 years of data (2024-2028)
update-holidays -start-year 2024 -end-year 2028
```

**Note**: The Google Calendar API limits requests to prevent abuse. Fetching more than 5 years at once may result in rate limiting.


## Contributing

We welcome contributions! Please see our [Contributing Guidelines](CONTRIBUTING.md) for details.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Add or update tests as needed
4. Commit your changes (`git commit -m 'Add amazing feature'`)
5. Push to your branch (`git push origin feature/amazing-feature`)
6. Open a Pull Request

For more details, see [DEVELOPMENT.md](DEVELOPMENT.md).


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Holiday data sourced from Google Calendar API (Japanese public holidays)
- Inspired by similar libraries in other languages

## Support

- [Report bugs](https://github.com/haruotsu/go-jpholiday/issues)
- [Request features](https://github.com/haruotsu/go-jpholiday/issues)
- [Read the docs](https://pkg.go.dev/github.com/haruotsu/go-jpholiday)

## Author

[@haruotsu](https://github.com/haruotsu)

---

**Note**: This library depends on the Google Calendar API for fetching holiday data. Please review the [Google Calendar API Terms of Service](https://developers.google.com/terms) and rate limits for commercial use.
