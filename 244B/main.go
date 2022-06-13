package main

import "fmt"

func main() {
	var (
		N int
		S string
	)
	fmt.Scan(&N)
	fmt.Scan(&S)

	var x int = 0
	var y int = 0
	var v int = 0

	for i := 0; i < N; i++ {
		if S[i] == 'S' {
			v := v % 4
			switch v {
			case 0:
				x++
			case 1:
				y--
			case 2:
				x--
			case 3:
				y++
			}

		} else {
			v++
		}
	}
	fmt.Println(x, y)

}
