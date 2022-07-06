package repository

import (
	"errors"
	"wmb_rest_api/model"

	"gorm.io/gorm"
)

type TransTypeRepository interface {
	Create(transType *model.TransType) error
	UpdateDescription(transType *model.TransType, by map[string]interface{}) error
	FindById(id string) (model.TransType, error)
	DeleteTransType(transType *model.TransType) error
	RetreiveLastestTransType() (model.TransType, error)
}

type transTypeRepository struct {
	db *gorm.DB
}

func (t *transTypeRepository) Create(transType *model.TransType) error {
	result := t.db.Create(transType).Error
	return result
}

func (t *transTypeRepository) UpdateDescription(transType *model.TransType, by map[string]interface{}) error {
	result := t.db.Model(transType).Updates(by).Error
	return result
}

func (t *transTypeRepository) FindById(id string) (model.TransType, error) {
	var transType model.TransType
	result := t.db.First(&transType, "id ilike ?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return transType, nil
		} else {
			return transType, err
		}
	}
	return transType, nil
}

func (t *transTypeRepository) DeleteTransType(transType *model.TransType) error {
	result := t.db.Delete(transType).Error
	return result
}

func (t *transTypeRepository) RetreiveLastestTransType() (model.TransType, error) {
	var transType model.TransType
	result := t.db.Last(&transType)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return transType, nil
		} else {
			return transType, err
		}
	}
	return transType, nil
}

func NewTransTypeRepository(db *gorm.DB) TransTypeRepository {
	repo := new(transTypeRepository)
	repo.db = db
	return repo
}
