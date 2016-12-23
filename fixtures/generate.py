#!/usr/bin/env python3
# This quick script generates fixtures for benchmarking.

import os
import random

# Number of words per line
MIN_LINE = 20
MAX_LINE = 100

# Words to randomly select to add to the line
WORDS    = ("fizz", "buzz", "foo", "bar", "baz")

# Paths of fixtures to create
BASEDIR  = os.path.dirname(__file__)
FIXTURES = {
    os.path.join(BASEDIR, "small.txt"): 100,
    os.path.join(BASEDIR, "medium.txt"): 1000,
    os.path.join(BASEDIR, "large.txt"): 10000,
}


def make_fixture(path, lines, words=WORDS, minlen=MIN_LINE, maxlen=MAX_LINE):
    """
    Writes a file to the specified path with the number of lines specified by
    randomly choosing between minlen and maxlen words and writing them.
    """

    with open(path, 'w') as f:
        for _ in range(lines):
            text = [
                random.choice(words)
                for _ in range(random.randint(minlen, maxlen))
            ]

            f.write(" ".join(text) + "\n")


if __name__ == '__main__':
    for path, lines in FIXTURES.items():
        make_fixture(path, lines)

    # Make the profiling fixture
    # make_fixture('jumbo.txt', 750000, minlen=100, maxlen=2000)
