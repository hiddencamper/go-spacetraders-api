package main

import (
	"encoding/json"
	"os"
)

type API_KEY struct {
	Key        string
	Expiration string
}

func LoadAPIKey() (*API_KEY, error) {
	file, err := os.ReadFile("api_key.txt")
	KeyData := new(API_KEY)
	if err != nil {
		return KeyData, err
	}
	err = json.Unmarshal(file, &KeyData)
	return KeyData, err
}

func SaveAPIKey(key string, expiration string) error {
	KeyData := API_KEY{
		Key:        key,
		Expiration: expiration,
	}
	file, err := json.Marshal(KeyData)
	if err != nil {
		return err
	}
	return os.WriteFile("api_key.txt", file, 0644)
}
