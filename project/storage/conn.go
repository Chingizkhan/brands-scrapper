package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"scrapper/config"
	"sync"
)

var conn *gorm.DB
var once sync.Once

func connect() {
	once.Do(func() {
		conf := config.Get().Postgres
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			conf.Host,
			conf.User,
			conf.Password,
			conf.DbName,
			conf.Port,
		)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		conn = db
	})
}

func Get() *gorm.DB {
	if conn == nil {
		connect()
	}
	return conn
}
