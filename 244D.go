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

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func group(A []string) bool {
	if (A[0] == "R" && A[1] == "G" && A[2] == "B") ||
		(A[0] == "G" && A[1] == "B" && A[2] == "R") ||
		(A[0] == "B" && A[1] == "R" && A[2] == "G") {
		return true
	} else {
		return false
	}
}

func main() {
	sc.Split(bufio.ScanWords) //nextInt(), nextStringのとき

	S := make([]string, 3)
	for i := 0; i < 3; i++ {
		S[i] = nextString()

	}

	T := make([]string, 3)
	for i := 0; i < 3; i++ {
		T[i] = nextString()
	}

	if (group(S) == true && group(T) == true) ||
		(group(S) == false && group(T) == false) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	return
}
