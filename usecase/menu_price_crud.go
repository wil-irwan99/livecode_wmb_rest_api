package usecase

import (
	"wmb_rest_api/model"
	"wmb_rest_api/repository"
)

type MenuPriceUseCase interface {
	CreateNewMenuPrice(menuPrice *model.MenuPrice) error
	FindMenuPriceById(id int) (model.MenuPrice, error)
	DeleteMenuPriceById(menuPrice *model.MenuPrice) error
	GetLastestMenuPrice() (model.MenuPrice, error)
	UpdateMenuPriceByID(menuPrice *model.MenuPrice, by map[string]interface{}) error
}

type menuPriceUseCase struct {
	menuPriceRepo repository.MenuPriceRepository
}

func (m *menuPriceUseCase) CreateNewMenuPrice(menuPrice *model.MenuPrice) error {
	return m.menuPriceRepo.Create(menuPrice)
}

func (m *menuPriceUseCase) FindMenuPriceById(id int) (model.MenuPrice, error) {
	return m.menuPriceRepo.FindById(id)
}

func (m *menuPriceUseCase) DeleteMenuPriceById(menuPrice *model.MenuPrice) error {
	return m.menuPriceRepo.Delete(menuPrice)
}

func (m *menuPriceUseCase) GetLastestMenuPrice() (model.MenuPrice, error) {
	return m.menuPriceRepo.RetreiveLastestMenuPrice()
}

func (m *menuPriceUseCase) UpdateMenuPriceByID(menuPrice *model.MenuPrice, by map[string]interface{}) error {
	return m.menuPriceRepo.UpdateByID(menuPrice, by)
}

func NewMenuPriceUseCase(menuPriceRepo repository.MenuPriceRepository) MenuPriceUseCase {
	return &menuPriceUseCase{
		menuPriceRepo: menuPriceRepo,
	}
}
