#!/usr/bin/env python3
# -*- coding: utf-8 -*-
#
# Copyright @ 2014-2024 Saeed Rasooli <saeed.gnu@gmail.com> (ilius)
#
# This library is free software; you can redistribute it and/or
# modify it under the terms of the GNU Lesser General Public
# License as published by the Free Software Foundation; either
# version 2.1 of the License, or (at your option) any later version.
#
# This library is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
# Lesser General Public License for more details.


import sys

MAX_NUM = 10**12 - 1

digit_text = {
	"00": "Zero",
	"01": "One",
	"02": "Two",
	"03": "Three",
	"04": "Four",
	"05": "Five",
	"06": "Six",
	"07": "Seven",
	"08": "Eight",
	"09": "Nine",
	"10": "Ten",
	"11": "Eleven",
	"12": "Twelve",
	"13": "Thirteen",
	"14": "Fourteen",
	"15": "Fifteen",
	"16": "Sixteen",
	"17": "Seventeen",
	"18": "Eighteen",
	"19": "Nineteen",
	"20": "Twenty",
	"30": "Thirty",
	"40": "Forty",
	"50": "Fifty",
	"60": "Sixty",
	"70": "Seventy",
	"80": "Eighty",
	"90": "Ninety",
}

cats = ["Billion", "Million", "Thousand"]


def convert_hundred(n):
	n = str(n)
	if len(n) < 2:
		n = "0" + n
	if n[0] == "1" or n[1] == "0":
		return digit_text[n]
	if int(n) < 10:
		return digit_text["0" + n[1]]
	return digit_text[n[0] + "0"] + " " + digit_text["0" + n[1]]


def convert_thousand(n):
	assert n < 1000
	if n < 10:
		return digit_text["0" + str(n)]
	if n < 100:
		return convert_hundred(n)
	n = str(n)
	if n[1:] == "00":
		return f"{digit_text['0' + n[0]]} Hundred"
	return f"{digit_text['0' + n[0]]} Hundred {convert_hundred(int(n[1:]))}"


def convert(n):
	n = str(n)
	if n == "0":
		return "Zero"
	n = n.zfill(12)
	s = []
	for i, cat in enumerate(cats):
		start, end = i * 3, (i + 1) * 3
		if int(n[start:end]) > 0:
			s.append(convert_thousand(int(n[start:end])) + " " + cat)
	if int(n[-3:]) > 0:
		s.append(convert_thousand(int(n[-3:])))
	if s:
		return ", ".join(s)
	return None


def testRandom():
	import random

	k = random.randrange(999999999999)
	print(k)
	print(convert(k))


if __name__ == "__main__":
	for arg in sys.argv[1:]:
		arg = arg.replace(",", "")
		try:
			k = int(arg)
		except ValueError:  # noqa: PERF203
			print(f"{arg}: non-numeric argument")
		else:
			if k > MAX_NUM:
				print(f"{k:,}: number can not be more than {MAX_NUM:,}")
			else:
				print(f"{k:,}\n{convert(k)}")
