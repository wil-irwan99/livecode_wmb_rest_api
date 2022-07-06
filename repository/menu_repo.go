package repository

import (
	"errors"
	"wmb_rest_api/model"

	"gorm.io/gorm"
)

type MenuRepository interface {
	Create(menu *model.Menu) error
	FindBy(by string, vals ...interface{}) ([]model.Menu, error)
	FindById(id int) (model.Menu, error)
	UpdateByID(menu *model.Menu, by map[string]interface{}) error
	DeleteMenu(menu *model.Menu) error
	RetreiveLastestMenu() (model.Menu, error)
}

type menuRepository struct {
	db *gorm.DB
}

func (m *menuRepository) Create(menu *model.Menu) error {
	result := m.db.Create(menu).Error
	return result
}

func (m *menuRepository) FindBy(by string, vals ...interface{}) ([]model.Menu, error) {
	var menu []model.Menu
	result := m.db.Where(by, vals...).Find(&menu)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return menu, nil
		} else {
			return menu, err
		}
	}
	return menu, nil
}

func (m *menuRepository) UpdateByID(menu *model.Menu, by map[string]interface{}) error {
	result := m.db.Model(menu).Updates(by).Error
	return result
}

func (m *menuRepository) DeleteMenu(menu *model.Menu) error {
	result := m.db.Delete(menu).Error
	return result
}

func (m *menuRepository) FindById(id int) (model.Menu, error) {
	var menu model.Menu
	result := m.db.First(&menu, "id = ?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return menu, nil
		} else {
			return menu, err
		}
	}
	return menu, nil
}

func (m *menuRepository) RetreiveLastestMenu() (model.Menu, error) {
	var menu model.Menu
	result := m.db.Last(&menu)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return menu, nil
		} else {
			return menu, err
		}
	}
	return menu, nil
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	repo := new(menuRepository)
	repo.db = db
	return repo
}
