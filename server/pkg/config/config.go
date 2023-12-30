package config

import (
	"log"
	"os"
	"sync"
)

type postgresConfig struct {
	DBName   string
	Username string
	Password string
	Host     string
	Port     string
}

type Config struct{
	PostgresConfig postgresConfig
	ResendApiKey string
}

var (
	singletonConfig *Config
	lock            = &sync.Mutex{}
)

func initConfig() *Config {
	var pgConf postgresConfig
	// configure postgres from env if local dev else secrets manager
		log.Println("fetching postgres db creds from local env")
		pgConf = postgresConfig{
			DBName:   os.Getenv("POSTGRES_DB_NAME"),
			Username: os.Getenv("POSTGRES_USERNAME"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Host:     os.Getenv("POSTGRES_HOST"),
			Port: 		os.Getenv("POSTGRES_PORT"),
		}
		ResendApiKey:=os.Getenv("RESEND_API_KEY")

		return &Config{PostgresConfig: pgConf,ResendApiKey: ResendApiKey}
}

// Get loads Config from env
func Get() *Config {
	if singletonConfig == nil {
		lock.Lock()
		defer lock.Unlock()
		if singletonConfig == nil {
			log.Println("creating singleton config instance")
			singletonConfig = initConfig()
		}
	}

	return singletonConfig
}
