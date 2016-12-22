// Package iterfile provides various mechanisms for reading a file one line
// at a time. These utilities aren't necessarily meant to be used as a library
// for use in production code  (though you're more than welcome to) but rather
// to profile and benchmark various iteration constructs.
package iterfile

import (
	"bufio"
	"os"
)

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
	fobj, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(fobj)
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	chnl := make(chan string)
	go func() {
		for scanner.Scan() {
			chnl <- scanner.Text()
		}
		close(chnl)
	}()

	return chnl, nil
}
