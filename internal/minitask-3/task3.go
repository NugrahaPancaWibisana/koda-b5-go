package minitask3

import "fmt"

func AddNumberToSlice() {
	data := []int{50, 75, 66, 20, 32, 90}

	data = append(data[:3], append([]int{88}, data[3:]...)...)

	for i, v := range data {
		fmt.Printf("Index %d: %d\n", i, v)
	}
}