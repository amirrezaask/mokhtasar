package config

import "os"

func getEnv(key string, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

var DatabaseHost = getEnv("MOKHTASAR_DATABASE_HOST", "127.0.0.1")
var DatabaseUser = getEnv("MOKHTASAR_DATABASE_USER", "postgres")
var DatabasePass = getEnv("MOKHTASAR_DATABASE_PASS", "admin")
var DatabaseName = getEnv("MOKHTASAR_DATABASE_NAME", "mokhtasar")
var DatabaseSSLMode = getEnv("MOKHTASAR_DATABASE_SSL_MODE", "disable")
