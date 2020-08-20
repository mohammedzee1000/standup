package config

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	configFile = "mytask_config.yaml"
	dataDir    = ".mytask"
	dataDirEnv = "MYTASK_DATA_DIR"
)

func getConfigFilePath() (string, error) {
	v, err := getDataDir()
	if err != nil {
		return "", err
	}
	fp := filepath.Join(v, configFile)
	return fp, nil
}

func getDataDir() (string, error) {
	dd := os.Getenv(dataDirEnv)
	if len(dd) <= 0 {
		u, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		dd = filepath.Join(u, dataDir)
	}
	_, err := os.Stat(dd)
	if err != nil {
		if os.IsNotExist(err) {
			err1 := os.Mkdir(dd, os.ModePerm)
			if err1 != nil {
				return "", fmt.Errorf("failed to create datadir %s: %w", dd, err1)
			}
		} else {
			return "", fmt.Errorf("failed to stat data dir %s: %w", dd, err)
		}
	}

	return dd, nil
}
