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
			input:    "Hello",
			expected: []string{"hello"},
		},
		{
			input:    "WoRlD",
			expected: []string{"world"},
		},
		{
			input:    "W o  R   l    D ",
			expected: []string{"w", "o", "r", "l", "d"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "    ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) => %#v (len=%d), expected %#v (len=%d)",
				c.input, actual, len(actual), c.expected, len(c.expected))
			continue
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("cleanInput(%q) index %d: got %q, expected %q",
					c.input, i, actual[i], c.expected[i])
			}
		}
	}
}
