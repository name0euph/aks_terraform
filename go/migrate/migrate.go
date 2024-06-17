package main

import (
	"fmt"
	"go-rest-api/db"
)

func main() {
	//データベース接続の初期化
	db.NewDB()
	defer fmt.Println("Successfully Migrated")
}
