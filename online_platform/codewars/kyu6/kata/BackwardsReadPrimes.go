// https://www.codewars.com/kata/5539fecef69c483c5a000015/train/go

/*
Backwards Read Primes are primes that when read backwards in base 10 (from right to left) are a different prime. (This rules out primes which are palindromes.)

Examples:
13 17 31 37 71 73 are Backwards Read Primes
13 is such because it's prime and read from right to left writes 31 which is prime too. Same for the others.

Task
Find all Backwards Read Primes between two positive given numbers (both inclusive), the second one always being greater than or equal to the first one. The resulting array or the resulting string will be ordered following the natural order of the prime numbers.

Examples (in general form):
backwardsPrime(2, 100) => [13, 17, 31, 37, 71, 73, 79, 97] backwardsPrime(9900, 10000) => [9923, 9931, 9941, 9967] backwardsPrime(501, 599) => []

See "Sample Tests" for your language.

Notes
Forth Return only the first backwards-read prime between start and end or 0 if you don't find any
Ruby Don't use Ruby Prime class, it's disabled.
*/

package kata

func reverse(x int) int {
	r := 0
	for x > 0 {
		r = r*10 + (x % 10)
		x /= 10
	}
	return r
}

var primes []int
var primesSet map[int]struct{}

func isPrime(x int) bool {
	_, ok := primesSet[x]
	if ok {
		return true
	}

	//for i := 2; i*i <= x; i++ {
	//	if x%i == 0 {
	//		return false
	//	}
	//}

	for _, p := range primes {
		if x%p == 0 {
			return false
		}
		if p*p > x {
			break
		}
	}

	primesSet[x] = struct{}{}

	return true
}

func BackwardsPrime(start int, stop int) []int {
	primes = make([]int, 0)
	primesSet = make(map[int]struct{})

	// your code
	for i := 2; i <= stop; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	answer := make([]int, 0)
	for i := start; i <= stop; i++ {
		r := reverse(i)
		if r == i {
			continue
		}

		if isPrime(i) && isPrime(r) {
			answer = append(answer, i)
		}
	}

	if len(answer) == 0 {
		return nil
	}
	return answer
}
