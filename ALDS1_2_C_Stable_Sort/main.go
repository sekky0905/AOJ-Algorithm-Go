package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func getTarget() (target []string, length int) {
	fmt.Scan(&length)

	target = make([]string, length, length)
	for i := range target {
		fmt.Scan(&target[i])
	}

	return target, length
}

func bubbleSort(a []string, n int) (sorted []string) {
	flag := true

	for i := 0; flag; i++ { // i は未ソートの部分の先頭を表す
		flag = false                    // ここでfalseにするが、内側のループでtrueになるので、入れ替えが生じている限りにおいては、ループは続く
		for j := n - 1; j >= i+1; j-- { //
			if isLargerNum(a[j-1], a[j]) {
				a[j], a[j-1] = a[j-1], a[j]
				flag = true
			}
		}
	}

	return a
}

func selectionSort(a []string, n int) (sorted []string) {
	for i := 0; i < n; i++ {
		minJ := i
		for j := i; j < n; j++ {
			if isLargerNum(a[minJ], a[j]) {
				minJ = j
			}
		}

		a[i], a[minJ] = a[minJ], a[i]
	}
	return a
}

func isLargerNum(a, b string) bool {
	return trimNum(a) > trimNum(b)
}

func trimNum(s string) int {
	const numIndex = 1
	num, _ := strconv.Atoi(s[numIndex:])
	return num
}

func isStable(stableSorted []string, target []string) bool {
	for i, stableV := range stableSorted {
		if stableV != target[i] {
			return false
		}
	}

	return true
}

func printLine(a []string) {
	var buf bytes.Buffer
	for i, v := range a {
		if i == len(a)-1 {
			buf.WriteString(fmt.Sprintf("%s", v))
			break
		}
		buf.WriteString(fmt.Sprintf("%s ", v))
	}
	fmt.Println(buf.String())
}

func printStability(stableSorted []string, target []string) {
	if isStable(stableSorted, target) {
		fmt.Println("Stable")
		return
	}
	fmt.Println("Not stable")
}

func main() {
	a, n := getTarget()
	b := make([]string, n)
	copy(b, a)

	bubbleSorted := bubbleSort(a, n)
	printLine(bubbleSorted)
	printStability(bubbleSorted, bubbleSorted)

	selectionSorted := selectionSort(b, n)
	printLine(selectionSorted)
	printStability(bubbleSorted, selectionSorted)
}
