package main

import (
	"log"
	"my-shortener/repository"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./model/query",
		ModelPkgPath: "./model",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	_ = godotenv.Load()

	db, err := repository.NewConnectDB()
	if err != nil {
		log.Fatalf("❌ เชื่อมต่อ DB ไม่ได้: %v", err)
	}

	g.UseDB(db)

	urlsdoc := g.GenerateModel("urls")

	g.ApplyBasic(urlsdoc)

	g.Execute()

}

var db *gorm.DB

func NewConnectDB() (*gorm.DB, error) {

	if db == nil {

		viper.SetConfigName("config")
		viper.AddConfigPath("./system")

		err := viper.ReadInConfig()
		if err != nil {
			return nil, err
		}

		err = godotenv.Load()
		if err != nil {
			return nil, err
		}

		dsn := viper.GetString("supabase.dsn")
		dsn = strings.ReplaceAll(dsn, "DB_HOST", os.Getenv("DB_HOST"))
		dsn = strings.ReplaceAll(dsn, "DB_PORT", os.Getenv("DB_PORT"))
		dsn = strings.ReplaceAll(dsn, "DB_USER", os.Getenv("DB_USER"))
		dsn = strings.ReplaceAll(dsn, "DB_PASSWORD", os.Getenv("DB_PASSWORD"))
		dsn = strings.ReplaceAll(dsn, "DB_NAME", os.Getenv("DB_NAME"))

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			PrepareStmt: true,
		})
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}
