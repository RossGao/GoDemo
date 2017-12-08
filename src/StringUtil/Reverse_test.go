package stringUtil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, wanted string
	}{
		{"!dlroW olleH", "Hello World!"},
		{"！界世 olleH", "Hello 世界！"},
		{"", ""},
	}

	for _, phrase := range cases {
		if phrase.wanted != Reverse(phrase.in) {
			t.Errorf("Reverse(%q) != %q, wanted %q", phrase.in, phrase.wanted, phrase.wanted)
		}
	}
}
