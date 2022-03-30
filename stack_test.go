package main

import "testing"

func TestStack(t *testing.T) {

	insertVals := []int{0, 1, 2, 3, 4, 5}

	var s Stack[int]

	for _, val := range insertVals {
		s = s.Push(val)
	}

	for i := len(insertVals) - 1; i > -1; i-- {
		if peekVal, isNotEmpty := s.Peek(); isNotEmpty {
			if peekVal != insertVals[i] {
				t.Logf("vals dont match: test: %d, stack: %d", insertVals[i], peekVal)
			} else {
				t.Logf("vals match: test: %d, stack: %d", insertVals[i], peekVal)
			}
			s = s.Pop()
		}
	}

}
