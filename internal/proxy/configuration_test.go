package proxy

import (
	"os"
	"testing"
)

func TestGetConfiguration(t *testing.T) {
	_ = os.Setenv(prefix+"CLIENT_ID", "test")
	_ = os.Setenv(prefix+"CLIENT_SECRET", "test")
	_ = os.Setenv(prefix+"WELL_KNOWN_URL", "test")
	_ = os.Setenv(prefix+"UPSTREAM_SERVER", "test")

	_, err := getConfiguration()
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestGetEnvWithDefaultValue(t *testing.T) {
	key := "_ABC_"
	value := "test"
	defaultValue := "default"
	returned := getEnvWithDefaultValue(key, defaultValue)
	if returned != defaultValue {
		t.Fatalf("Expected: %s, Got: %s", defaultValue, returned)
	}
	_ = os.Setenv(key, value)
	returned = getEnvWithDefaultValue(key, value)
	if returned != value {
		t.Fatalf("Expected: %s, Got: %s", value, returned)
	}
}
