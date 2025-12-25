package minitask8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

	return fmt.Sprintf("Jumlah Pembayaran fiktif: %d", pay)
}

type IPayment interface {
	Pay(list []int) string
}

func Payment() {
	scanner := bufio.NewScanner(os.Stdin)
	listFiktif := []string{}

	for {
		fmt.Println("======================================================")
		fmt.Println("|| Payment                                          ||")
		fmt.Println("||                                                  ||")
		fmt.Println("||  Ketik exit untuk keluar                         ||")
		fmt.Println("||  Ketik 1 Bank Payment                            ||")
		fmt.Println("||  Ketik 2 Online Payment                          ||")
		fmt.Println("||  Ketik 3 Fiktif Payment                          ||")
		fmt.Println("||                                                  ||")
		fmt.Println("======================================================")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Print("Masukan opsi: ")

		scanner.Scan()

		switch scanner.Text() {
		case "exit":
			fmt.Println("List pembayaran fiktif")
			for _, v := range listFiktif {
				fmt.Println(v)
			}
			return
		case "1":
			list := []int{}

			for {
				fmt.Println("Ketik exit untuk keluar")
				fmt.Println("")
				fmt.Println("")
				fmt.Println("")
				fmt.Print("Masukan harga: ")
				scanner.Scan()

				if scanner.Text() == "exit" {
					fmt.Println(IPayment.Pay(Bank{}, list))
					break
				}

				v, e := strconv.Atoi(scanner.Text())

				if e != nil {
					fmt.Println(e)
				}

				if v <= 0 {
					fmt.Println("Pembayaran tidak boleh 0")
				}

				list = append(list, v)
			}
		case "2":
			list := []int{}

			for {
				fmt.Println("Ketik exit untuk keluar")
				fmt.Println("")
				fmt.Println("")
				fmt.Println("")
				fmt.Print("Masukan harga: ")
				scanner.Scan()

				if scanner.Text() == "exit" {
					fmt.Println(IPayment.Pay(Online{}, list))
					break
				}

				v, e := strconv.Atoi(scanner.Text())

				if e != nil {
					fmt.Println(e)
				}

				if v <= 0 {
					fmt.Println("Pembayaran tidak boleh 0")
				}

				list = append(list, v)
			}
		case "3":
			list := []int{}

			for {
				fmt.Println("Ketik exit untuk keluar")
				fmt.Println("")
				fmt.Println("")
				fmt.Println("")
				fmt.Print("Masukan harga: ")
				scanner.Scan()

				if scanner.Text() == "exit" {
					listFiktif = append(listFiktif, IPayment.Pay(Fiktif{}, list))
					break
				}

				v, e := strconv.Atoi(scanner.Text())

				if e != nil {
					fmt.Println(e)
				}

				if v <= 0 {
					fmt.Println("Pembayaran tidak boleh 0")
				}

				list = append(list, v)
			}
		}
	}
}
