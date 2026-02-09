package config

type ObservabilityConfig struct {
	Environment string
	ServiceName string
	LogLevel    string
}

func (c *ObservabilityConfig) IsProduction() bool {
	return c.Environment != "dev"
}

func (c *ObservabilityConfig) GetLogLevel() string {
	switch c.Environment {
	case "dev":
		if c.LogLevel == "" {
			return "debug"
		}
	case "prod":
		if c.LogLevel == "" {
			return "info"
		}
	}

	return c.LogLevel
}
