package model

type Config struct {
	StartYear int
	EndYear   int
	CacheFile string
	DryRun    bool
	Debug     bool
	ShowHelp  bool
	ShowVersion bool
}
