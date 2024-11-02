package main

import (
	"encoding/json"
	"os"
)

var API_KEY string

func LoadAPIKey() error {
	file, err := os.ReadFile("api_key.txt")
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &API_KEY)
	return err
}

func SaveAPIKey(key string) error {
	file, err := json.Marshal(key)
	if err != nil {
		return err
	}
	return os.WriteFile("api_key.txt", file, 0644)
}
