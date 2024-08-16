package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	// .envファイルを読み込む
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	// PostgresSQLに接続するためのDSN
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))

	// DB接続
	sqlDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully connected to DB")

	return sqlDb
}

func CloseDB(db *gorm.DB) {
	sqlDb, _ := db.DB()

	// DB接続のクローズ
	if err := sqlDb.Close(); err != nil {
		log.Fatalln(err)
	}
}
