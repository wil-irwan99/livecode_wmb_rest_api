package usecase

import (
	"wmb_rest_api/model"
	"wmb_rest_api/repository"
)

type MenuUseCase interface {
	CreateNewMenu(menu *model.Menu) error
	FindMenuById(id int) (model.Menu, error)
	UpdateMenuByID(menu *model.Menu, by map[string]interface{}) error
	DeleteMenuByID(menu *model.Menu) error
	GetLastestMenuData() (model.Menu, error)
}

type menuUseCase struct {
	menuRepo repository.MenuRepository
}

func (m *menuUseCase) CreateNewMenu(menu *model.Menu) error {
	return m.menuRepo.Create(menu)
}

func (m *menuUseCase) FindMenuById(id int) (model.Menu, error) {
	return m.menuRepo.FindById(id)
}

func (m *menuUseCase) UpdateMenuByID(menu *model.Menu, by map[string]interface{}) error {
	return m.menuRepo.UpdateByID(menu, by)
}

func (m *menuUseCase) DeleteMenuByID(menu *model.Menu) error {
	return m.menuRepo.DeleteMenu(menu)
}

func (m *menuUseCase) GetLastestMenuData() (model.Menu, error) {
	return m.menuRepo.RetreiveLastestMenu()
}

func NewMenuUseCase(menuRepo repository.MenuRepository) MenuUseCase {
	return &menuUseCase{
		menuRepo: menuRepo,
	}
}
