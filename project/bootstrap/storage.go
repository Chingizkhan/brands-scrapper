package bootstrap

import (
	"github.com/fatih/color"
	"scrapper/storage"
	"scrapper/storage/model/image"
	"scrapper/storage/model/product"
)

func InitStorage() {
	color.Yellow("Start migration")
	err := storage.Get().AutoMigrate(&product.Product{})
	err = storage.Get().AutoMigrate(&image.Image{})
	if err != nil {
		panic("failed to migrate")
	}
	color.Green("Migration finished")
}
