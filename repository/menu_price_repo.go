package repository

import (
	"errors"
	"wmb_rest_api/model"

	"gorm.io/gorm"
)

type MenuPriceRepository interface {
	Create(menuPrice *model.MenuPrice) error
	FindById(id int) (model.MenuPrice, error)
	Delete(menuPrice *model.MenuPrice) error
	RetreiveLastestMenuPrice() (model.MenuPrice, error)
	UpdateByID(menuPrice *model.MenuPrice, by map[string]interface{}) error
}

type menuPriceRepository struct {
	db *gorm.DB
}

func (m *menuPriceRepository) Create(menuPrice *model.MenuPrice) error {
	result := m.db.Create(menuPrice).Error
	return result
}

func (m *menuPriceRepository) FindById(id int) (model.MenuPrice, error) {
	var menuPrice model.MenuPrice
	result := m.db.First(&menuPrice, "id = ?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return menuPrice, nil
		} else {
			return menuPrice, err
		}
	}
	return menuPrice, nil
}

func (m *menuPriceRepository) Delete(menuPrice *model.MenuPrice) error {
	result := m.db.Delete(menuPrice).Error
	return result
}

func (m *menuPriceRepository) RetreiveLastestMenuPrice() (model.MenuPrice, error) {
	var menuPrice model.MenuPrice
	result := m.db.Last(&menuPrice)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return menuPrice, nil
		} else {
			return menuPrice, err
		}
	}
	return menuPrice, nil
}

func (m *menuPriceRepository) UpdateByID(menuPrice *model.MenuPrice, by map[string]interface{}) error {
	result := m.db.Model(menuPrice).Updates(by).Error
	return result
}

func NewMenuPriceRepository(db *gorm.DB) MenuPriceRepository {
	repo := new(menuPriceRepository)
	repo.db = db
	return repo
}
