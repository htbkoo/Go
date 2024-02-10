// https://www.codewars.com/kata/515f51d438015969f7000013/train/go

// Write a function that when given a number >= 0, returns an Array of ascending length subarrays.
//
// pyramid(0) => [ ]
// pyramid(1) => [ [1] ]
// pyramid(2) => [ [1], [1, 1] ]
// pyramid(3) => [ [1], [1, 1], [1, 1, 1] ]
//
// Note: the subarrays should be filled with 1s

package kata

func Pyramid(n int) [][]int {
	// your code here
	var pyramid = make([][]int, n)
	for i := 0; i < n; i++ {
		pyramid[i] = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			pyramid[i][j] = 1
		}
	}

	return pyramid
}
