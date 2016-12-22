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
			t.Errorf("ChanReadlines expected %d lines got %d lines", lt.lines, lines)
		}

		if words != lt.words {
			t.Errorf("ChanReadlines expected %d words got %d words", lt.words, words)
		}

		if chars != lt.chars {
			t.Errorf("ChanReadlines expected %d chars got %d chars", lt.chars, chars)
		}

	}
}

func benchmarkChanReadlines(path string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		var chars int
		reader, _ := ChanReadlines(path)
		for line := range reader {
			chars += len(line)
		}
	}
}

func BenchmarkChanReadlinesSmall(b *testing.B)  { benchmarkChanReadlines(lineTests[0].path, b) }
func BenchmarkChanReadlinesMedium(b *testing.B) { benchmarkChanReadlines(lineTests[1].path, b) }
func BenchmarkChanReadlinesLarge(b *testing.B)  { benchmarkChanReadlines(lineTests[2].path, b) }

func TestCallbackReadlines(t *testing.T) {
	for _, lt := range lineTests {

		var (
			lines int
			words int
			chars int
		)

		err := CallbackReadlines(lt.path, func(line string) error {
			lines++            // count the number of lines
			chars += len(line) // count the number of chars

			tokens := strings.Split(line, " ")
			words += len(tokens)

			return nil
		})

		if err != nil {
			t.Fatalf("could not test CallbackReadlines on %s", lt.path)
		}

		// Compare the counts to actual values.
		if lines != lt.lines {
			t.Errorf("CallbackReadlines expected %d lines got %d lines", lt.lines, lines)
		}

		if words != lt.words {
			t.Errorf("CallbackReadlines expected %d words got %d words", lt.words, words)
		}

		if chars != lt.chars {
			t.Errorf("CallbackReadlines expected %d chars got %d chars", lt.chars, chars)
		}

	}
}

func benchmarkCallbackReadlines(path string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		var chars int
		CallbackReadlines(path, func(line string) error {
			chars += len(line)
			return nil
		})
	}
}

func BenchmarkChallbackReadlinesSmall(b *testing.B)  { benchmarkCallbackReadlines(lineTests[0].path, b) }
func BenchmarkChallbackReadlinesMedium(b *testing.B) { benchmarkCallbackReadlines(lineTests[1].path, b) }
func BenchmarkChallbackReadlinesLarge(b *testing.B)  { benchmarkCallbackReadlines(lineTests[2].path, b) }

func TestGeneratorReadlines(t *testing.T) {
	t.Skip("generator readlines function is not working right now")

	for _, lt := range lineTests {

		var (
			lines int
			words int
			chars int
		)

		var line string
		for gen, next, err := GeneratorReadlines(lt.path); next; line, next, err = gen() {
			if err != nil {
				t.Fatalf("could not test GeneratorReadlines on %s", lt.path)
				break
			}
			lines++            // count the number of lines
			chars += len(line) // count the number of chars

			tokens := strings.Split(line, " ")
			words += len(tokens)
		}

		// Compare the counts to actual values.
		if lines != lt.lines {
			t.Errorf("GeneratorReadlines expected %d lines got %d lines", lt.lines, lines)
		}

		if words != lt.words {
			t.Errorf("GeneratorReadlines expected %d words got %d words", lt.words, words)
		}

		if chars != lt.chars {
			t.Errorf("GeneratorReadlines expected %d chars got %d chars", lt.chars, chars)
		}

	}
}

func TestIteratorReadlines(t *testing.T) {
	for _, lt := range lineTests {

		var (
			lines int
			words int
			chars int
		)

		reader, err := IteratorReadlines(lt.path)
		if err != nil {
			t.Fatalf("could not test IteratorReadlines on %s", lt.path)

		}
		for reader.Next() {
			line := reader.Line()
			lines++            // count the number of lines
			chars += len(line) // count the number of chars

			tokens := strings.Split(line, " ")
			words += len(tokens)
		}

		// Compare the counts to actual values.
		if lines != lt.lines {
			t.Errorf("IteratorReadlines expected %d lines got %d lines", lt.lines, lines)
		}

		if words != lt.words {
			t.Errorf("IteratorReadlines expected %d words got %d words", lt.words, words)
		}

		if chars != lt.chars {
			t.Errorf("IteratorReadlines expected %d chars got %d chars", lt.chars, chars)
		}

	}
}

func benchmarkIteratorReadlines(path string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		var chars int
		reader, _ := IteratorReadlines(path)
		for reader.Next() {
			chars += len(reader.Line())
		}
	}
}

func BenchmarkIteratorReadlinesSmall(b *testing.B)  { benchmarkIteratorReadlines(lineTests[0].path, b) }
func BenchmarkIteratorReadlinesMedium(b *testing.B) { benchmarkIteratorReadlines(lineTests[1].path, b) }
func BenchmarkIteratorReadlinesLarge(b *testing.B)  { benchmarkIteratorReadlines(lineTests[2].path, b) }
