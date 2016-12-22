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
