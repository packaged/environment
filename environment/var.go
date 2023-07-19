package environment

import "os"

// Variable returns the value of an environment variable, or a default value if not set
func Variable(name, defaultValue string) string {
	if val := os.Getenv(name); val != "" {
		return val
	}
	return defaultValue
}
