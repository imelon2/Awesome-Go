package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Provider map[string]string `yaml:"provider"`
}

// YAML 파일에 데이터를 저장하는 함수
func saveConfig(config Config, filename string) error {
	file, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, file, 0644)
}

// YAML 파일에서 데이터를 읽는 함수
func loadConfig(filename string) (Config, error) {
	var config Config
	file, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// 파일이 없으면 빈 Config 반환
			return Config{Provider: make(map[string]string)}, nil
		}
		return config, err
	}
	err = yaml.Unmarshal(file, &config)
	return config, err
}

// 데이터를 추가하는 함수
func addItem(config *Config, key, value string) {
	config.Provider[key] = value
}

// 데이터를 삭제하는 함수
func deleteItem(config *Config, key string) {
	delete(config.Provider, key)
}

func addProvider(config *Config, key, value string) {
	if config.Provider == nil {
		config.Provider = make(map[string]string)
	}
	config.Provider[key] = value
}

func deleteProvider(config *Config, key string) {
	delete(config.Provider, key)
}

func main() {
	config, _ := loadConfig("./config.yml")

	// 데이터 추가
	addProvider(&config, "arbitrum", "http://localhost:8547")
	saveConfig(config, "config.yml")

	// provider 섹션의 항목들을 순차적으로 조회
	for key, value := range config.Provider {
		fmt.Printf("Provider: %s, URL: %s\n", key, value)
	}

	// 데이터 삭제
	deleteProvider(&config, "ethereum")
	// 삭제 후 저장
	saveConfig(config, "config.yml")

	for key, value := range config.Provider {
		fmt.Printf("Provider: %s, URL: %s\n", key, value)
	}
}
