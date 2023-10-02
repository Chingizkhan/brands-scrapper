package config

import (
	"log"
	"scrapper/utils/env"
	"sync"
)

var once sync.Once
var config *Config

type Config struct {
	Postgres
	UserAgent                            string `env:"USER_AGENT"`
	MichaelKorsWomenHandbagsCrossbodyUrl string `env:"MICHAEL_KORS_WOMEN_HANDBAGS_CROSSBODY_URL"`
}

type Postgres struct {
	Host     string `env:"POSTGRES_HOST"`
	DbName   string `env:"POSTGRES_DB"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Port     string `env:"POSTGRES_PORT"`
}

func Get() *Config {
	once.Do(func() {
		conf := &Config{}
		if err := env.Unmarshal(conf); err != nil {
			log.Fatalln(err.Error())
		}
		config = conf
	})
	return config
}
