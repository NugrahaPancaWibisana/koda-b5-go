package main

import (
	"bufio"
	"fmt"
	"os"

	minitask1 "github.com/NugrahaPancaWibisana/koda-b5-go/internal/minitask-1"
	minitask2 "github.com/NugrahaPancaWibisana/koda-b5-go/internal/minitask-2"
	minitask3 "github.com/NugrahaPancaWibisana/koda-b5-go/internal/minitask-3"
	minitask4 "github.com/NugrahaPancaWibisana/koda-b5-go/internal/minitask-4"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("======================================================")
		fmt.Println("|| Minitask                                         ||")
		fmt.Println("||                                                  ||")
		fmt.Println("||  Ketik 0 untuk keluar                            ||")
		fmt.Println("||  Ketik 1 untuk minitask 1                        ||")
		fmt.Println("||  Ketik 2 untuk minitask 2                        ||")
		fmt.Println("||  Ketik 3 untuk minitask 3                        ||")
		fmt.Println("||  Ketik 4 untuk minitask 4                        ||")
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
			return;
		case "1":
			fmt.Println(minitask1.HitungLingkaran(2))
		case "2":
			minitask2.SegitigaSikuSiku(10)
		case "3":
			minitask3.AddNumberToSlice()
		case "4":
			minitask4.DataDiri()
		default:
			fmt.Println("Input invalid!")
		}
	}
}
