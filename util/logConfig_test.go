package util

import (
	"os"
	"testing"
)

func TestLogConfig(t *testing.T) {
	path := "./testConfig.json"

	// Check if the config file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fatalf("Config file does not exist: %s", path)
	}

	err := ParseLoggingConfig(path)
	if err != nil {
		t.Fatalf("Failed to parse logging config: %v", err)
	}

	cfg := GetLoggingConfig()
	if cfg == nil {
		t.Fatal("GetLoggingConfig returned nil")
	}

	if level, ok := cfg.levels["config.Config"]; !ok || level == 0 {
		t.Fatalf("Config did not parse correctly: %v", cfg.levels)
	} else {
		t.Logf("Parsed log level: %d", cfg.levels["config.Config"])
	}
}
