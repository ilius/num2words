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


# Based on https://github.com/ahmadRagheb/number2word-Arabic

import locale
import sys
import typing

ar_zero = "صفر"


class SmallWord(typing.NamedTuple):
	male: str
	female: str

	def get(self, mf: str) -> str:
		return getattr(self, mf)

# words = list[
# 	{
# 		male=list[
# 			{
# 				"0": "",
# 				"1": "واحد",
# 				"2": "اثنان",
# 				"3": "ثلاثة",
# 				"4": "أربعة",
# 				"5": "خمسة",
# 				"6": "ستة",
# 				"7": "سبعة",
# 				"8": "ثمانية",
# 				"9": "تسعة",
# 				"10": "عشرة",
# 				"11": "أحد عشر",
# 				"12": "اثنا عشر",
# 				"13": "ثلاثة عشر",
# 				"14": "أربعة عشر",
# 				"15": "خمسة عشر",
# 				"16": "ستة عشر",
# 				"17": "سبعة عشر",
# 				"18": "ثمانية عشر",
# 				"19": "تسعة عشر",
# 				"20": "عشرون",
# 				"30": "ثلاثون",
# 				"40": "أربعون",
# 				"50": "خمسون",
# 				"60": "ستون",
# 				"70": "سبعون",
# 				"80": "ثمانون",
# 				"90": "تسعون",
# 				"100": "مئة",
# 				"200": "مئتان",
# 				"300": "ثلاثمئة",
# 				"400": "أربعمئة",
# 				"500": "خمسمئة",
# 				"600": "ستمئة",
# 				"700": "سبعمئة",
# 				"800": "ثمانمئة",
# 				"900": "تسعمئة",
# 			}
# 		],
# 		female=list[
# 			{
# 				"0": "",
# 				"1": "واحدة",
# 				"2": "اثنتان",
# 				"3": "ثلاث",
# 				"4": "أربع",
# 				"5": "خمس",
# 				"6": "ست",
# 				"7": "سبع",
# 				"8": "ثمان",
# 				"9": "تسع",
# 				"10": "عشر",
# 				"11": "إحدى عشرة",
# 				"12": "ثنتا عشرة",
# 				"13": "ثلاث عشرة",
# 				"14": "أربع عشرة",
# 				"15": "خمس عشرة",
# 				"16": "ست عشرة",
# 				"17": "سبع عشرة",
# 				"18": "ثمان عشرة",
# 				"19": "تسع عشرة",
# 				"20": "عشرون",
# 				"30": "ثلاثون",
# 				"40": "أربعون",
# 				"50": "خمسون",
# 				"60": "ستون",
# 				"70": "سبعون",
# 				"80": "ثمانون",
# 				"90": "تسعون",
# 				"100": "مئة",
# 				"200": "مئتان",
# 				"300": "ثلاثمئة",
# 				"400": "أربعمئة",
# 				"500": "خمسمئة",
# 				"600": "ستمئة",
# 				"700": "سبعمئة",
# 				"800": "ثمانمئة",
# 				"900": "تسعمئة",
# 			}
# 		],
# 	}
# ]


words = {
	"0": SmallWord(
		male="",
		female="",
	),
	"1": SmallWord(
		female="واحدة",
		male="واحد",
	),
	"2": SmallWord(
		female="اثنتان",
		male="اثنان",
	),
	"3": SmallWord(
		female="ثلاث",
		male="ثلاثة",
	),
	"4": SmallWord(
		female="أربع",
		male="أربعة",
	),
	"5": SmallWord(
		female="خمس",
		male="خمسة",
	),
	"6": SmallWord(
		female="ست",
		male="ستة",
	),
	"7": SmallWord(
		female="سبع",
		male="سبعة",
	),
	"8": SmallWord(
		female="ثمان",
		male="ثمانية",
	),
	"9": SmallWord(
		female="تسع",
		male="تسعة",
	),
	"10": SmallWord(
		female="عشر",
		male="عشرة",
	),
	"11": SmallWord(
		female="إحدى عشرة",
		male="أحد عشر",
	),
	"12": SmallWord(
		female="ثنتا عشرة",
		male="اثنا عشر",
	),
	"13": SmallWord(
		female="ثلاث عشرة",
		male="ثلاثة عشر",
	),
	"14": SmallWord(
		female="أربع عشرة",
		male="أربعة عشر",
	),
	"15": SmallWord(
		female="خمس عشرة",
		male="خمسة عشر",
	),
	"16": SmallWord(
		female="ست عشرة",
		male="ستة عشر",
	),
	"17": SmallWord(
		female="سبع عشرة",
		male="سبعة عشر",
	),
	"18": SmallWord(
		female="ثمان عشرة",
		male="ثمانية عشر",
	),
	"19": SmallWord(
		female="تسع عشرة",
		male="تسعة عشر",
	),
	"20": SmallWord(
		female="عشرون",
		male="عشرون",
	),
	"30": SmallWord(
		female="ثلاثون",
		male="ثلاثون",
	),
	"40": SmallWord(
		female="أربعون",
		male="أربعون",
	),
	"50": SmallWord(
		female="خمسون",
		male="خمسون",
	),
	"60": SmallWord(
		female="ستون",
		male="ستون",
	),
	"70": SmallWord(
		female="سبعون",
		male="سبعون",
	),
	"80": SmallWord(
		female="ثمانون",
		male="ثمانون",
	),
	"90": SmallWord(
		female="تسعون",
		male="تسعون",
	),
	"100": SmallWord(
		female="مائة",
		male="مائة",
	),
	"200": SmallWord(
		female="مئتان",
		male="مئتان",
	),
	"300": SmallWord(
		female="ثلاثمائة",
		male="ثلاثمائة",
	),
	"400": SmallWord(
		female="أربعمائة",
		male="أربعمائة",
	),
	"500": SmallWord(
		female="خمسمائة",
		male="خمسمائة",
	),
	"600": SmallWord(
		female="ستمائة",
		male="ستمائة",
	),
	"700": SmallWord(
		female="سبعمائة",
		male="سبعمائة",
	),
	"800": SmallWord(
		female="ثمانمائة",
		male="ثمانمائة",
	),
	"900": SmallWord(
		female="تسعمائة",
		male="تسعمائة",
	),
}


class number2word:
	def __init__(self, number):
		self.number = number

	def to_string(self):
		returnmsg = ""
		# // convert number into array of(string) number each case
		# // -------number: 121210002876 - ---------- //
		# //   0          1          2          3 //
		# // '121'      '210'      '002'      '876'

		# to format number like 10012 became 100,12
		my_number = self.number
		english_format_number = self.convert_number(my_number)
		# we split it into array
		array_number = english_format_number.split(",")
		# array_number is type of list
		# frappe.throw(array_number)
		# convert each number(hundred) to arabic
		for i in range(len(array_number)):
			place = len(array_number) - i
			returnmsg = returnmsg + self.convert(array_number[i], place)
			# if array_number[i+1] and array_number[i+1] > 0:
			if 0 <= i < len(array_number) - 1:
				returnmsg = returnmsg + " و "
		return returnmsg.strip()

	@staticmethod
	def convert_number(number):
		locale.setlocale(locale.LC_ALL, "")
		x = number
		x = int(float(x))
		x1 = locale.format_string("%d", x, grouping=True)
		# frappe.throw(type(x).__name__)
		if x < 0 or x > 999999999999:
			raise ValueError("Value out of range")
		return x1

	@staticmethod
	def convert(number: str | int, place: int):
		number = str(number)
		if number in ("", "0"):
			return ""

		# take in charge the sex of NUMBERED
		#  sex = self.sex
		returnmsg = ""

		# sex = "male"
		# the number word in arabic for masculine and feminine

		# take in charge the different way of writing the thousands and millions ...
		# mil = list[
		#     '2' : list['1' : 'ألف', '2' : 'ألفان', '3' : 'آلاف'],
		#     '3' : list['1' : 'مليون', '2' : 'مليونان', '3' : 'ملايين'],
		#     '4' : list['1' : 'مليار', '2' : 'ملياران', '3' : 'مليارات'] ]

		mf = {
			1: "male",
			2: "male",
			3: "male",
			4: "male",
		}
		number_length = len(number)

		# we are dealing with 3 digits number in each loop the main method calls convert
		# method and pass a string with tree digit in it ...
		# for example 123 or 19 or 3 ..

		# we will clean left zero for example 001 will be 1 ,,
		# 012 will be 12

		if number[0] == 0:
			if number[1] == 0:
				return int(number[:-1])
			return int(number[:-2])

		# switching number length
		# if number have on digits like "1"
		returnmsg = ""
		if number_length == 1:
			# number=number+'one'
			if place == 1:
				returnmsg = returnmsg + words[number].get(mf[place])
			if place == 2:
				if int(number) == 1:
					returnmsg = " ألف"
				elif int(number) == 2:
					returnmsg = " ألفان"
				else:
					returnmsg = returnmsg + words[number].get(mf[place]) + " آلاف"
			if place == 3:
				if int(number) == 1:
					returnmsg = returnmsg + " مليون"
				elif int(number) == 2:
					returnmsg = returnmsg + " مليونان"
				else:
					returnmsg = returnmsg + words[number].get(mf[place]) + " ملايين"
			if place == 4:
				if int(number) == 1:
					returnmsg = returnmsg + " مليار"
				elif int(number) == 2:
					returnmsg = returnmsg + " ملياران"
				else:
					returnmsg = returnmsg + words[number].get(mf[place]) + " مليارات"

		elif number_length == 2:
			# number=number+'two'

			# if words[number].get(mf[place]):
			if number in words:
				returnmsg = returnmsg + words[number].get(mf[place])
			else:
				twoy = int(number[0]) * 10
				ony = number[1]
				returnmsg = (
					returnmsg
					+ words[ony].get(mf[place])
					+ " و "
					+ words[str(twoy)].get(mf[place])
				)

			if place == 2:
				returnmsg = returnmsg + " ألف"
			if place == 3:
				returnmsg = returnmsg + " مليون"
			if place == 4:
				returnmsg = returnmsg + " مليار"

		elif number_length == 3:
			# number=number+'three'
			# if words[number].get(mf[place]):
			if number in words:
				returnmsg = returnmsg + words[str(number)].get(mf[place])

				if number == "200":
					returnmsg = " مئتا"

				if place == 2:
					returnmsg = returnmsg + " ألف"
				if place == 3:
					returnmsg = returnmsg + " مليون"
				if place == 4:
					returnmsg = returnmsg + " مليار"

				return returnmsg

			threey = int(number[0]) * 100
			threey = str(threey)
			if words[threey].get(mf[place]):
				returnmsg = returnmsg + words[threey].get(mf[place])

			twoyony = (int(number[1]) * 10) + int(number[2])
			if int(twoyony) == 2:
				if place == 1:
					twoyony = words["2"].get(mf[place])
				if place == 2:
					twoyony = " ألفان"
				if place == 3:
					twoyony = " مليونان"
				if place == 4:
					twoyony = " ملياران"

				if int(threey) != 0:
					twoyony = "و " + str(twoyony)

				returnmsg = returnmsg + " " + twoyony

			elif int(twoyony) == 1:
				twoyony = str(twoyony)
				if place == 1:
					twoyony = words["1"].get(mf[place])
				if place == 2:
					twoyony = "ألف"
				if place == 3:
					twoyony = "مليون"
				if place == 4:
					twoyony = "مليار"

				if int(threey) != 0:
					twoyony = "و " + str(twoyony)

				returnmsg = returnmsg + " " + twoyony

			else:
				# if words[twoyony].get(mf[place]):
				twoyony = str(twoyony)
				if twoyony in words:
					# if words.has_key(twoyony):
					twoyony = words[twoyony].get(mf[place])
				else:
					twoy = int(number[1]) * 10
					twoy = str(twoy)
					ony = number[2]
					twoyony = (
						words[ony].get(mf[place]) + " و " + words[twoy].get(mf[place])
					)
				if twoyony and int(threey) != 0:
					returnmsg = returnmsg + " و " + twoyony
				else:
					returnmsg = returnmsg + " " + twoyony

				if place == 2:
					returnmsg = returnmsg + " ألف"
				if place == 3:
					returnmsg = returnmsg + " مليون"
				if place == 4:
					returnmsg = returnmsg + " مليار"

		return returnmsg


def convert(st: str) -> str:
	if st == "0":
		return ar_zero
	return number2word(int(st)).to_string()


if __name__ == "__main__":
	for arg in sys.argv[1:]:
		number = int(arg)
		num = number2word(number)
		print(num.to_string())
