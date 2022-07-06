package usecase

import (
	"wmb_rest_api/model"
	"wmb_rest_api/repository"
)

type CustomerRegistration interface {
	RegistrasiCust(nama string, no_hp string) error
}

type customerRegistration struct {
	custRepo repository.CustomerRepository
}

func (c *customerRegistration) RegistrasiCust(nama string, no_hp string) error {

	customer := model.Customer{
		CustomerName:  nama,
		MobilePhoneNo: no_hp,
		IsMember:      false,
	}

	err := c.custRepo.Create(&customer)
	if err != nil {
		return err
	}

	return nil
}

func NewCustomerRegistration(custRepo repository.CustomerRepository) CustomerRegistration {
	return &customerRegistration{
		custRepo: custRepo,
	}
}
