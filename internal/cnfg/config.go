package cnfg

// Config конфигурация бенчмарка
type Config struct {
	GinPort int
	// FiberPort           int
	// Host                string
	// ResultsDir          string
	// WarmupRequests      int
	// Measurements        int
	// ConcurrentLevels    []int
	MetricsUpdateTimeMS int
	// LogsFilename        string
}

// NewDefaultConfig создает конфигурацию по умолчанию
func NewDefaultConfig() *Config {
	return &Config{
		GinPort: 8080,
		// FiberPort:           8082,
		// Host:                "localhost",
		// ResultsDir:          "results",
		// WarmupRequests:      10,
		// Measurements:        100,
		// ConcurrentLevels:    []int{1, 5, 10, 25, 50, 75, 100, 150, 200, 300},
		MetricsUpdateTimeMS: 1000,
		// LogsFilename:        "./metrics_data/logs/metrics_log.msgpack",
	}
}
