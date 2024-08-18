package usecase

import (
	"backend-namecard/domain"
)

// NamecardUsecaseのインターフェースを定義
type INamecardUsecase interface {
	GetAll(userId uint) ([]domain.Namecard, error)
	GetById(namecardId uint) (domain.Namecard, error)
	Create(namecard domain.Namecard) (domain.Namecard, error)
	Update(namecard domain.Namecard, namecardId uint) (domain.Namecard, error)
	Delete(namecardId uint) error
}

type namecardUsecase struct {
	nr domain.INamecardRepository
}

// インスタンス生成
func NewNamecardUsecase(nr domain.INamecardRepository) INamecardUsecase {
	return &namecardUsecase{nr}
}

func (nu *namecardUsecase) GetAll(userId uint) ([]domain.Namecard, error) {
	// 空のスライス
	namecards := []domain.Namecard{}

	// スライスのポインタを引数に渡してタスクを取得
	if err := nu.nr.FindAll(&namecards, userId); err != nil {
		return []domain.Namecard{}, err
	}
	return namecards, nil
}

func (nu *namecardUsecase) GetById(namecardId uint) (domain.Namecard, error) {
	// 空の構造体
	namecard := domain.Namecard{}

	// 構造体のポインタを引数に渡してタスクを取得
	if err := nu.nr.FindById(&namecard, namecardId); err != nil {
		return domain.Namecard{}, err
	}
	return namecard, nil
}

func (nu *namecardUsecase) Create(namecard domain.Namecard) (domain.Namecard, error) {
	if err := nu.nr.Create(&namecard); err != nil {
		return domain.Namecard{}, err
	}
	return namecard, nil
}

func (nu *namecardUsecase) Update(namecard domain.Namecard, namecardId uint) (domain.Namecard, error) {
	if err := nu.nr.Update(&namecard, namecardId); err != nil {
		return domain.Namecard{}, err
	}
	return namecard, nil
}

func (nu *namecardUsecase) Delete(namecardId uint) error {
	if err := nu.nr.Delete(namecardId); err != nil {
		return err
	}
	return nil
}
