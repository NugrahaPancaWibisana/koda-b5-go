package minitask1

import "fmt"

func HitungLingkaran(r int) string {
	return fmt.Sprintf("Luas = %d | Keliling = %d", luasLingkaran(r), kelilingLingkaran(r))
}

func kelilingLingkaran(r int) int {
	var pi float64 = 3.14
	var x float64 = float64(2) * pi * float64(r)

	return int(x)
}

func luasLingkaran(r int) int {
	var pi float64 = 3.14
	var x float64 = pi * float64(r*r)

	return int(x)
}
