package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bbengfort/iterfile"
	"github.com/urfave/cli"
)

func profileChanReadLines(path string) (int, error) {
	var chars int
	reader, err := iterfile.ChanReadlines(path)
	if err != nil {
		return chars, err
	}

	for line := range reader {
		chars += len(line)
	}

	return chars, nil
}

func profileCallbackReadlines(path string) (int, error) {
	var chars int
	err := iterfile.CallbackReadlines(path, func(line string) error {
		chars += len(line)
		return nil
	})

	return chars, err
}

func profileIteratorReadlines(path string) (int, error) {
	var chars int
	reader, err := iterfile.IteratorReadlines(path)
	if err != nil {
		return chars, err
	}

	for reader.Next() {
		chars += len(reader.Line())
	}

	return chars, nil
}

var profilers = map[string]func(string) (int, error){
	"channel":  profileChanReadLines,
	"callback": profileCallbackReadlines,
	"iterator": profileIteratorReadlines,
}

func main() {

	// Define the CLI app
	app := cli.NewApp()
	app.Name = "readline"
	app.Usage = "Counts the number of characters without newlines"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "func, f",
			Value: "channel",
			Usage: "Specify the readlines function to profile",
		},
	}

	// The main method or action of the app
	app.Action = func(c *cli.Context) error {
		if c.NArg() < 1 {
			return cli.NewExitError("specify files to read for profiling", 1)
		}

		profiler, ok := profilers[c.String("func")]
		if !ok {
			return cli.NewExitError("specify a correct profiler function", 1)
		}

		chars := 0
		start := time.Now()
		for _, path := range c.Args() {
			charc, err := profiler(path)
			if err != nil {
				return cli.NewExitError(err, 2)
			}

			chars += charc
		}

		elapsed := time.Since(start)
		fmt.Printf("Counted %d characters in %s\n", chars, elapsed)

		return nil
	}

	// Run the CLI App
	app.Run(os.Args)
}
