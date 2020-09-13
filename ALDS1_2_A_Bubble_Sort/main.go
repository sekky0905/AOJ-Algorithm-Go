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

func bubbleSort(a []int, n int) (sorted []int, swapN int) {
	flag := true

	for i := 0; flag; i++ { // i は未ソートの部分の先頭を表す
		flag = false                  // ここでfalseにするが、内側のループでtrueになるので、入れ替えが生じている限りにおいては、ループは続く
		for j := n - 1; j >= i + 1; j-- { //
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
				swapN++
				flag = true
			}
		}
	}

	return a, swapN
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

func main() {
	a, n := getTarget()
	sorted, swapN := bubbleSort(a, n)
	printLine(sorted, n)
	fmt.Println(swapN)
}
