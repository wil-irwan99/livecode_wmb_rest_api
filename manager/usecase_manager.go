package manager

import (
	"wmb_rest_api/usecase"
)

type UseCaseManager interface {
	CetakBill() usecase.CetakBill
	CustomerRegistration() usecase.CustomerRegistration
	InsertMenuPrice() usecase.InsertMenuPrice
	MemberActivation() usecase.MemberActivation
	TableValidation() usecase.TableValidation
	CustomerUseCase() usecase.CustomerUseCase
	MenuUseCase() usecase.MenuUseCase
	MenuPriceUseCase() usecase.MenuPriceUseCase
	TableUseCase() usecase.TableUseCase
	TransTypeUseCase() usecase.TransTypeUseCase
	DiscountUseCase() usecase.DiscountUseCase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) CetakBill() usecase.CetakBill {
	return usecase.NewCetakBill(u.repoManager.BillDetailRepo(), u.repoManager.TableRepo(), u.repoManager.MenuPriceRepo(), u.repoManager.MenuRepo(), u.repoManager.CustomerRepo(), u.repoManager.BillRepo())
}

func (u *useCaseManager) CustomerRegistration() usecase.CustomerRegistration {
	return usecase.NewCustomerRegistration(u.repoManager.CustomerRepo())
}

func (u *useCaseManager) InsertMenuPrice() usecase.InsertMenuPrice {
	return usecase.NewInsertMenuPrice(u.repoManager.MenuRepo(), u.repoManager.MenuPriceRepo())
}

func (u *useCaseManager) MemberActivation() usecase.MemberActivation {
	return usecase.NewMemberActivation(u.repoManager.CustomerRepo(), u.repoManager.DiscountRepo())
}

func (u *useCaseManager) TableValidation() usecase.TableValidation {
	return usecase.NewTableValidation(u.repoManager.CustomerRepo(), u.repoManager.TransTypeRepo(), u.repoManager.TableRepo(), u.repoManager.BillRepo(), u.repoManager.MenuPriceRepo(), u.repoManager.BillDetailRepo())
}

func (u *useCaseManager) CustomerUseCase() usecase.CustomerUseCase {
	return usecase.NewCustomerUseCase(u.repoManager.CustomerRepo())
}

func (u *useCaseManager) MenuUseCase() usecase.MenuUseCase {
	return usecase.NewMenuUseCase(u.repoManager.MenuRepo())
}

func (u *useCaseManager) MenuPriceUseCase() usecase.MenuPriceUseCase {
	return usecase.NewMenuPriceUseCase(u.repoManager.MenuPriceRepo())
}

func (u *useCaseManager) TableUseCase() usecase.TableUseCase {
	return usecase.NewTableUseCase(u.repoManager.TableRepo())
}

func (u *useCaseManager) TransTypeUseCase() usecase.TransTypeUseCase {
	return usecase.NewTransTypeUseCase(u.repoManager.TransTypeRepo())
}

func (u *useCaseManager) DiscountUseCase() usecase.DiscountUseCase {
	return usecase.NewDiscountUseCase(u.repoManager.DiscountRepo())
}

func (u *useCaseManager) PenjualanHarian() usecase.DailySell {
	return usecase.NewDailySell(u.repoManager.BillDetailRepo(), u.repoManager.MenuPriceRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
