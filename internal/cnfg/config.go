package cnfg

// Config конфигурация бенчмарка
type Config struct {
	GinPort             int
	FiberPort           int
	MetricsUpdateTimeMS int
	LogsFilename        string
}

// NewDefaultConfig создает конфигурацию по умолчанию
func NewDefaultConfig() *Config {
	return &Config{
		GinPort:             8080,
		FiberPort:           8080,
		MetricsUpdateTimeMS: 1000,
		LogsFilename:        "./metrics_data/logs/logs_time_series.txt",
	}
}
