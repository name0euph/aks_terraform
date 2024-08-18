package domain

import (
	"gorm.io/gorm"
)

type INamecardRepository interface {
	FindAll(namecards *[]Namecard, userId uint) error
	FindById(namecard *Namecard, namecardId uint) error
	Create(namecard *Namecard) error
	Update(namecard *Namecard, namecardId uint) error
	Delete(namecardId uint) error
}

type NamecardRepository struct {
	db *gorm.DB
}

// データベースの接続を外側から受け取る
func NewTaskRepository(db *gorm.DB) INamecardRepository {
	return &NamecardRepository{db}
}

// ユーザの全てのNamecardを取得
func (nr *NamecardRepository) FindAll(namecards *[]Namecard, userId uint) error {
	// ユーザIDを元にNamecardを取得し、CreatedAtでソート
	if err := nr.db.Where("user_id=?", userId).Order("created_at").Find(namecards).Error; err != nil {
		return err
	}
	return nil
}

func (nr *NamecardRepository) FindById(namecard *Namecard, namecardId uint) error {
	// NamecardIDを元にNamecardを取得
	if err := nr.db.Where("id=?", namecardId).First(namecard).Error; err != nil {
		return err
	}
	return nil
}

func (nr *NamecardRepository) Create(namecard *Namecard) error {
	// Namecardを作成
	if err := nr.db.Create(namecard).Error; err != nil {
		return err
	}
	return nil
}

func (nr *NamecardRepository) Update(namecard *Namecard, namecardId uint) error {
	// 更新するフィールドをマッピング
	updates := map[string]interface{}{
		"name":     namecard.Name,
		"company":  namecard.Company,
		"position": namecard.Position,
		"email":    namecard.Email,
		"phone":    namecard.Phone,
	}

	// NamecardIDを元にNamecardを取得し、更新
	result := nr.db.Model(namecard).Where("id=?", namecardId).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (nr *NamecardRepository) Delete(namecardId uint) error {
	// NamecardIDを元にNamecardを取得し、削除
	result := nr.db.Where("id=?", namecardId).Delete(&Namecard{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
