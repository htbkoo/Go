// https://www.codewars.com/kata/5b077ebdaf15be5c7f000077/train/go

// If you can't sleep, just count sheep!!
// Task:
//
// Given a non-negative integer, 3 for example, return a string with a murmur: "1 sheep...2 sheep...3 sheep...". Input will always be valid, i.e. no negative integers.

package kata

import (
	"strconv"
	"strings"
)

func countSheep(num int) string {
	// Your code here!
	var str_slice = []string{}

	for i := 0; i < num; i++ {
		str_slice = append(str_slice, strconv.Itoa(i+1)+" sheep...")
	}

	return strings.Join(str_slice, "")
}
