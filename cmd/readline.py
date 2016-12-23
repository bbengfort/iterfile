#!/usr/bin/env python 3

import sys
import time

def countchars(path):
    with open(path, 'r') as f:
        return sum(
            len(line) for line in f
        )


if __name__ == '__main__':
    chars = 0
    start = time.time()

    for path in sys.argv[1:]:
        chars += countchars(path)

    elapsed = time.time() - start
    print("Counted {} characters in {} seconds".format(chars, elapsed))
