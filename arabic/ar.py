#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# Copyright (c) 2024 Saeed Rasooli
# Copyright (c) 2017 ahmadRagheb
# Copyright (c) 2015, Frappe Technologies Pvt. Ltd. and contributors
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.


import sys
import typing
from dataclasses import dataclass

ar_and = " و "
ar_zero = "صفر"


class SmallWord(typing.NamedTuple):
	male: str
	female: str

	def get(self, mf: str) -> str:
		return getattr(self, mf)


small_words = {
	0: SmallWord(
		male="",
		female="",
	),
	1: SmallWord(
		female="واحدة",
		male="واحد",
	),
	2: SmallWord(
		female="اثنتان",
		male="اثنان",
	),
	3: SmallWord(
		female="ثلاث",
		male="ثلاثة",
	),
	4: SmallWord(
		female="أربع",
		male="أربعة",
	),
	5: SmallWord(
		female="خمس",
		male="خمسة",
	),
	6: SmallWord(
		female="ست",
		male="ستة",
	),
	7: SmallWord(
		female="سبع",
		male="سبعة",
	),
	8: SmallWord(
		female="ثمان",
		male="ثمانية",
	),
	9: SmallWord(
		female="تسع",
		male="تسعة",
	),
	10: SmallWord(
		female="عشر",
		male="عشرة",
	),
	11: SmallWord(
		female="إحدى عشرة",
		male="أحد عشر",
	),
	12: SmallWord(
		female="ثنتا عشرة",
		male="اثنا عشر",
	),
	13: SmallWord(
		female="ثلاث عشرة",
		male="ثلاثة عشر",
	),
	14: SmallWord(
		female="أربع عشرة",
		male="أربعة عشر",
	),
	15: SmallWord(
		female="خمس عشرة",
		male="خمسة عشر",
	),
	16: SmallWord(
		female="ست عشرة",
		male="ستة عشر",
	),
	17: SmallWord(
		female="سبع عشرة",
		male="سبعة عشر",
	),
	18: SmallWord(
		female="ثمان عشرة",
		male="ثمانية عشر",
	),
	19: SmallWord(
		female="تسع عشرة",
		male="تسعة عشر",
	),
	20: SmallWord(
		female="عشرون",
		male="عشرون",
	),
	30: SmallWord(
		female="ثلاثون",
		male="ثلاثون",
	),
	40: SmallWord(
		female="أربعون",
		male="أربعون",
	),
	50: SmallWord(
		female="خمسون",
		male="خمسون",
	),
	60: SmallWord(
		female="ستون",
		male="ستون",
	),
	70: SmallWord(
		female="سبعون",
		male="سبعون",
	),
	80: SmallWord(
		female="ثمانون",
		male="ثمانون",
	),
	90: SmallWord(
		female="تسعون",
		male="تسعون",
	),
	100: SmallWord(
		female="مائة",
		male="مائة",
	),
	200: SmallWord(
		female="مئتان",
		male="مئتان",
	),
	300: SmallWord(
		female="ثلاثمائة",
		male="ثلاثمائة",
	),
	400: SmallWord(
		female="أربعمائة",
		male="أربعمائة",
	),
	500: SmallWord(
		female="خمسمائة",
		male="خمسمائة",
	),
	600: SmallWord(
		female="ستمائة",
		male="ستمائة",
	),
	700: SmallWord(
		female="سبعمائة",
		male="سبعمائة",
	),
	800: SmallWord(
		female="ثمانمائة",
		male="ثمانمائة",
	),
	900: SmallWord(
		female="تسعمائة",
		male="تسعمائة",
	),
}


class GroupWord(typing.NamedTuple):
	normal: str
	genitive: str
	appended: str
	plural: str

	def get(self, mf: str) -> str:
		return getattr(self, mf)


group_words: list[GroupWord] = [
	GroupWord(  # 10^2 Hundred
		normal="مائة",
		genitive="مئتا",
		appended="",
		plural="",
	),
	GroupWord(  # 10^3 Thousand
		normal="ألف",
		genitive="ألفا",
		appended="ألفاً",
		plural="آلاف",
	),
	GroupWord(  # 10^6 Million
		normal="مليون",
		genitive="مليونا",
		appended="مليوناً",
		plural="ملايين",
	),
	GroupWord(  # 10^9 Billion
		normal="مليار",
		genitive="مليارا",
		appended="ملياراً",
		plural="مليارات",
	),
	GroupWord(  # 10^12 Trillion
		normal="تريليون",
		genitive="تريليونا",
		appended="تريليوناً",
		plural="تريليونات",
	),
	GroupWord(  # 10^15 Quadrillion
		normal="كوادريليون",
		genitive="كوادريليونا",
		appended="كوادريليوناً",
		plural="كوادريليونات",
	),
	GroupWord(  # 10^18 Quintillion
		normal="كوينتليون",
		genitive="كوينتليونا",
		appended="كوينتليوناً",
		plural="كوينتليونات",
	),
	GroupWord(  # 10^21 Sextillion
		normal="سكستيليون",
		genitive="سكستيليونا",
		appended="سكستيليوناً",
		plural="سكستيليونات",
	),
]


@dataclass
class Group:
	level: int
	number: int


def extractGroupsByString(st) -> list[Group]:
	n = len(st)
	d, m = divmod(n, 3)
	numbers = [int(st[n - 3 * i - 3 : n - 3 * i]) for i in range(d)]
	if m > 0:
		numbers.append(int(st[:m]))
	return [Group(level=level, number=number) for level, number in enumerate(numbers)]


def bigIntCountDigits(bn: int) -> int:
	if bn == 0:
		return 1
	count = 0
	while bn != 0:
		bn = bn // 10
		count += 1
	return count


def extractGroupsByBigInt(bn: int, digitCount: int) -> list[Group]:
	groupCount = digitCount // 3
	numbers = [0] * groupCount
	for i in range(groupCount):
		div, mod = divmod(bn, 1000)
		numbers[i] = mod
		bn = div
	m = digitCount % 3
	if m > 0:
		numbers.append(bn)
	return [Group(level=level, number=number) for level, number in enumerate(numbers)]


def getDigitWord(digit: int, groupLevel: int, feminine: bool) -> str:
	if feminine and groupLevel == 0:
		return small_words[digit].female
	return small_words[digit].male


def processTens(tens: int, hundreds: int, groupLevel: int, feminine: bool) -> str:
	if tens < 20:
		# if we are processing under 20 numbers
		if tens == 2 and hundreds == 0 and groupLevel > 0:
			# This is special case for number 2 when it comes alone in the group
			# In the case of individuals
			return group_words[groupLevel].genitive + "ن"
		# General case
		if tens == 1 and groupLevel > 0:
			return group_words[groupLevel].normal
		# Get Feminine status for this digit
		return getDigitWord(tens, groupLevel, feminine)

	ones = tens % 10
	if ones == 0:
		return small_words[tens].male

	return (
		getDigitWord(ones, groupLevel, feminine)
		+ ar_and
		+ small_words[int(tens / 10) * 10].male
	)


def processGroup(group: Group, feminine: bool) -> str:
	tens = group.number % 100
	hundreds = int(group.number / 100) * 100
	if hundreds == 0:
		return processTens(tens, hundreds, group.level, feminine)

	if tens == 0:
		if hundreds == 200 and group.level > 0:
			# genitive case - حالة المضاف
			return group_words[0].genitive
		return small_words[hundreds].male

	# normal case - الحالة العادية
	return (
		small_words[hundreds].male
		+ ar_and
		+ processTens(tens, hundreds, group.level, feminine)
	)


# groupNumber < 1000
def convertGroup(group: Group, feminine: bool, appending: bool) -> str:
	# convert group into its text
	groupDescription = processGroup(group, feminine)
	if group.level == 0:
		return groupDescription

	if group.number != 2 and group.number % 100 != 1:
		if group.number >= 3 and group.number <= 10:
			# for numbers between 3 and 9 we use plural name
			return groupDescription + " " + group_words[group.level].plural
		if appending:
			# use appending case
			return groupDescription + " " + group_words[group.level].appended
		# use normal case
		return groupDescription + " " + group_words[group.level].normal

	return groupDescription


def convert_int(num: int) -> str:
	if num == 0:
		return ar_zero
	groups = extractGroupsByBigInt(num, bigIntCountDigits(num))
	result: list[str] = []
	for group in groups:
		if group.number == 0:
			continue
		groupResult = convertGroup(group, False, len(result) > 0)
		result.insert(0, groupResult)
	return ar_and.join(result)


def convert_string(st: str) -> str:
	if st == "0":
		return ar_zero
	groups = extractGroupsByString(st)
	result: list[str] = []
	for group in groups:
		if group.number == 0:
			continue
		groupResult = convertGroup(group, False, len(result) > 0)
		result.insert(0, groupResult)
	return ar_and.join(result)


if __name__ == "__main__":
	for arg in sys.argv[1:]:
		number = int(arg)
		print(convert_int(number))
