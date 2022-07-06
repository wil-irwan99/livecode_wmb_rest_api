package repository

import (
	"errors"
	"wmb_rest_api/model"

	"gorm.io/gorm"
)

type BillRepository interface {
	Create(bill *model.Bill) error
	FindById(id int) (model.Bill, error)
	UpdateByID(bill *model.Bill, by map[string]interface{}) error
	UpdateByModel(payload *model.Bill) error
	FindAllWithPreload(preload string) ([]model.Bill, error)
	RetreiveLastedBillID() (model.Bill, error)
}

type billRepository struct {
	db *gorm.DB
}

func (b *billRepository) Create(bill *model.Bill) error {
	result := b.db.Create(bill).Error
	return result
}

func (b *billRepository) FindById(id int) (model.Bill, error) {
	var bill model.Bill
	result := b.db.First(&bill, "id = ?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return bill, nil
		} else {
			return bill, err
		}
	}
	return bill, nil
}

func (b *billRepository) RetreiveLastedBillID() (model.Bill, error) {
	var bill model.Bill
	result := b.db.Last(&bill)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return bill, nil
		} else {
			return bill, err
		}
	}
	return bill, nil
}

func (b *billRepository) UpdateByID(bill *model.Bill, by map[string]interface{}) error {
	result := b.db.Model(bill).Updates(by).Error
	return result
}

func (b *billRepository) UpdateByModel(payload *model.Bill) error {
	result := b.db.Model(&payload).Updates(payload).Error //updates bisa untuk struct atau map[string]interface
	return result
}

func (b *billRepository) FindAllWithPreload(preload string) ([]model.Bill, error) {
	var bill []model.Bill
	result := b.db.Preload(preload).Find(&bill)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return bill, nil
		} else {
			return bill, err
		}
	}
	return bill, nil
}

func NewBillRepository(db *gorm.DB) BillRepository {
	repo := new(billRepository)
	repo.db = db
	return repo
}
