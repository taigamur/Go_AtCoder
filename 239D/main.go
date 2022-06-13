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
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func judge(n int) bool {

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords) //nextInt()のとき
	var a, b, c, d int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)

	for i := a; i <= b; i++ {
		f := false
		for j := c; j <= d; j++ {
			if judge(i + j) {
				f = true
			}
		}
		if f == false {
			fmt.Println("Takahashi")
			return
		}
	}
	fmt.Println("Aoki")
}
