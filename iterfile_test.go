package iterfile

import (
	"strings"
	"testing"
)

type lineTest struct {
	path  string // Input to readlines file
	lines int    // The expected number of lines
	words int    // The expected number of words
	chars int    // The expected number of bytes
}

var lineTests = []lineTest{
	{"fixtures/small.txt", 100, 6030, 26419},
	{"fixtures/medium.txt", 1000, 59284, 259816},
	{"fixtures/large.txt", 10000, 598305, 2622952},
}

func TestChanReadlines(t *testing.T) {
	for _, lt := range lineTests {

		var (
			lines int
			words int
			chars int
		)

		// Create the reader and fail if it can't connect
		reader, err := ChanReadlines(lt.path)
		if err != nil {
			t.Fatalf("could not test ChanReadlines on %s", lt.path)
		}

		// Iterate over all lines and compute the counts
		for line := range reader {
			lines++            // count the number of lines
			chars += len(line) // count the number of chars

			tokens := strings.Split(line, " ")
			words += len(tokens)

		}

		// Compare the counts to actual values.
		if lines != lt.lines {
			t.Errorf("expected %d lines got %d lines", lt.lines, lines)
		}

		if words != lt.words {
			t.Errorf("expected %d words got %d words", lt.words, words)
		}

		if chars != lt.chars {
			t.Errorf("expected %d chars got %d chars", lt.chars, chars)
		}

	}
}
