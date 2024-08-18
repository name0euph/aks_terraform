package domain_test

import (
	"backend-namecard/domain"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var mockNamecard = domain.Namecard{
	ID:        1,
	Name:      "Alice",
	Company:   "Company A",
	Position:  "Manager",
	Email:     "alice@example.com",
	Phone:     "123-456-7890",
	CreatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	UpdatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	UserId:    1,
}

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

func TestFindAll(t *testing.T) {
	// gorm.DBとsqlmockのインスタンスを取得
	db, mock, err := getDBMock()
	if err != nil {
		t.Fatal(err)
	}

	r := domain.NewTaskRepository(db) // 初期化

	// モックの設定
	rows := sqlmock.NewRows([]string{"id", "name", "company", "position", "email", "phone", "created_at", "updated_at", "user_id"}).
		AddRow(mockNamecard.ID, mockNamecard.Name, mockNamecard.Company, mockNamecard.Position, mockNamecard.Email, mockNamecard.Phone,
			mockNamecard.CreatedAt, mockNamecard.UpdatedAt, mockNamecard.UserId)

	// ExpextQueryで期待するSQLクエリを定義
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "namecards" WHERE user_id=$1 ORDER BY created_at`)).
		WithArgs(1).
		WillReturnRows(rows)

	var namecards []domain.Namecard
	if err := r.FindAll(&namecards, 1); err != nil {
		t.Fatal(err)
	}

	// 使用されたモックDBが期待通りの値を持っていることを検証
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFindById(t *testing.T) {
	// sqlmockのインスタンスを作成
	db, mock, err := getDBMock()
	if err != nil {
		t.Fatal(err)
	}

	r := domain.NewTaskRepository(db) // 初期化

	// モックの設定
	rows := sqlmock.NewRows([]string{"id", "name", "company", "position", "email", "phone", "created_at", "updated_at", "user_id"}).
		AddRow(mockNamecard.ID, mockNamecard.Name, mockNamecard.Company, mockNamecard.Position, mockNamecard.Email, mockNamecard.Phone,
			mockNamecard.CreatedAt, mockNamecard.UpdatedAt, mockNamecard.UserId)

	// ExpextQueryで期待するSQLクエリを定義
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "namecards" WHERE id=$1 ORDER BY "namecards"."id" LIMIT $2`)).
		WithArgs(1, 1).
		WillReturnRows(rows)

	var namecard domain.Namecard
	if err := r.FindById(&namecard, 1); err != nil {
		t.Fatal(err)
	}

	// 使用されたモックDBが期待通りの値を持っていることを検証
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreate(t *testing.T) {
	// sqlmockのインスタンスを作成
	db, mock, err := getDBMock()
	if err != nil {
		t.Fatal(err)
	}

	r := domain.NewTaskRepository(db) // 初期化

	// テストデータ
	newNamecard := domain.Namecard{
		Name:      "Alice",
		Company:   "Company A",
		Position:  "Manager",
		Email:     "alice@example.com",
		Phone:     "123-456-7890",
		CreatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		UserId:    1,
	}

	// ExpextQueryで期待するSQLクエリを定義
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "namecards"`)). // namecardsテーブルにデータを挿入するクエリを実行した際に返すデータを設定
		WithArgs(newNamecard.Name, newNamecard.Company, newNamecard.Position, newNamecard.Email, newNamecard.Phone,
										newNamecard.CreatedAt, newNamecard.UpdatedAt, newNamecard.UserId).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1)) // 1件のデータが追加されたことを検証
	mock.ExpectCommit()

	if err := r.Create(&newNamecard); err != nil {
		t.Fatal(err)
	}

	// 使用されたモックDBが期待通りの値を持っていることを検証
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdate(t *testing.T) {
	// sqlmockのインスタンスを作成
	db, mock, err := getDBMock()
	if err != nil {
		t.Fatal(err)
	}

	r := domain.NewTaskRepository(db) // 初期化

	// テストデータ
	updatedNamecard := domain.Namecard{
		Name:      "Bob",
		Company:   "Company B",
		Position:  "Developer",
		Email:     "bob@example.com",
		Phone:     "123-456-7890",
		CreatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2024, 1, 12, 0, 0, 0, 0, time.UTC),
		UserId:    1,
	}

	// ExpextQueryで期待するSQLクエリを定義
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE "namecards" SET "company"=$1,"email"=$2,"name"=$3,"phone"=$4,"position"=$5,"updated_at"=$6 WHERE id=$7`)).
		WithArgs(updatedNamecard.Company, updatedNamecard.Email, updatedNamecard.Name, updatedNamecard.Phone, updatedNamecard.Position,
								updatedNamecard.UpdatedAt, 1).
		WillReturnResult(sqlmock.NewResult(0, 1)) // 1件のデータが更新されたことを検証
	mock.ExpectCommit()

	if err := r.Update(&updatedNamecard, 1); err != nil {
		t.Fatal(err)
	}

	// 使用されたモックDBが期待通りの値を持っていることを検証
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
