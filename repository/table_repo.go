package repository

import (
	"errors"
	"wmb_rest_api/model"

	"gorm.io/gorm"
)

type TableRepository interface {
	Create(table *model.Table) error
	FindBy(by string, vals ...interface{}) ([]model.Table, error)
	FindById(id int) (model.Table, error)
	UpdateAvailable(menu *model.Table, by map[string]interface{}) error
	DeleteTable(table *model.Table) error
	RetreiveLastestTable() (model.Table, error)
}

type tableRepository struct {
	db *gorm.DB
}

func (t *tableRepository) Create(table *model.Table) error {
	result := t.db.Create(table).Error
	return result
}

func (t *tableRepository) FindBy(by string, vals ...interface{}) ([]model.Table, error) {
	var table []model.Table
	result := t.db.Where(by, vals...).Find(&table)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return table, nil
		} else {
			return table, err
		}
	}
	return table, nil
}

func (t *tableRepository) FindById(id int) (model.Table, error) {
	var table model.Table
	result := t.db.First(&table, "id = ?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return table, nil
		} else {
			return table, err
		}
	}
	return table, nil
}

func (t *tableRepository) UpdateAvailable(menu *model.Table, by map[string]interface{}) error {
	result := t.db.Model(menu).Updates(by).Error
	return result
}

func (t *tableRepository) DeleteTable(table *model.Table) error {
	result := t.db.Delete(table).Error
	return result
}

func (t *tableRepository) RetreiveLastestTable() (model.Table, error) {
	var table model.Table
	result := t.db.Last(&table)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return table, nil
		} else {
			return table, err
		}
	}
	return table, nil
}

func NewTableRpository(db *gorm.DB) TableRepository {
	repo := new(tableRepository)
	repo.db = db
	return repo
}
