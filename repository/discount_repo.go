package repository

import (
	"errors"
	"wmb_rest_api/model"

	"gorm.io/gorm"
)

type DiscountRepository interface {
	Create(discount *model.Discount) error
	UpdateDiscount(discount *model.Discount, by map[string]interface{}) error
	FindBy(by string, vals ...interface{}) ([]model.Discount, error)
	FindById(id int) (model.Discount, error)
	RetreiveLastestDiscount() (model.Discount, error)
	DeleteDisc(disc *model.Discount) error
}

type discountRepository struct {
	db *gorm.DB
}

func (d *discountRepository) Create(discount *model.Discount) error {
	result := d.db.Create(discount).Error
	return result
}

func (d *discountRepository) UpdateDiscount(discount *model.Discount, by map[string]interface{}) error {
	result := d.db.Model(discount).Updates(by).Error
	return result
}

func (d *discountRepository) FindBy(by string, vals ...interface{}) ([]model.Discount, error) {
	var discount []model.Discount
	result := d.db.Where(by, vals...).Find(&discount)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return discount, nil
		} else {
			return discount, err
		}
	}
	return discount, nil
}

func (d *discountRepository) FindById(id int) (model.Discount, error) {
	var discount model.Discount
	result := d.db.First(&discount, "id = ?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return discount, nil
		} else {
			return discount, err
		}
	}
	return discount, nil
}

func (d *discountRepository) RetreiveLastestDiscount() (model.Discount, error) {
	var disc model.Discount
	result := d.db.Last(&disc)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return disc, nil
		} else {
			return disc, err
		}
	}
	return disc, nil
}

func (d *discountRepository) DeleteDisc(disc *model.Discount) error {
	result := d.db.Delete(disc).Error
	return result
}

func NewDiscountRepository(db *gorm.DB) DiscountRepository {
	repo := new(discountRepository)
	repo.db = db
	return repo
}
