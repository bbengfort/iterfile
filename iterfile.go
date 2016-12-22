// Package iterfile provides various mechanisms for reading a file one line
// at a time. These utilities aren't necessarily meant to be used as a library
// for use in production code  (though you're more than welcome to) but rather
// to profile and benchmark various iteration constructs.
package iterfile

import (
	"bufio"
	"os"
)

// Helper function to open a file and create a line scanner, code that will
// probably be used in every single iterfile function that I write!
func openFile(path string) (*bufio.Scanner, *os.File, error) {
	fobj, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	scanner := bufio.NewScanner(fobj)
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return scanner, fobj, nil
}

// ChanReadlines returns an channel that can be used in conjunction with the
// range keyword for looping over every line in the file. Basic usage is:
//
//     reader, err := ChanReadlines("myfile.txt")
//     for line := range reader {
//         // do something with the line
//     }
//
// The channel will be closed by the reader when the entire file is read.
func ChanReadlines(path string) (<-chan string, error) {
	scanner, file, err := openFile(path)
	if err != nil {
		return nil, err
	}

	chnl := make(chan string)
	go func() {
		for scanner.Scan() {
			chnl <- scanner.Text()
		}
		file.Close()
		close(chnl)
	}()

	return chnl, nil
}

// CallbackReadlines allows the caller to specify a callback function whose
// input is the line being read. The callback is then called on each line in
// the file. Note that the CallbackReadlines function returns an error that
// should be checked along with the results of the callbacks. Basic usage is:
//
//     func linecb(line string) error {
//         // do something with the line
//     }
//
//     err := CallbackReadlines("myfile.txt", linecb)
//
// Note that the callback function can return an error, which if detected will
// cause the loop to break and return the error from the callback.
func CallbackReadlines(path string, cb func(string) error) error {
	scanner, file, err := openFile(path)
	if err != nil {
		return err
	}

	defer file.Close()

	for scanner.Scan() {
		if err := cb(scanner.Text()); err != nil {
			return err
		}
	}

	return nil
}

// GeneratorReadlines returns a closure that can be called multiple times as
// though it were a generator function, creating kind of an interesting for
// expression construct that can be fit into a single line. Basic usage is:
//
//     for gen, next, err := GeneratorReadlines("myfile.txt"); next; line, next, err = gen() {
//         // do something with the line
//     }
//
// The loop stops when the generator next bool returns false.
func GeneratorReadlines(path string) (func() (string, bool, error), bool, error) {
	scanner, file, err := openFile(path)
	if err != nil {
		return nil, false, err
	}

	next := scanner.Scan()
	line := scanner.Text()

	return func() (string, bool, error) {
		prevLine := line
		next = scanner.Scan()
		line = scanner.Text()

		if !next {
			file.Close()
		}

		return prevLine, next, nil

	}, next, nil
}

// LineIterator specifies how an iterable object over file lines should work.
type LineIterator interface {
	Next() bool   // Advances the iterator to the next line
	Line() string // Returns the current line of iteration
}

type lineIterator struct {
	file    *os.File       // Reference to the open file
	scanner *bufio.Scanner // Reference to the scanner object
	current string         // The current line for multiple Line() calls
}

// Line returns the current line of the line iterator.
func (lit *lineIterator) Line() string {
	return lit.current
}

// Next advances the iterator with the scanner and returns whether or not
// another line exists on the stateful iterator object.
func (lit *lineIterator) Next() bool {
	hasNext := lit.scanner.Scan()
	lit.current = lit.scanner.Text()

	if !hasNext {
		lit.file.Close()
	}

	return hasNext
}

// IteratorReadlines returns a LineIterator to loop over by calling its Next()
// method and obtaining the value with its Line() method. Basic usage is:
//
//     reader, err := IteratorReadlines("myfile.txt")
//     for reader.Next() {
//          line := reader.Line()
//          // do something with the line
//     }
//
// Once the LineIterator is exahusted it cannot be used again or reset.
func IteratorReadlines(path string) (LineIterator, error) {
	scanner, file, err := openFile(path)
	if err != nil {
		return nil, err
	}

	return &lineIterator{
		scanner: scanner,
		file:    file,
		current: "",
	}, nil
}
