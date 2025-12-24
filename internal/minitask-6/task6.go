package minitask6

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func OpenFile() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukan file path: ")
	scanner.Scan()

	file, e := os.Open(scanner.Text())
	if e != nil {
		fmt.Println(e)
		return
	}

	readFile(file)
}

func readFile(file *os.File) {
	defer file.Close() 

	data, e := io.ReadAll(file)
	if e != nil {
		panic(e)
	}

	fmt.Println(string(data))
}