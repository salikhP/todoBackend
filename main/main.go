package main

import (
	"bufio"
	"fmt"
	"os"
)

func canZeroArray(a []int, k int) bool {
	n := len(a)
	diff := make([]int, n+2)
	ops := 0

	for i := 0; i < n; i++ {
		ops += diff[i]
		curr := a[i] + ops
		if curr > 0 {
			if i+k > n {
				return false
			}
			diff[i] -= curr
			diff[i+k] += curr
			ops -= curr
		}
	}

	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for test := 0; test < t; test++ {
		var n, k int
		fmt.Fscan(in, &n, &k)

		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}

		if canZeroArray(a, k) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
