package main

import (
	"bufio"
	"fmt"
	"os"

	minitask1 "github.com/NugrahaPancaWibisana/koda-b5-go/internal/minitask-1"
	minitask2 "github.com/NugrahaPancaWibisana/koda-b5-go/internal/minitask-2"
	minitask3 "github.com/NugrahaPancaWibisana/koda-b5-go/internal/minitask-3"
	minitask4 "github.com/NugrahaPancaWibisana/koda-b5-go/internal/minitask-4"
	minitask6 "github.com/NugrahaPancaWibisana/koda-b5-go/internal/minitask-6"
	// minitask7 "github.com/NugrahaPancaWibisana/koda-b5-go/internal/minitask-7"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// person := minitask7.NewPerson("Nugraha Panca Wibisana", "Indonesia", "08123456789")

	for {
		fmt.Println("======================================================")
		fmt.Println("|| Minitask                                         ||")
		fmt.Println("||                                                  ||")
		fmt.Println("||  Ketik 0 untuk keluar                            ||")
		fmt.Println("||  Ketik 1 untuk minitask 1                        ||")
		fmt.Println("||  Ketik 2 untuk minitask 2                        ||")
		fmt.Println("||  Ketik 3 untuk minitask 3                        ||")
		fmt.Println("||  Ketik 4 untuk minitask 4                        ||")
		fmt.Println("||  Ketik 5 untuk minitask 6                        ||")
		// fmt.Println("||  Ketik 6 untuk minitask 7 Print func             ||")
		// fmt.Println("||  Ketik 7 untuk minitask 7 Greet func             ||")
		// fmt.Println("||  Ketik 8 untuk minitask 7 UpdatePersonName func  ||")
		fmt.Println("||                                                  ||")
		fmt.Println("======================================================")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Print("Masukan nomor: ")

		scanner.Scan()

		switch scanner.Text() {
		case "0":
			fmt.Println("Anda telah keluar")
			return
		case "1":
			fmt.Println(minitask1.HitungLingkaran(2))
		case "2":
			minitask2.SegitigaSikuSiku(10)
		case "3":
			minitask3.AddNumberToSlice()
		case "4":
			minitask4.DataDiri()
		case "5":
			minitask6.OpenFile()
		// case "6":
		// 	person.Print()
		// case "7":
		// 	person.Greet()
		// case "8":
		// 	fmt.Print("Masukan nama: ")

		// 	scanner.Scan()

		// 	person.UpdatePersonName(scanner.Text())
		default:
			fmt.Println("Input invalid!")
		}
	}
}
