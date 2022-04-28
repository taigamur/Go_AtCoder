package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	N := nextInt()
	A := make([]int, N)
	for i := 0; i < 4*N-1; i++ {
		A[nextInt()-1] += 1
	}
	for i := 0; i < N; i++ {
		if A[i] == 3 {
			fmt.Println(i + 1)
			return
		}
	}
	return
}
