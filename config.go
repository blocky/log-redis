package main

import (
	"fmt"
	"io"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	RedisAddress  string
	RedisPassword string
	RedisDatabase int
}

func ReadConfig() (Config, error) {
	file, err := os.Open("config.toml")
	if err != nil {
		return Config{}, fmt.Errorf("error opening config file: %w", err)
	}
	defer func() { _ = file.Close() }()
	data, err := io.ReadAll(file)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	err = toml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, fmt.Errorf("error unmarshalling config file: %w", err)
	}
	return config, nil
}
