package usecase

import (
	"wmb_rest_api/model"
	"wmb_rest_api/repository"
)

type TableUseCase interface {
	DeleteTableByID(table *model.Table) error
	FindTableById(id int) (model.Table, error)
	CreateTable(table *model.Table) error
	UpdateTableAvailable(table *model.Table, by map[string]interface{}) error
	GetLastestTableData() (model.Table, error)
}

type tableUseCase struct {
	tableRepo repository.TableRepository
}

func (t *tableUseCase) DeleteTableByID(table *model.Table) error {
	return t.tableRepo.DeleteTable(table)
}

func (t *tableUseCase) FindTableById(id int) (model.Table, error) {
	return t.tableRepo.FindById(id)
}

func (t *tableUseCase) CreateTable(table *model.Table) error {
	return t.tableRepo.Create(table)
}

func (t *tableUseCase) UpdateTableAvailable(table *model.Table, by map[string]interface{}) error {
	return t.tableRepo.UpdateAvailable(table, by)
}

func (t *tableUseCase) GetLastestTableData() (model.Table, error) {
	return t.tableRepo.RetreiveLastestTable()
}

func NewTableUseCase(tableRepo repository.TableRepository) TableUseCase {
	return &tableUseCase{
		tableRepo: tableRepo,
	}
}
