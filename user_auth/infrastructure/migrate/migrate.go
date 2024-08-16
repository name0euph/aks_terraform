package main

import (
	"user_auth/domain"
	"user_auth/infrastructure"

	"fmt"
)

func main() {
	dbConn := infrastructure.NewDB()     // データベースに接続
	defer infrastructure.CloseDB(dbConn) //migrationが終わったらデータベースを閉じる
	defer fmt.Println("Successfully Migrated")
	dbConn.AutoMigrate(&domain.User{}) // マイグレーション
}
