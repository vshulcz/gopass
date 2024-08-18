package internal

import (
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	DatabasePath  string
	MasterKeyPath string
}

func LoadConfig() *Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get home directory: %v", err)
	}

	config := &Config{
		DatabasePath:  filepath.Join(homeDir, ".gopass_manager.db"),
		MasterKeyPath: filepath.Join(homeDir, ".gopass_key"),
	}

	return config
}
