package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack struct {
	max int
	top int
	s   []int
}

func newStack(max int) *stack {
	const empty = 0
	return &stack{max: max, top: empty, s: make([]int, max)}
}

func (s stack) isEmpty() bool {
	return s.top == 0
}

func (s stack) isFull() bool {
	return s.top == s.max
}

func (s *stack) push(v int) {
	if s.isFull() {
		panic("over flow")
	}

	s.top++
	s.s[s.top] = v
}

func (s *stack) pop() int {
	if s.isEmpty() {
		panic("under flow")
	}

	s.top--
	return s.s[s.top+1]
}

func getTarget() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	return in.Text()
}

func main() {
	const max = 200
	st := newStack(max)

	target := getTarget()
	s := strings.Split(target, " ")
	res := calc(st, s)
	v := res.pop()
	fmt.Println(v)
}

func calc(st *stack, target []string) *stack {
	for _, v := range target {
		switch v {
		case "-":
			a, b := getCalcTarget(st)

			st.push(a - b)
		case "+":
			a, b := getCalcTarget(st)
			st.push(a + b)
		case "*":
			a, b := getCalcTarget(st)
			st.push(a * b)
		default:
			vi, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			st.push(vi)
		}
	}
	return st
}

func getCalcTarget(st *stack) (int, int) {
	// stack なのでこの順番
	b := st.pop()
	a := st.pop()

	return a, b
}
