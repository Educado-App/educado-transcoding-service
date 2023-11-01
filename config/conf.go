package config

import (
	"os"
)

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

var (
	// Environment Configs
	Environment = getEnv("ENV", "development")

	// General Server Configs
	APIPort = getEnv("API_PORT", "8080")

	// Cloud Storage (GCP)
	GCPBucketName  = getEnv("BUCKET_NAME", "")
	GCPCredentials = getEnv("BUCKET_CREDS", "")
)
