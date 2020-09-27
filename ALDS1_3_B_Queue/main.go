package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type queue struct {
	max  int
	head int
	tail int
	s    []process
}

func newQueue(max int, s []process) *queue {
	const initNum = 0

	que := &queue{max: max, head: initNum, tail: initNum, s: make([]process, max)}

	for _, v := range s {
		que.enqueue(v)
	}

	return que
}

func (q queue) isFull() bool {
	// eg.) head = 0, tail 9 = max = 10
	// head = 3, tail = 2 max = 10
	return q.head == (q.tail+1)%q.max
}

func (q queue) isEmpty() bool {
	return q.head == q.tail
}

func (q *queue) enqueue(v process) {
	if q.isFull() {
		panic("over flow")
	}

	q.s[q.tail] = v

	if q.tail+1 == q.max {
		// リングバッファが一周したので、最初に戻す
		q.tail = 0
		return
	}

	q.tail++
}

func (q *queue) dequeue() process {
	if q.isEmpty() {
		panic("under flow")
	}

	x := q.s[q.head]

	if q.head+1 == q.max {
		// リングバッファが一周したので、最初に戻す
		q.head = 0
		return x
	}

	q.head++
	return x
}

type process struct {
	name         string
	completeTime int
}

func newProcess(name string, time int) *process {
	return &process{name: name, completeTime: time}
}

func (p *process) consumeTime(qTime int) (remainingTime int) {
	const empty = 0

	beforeMinus := p.completeTime
	remainingCompleteTime := p.completeTime - qTime

	if remainingCompleteTime <= empty {
		p.completeTime = empty
		return qTime - beforeMinus
	}

	p.completeTime = remainingCompleteTime
	return 0
}

func (p *process) isCompleted() bool {
	const completed = 0
	return p.completeTime == completed
}

var sc = bufio.NewScanner(os.Stdin)

const sep = " "

type firstLine struct {
	n     int
	qTime int
}

func newFirstLine(n int, qTime int) *firstLine {
	return &firstLine{n: n, qTime: qTime}
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func getFirstLine(s string) *firstLine {
	const (
		nIdx     = 0
		qTimeIdx = 1
	)

	split := strings.Split(s, sep)

	n, _ := strconv.Atoi(split[nIdx])
	qTime, _ := strconv.Atoi(split[qTimeIdx])

	return newFirstLine(n, qTime)
}

func getProcess(s string) *process {
	const (
		nameIdx         = 0
		completeTimeIdx = 1
	)

	split := strings.Split(s, sep)

	name := split[nameIdx]
	ct, _ := strconv.Atoi(split[completeTimeIdx])

	return newProcess(name, ct)
}

func calc(que queue, qTime int) {
	totalConsumedTime := 0

	for {
		if que.isEmpty() {
			break
		}

		p := que.dequeue()

		remainingTime := p.consumeTime(qTime)
		consumedTime := qTime - remainingTime
		totalConsumedTime = totalConsumedTime + consumedTime

		if !p.isCompleted() {
			que.enqueue(p)
			continue
		}

		fmt.Printf("%s %d\n", p.name, totalConsumedTime)
	}
}

func main() {
	firstLineStr := nextLine()
	fl := getFirstLine(firstLineStr)

	processes := make([]process, fl.n)

	for i := 0; i < fl.n; i++ {
		ls := nextLine()
		p := getProcess(ls)

		processes[i] = *p
	}

	const max = 100000
	que := newQueue(max, processes)
	calc(*que, fl.qTime)
}
