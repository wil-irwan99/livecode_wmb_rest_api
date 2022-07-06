package usecase

import (
	"wmb_rest_api/model"
	"wmb_rest_api/repository"
)

type InsertMenuPrice interface {
	InsertHargaMenu(id int, harga int) error
}

type insertMenuPrice struct {
	menuRepo      repository.MenuRepository
	menuPriceRepo repository.MenuPriceRepository
}

func (i *insertMenuPrice) InsertHargaMenu(id int, harga int) error {

	_, err := i.menuRepo.FindBy("id = ?", id)
	if err != nil {
		return err
	}

	menuExisted := model.MenuPrice{
		MenuID: id,
		Price:  harga,
	}
	err = i.menuPriceRepo.Create(&menuExisted)
	if err != nil {
		return err
	}

	return nil
}

func NewInsertMenuPrice(menuRepo repository.MenuRepository, menuPriceRepo repository.MenuPriceRepository) InsertMenuPrice {
	return &insertMenuPrice{
		menuRepo:      menuRepo,
		menuPriceRepo: menuPriceRepo,
	}
}
