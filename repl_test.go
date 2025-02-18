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
                        input:    "This is a test",
                        expected: []string{"this", "is", "a", "test"},
                },
                {
                        input:    "the Quick brown Fox",
                        expected: []string{"the", "quick", "brown", "fox"},
                },
                {
                        input:    "jumped over the lazy dog",
                        expected: []string{"jumped", "over", "the", "lazy", "dog"},
                },
                {
                        input:    "Uno Duo Tres Quatro Sinco sinco ses",
                        expected: []string{"uno", "duo", "tres", "quatro", "sinco", "sinco", "ses"},
                },
                {
                        input:    "thisisareallylongsentancewithnospaces",
                        expected: []string{"thisisareallylongsentancewithnospaces"},
                },
                {
                        input:    "",
                        expected: []string{},
                },
		// add more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		if len(actual) != len(c.expected) {
			// t.Errorf("%s does not match %s", actual, c.expected)
			t.Errorf("got %d words, want %d words for input '%s'", len(actual), len(c.expected), c.input)
		}
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			if word != expectedWord {
				//t.Errorf("%s is not in the expected list of words", word)
				t.Errorf("word at position %d: got '%s', want '%s' for input '%s'", i, word, expectedWord, c.input)
			}
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}


}
