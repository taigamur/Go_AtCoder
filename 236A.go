package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)
	var a int
	var b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	// とりあえずruneスライスをつくる
	str := []rune(S)

	tmp := str[a-1]
	str[a-1] = str[b-1]
	str[b-1] = tmp

	for i := 0; i < len(S); i++ {
		fmt.Print(string(str[i]))
	}
	fmt.Println()
}
