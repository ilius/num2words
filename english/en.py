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

en_and     = ", "
en_zero    = "Zero"
en_hundred = "Hundred"

small_words = {
	0:  en_zero,
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
	11: "Eleven",
	12: "Twelve",
	13: "Thirteen",
	14: "Fourteen",
	15: "Fifteen",
	16: "Sixteen",
	17: "Seventeen",
	18: "Eighteen",
	19: "Nineteen",
	20: "Twenty",
	30: "Thirty",
	40: "Forty",
	50: "Fifty",
	60: "Sixty",
	70: "Seventy",
	80: "Eighty",
	90: "Ninety",
}

big_words = [
	"One",
	"Thousand",
	"Million",
	"Billion",
]

def extractGroupsByString(numStr: str) -> list[int]:
	digitCount = len(numStr)
	groupCount = digitCount // 3
	groups = [
		int(numStr[digitCount-3*i-3:digitCount-3*i])
		for i in range(groupCount)
	]
	m = digitCount % 3
	if m > 0:
		groups.append(int(numStr[:m]))
	return groups


def bigIntCountDigits(bn: int) -> int:
	if bn == 0:
		return 1
	count = 0
	while bn != 0:
		bn = bn // 10
		count += 1
	return count


def extractGroupsByBigInt(bn: int, digitCount: int) -> list[int]:
	groupCount = digitCount // 3
	groups = [0] * groupCount
	for i in range(groupCount):
		div, mod = divmod(bn, 1000)
		groups[i] = mod
		bn = div
	m = digitCount % 3
	if m > 0:
		groups.append(bn)
	return groups



# n >= 1000
def convertLarge(groups: list[int]) -> str:
	k = len(groups)
	w_groups: list[str] = []
	for i in range(k):
		p = groups[i]
		if p == 0:
			continue
		if i == 0:
			w_groups.append(convertSmall(p))
			continue
		order = ""
		if i < len(big_words):
			order = big_words[i]
		else:
			order = ""
			d = i // 3
			m = i % 3
			t9 = big_words[3]
			for j in range(d):
				if j > 0:
					order += " "
				order += t9
			if m != 0:
				if order != "":
					order = " " + order
				order = big_words[m] + order
		w_groups.append(convertSmall(p)+" "+order)
	return en_and.join(reversed(w_groups))


# num < 1000
def convertSmall(num: int) -> str:
	if num in small_words:
		return small_words[num]
	ones = num % 10
	tens = (num % 100) // 10
	hundreds = num // 100
	result = ""
	if hundreds != 0:
		result += small_words[hundreds] + " " + en_hundred
		if tens != 0 or ones != 0:
			result += " "
	if tens != 0:
		word = small_words.get(num%100)
		if word is not None:
			result += word
			return result # OK, Done
		result += small_words[tens*10]
		if ones != 0:
			result += " "
	if ones != 0:
		result += small_words[ones]
	return result


# ConvertString: only for non-negative integers
def convert_string(st: str) -> str:
	if len(st) <= 3: # n <= 999
		return convertSmall(int(st))
	# n >= 1000
	return convertLarge(extractGroupsByString(st))


# convert_int: only for non-negative integers
def convert_int(bn: int) -> str:
	if bn < 0:
		return "Negative " + convert_int(abs(bn))
	digitCount = bigIntCountDigits(bn)
	if digitCount <= 3: # n <= 999
		return convertSmall(bn)
	# n >= 1000
	return convertLarge(extractGroupsByBigInt(bn, digitCount))


if __name__ == "__main__":
	for arg in sys.argv[1:]:
		arg = arg.replace(",", "")
		print(f"{arg}\n{convert_string(arg)}")
		try:
			k = int(arg)
		except ValueError:  # noqa: PERF203
			print(f"{arg}: non-numeric argument")
		else:
			print(f"{k:,}\n{convert_int(k)}")
