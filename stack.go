package main

import "fmt"

type Stack[T any] []T

func (s *Stack[T]) Push(item T) Stack[T] {
	return append(*s, item)

}

func (s *Stack[T]) Peek() (T, bool) {
	if len(*s) == 0 {
		var zero T
		return zero, false
	}
	return (*s)[len(*s)-1], true
}

func (s *Stack[T]) Pop() Stack[T] {

	if len(*s) != 0 {
		return (*s)[:len(*s)-1]
	}

	return Stack[T]{}

}

func (s Stack[T]) Print() {

	for _, val := range s {
		fmt.Print(val, " ")
	}

}
