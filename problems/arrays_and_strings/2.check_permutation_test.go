package arraysandstrings

import "testing"

func TestCheckPermutation(t *testing.T){
	tests := []struct {
		name     string
		inputA   string
		inputB   string
		expected bool
	}{
		// Basic cases
		{"Basic Permutation", "cat", "tac", true},
		{"Basic Non-Permutation", "cat", "ttc", false},

		// Case sensitivity
		{"Case Sensitive - Different Case", "Cat", "tac", false},
		{"Case Sensitive - Another Case", "Dog", "God", false},

		// Different lengths
		{"Different Lengths", "abc", "abcd", false},
		{"Different Lengths - Shorter Second", "hello", "hell", false},

		// Empty strings
		{"Both Empty Strings", "", "", true},
		{"One Empty String", "a", "", false},

		// Same characters, different frequencies
		{"Same Characters, Different Frequency", "aabb", "abab", true},
		{"Different Character Counts", "aabb", "aabc", false},

		// Unicode characters
		{"Unicode - Chinese", "你好", "好你", true},
		{"Unicode - Japanese", "こんにちは", "にちはこ", false},

		// Whitespace significance
	 	{"Whitespace Included", "dog ", " god", true},
		{"Whitespace Different", "dog", "god ", false},

		// Single-character strings
		{"Single Character - Same", "a", "a", true},
		{"Single Character - Different", "a", "b", false},

		// Large inputs (stress test)
		{"Long Strings - Valid Permutation", "mnopqrstuvwxyzabcdefghijkl", "mnopqrstuvwxyzabcdefglhijk", true},
		{"Long Strings - Invalid Permutation", "abcdefghijklmnoqrstuvwxyz", "mnopqrstuvwxyzabcdefghiij", false}, 
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() { // Recover from panic and print the test name
				if r := recover(); r != nil {
					t.Errorf("Test %q panicked: %v", test.name, r)
				}
			}()

			result := CheckPermutation(test.inputA, test.inputB)
			if result != test.expected {
				t.Errorf("%s: CheckPermutation(%q, %q) = %v; want %v",
					test.name, test.inputA, test.inputB, result, test.expected)
			}
		})
	}
}