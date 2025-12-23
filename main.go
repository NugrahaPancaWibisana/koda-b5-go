package main

import "fmt"

func main() {
	fmt.Println(HitungLingkaran(2))
	SegitigaSikuSiku(10)
}

func HitungLingkaran(r int) string {
	return fmt.Sprintf("Luas = %d | Keliling = %d", LuasLingkaran(r), KelilingLingkaran(r))
}

func KelilingLingkaran(r int) int {
	var pi float64 = 3.14
	var x float64 = float64(2) * pi * float64(r)

	return int(x)
}

func LuasLingkaran(r int) int {
	var pi float64 = 3.14
	var x float64 = pi * float64(r*r)

	return int(x)
}

func SegitigaSikuSiku(height int) {
	// normal for
	var starFor string = "*"

	for i := 1; i < height; i++ {
		fmt.Println(starFor)
		starFor = starFor + "*"
	}

	fmt.Println("")

	// while
	var starWhile string = "*"

	var j int = 1
	for j < height {
		fmt.Println(starWhile)
		starWhile = starWhile + "*"
		j++
	}

	fmt.Println("")

	// range
	var starRange string = "*"

	for range height {
		fmt.Println(starRange)
		starRange = starRange + "*"
	}
}
