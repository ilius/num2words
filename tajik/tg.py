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

tg_and = "у "
tg_zero = "сифр"

faBaseNum = {
	1: "Як",
	2: "ду",
	3: "Се",
	4: "Чаҳор",
	5: "Пань",
	6: "Шаш",
	7: "Ҳафт",
	8: "Ҳашт",
	9: "Ну",
	10: "даҳ",
	11: "Ёздаҳ",
	12: "дувоздаҳ",
	13: "Сенздаҳ",
	14: "Чаҳорум",
	15: "Понздаҳ",
	16: "Шонздаҳ",
	17: "Ҳабдаҳ",
	18: "Ҳаждаҳ",
	19: "Нуздаҳ",
	20: "Бист",
	30: "Сӣ",
	40: "Чил",
	50: "Панҷоҳ",
	60: "Шаст",
	70: "Ҳафтод",
	80: "Ҳаштод",
	90: "Навад",
	100: "Сад",
	200: "Дусад",
	300: "Сесад",
	500: "Панҷсад",
}
faBaseNumKeys = set(faBaseNum.keys())

faBigNumFirst = ["Як", "Ҳазор", "Миллион"]

# European
# faBigNumEU = faBigNumFirst + ["میلیارد", "بیلیون", "بیلیارد", "تریلیون", "تریلیارد"]

# American
# faBigNumUS = faBigNumFirst + [
# "بیلیون",
# "تریلیون",
# "کوآدریلیون",
# "کوینتیلیون",
# "سکستیلیون",
# ]

# Common in Iran (the rest are uncommon or mistaken)
faBigNumIran = faBigNumFirst + [
	"Миллиард",  # Milliard
	"Триллион",  # Trillion
]


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
	return tg_and.join(reversed(w_groups))


# num < 1000
def convertSmall(n: int) -> str:
	if n == 0:
		return tg_zero
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
			fa += tg_and
	if d != 0:
		if dy in faBaseNumKeys:
			fa += faBaseNum[dy]
			return fa
		fa += faBaseNum[d * 10]
		if y != 0:
			fa += tg_and
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


def _addOrdinalSuffix(norm_tg: str) -> str:
	if not norm_tg:
		return ""
	if norm_tg.endswith("ӣ"):
		return norm_tg[:-1] + "‌юм"
	elif norm_tg.endswith("се") or norm_tg.endswith("Се"):
		return norm_tg + "вум"
	return norm_tg + "юм"


def convert_str_ordinal(st: str):
	if st == "1":
		return "якум"
	if st == "10":
		return "даҳум"
	norm_tg = convert_str(st)
	return _addOrdinalSuffix(norm_tg)


def convert_int_ordinal(num):
	if num == 1:
		return "якум"
	if num == 10:
		return "даҳум"
	norm_tg = convert_int(num)
	return _addOrdinalSuffix(norm_tg)


if __name__ == "__main__":
	for arg in sys.argv[1:]:
		arg = arg.replace(",", "")
		try:
			k = int(arg)
		except ValueError:  # noqa: PERF203
			print(f"{arg}: non-numeric argument")
		else:
			print(f"{k:,}\n{convert_int(k)}\n{convert_int_ordinal(k)}\n")
