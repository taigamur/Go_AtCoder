package main

import "fmt"

func main() {
	var N int
	var S string

	fmt.Scanf("%d", &N)
	fmt.Scan(&S)
	fmt.Println(string(S[N-1]))
}
