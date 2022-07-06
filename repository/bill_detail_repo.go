package repository

import (
	"errors"
	"time"
	"wmb_rest_api/model"

	"gorm.io/gorm"
)

type BillDetailRepository interface {
	Create(billDetail *model.BillDetail) error
	FindById(id int) (model.BillDetail, error)
	UpdateByID(billDetail *model.BillDetail, by map[string]interface{}) error
	UpdateByModel(payload *model.BillDetail) error
	FindByBillId(bill_id int) ([]model.BillDetail, error)
	FindByDate(yesterday time.Time, tomorrow time.Time) ([]model.BillDetail, error)
}

type billDetailRepository struct {
	db *gorm.DB
}

func (b *billDetailRepository) FindByDate(yesterday time.Time, tomorrow time.Time) ([]model.BillDetail, error) {
	var billDetail []model.BillDetail
	result := b.db.Preload("MenuPrice").Preload("Bill").Where("created_at < ? AND created_at > ?", tomorrow, yesterday).Find(&billDetail)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return billDetail, nil
		} else {
			return billDetail, err
		}
	}
	return billDetail, nil
}

func (b *billDetailRepository) FindByBillId(bill_id int) ([]model.BillDetail, error) {
	var billDetail []model.BillDetail
	result := b.db.Preload("MenuPrice").Preload("Bill").Where("bill_id = ?", bill_id).Find(&billDetail)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return billDetail, nil
		} else {
			return billDetail, err
		}
	}
	return billDetail, nil
}

func (b *billDetailRepository) Create(billDetail *model.BillDetail) error {
	result := b.db.Create(billDetail).Error
	return result
}

func (b *billDetailRepository) FindById(id int) (model.BillDetail, error) {
	var billDetail model.BillDetail
	result := b.db.First(&billDetail, "id = ?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return billDetail, nil
		} else {
			return billDetail, err
		}
	}
	return billDetail, nil
}

func (b *billDetailRepository) UpdateByID(billDetail *model.BillDetail, by map[string]interface{}) error {
	result := b.db.Model(billDetail).Updates(by).Error
	return result
}

func (b *billDetailRepository) UpdateByModel(payload *model.BillDetail) error {
	result := b.db.Model(&payload).Updates(payload).Error //updates bisa untuk struct atau map[string]interface
	return result
}

func NewBillDetailRepository(db *gorm.DB) BillDetailRepository {
	repo := new(billDetailRepository)
	repo.db = db
	return repo
}
