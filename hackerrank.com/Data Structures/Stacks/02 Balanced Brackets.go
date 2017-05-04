// my submission for:
// https://www.hackerrank.com/challenges/balanced-brackets
// by Shengyao, 2017

package main

import "fmt"

type Stack struct {
	data []byte
}

func (s *Stack) Push(t byte) {
	s.data = append(s.data, t)
}

func (s *Stack) Pop() byte {
	var tmp byte = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return tmp
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) Empty() bool {
	if s.Size() == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var N int
	fmt.Scanf("%d", &N)
	for i := 0; i < N; i++ {
		var s Stack
		var line string
		fmt.Scanf("%s", &line)
		var j int
		for j = 0; j < len(line); j++ {
			var tmp byte = line[j]
			if tmp == '(' || tmp == '[' || tmp == '{' {
				s.Push(line[j])
			} else if tmp == ')' {
				if s.Empty() || s.Pop() != '(' {
					break
				}
			} else if tmp == ']' {
				if s.Empty() || s.Pop() != '[' {
					break
				}
			} else if tmp == '}' {
				if s.Empty() || s.Pop() != '{' {
					break
				}
			}
		}
		if j == len(line) && s.Empty() {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
