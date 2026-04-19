package repository

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func NewConnectDB() (*gorm.DB, error) {

	if db == nil {

		viper.SetConfigName("config")
		viper.AddConfigPath("./system")

		err := viper.ReadInConfig()
		if err != nil {
			return nil, err
		}

		godotenv.Load()

		dsn := viper.GetString("supabase.dsn")
		dsn = strings.ReplaceAll(dsn, "DB_HOST", os.Getenv("DB_HOST"))
		dsn = strings.ReplaceAll(dsn, "DB_PORT", os.Getenv("DB_PORT"))
		dsn = strings.ReplaceAll(dsn, "DB_USER", os.Getenv("DB_USER"))
		dsn = strings.ReplaceAll(dsn, "DB_PASSWORD", os.Getenv("DB_PASSWORD"))
		dsn = strings.ReplaceAll(dsn, "DB_NAME", os.Getenv("DB_NAME"))

		log.Printf("DSN = %s", dsn)

		database, err := gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		}), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		db = database
	}
	return db, nil
}
