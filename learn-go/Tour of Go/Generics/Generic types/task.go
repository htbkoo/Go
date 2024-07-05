package main

import (
	"fmt"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

//func (l *List[T]) Find(x T) int {
//	i := 0
//	for l != nil {
//		if l.val == x {
//			return i
//		}
//	}
//
//}

func (l *List[T]) Length() int {
	i := 0
	for l != nil && l.next != nil {
		i++
		l = l.next
	}
	return i
}

func (l *List[T]) Append(val T) {
	for l != nil && l.next != nil {
		l = l.next
	}
	l.next = &List[T]{nil, val}
}

func main() {
	var l List[int]
	fmt.Println(l.Length())
	l.Append(1)
	fmt.Println(l.Length())

}
