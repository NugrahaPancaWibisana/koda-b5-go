package minitask2

import "fmt"

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
