package main

import (
	"bytes"
	"fmt"
)

func getTarget() (target []int, length int) {
	fmt.Scan(&length)

	target = make([]int, length, length)
	for i := range target {
		fmt.Scan(&target[i])
	}

	return target, length
}

func selectionSort(a []int, n int) (sorted []int, swapN int) {
	for i := 0; i < n; i++ {
		minJ := i
		for j := i; j < n; j++ {
			if a[j] < a[minJ] {
				minJ = j
			}
		}

		a[i], a[minJ] = a[minJ], a[i]

		if a[i] != a[minJ] { // i と minj が異なり実際に交換が行われた回数を出力するため
			swapN++
		}
	}
	return a, swapN
}

func printLine(a []int) {
	var buf bytes.Buffer
	for i, v := range a {
		if i == len(a)-1 {
			buf.WriteString(fmt.Sprintf("%d", v))
			break
		}
		buf.WriteString(fmt.Sprintf("%d ", v))
	}
	fmt.Println(buf.String())
}

func main() {
	a, n := getTarget()
	sorted, swapN := selectionSort(a, n)
	printLine(sorted)
	fmt.Println(swapN)
}
