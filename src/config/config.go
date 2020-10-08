package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

type config struct {
	ORM orm `json:"orm"`
}

type orm struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DB       string `json:"db"`
}

var instance *config
var err error
var once sync.Once

const configPath string = "./config/config.json"

func GetInstance() (*config, error) {
	once.Do(func() {
		instance, err = fillConfig()
	})
	return instance, err
}

func fillConfig() (*config, error) {
	byt, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error read file config %v: %v", configPath, err)
	}
	config := config{}
	if err := json.Unmarshal(byt, &config); err != nil {
		return nil, fmt.Errorf("file config is not valid %v: %v", configPath, err)
	}
	return &config, nil
}
