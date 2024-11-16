#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# File: num2words/fa.py
#
# Author: Saeed Rasooli <saeed.gnu@gmail.com>    (ilius)
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

fa_zero = "صفر"

faBaseNum = {
	1: "یک",
	2: "دو",
	3: "سه",
	4: "چهار",
	5: "پنج",
	6: "شش",
	7: "هفت",
	8: "هشت",
	9: "نه",
	10: "ده",
	11: "یازده",
	12: "دوازده",
	13: "سیزده",
	14: "چهارده",
	15: "پانزده",
	16: "شانزده",
	17: "هفده",
	18: "هجده",
	19: "نوزده",
	20: "بیست",
	30: "سی",
	40: "چهل",
	50: "پنجاه",
	60: "شصت",
	70: "هفتاد",
	80: "هشتاد",
	90: "نود",
	100: "صد",
	200: "دویست",
	300: "سیصد",
	500: "پانصد",
}
faBaseNumKeys = set(faBaseNum.keys())

faBigNumFirst = ["یک", "هزار", "میلیون"]

# European
faBigNumEU = faBigNumFirst + ["میلیارد", "بیلیون", "بیلیارد", "تریلیون", "تریلیارد"]

# American
faBigNumUS = faBigNumFirst + [
	"بیلیون",
	"تریلیون",
	"کوآدریلیون",
	"کوینتیلیون",
	"سکستیلیون",
]

# Common in Iran (the rest are uncommon or mistaken)
faBigNumIran = faBigNumFirst + ["میلیارد", "تریلیون"]


faBigNum = faBigNumIran


def extractGroupsByString(st):
	n = len(st)
	d, m = divmod(n, 3)
	parts = [int(st[n - 3 * i - 3 : n - 3 * i]) for i in range(d)]
	if m > 0:
		parts.append(int(st[:m]))
	return parts


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
	w_groups = []
	for i in range(k):
		faOrder = ""
		p = groups[i]
		if p == 0:
			continue
		if i == 0:
			wpart = convertSmall(p)
		else:
			if i < len(faBigNum):
				faOrder = faBigNum[i]
			else:
				faOrder = ""
				d, m = divmod(i, 3)
				t9 = faBigNum[3]
				for j in range(d):
					if j > 0:
						faOrder += "‌"
					faOrder += t9
				if m != 0:
					if faOrder != "":
						faOrder = "‌" + faOrder
					faOrder = faBigNum[m] + faOrder
			wpart = faOrder if i == 1 and p == 1 else convertSmall(p) + " " + faOrder
		w_groups.append(wpart)
	return " و ".join(reversed(w_groups))


# num < 1000
def convertSmall(n: int) -> str:
	if n == 0:
		return fa_zero
	if n in faBaseNumKeys:
		return faBaseNum[n]
	y = n % 10
	d = int((n % 100) / 10)
	s = int(n / 100)
	# print s, d, y
	dy = 10 * d + y
	fa = ""
	if s != 0:
		if s * 100 in faBaseNumKeys:
			fa += faBaseNum[s * 100]
		else:
			fa += faBaseNum[s] + faBaseNum[100]
		if d != 0 or y != 0:
			fa += " و "
	if d != 0:
		if dy in faBaseNumKeys:
			fa += faBaseNum[dy]
			return fa
		fa += faBaseNum[d * 10]
		if y != 0:
			fa += " و "
	if y != 0:
		fa += faBaseNum[y]
	return fa


def convert_str(st):
	if len(st) > 3:
		return convertLarge(extractGroupsByString(st))

	# now assume that n <= 999
	return convertSmall(int(st))


# convert_int: only for non-negative integers
def convert_int(bn: int) -> str:
	if bn < 0:
		return "Negative " + convert_int(abs(bn))
	digitCount = bigIntCountDigits(bn)
	if digitCount <= 3:  # n <= 999
		return convertSmall(bn)
	# n >= 1000
	return convertLarge(extractGroupsByBigInt(bn, digitCount))


def convert_str_ordinal(st: str):
	if st == "1":
		return "اول"  # or "یکم"
	if st == "10":
		return "دهم"
	norm_fa = convert_str(st)
	if not norm_fa:
		return ""
	if norm_fa.endswith("ی"):
		norm_fa += "‌ام"
	elif norm_fa.endswith("سه"):
		norm_fa = norm_fa[:-1] + "وم"
	else:
		norm_fa += "م"
	return norm_fa


def convert_int_ordinal(num):
	if num == 1:
		return "اول"  # or "یکم"
	if num == 10:
		return "دهم"
	norm_fa = convert_int(num)
	if not norm_fa:
		return ""
	if norm_fa.endswith("ی"):
		norm_fa += "‌ام"
	elif norm_fa.endswith("سه"):
		norm_fa = norm_fa[:-1] + "وم"
	else:
		norm_fa += "م"
	return norm_fa


if __name__ == "__main__":
	for arg in sys.argv[1:]:
		arg = arg.replace(",", "")
		try:
			k = int(arg)
		except ValueError:  # noqa: PERF203
			print(f"{arg}: non-numeric argument")
		else:
			print(f"{k:,}\n{convert_int(k)}\n{convert_int_ordinal(k)}\n")
