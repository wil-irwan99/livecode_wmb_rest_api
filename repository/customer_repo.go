package repository

import (
	"errors"
	"wmb_rest_api/model"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *model.Customer) error
	FindById(id int) (model.Customer, error)
	UpdateByID(customer *model.Customer, by map[string]interface{}) error
	UpdateByModel(payload *model.Customer) error
	RetreiveLastestCustomer() (model.Customer, error)
}

type customerRepository struct {
	db *gorm.DB
}

func (c *customerRepository) Create(customer *model.Customer) error {
	result := c.db.Create(customer).Error
	return result
}

func (c *customerRepository) FindById(id int) (model.Customer, error) {
	var customer model.Customer
	result := c.db.Preload("Discounts").First(&customer, "id = ?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}

func (c *customerRepository) UpdateByID(customer *model.Customer, by map[string]interface{}) error {
	result := c.db.Model(customer).Updates(by).Error
	return result
}

func (c *customerRepository) UpdateByModel(payload *model.Customer) error {
	result := c.db.Model(&payload).Updates(payload).Error //updates bisa untuk struct atau map[string]interface
	return result
}

func (c *customerRepository) RetreiveLastestCustomer() (model.Customer, error) {
	var cust model.Customer
	result := c.db.Last(&cust)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return cust, nil
		} else {
			return cust, err
		}
	}
	return cust, nil
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	repo := new(customerRepository)
	repo.db = db
	return repo
}
