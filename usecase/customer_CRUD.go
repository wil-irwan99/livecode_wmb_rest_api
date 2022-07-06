package usecase

import (
	"wmb_rest_api/model"
	"wmb_rest_api/repository"
)

type CustomerUseCase interface {
	FindCustomerById(id int) (model.Customer, error)
	GetLastDataCustomerList() (model.Customer, error)
}

type customerUseCase struct {
	custRepo repository.CustomerRepository
}

func (c *customerUseCase) FindCustomerById(id int) (model.Customer, error) {
	return c.custRepo.FindById(id)
}

func (c *customerUseCase) GetLastDataCustomerList() (model.Customer, error) {
	return c.custRepo.RetreiveLastestCustomer()
}

func NewCustomerUseCase(custRepo repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{
		custRepo: custRepo,
	}
}
