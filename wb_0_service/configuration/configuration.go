package configuration

type Config struct {
	// Database
	DbName     string `env:"DB_NAME"`
	DbHost     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DbPort     string `env:"DB_PORT" envDefault:"5432"`
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`

	// Server
	Port string `env:"PORT" envDefault:"8080"`

	// Cache configuration
	OrderCacheCapacity int

	// Logging
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`

	// NATS
	NatsUrl string `env:"NATS_URL" envDefault:"nats://127.0.0.1:4222"`

	// Other
	Environment string `env:"ENVIRONMENT" envDefault:"dev"`
}
