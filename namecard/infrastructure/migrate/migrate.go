package main

import (
	"backend-namecard/domain"
	"backend-namecard/infrastructure"

	"fmt"
)

func main() {
	dbConn := infrastructure.NewDB()     // データベースに接続
	defer infrastructure.CloseDB(dbConn) //migrationが終わったらデータベースを閉じる
	defer fmt.Println("Successfully Migrated")
	dbConn.AutoMigrate(&domain.Namecard{}) // マイグレーション
}
