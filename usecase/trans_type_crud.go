package usecase

import (
	"wmb_rest_api/model"
	"wmb_rest_api/repository"
)

type TransTypeUseCase interface {
	CreateTransType(transType *model.TransType) error
	FindTransTypeById(id string) (model.TransType, error)
	UpdateTransTypeDescription(transType *model.TransType, by map[string]interface{}) error
	DeleteTransTypeById(transType *model.TransType) error
	GetLastestTransTypeData() (model.TransType, error)
}

type transTypeUseCase struct {
	transTypeRepo repository.TransTypeRepository
}

func (t *transTypeUseCase) CreateTransType(transType *model.TransType) error {
	return t.transTypeRepo.Create(transType)
}

func (t *transTypeUseCase) FindTransTypeById(id string) (model.TransType, error) {
	return t.transTypeRepo.FindById(id)
}

func (t *transTypeUseCase) UpdateTransTypeDescription(transType *model.TransType, by map[string]interface{}) error {
	return t.transTypeRepo.UpdateDescription(transType, by)
}

func (t *transTypeUseCase) DeleteTransTypeById(transType *model.TransType) error {
	return t.transTypeRepo.DeleteTransType(transType)
}

func (t *transTypeUseCase) GetLastestTransTypeData() (model.TransType, error) {
	return t.transTypeRepo.RetreiveLastestTransType()
}

func NewTransTypeUseCase(transTypeRepo repository.TransTypeRepository) TransTypeUseCase {
	return &transTypeUseCase{
		transTypeRepo: transTypeRepo,
	}
}
