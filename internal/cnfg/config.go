package cnfg

// Config конфигурация бенчмарка
type Config struct {
	GinPort          int
	FiberPort        int
	Host             string
	ResultsDir       string
	WarmupRequests   int
	Measurements     int
	ConcurrentLevels []int
}

// NewDefaultConfig создает конфигурацию по умолчанию
func NewDefaultConfig() *Config {
	return &Config{
		GinPort:          8081,
		FiberPort:        8082,
		Host:             "localhost",
		ResultsDir:       "results",
		WarmupRequests:   10,
		Measurements:     100,
		ConcurrentLevels: []int{1, 5, 10, 25, 50, 75, 100, 150, 200, 300},
	}
}
