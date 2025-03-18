package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "    ",
			expected: []string{},
		},
		{
			input:    "single",
			expected: []string{"single"},
		},
		{
			input:    "   leading spaces",
			expected: []string{"leading", "spaces"},
		},
		{
			input:    "trailing spaces   ",
			expected: []string{"trailing", "spaces"},
		},
		{
			input:    "multiple   spaces   between   words",
			expected: []string{"multiple", "spaces", "between", "words"},
		},
		{
			input:    "tabs\t\tbetween\twords",
			expected: []string{"tabs", "between", "words"},
		},
		{
			input:    "mixed   \t  whitespace",
			expected: []string{"mixed", "whitespace"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(c.expected) != len(actual) {
			t.Errorf("Length mismatch: got %d, want %d", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if expectedWord != word {
				t.Errorf("Word mismatch at position %d: got %s, want %s", i, word, expectedWord)
			}
		}
	}
}
