package configuration

type Config struct {
	// Database
	DbName     string `env:"DB_NAME"`
	DbIp       string `env:"DB_IP" envDefault:"127.0.0.1"`
	DbPort     string `env:"DB_PORT" envDefault:"5432"`
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`

	// Server
	Host    string
	Port    string `env:"PORT" envDefault:"8080"`
	UseHttp bool

	// Cache configuration fields
	OrderCacheCapacity int

	// Logging
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`

	// NATS
	NatsUrl string `env:"NATS_URL" envDefault:"nats://127.0.0.1:4222"`

	// Other
	Environment string `env:ENVIRONMENT envDefault:"dev"`
}
