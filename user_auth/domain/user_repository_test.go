package domain_test

import (
	"user_auth/domain"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	// sqlmockのインスタンスを作成
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	// gorm.DBのインスタンスを作成
	gbd, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	return gbd, mock, nil
}

// CreateUserメソッドのテスト
func TestCreateUser(t *testing.T) {
	// gorm.DBとsqlmockのインスタンスを取得
	db, mock, err := getDBMock()
	if err != nil {
		t.Fatal(err)
	}

	r := domain.NewUserRepository(db)

	// モックの設定
	newUser := &domain.User{
		Email:    "test@test.com",
		Password: "hogehoge",
	}

	// モックの設定
	mock.ExpectBegin()

	// ユーザを作成するクエリを実行した際に返すデータを設定
	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users"`)).                                  // usersテーブルにデータを挿入するクエリを実行した際に返すデータを設定
		WithArgs(newUser.Email, newUser.Password).                // クエリにバインドされる引数を設定
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1)) // 1件のデータが追加されたことを検証

	mock.ExpectCommit()

	err = r.CreateUser(newUser)
	if err != nil {
		t.Fatal(err)
	}

	// 使用されたモックDBが期待通りの値を持っていることを検証
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// GetUserByEmailメソッドのテスト
func TestGetUserByEmail(t *testing.T) {
	// gorm.DBとsqlmockのインスタンスを取得
	db, mock, err := getDBMock()
	if err != nil {
		t.Fatal(err)
	}

	r := domain.NewUserRepository(db)

	// モックの設定
	rows := sqlmock.NewRows([]string{"id", "email", "password"}).
		AddRow(1, "test1@test.com", "password1").
		AddRow(1, "test2@test.com", "password2")

	// メールアドレスを元にユーザを取得するクエリを実行した際に返すデータを設定
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."id" LIMIT $2`)).
		WithArgs("test1@test.com", 1).
		WillReturnRows(rows)

	var user domain.User
	err = r.GetUserByEmail(&user, "test1@test.com")
	if err != nil {
		t.Fatal(err)
	}

	// 使用されたモックDBが期待通りの値を持っていることを検証
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}