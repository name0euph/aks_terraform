package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// Userのバリデーションを行うinterface
type IUserValidator interface {
	UserValidate(user model.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) UserValidate(user model.User) error {
	// バリデーションのルールを定義
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email, // Emailフィールドのバリデーション
			validation.Required.Error("email is required"),
			validation.Length(1, 30).Error("limited max 30 char"),
			is.Email.Error("invalid email format"),
		),
		validation.Field(
			&user.Password, // Passwordフィールドのバリデーション
			validation.Required.Error("password is required"),
			validation.Length(6, 30).Error("limited mix 6 max 30 char"),
		),
	)
}
