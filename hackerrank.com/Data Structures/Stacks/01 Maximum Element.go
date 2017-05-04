// my submission for:
// https://www.hackerrank.com/challenges/maximum-element
// by Shengyao, 2017

package main

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) Push(t int) {
	s.data = append(s.data, t)
}

func (s *Stack) Pop() {
	s.data = s.data[:len(s.data)-1]
}

func (s *Stack) Max() int {
	var max int
	for i := 0; i < len(s.data); i++ {
		if s.data[i] > max {
			max = s.data[i]
		}
	}
	return max
}

func main() {
	var N, i, t, x int
	fmt.Scanf("%d", &N)
	var st Stack
	for i = 0; i < N; i++ {
		fmt.Scanf("%d", &t)
		if t == 1 {
			fmt.Scanf("%d", &x)
			st.Push(x)
		} else if t == 2 {
			st.Pop()
		} else if t == 3 {
			fmt.Println(st.Max())
		}
	}
}
