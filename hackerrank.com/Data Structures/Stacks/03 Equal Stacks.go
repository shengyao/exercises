// my submission for:
// https://www.hackerrank.com/challenges/equal-stacks
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
	if s.Size() > 1 {
		s.data = s.data[1:]
	} else {
		s.data = []int{}
	}
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) Sum() int {
	sum := 0
	for _, value := range s.data {
		sum += value
	}
	return sum
}

func (s *Stack) Have(t int) bool {
	for _, value := range s.data {
		if t == value {
			return true
		}
	}
	return false
}

func solutionSlow() {
	var n1, n2, n3 int
	fmt.Scanf("%d", &n1)
	fmt.Scanf("%d", &n2)
	fmt.Scanf("%d", &n3)
	var s1, s2, s3 Stack
	for i := 0; i < n1; i++ {
		var tmp int
		fmt.Scanf("%d", &tmp)
		s1.Push(tmp)
	}
	for i := 0; i < n2; i++ {
		var tmp int
		fmt.Scanf("%d", &tmp)
		s2.Push(tmp)
	}
	for i := 0; i < n3; i++ {
		var tmp int
		fmt.Scanf("%d", &tmp)
		s3.Push(tmp)
	}
	for s1.Sum() != s2.Sum() || s2.Sum() != s3.Sum() {
		var tmp *Stack
		var max int
		if max < s1.Sum() {
			max = s1.Sum()
			tmp = &s1
		}
		if max < s2.Sum() {
			max = s2.Sum()
			tmp = &s2
		}
		if max < s3.Sum() {
			max = s3.Sum()
			tmp = &s3
		}
		(*tmp).Pop()
	}
	fmt.Println(s1.Sum())
}

func solutionFast() {
	var n1, n2, n3 int
	fmt.Scanf("%d", &n1)
	fmt.Scanf("%d", &n2)
	fmt.Scanf("%d", &n3)

	var s1, s2, s3 Stack

	for i := 0; i < n1; i++ {
		var tmp int
		fmt.Scanf("%d", &tmp)
		s1.Push(tmp)
	}
	for i := 0; i < n2; i++ {
		var tmp int
		fmt.Scanf("%d", &tmp)
		s2.Push(tmp)
	}
	for i := 0; i < n3; i++ {
		var tmp int
		fmt.Scanf("%d", &tmp)
		s3.Push(tmp)
	}

	for i := n1 - 2; i >= 0; i-- {
		s1.data[i] += s1.data[i+1]
	}
	for i := n2 - 2; i >= 0; i-- {
		s2.data[i] += s2.data[i+1]
	}
	for i := n3 - 2; i >= 0; i-- {
		s3.data[i] += s3.data[i+1]
	}

	var tmp, tmp2, tmp3 *Stack
	var min int
	if s1.Size() < s2.Size() {
		min = s1.Size()
		tmp = &s1
		tmp2 = &s2
		tmp3 = &s3
	} else {
		min = s2.Size()
		tmp = &s2
		tmp2 = &s1
		tmp3 = &s3
	}
	if (*tmp).Size() < s3.Size() {
		min = s3.Size()
		tmp = &s3
		tmp2 = &s1
		tmp3 = &s2
	}

	var result int
	for i := 0; i < min; i++ {
		result = (*tmp).data[i]
		if (*tmp2).Have(result) && (*tmp3).Have(result) {
			break
		} else {
			result = 0
		}
	}
	fmt.Println(result)
}

func main() {
	//solutionSlow() // works on small cases, but will timeout on large cases
	solutionFast() // works on all cases
}
