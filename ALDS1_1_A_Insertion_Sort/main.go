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

func main() {
	a, n := getTarget()
	insertionSort(a, n)
}

func insertionSort(a []int, n int) {
	for i := 0; i < n; i++ {
		v := a[i]
		// 1つindexが小さいやつをjとする
		j := i - 1

		for j >= 0 && a[j] > v { // jが0以上かつ、jのやつの方が大きい場合
			a[j+1] = a[j]
			j--
		}
		// ここの時点では-1か、a[j] <= v の場合なので、j+1に格納する
		a[j+1] = v
		printLine(a, n)
	}
}

func printLine(a []int, n int) {
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
