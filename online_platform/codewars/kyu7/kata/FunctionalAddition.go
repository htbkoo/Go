package kata

// https://www.codewars.com/kata/538835ae443aae6e03000547/train/go

// Create a function add(n)/Add(n) which returns a function that always adds n to any number
//
// Note for Java: the return type and methods have not been provided to make it a bit more challenging.
//
// var addOne = Add(1)
// addOne(3) // 4

func Add(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}
