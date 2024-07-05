// https://www.codewars.com/kata/59f08f89a5e129c543000069/train/go

/*
In this Kata, you will be given an array of strings and your task is to remove all consecutive duplicate letters from each string in the array.

For example:

dup(["abracadabra","allottee","assessee"]) = ["abracadabra","alote","asese"].

dup(["kelless","keenness"]) = ["keles","kenes"].

Strings will be lowercase only, no spaces. See test cases for more examples.

Good luck!

If you like this Kata, please try:

Alternate capitalization

Vowel consonant lexicon
*/

package kata

import "strings"

func dedup(s string) string {
	sb := strings.Builder{}
	var prev int32
	for _, ch := range s {
		if ch != prev {
			prev = ch
			sb.WriteString(string(ch))
		}
	}

	return sb.String()
}

func Dup(arr []string) []string {
	//..

	answer := make([]string, len(arr))
	for i, s := range arr {
		answer[i] = dedup(s)
	}
	return answer
}
