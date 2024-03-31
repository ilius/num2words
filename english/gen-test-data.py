#!/usr/bin/env python3

import gzip

from en import MAX_NUM, convert

# my select of prime numbers: 7, 71, 719, 7121, 71171, 711121, 7113221

with gzip.open("test-data.gz", "wt", encoding="utf-8") as _file:

	def add(n: int):
		if n > MAX_NUM:
			return
		n_st = str(n)
		w_st = convert(n_st)
		_file.write(f"{n_st}\t{w_st}\n")

	for n in range(100):
		add(n)
	for n in range(100, 1000, 7):
		add(n)
	for n in range(1000, 10_000, 71):
		add(n)
	for n in range(10_000, 100_000, 719):
		add(n)
	for n in range(100_000, 1_000_00, 7_121):
		add(n)
	for n in range(1_000_00, 10_000_000, 71_171):
		add(n)
	for n in range(10_000_000, 100_000_000, 711_121):
		add(n)
	for n in range(100_000_00, 1_000_000_000, 7_113_221):
		add(n)
	for n in range(10_000_000, 10_100_000, 71):
		add(n)
	for n in range(1000, 10_000, 71):
		add(n * 1_001_001)
	for n in range(1000, 10_000, 71):
		add(n * 1_001_001_001)
	for n in range(1000, 10_000, 71):
		add(n * 1_001_001_001_001)
	for n in range(1000, 10_000, 71):
		add(n * 1_001_001_001_001_001)
