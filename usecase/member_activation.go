package usecase

import (
	"wmb_rest_api/repository"
)

type MemberActivation interface {
	AktivasiMember(id_cust int, id_disc int) error
}

type memberActivation struct {
	custRepo repository.CustomerRepository
	discRepo repository.DiscountRepository
}

func (m *memberActivation) AktivasiMember(id_cust int, id_disc int) error {

	customerExist, err := m.custRepo.FindById(id_cust)
	if err != nil {
		return err
	}

	discExist, err := m.discRepo.FindById(id_disc)
	if err != nil {
		return err
	}

	customerExist.Discounts = append(customerExist.Discounts, &discExist)

	customerExist.IsMember = true

	err = m.custRepo.UpdateByModel(&customerExist)
	if err != nil {
		return err
	}

	return nil
}

func NewMemberActivation(custRepo repository.CustomerRepository, discRepo repository.DiscountRepository) MemberActivation {
	return &memberActivation{
		custRepo: custRepo,
		discRepo: discRepo,
	}
}
