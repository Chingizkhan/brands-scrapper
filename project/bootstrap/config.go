package bootstrap

import (
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"log"
	"scrapper/config"
)

func InitConfig() {
	color.Yellow("Init config")
	config.Get()
	color.Green("Config initialized")
}

func InitEnv(envDir string) {
	color.Yellow("Loading .env file...")
	err := godotenv.Load(envDir)
	if err != nil {
		log.Panicf("loading .env file failed")
	}
	color.Green("Initialized .env file")
}
