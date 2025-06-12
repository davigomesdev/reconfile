package env_config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	AppEnv              string
	Port                string
	DatabaseHost        string
	DatabasePort        string
	DatabaseName        string
	DatabaseUser        string
	DatabasePassword    string
	DatabaseSSLMode     string
	DatabaseTimezone    string
	JWTSecretKey        string
	JWTExpiresAccessIn  time.Duration
	JWTExpiresRefreshIn time.Duration
}

func LoadConfig() *EnvConfig {
	env := os.Getenv("APP_ENV")

	if env == "" {
		env = "development"
	}

	envFile := ".env." + env

	if err := godotenv.Load(envFile); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &EnvConfig{
		AppEnv:              env,
		Port:                getEnv("PORT", "3000"),
		DatabaseHost:        os.Getenv("DATABASE_HOST"),
		DatabasePort:        os.Getenv("DATABASE_PORT"),
		DatabaseName:        os.Getenv("DATABASE_NAME"),
		DatabaseUser:        os.Getenv("DATABASE_USER"),
		DatabasePassword:    os.Getenv("DATABASE_PASSWORD"),
		DatabaseSSLMode:     getEnv("DATABASE_SSLMODE", "disable"),
		DatabaseTimezone:    getEnv("DATABASE_TIMEZONE", "UTC"),
		JWTSecretKey:        os.Getenv("JWT_SECRET_KEY"),
		JWTExpiresAccessIn:  parseDuration("JWT_EXPIRES_ACCESS_IN", "15m"),
		JWTExpiresRefreshIn: parseDuration("JWT_EXPIRES_REFRESH_IN", "168h"),
	}
}

func getEnv(key string, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}

func parseDuration(key string, defaultVal string) time.Duration {
	val := os.Getenv(key)
	if val == "" {
		val = defaultVal
	}
	dur, err := time.ParseDuration(val)
	if err != nil {
		log.Printf("Erro ao fazer parse de %s: %v, usando valor padrão: %s", key, err, defaultVal)
		dur, _ = time.ParseDuration(defaultVal)
	}
	return dur
}
