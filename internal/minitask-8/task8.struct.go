package minitask8

import "fmt"

type Bank struct{}

func (b Bank) Pay(list []int) string {
	var pay int = 0

	for _, v := range list {
		pay += v
	}

	return fmt.Sprintf("Jumlah Pembayaran Bank: %d \n", pay)
}

type Online struct{}

func (o Online) Pay(list []int) string {
	var pay int = 0

	for _, v := range list {
		pay += v
	}

	return fmt.Sprintf("Jumlah Pembayaran Online: %d \n", pay)
}

type Fiktif struct{}

func (f Fiktif) Pay(list []int) string {
	var pay int = 0

	for _, v := range list {
		pay += v
	}

	return fmt.Sprintf("Jumlah Pembayaran Online: %d \n", pay)
}

type IPayment interface {
	Pay(list []int) string
}

func Payment()  {
	
}