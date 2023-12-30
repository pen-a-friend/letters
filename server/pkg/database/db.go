package database

import (
	"fmt"
	"log"
	"pen-a-friend/pkg/config"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	singletonDB *gorm.DB
	lock        = &sync.Mutex{}
)

func initDB() *gorm.DB {
	var err error
	dbConf := config.Get().PostgresConfig
	dsn := fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConf.Port, dbConf.Host, dbConf.Username, dbConf.Password, dbConf.DBName)
	singletonDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = singletonDB.AutoMigrate(&Users{},&Mail{})
	if err != nil {
		panic("failed to migrate the database.")
	}

	return singletonDB
}

func GetDB() *gorm.DB {
	if singletonDB == nil {
		lock.Lock()
		defer lock.Unlock()
		if singletonDB == nil {
			log.Println("creating singleton gorm db instance")
			singletonDB = initDB()
		}
	}
	return singletonDB
}
