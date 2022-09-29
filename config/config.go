package config

import "os"

type db struct {
	Name     string
	Hostname string
	Port     string
	Username string
	Pass     string
	SSLMode  string
}

type Config struct {
	Hostname   string
	Port       string
	DB         db
}

func NewConfig() *Config {
	return &Config{
		DB: db{
			Name: getEnv("COURSE_API_DB_NAME", "courses"),
			Hostname: getEnv("COURSE_API_DB_HOSTNAME", "localhost"),
			Port: getEnv("COURSE_API_DB_PORT", "5432"),
			Username: getEnv("COURSE_API_DB_USER", "courses_api"),
			Pass: getEnv("COURSE_API_DB_PASSWORD", "password"),
			SSLMode: getEnv("COURSE_API_DB_SSLMODE", "disable"),
		},
		Hostname: getEnv("COURSE_API_HOSTNAME", "localhost"),
		Port: getEnv("COURSE_API_PORT", "8000"),
	}
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}

	return defaultValue
}
