package usecase

import (
	"wmb_rest_api/model"
	"wmb_rest_api/repository"
)

type DiscountUseCase interface {
	CreateDiscount(discount *model.Discount) error
	UpdateDiscountById(discount *model.Discount, by map[string]interface{}) error
	FindDiscountById(id int) (model.Discount, error)
	GetLastestDiscountData() (model.Discount, error)
	DeleteDiscountById(disc *model.Discount) error
}

type discountUseCase struct {
	discountRepo repository.DiscountRepository
}

func (d *discountUseCase) CreateDiscount(discount *model.Discount) error {
	return d.discountRepo.Create(discount)
}

func (d *discountUseCase) UpdateDiscountById(discount *model.Discount, by map[string]interface{}) error {
	return d.discountRepo.UpdateDiscount(discount, by)
}

func (d *discountUseCase) FindDiscountById(id int) (model.Discount, error) {
	return d.discountRepo.FindById(id)
}

func (d *discountUseCase) GetLastestDiscountData() (model.Discount, error) {
	return d.discountRepo.RetreiveLastestDiscount()
}

func (d *discountUseCase) DeleteDiscountById(disc *model.Discount) error {
	return d.discountRepo.DeleteDisc(disc)
}

func NewDiscountUseCase(discountRepo repository.DiscountRepository) DiscountUseCase {
	return &discountUseCase{
		discountRepo: discountRepo,
	}
}
