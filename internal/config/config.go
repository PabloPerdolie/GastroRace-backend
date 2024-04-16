package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

var CONFIG Config

type Config struct {
	Env    string `yaml:"env"`
	Server struct {
		URL string `yaml:"address" env-default:"localhost:8080"`
	} `yaml:"server"`
	DB struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		Host     string `yaml:"host"`
		Port     int64  `yaml:"port"`
	} `yaml:"db"`
}

func InitConfig() error {
	if err := cleanenv.ReadConfig("config/config.yml", &CONFIG); err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("Successfully initialized config")
	return nil
}
