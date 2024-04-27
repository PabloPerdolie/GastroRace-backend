package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

var CONFIG Config

type Config struct {
	Env    string `yaml:"env"`
	Server struct {
		URL string `yaml:"url" env-default:"0.0.0.0:3001"`
	} `yaml:"server"`
	DB struct {
		Name string `yaml:"name"`
		Url  string `yaml:"url"`
	} `yaml:"db"`
}

func InitConfig() error {
	if err := cleanenv.ReadConfig("usr/local/src/config/config.yml", &CONFIG); err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("Successfully initialized config")
	return nil
}
