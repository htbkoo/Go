package kata

// https://www.codewars.com/kata/53369039d7ab3ac506000467/train/go
// Complete the method that takes a boolean value and return a "Yes" string for true, or a "No" string for false.

func BoolToWord(word bool) string {
	if word {
		return "Yes"
	} else {
		return "No"
	}
}
