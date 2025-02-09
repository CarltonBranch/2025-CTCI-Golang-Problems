package arraysandstrings

import "testing"

func TestIsUnique(t *testing.T) {
    tests := []struct {
        input    string
        expected bool
    }{
        {"abcd", true},
        {"hello", false},
        {"", true},
        {"a", true},
    }

    for _, test := range tests {
        result := IsUnique(test.input)
        if result != test.expected {
            t.Errorf("IsUnique(%q) = %v; want %v", test.input, result, test.expected)
        }
    }
}
