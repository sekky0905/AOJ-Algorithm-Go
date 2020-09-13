package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func getTarget(sc *bufio.Scanner) int {
	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())
	return num
}

// 条件
// 後 - 前
func main() {
	sc := bufio.NewScanner(os.Stdin)
	n := getTarget(sc)

	minV := getTarget(sc)
	maxV := math.MinInt32

	const startWithoutFirstValue = 1

	for i := startWithoutFirstValue; i < n; i++ {
		v := getTarget(sc)

		// max は、今のmaxと今の値-minのでかい方
		maxV = maxValue(maxV, v-minV)
		minV = minValue(minV, v)
	}

	fmt.Println(maxV)
}

func maxValue(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minValue(x, y int) int {
	if x > y {
		return y
	}
	return x
}
