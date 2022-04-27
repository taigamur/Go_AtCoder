package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	S := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&S[i])
	}

	var ans = false

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if (i + 5) < N {
				var cnt = 0
				for k := 0; k < 6; k++ {
					if S[i+k][j] == '#' {
						cnt++
					}
				}
				if cnt >= 4 {
					ans = true
				}
			}
			if (j + 5) < N {
				var cnt = 0
				for k := 0; k < 6; k++ {
					if S[i][j+k] == '#' {
						cnt++
					}
					if cnt >= 4 {
						ans = true
					}
				}
			}
			if (i+5) < N && (j+5) < N {
				var cnt = 0
				for k := 0; k < 6; k++ {
					if S[i+k][j+k] == '#' {
						cnt++
					}
					if cnt >= 4 {
						ans = true
					}
				}
			}
			if 0 <= (i-5) && (j+5) < N {
				var cnt = 0
				for k := 0; k < 6; k++ {
					if S[i-k][j+k] == '#' {
						cnt++
					}
					if cnt >= 4 {
						ans = true
					}
				}
			}
		}
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
