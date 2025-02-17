package arraysandstrings

func CheckPermutation(a string, b string) bool {

	if len([]rune(a)) != len([]rune(b)) {
		return false
	}

	chMapA := make(map[rune]int)
	//range is smart enough to handle unicode characters as runes not byte slices
	for _, ch := range a {
		chMapA[ch]++
	}

	for _, ch := range b {
		chMapA[ch]--
		//no need to store and compare a second map of b
		//just decrement the matching a values and if negative - not permutation
		if chMapA[ch] < 0 {
			return false
		}
	}

	return true
}