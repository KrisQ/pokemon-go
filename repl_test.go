package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "Hello World",
			expected: []string{
				"hello", "world",
			},
		},
		{
			input: "Il etait un petit navire",
			expected: []string{
				"il", "etait", "un", "petit", "navire",
			},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test

		if len(actual) != len(c.expected) {
			t.Errorf("The length of actual: %v doesn't match the length of expted: %v", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test:
			if word != expectedWord {
				t.Errorf("Words don't match. Expected: %v Actual: %v", word, expectedWord)
			}
		}
	}
}
