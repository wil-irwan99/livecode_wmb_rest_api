package usecase

import (
	"fmt"
	"wmb_rest_api/dto"
	"wmb_rest_api/repository"
)

type CetakBill interface {
	PrintBill(bill_id int) (dto.PrintBill, error)
}

type cetakBill struct {
	billDetailRepo repository.BillDetailRepository
	tableRepo      repository.TableRepository
	menuPriceRepo  repository.MenuPriceRepository
	menuRepo       repository.MenuRepository
	customerRepo   repository.CustomerRepository
	billRepo       repository.BillRepository
}

func (b *cetakBill) PrintBill(bill_id int) (dto.PrintBill, error) {

	var billTransaksi dto.PrintBill

	result, err := b.billDetailRepo.FindByBillId(bill_id)
	if err != nil {
		return dto.PrintBill{}, err
	}

	mejaKosong := result[0].Bill.TableID
	tableExist, _ := b.tableRepo.FindById(mejaKosong)
	b.tableRepo.UpdateAvailable(&tableExist, map[string]interface{}{
		"is_available": true,
	})

	billExist, _ := b.billRepo.FindById(bill_id)
	custIdExist := billExist.CustomerID

	custExist, _ := b.customerRepo.FindById(custIdExist)

	memberStatus := custExist.IsMember

	// var menuSlice []string
	// var menuPriceSlice []int
	// var qtySlice []int

	var totalPrice float32
	var maxDisc int = 0

	for _, v := range result {
		menuPriceExist, err := b.menuPriceRepo.FindById(v.MenuPriceID)
		if err != nil {
			return dto.PrintBill{}, err
		}

		menuExist, err := b.menuRepo.FindById(menuPriceExist.MenuID)
		if err != nil {
			return dto.PrintBill{}, err
		}

		totalPrice += float32(v.MenuPrice.Price) * float32(v.Qty)

		// menuSlice = append(menuSlice, menuExist.MenuName)
		// menuPriceSlice = append(menuPriceSlice, v.MenuPrice.Price)
		// qtySlice = append(qtySlice, v.Qty)

		billTransaksi.Orders = append(billTransaksi.Orders, dto.OrderList{
			Name:  menuExist.MenuName,
			Qty:   v.Qty,
			Price: v.MenuPrice.Price,
		})

	}

	// for i := range menuSlice {
	// 	fmt.Println(menuSlice[i], " : ", menuPriceSlice[i], "x", qtySlice[i])
	// }

	if memberStatus == true {
		for i := 0; i < len(custExist.Discounts); i++ {
			maxDisc = custExist.Discounts[0].Pct
			if custExist.Discounts[i].Pct > maxDisc {
				maxDisc = custExist.Discounts[i].Pct
			}
		}
		fmt.Println("Member Discount : ", maxDisc)
		totalPrice -= totalPrice * (float32(maxDisc) / 100)
	}

	billTransaksi.BillId = bill_id
	billTransaksi.TransDate = billExist.TransDate
	billTransaksi.Discount = maxDisc
	billTransaksi.TotalPrice = totalPrice

	return billTransaksi, nil
}

func NewCetakBill(billDetailRepo repository.BillDetailRepository, tableRepo repository.TableRepository, menuPriceRepo repository.MenuPriceRepository, menuRepo repository.MenuRepository, customerRepo repository.CustomerRepository, billRepo repository.BillRepository) CetakBill {
	return &cetakBill{
		billDetailRepo: billDetailRepo,
		tableRepo:      tableRepo,
		menuPriceRepo:  menuPriceRepo,
		menuRepo:       menuRepo,
		customerRepo:   customerRepo,
		billRepo:       billRepo,
	}
}
