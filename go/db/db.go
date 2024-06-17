package db

import (
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/joho/godotenv"
)

func NewDB() *azcosmos.Client {
	if os.Getenv("GO_ENV") == "dev" {
		// GO_ENVがdevの場合は、.envファイルを読み込む
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}

	// Azure Cosmos DBに接続
	endpoint := os.Getenv("COSMOS_DB_ENDPOINT")
	key := os.Getenv("COSMOS_DB_KEY")
	cred, _ := azcosmos.NewKeyCredential(key)
	client, err := azcosmos.NewClientWithKey(endpoint, cred, nil)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected")
	return client
}

/*
	// DBに接続
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
*/
