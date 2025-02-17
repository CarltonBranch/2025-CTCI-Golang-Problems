package arraysandstrings

// IsUnique checks if a string has all unique characters
func IsUnique(s string) bool {
    charMap:= make(map [rune]bool)
    for _, ch:= range s{
        if charMap[ch]{
            return false
        }
        charMap[ch] = true
    }
    return true
}