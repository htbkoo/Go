package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func walkImpl(t *tree.Tree, ch chan int) {
	/* TODO */

	if t == nil {
		return
	}

	walkImpl(t.Left, ch)
	ch <- t.Value
	walkImpl(t.Right, ch)

	return
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	/* TODO */

	walkImpl(t, ch)
	close(ch)
	return
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	/* TODO */

	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if !ok1 && !ok2 {
			return true
		}

		if !ok1 || !ok2 {
			return false
		}

		if v1 != v2 {
			return false
		}
	}
}

func main() {
	fmt.Print("tree.New(1) == tree.New(1): ")
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}

	fmt.Print("tree.New(1) != tree.New(2): ")
	if !Same(tree.New(1), tree.New(2)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}
}
