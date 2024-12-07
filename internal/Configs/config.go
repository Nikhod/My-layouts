package Configs

import (
	"Nikcase/pkg/models"
	"encoding/json"
	"os"
)

func InitConfigs() (*models.Configs, error) {
	bytes, err := os.ReadFile("./internal/Configs/config.json")
	if err != nil {
		return nil, err
	}

	var config models.Configs
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
