# File Iteration Benchmarks  

[![Build Status](https://travis-ci.org/bbengfort/iterfile.svg?branch=master)](https://travis-ci.org/bbengfort/iterfile)
[![Go Report Card](https://goreportcard.com/badge/github.com/bbengfort/iterfile)](https://goreportcard.com/report/github.com/bbengfort/iterfile)
[![GoDoc](https://godoc.org/github.com/bbengfort/iterfile?status.svg)](https://godoc.org/github.com/bbengfort/iterfile)

**Benchmarking for various file iteration utilities**

[![Lines & Curves](fixtures/lines.jpg)](https://flic.kr/p/iaVByW)

This small library provides various mechanisms for reading a file one line at a time. These utilities aren't necessarily meant to be used as a library for use in production code  (though you're more than welcome to) but rather to profile and benchmark various iteration constructs. See [Benchmarking Readline Iterators](https://bbengfort.github.io/programmer/2016/12/22/benchmarking-readlines.html) for a complete post about this repository.

> Read more at [Benchmarking Readline Iterators](https://bbengfort.github.io/programmer/2016/12/22/benchmarking-readlines.html) and [Yielding Functions for Iteration in Go](http://bbengfort.github.io/snippets/2016/12/22/yielding-functions-for-iteration-golang.html).

## Usage

All of the functions in this library are `Readlines` functions; that is they take as input at least the path to a file, and then provide some iterable context with which to handle one line of the file at a time. The examples for usage here will simply be line counts, the testing methodology uses line, word, and character counts (less the newline characters). Currently we have implemented:

- `ChanReadlines`: returns a channel to `range` on.

### Channel Readlines

Use the channel based readlines iterator as follows:

```go
// construct the reader and the line count.
var lines int
reader, err := ChanReadlines("fixtures/small.txt")

// check if there was an error opening the file or scanning.
if err != nil {
    log.Fatal(err)
}

// iterate over the lines using range
for line := range reader {
    lines++
}
```

Variants of this reader would not require the error checking at the beginning, but would rather yield errors in iteration along with the line.

## Benchmarks

Benchmarks can be run with the `go test -bench=.` command. The current benchmarks are as follows:

```
BenchmarkChanReadlinesSmall-8    	   20000	     72816 ns/op
BenchmarkChanReadlinesMedium-8   	    2000	    636396 ns/op
BenchmarkChanReadlinesLarge-8    	     200	   6236999 ns/op
```

We benchmark each line count function on small (100 lines), medium (1000 lines) and large (10000 lines) text files.  

## About

Learning a new programming language often means that you want to explore everything as completely as possible. That's what this small repository is about for me, learning to write benchmarking code and to write quality iterators that are Go idiomatic. Of course, then the repository gets out of control with Repo images, etc. But hey - if you're not having fun, why are you programming?

### Acknowledgements

Table based testing inspired by Dave Chaney's [Writing table driven tests in Go](https://dave.cheney.net/2013/06/09/writing-table-driven-tests-in-go) blog post. Benchmarking was similarly inspired by [How to write benchmarks in Go](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go). Check those posts out if you haven't already.

The banner image used in this README, [&ldquo;lines & curves&rdquo;](https://flic.kr/p/iaVByW) by [Josef Stuefer](https://www.flickr.com/photos/josefstuefer/) is used by a Creative Commons [BY-NC-ND](https://creativecommons.org/licenses/by-nc-nd/2.0/) license.
