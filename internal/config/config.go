package config

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"io"
	"os"
)

func Read() Config {
	var cfg Config
	
	fullPath, err := getConfigFilePath()
	if err != nil {
		fmt.Println(err)	
		return Config{}
	}

	jsonFile, err := os.Open(fullPath)	
	if err != nil {
		fmt.Println(err)
		return Config{}
	}

	defer jsonFile.Close()

	byteVal, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return Config{}
	}

	err = json.Unmarshal(byteVal, &cfg)
	if err != nil {
		fmt.Println(err)
		return Config{}
	}

	return cfg
}

func (c *Config) SetUser() {

}

func getConfigFilePath() (string, error) {

	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(homeDir, configFileName)	

	return fullPath, nil
}

func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		fmt.Println(err)	
		return Config{}
	}


}
