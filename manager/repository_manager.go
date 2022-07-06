package manager

import (
	"wmb_rest_api/repository"
)

type RepositoryManager interface {
	BillRepo() repository.BillRepository
	BillDetailRepo() repository.BillDetailRepository
	CustomerRepo() repository.CustomerRepository
	DiscountRepo() repository.DiscountRepository
	MenuRepo() repository.MenuRepository
	MenuPriceRepo() repository.MenuPriceRepository
	TableRepo() repository.TableRepository
	TransTypeRepo() repository.TransTypeRepository
}

type repositoryManager struct {
	infra Infra
}

func (r *repositoryManager) BillRepo() repository.BillRepository {
	return repository.NewBillRepository(r.infra.SqlDb())
}

func (r *repositoryManager) BillDetailRepo() repository.BillDetailRepository {
	return repository.NewBillDetailRepository(r.infra.SqlDb())
}

func (r *repositoryManager) CustomerRepo() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infra.SqlDb())
}

func (r *repositoryManager) DiscountRepo() repository.DiscountRepository {
	return repository.NewDiscountRepository(r.infra.SqlDb())
}

func (r *repositoryManager) MenuRepo() repository.MenuRepository {
	return repository.NewMenuRepository(r.infra.SqlDb())
}

func (r *repositoryManager) MenuPriceRepo() repository.MenuPriceRepository {
	return repository.NewMenuPriceRepository(r.infra.SqlDb())
}

func (r *repositoryManager) TableRepo() repository.TableRepository {
	return repository.NewTableRpository(r.infra.SqlDb())
}

func (r *repositoryManager) TransTypeRepo() repository.TransTypeRepository {
	return repository.NewTransTypeRepository(r.infra.SqlDb())
}

func NewRepositoryManager(infra Infra) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
