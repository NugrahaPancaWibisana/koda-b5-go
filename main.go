package main

import "fmt"

func main() {
	fmt.Println(HitungLingkaran(2))
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
