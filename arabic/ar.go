package arabic

// based on https://github.com/bluemix/NumberToArabicWords

import (
	"math/big"
	"strings"
)

var (
	big_0    = big.NewInt(0)
	big_1    = big.NewInt(1)
	big_1000 = big.NewInt(1000)
)

const (
	ar_and  = " و "
	ar_zero = "صفر"
)

var small_words = map[int][2]string{
	1: {
		"واحد",
		"إحدى",
	},
	2: {
		"اثنان",
		"اثنتان",
	},
	3: {
		"ثلاثة",
		"ثلاث",
	},
	4: {
		"أربعة",
		"أربع",
	},
	5: {
		"خمسة",
		"خمس",
	},
	6: {
		"ستة",
		"ست",
	},
	7: {
		"سبعة",
		"سبع",
	},
	8: {
		"ثمانية",
		"ثمان",
	},
	9: {
		"تسعة",
		"تسع",
	},
	10: {
		"عشرة",
		"عشر",
	},
	11: {
		"أحد عشر",
		"إحدى عشرة",
	},
	12: {
		"اثنا عشر",
		"اثنتا عشرة",
	},
	13: {
		"ثلاثة عشر",
		"ثلاث عشرة",
	},
	14: {
		"أربعة عشر",
		"أربع عشرة",
	},
	15: {
		"خمسة عشر",
		"خمس عشرة",
	},
	16: {
		"ستة عشر",
		"ست عشرة",
	},
	17: {
		"سبعة عشر",
		"سبع عشرة",
	},
	18: {
		"ثمانية عشر",
		"ثماني عشرة",
	},
	19: {
		"تسعة عشر",
		"تسع عشرة",
	},
	20: {
		"عشرون",
		"عشرون",
	},
	30: {
		"ثلاثون",
		"ثلاثون",
	},
	40: {
		"أربعون",
		"أربعون",
	},
	50: {
		"خمسون",
		"خمسون",
	},
	60: {
		"ستون",
		"ستون",
	},
	70: {
		"سبعون",
		"سبعون",
	},
	80: {
		"ثمانون",
		"ثمانون",
	},
	90: {
		"تسعون",
		"تسعون",
	},
	100: {
		"مائة",
		"مائة",
	},
	200: {
		"مئتان",
		"مئتان",
	},
	300: {
		"ثلاثمائة",
		"ثلاثمائة",
	},
	400: {
		"أربعمائة",
		"أربعمائة",
	},
	500: {
		"خمسمائة",
		"خمسمائة",
	},
	600: {
		"ستمائة",
		"ستمائة",
	},
	700: {
		"سبعمائة",
		"سبعمائة",
	},
	800: {
		"ثمانمائة",
		"ثمانمائة",
	},
	900: {
		"تسعمائة",
		"تسعمائة",
	},
}

type GroupWord struct {
	Normal   string
	Genitive string
	Appended string
	Plural   string
}

var group_words = []GroupWord{
	{ // 10^2 Hundred
		Normal:   "مائة",
		Genitive: "مئتا",
		Appended: "",
		Plural:   "",
	},
	{ // 10^3 Thousand
		Normal:   "ألف",
		Genitive: "ألفا",
		Appended: "ألفاً",
		Plural:   "آلاف",
	},
	{ // 10^6 Million
		Normal:   "مليون",
		Genitive: "مليونا",
		Appended: "مليوناً",
		Plural:   "ملايين",
	},
	{ // 10^9 Billion
		Normal:   "مليار",
		Genitive: "مليارا",
		Appended: "ملياراً",
		Plural:   "مليارات",
	},
	{ // 10^12 Trillion
		Normal:   "تريليون",
		Genitive: "تريليونا",
		Appended: "تريليوناً",
		Plural:   "تريليونات",
	},
	{ // 10^15 Quadrillion
		Normal:   "كوادريليون",
		Genitive: "كوادريليونا",
		Appended: "كوادريليوناً",
		Plural:   "كوادريليونات",
	},
	{ // 10^18 Quintillion
		Normal:   "كوينتليون",
		Genitive: "كوينتليونا",
		Appended: "كوينتليوناً",
		Plural:   "كوينتليونات",
	},
	{ // 10^21 Sextillion
		Normal:   "سكستيليون",
		Genitive: "سكستيليونا",
		Appended: "سكستيليوناً",
		Plural:   "سكستيليونات",
	},
}

func ConvertString(number string) string {
	num_big := &big.Int{}
	num_big.SetString(number, 10)
	return convertBigInt(num_big, false)
}

func ConvertBigInt(number *big.Int) string {
	return convertBigInt(number, false)
}

// // num < 1000
// func convertSmall(num uint16) string {

// }

func convertBigInt(numberOrig *big.Int, feminine bool) string {
	if numberOrig.Cmp(big_0) == 0 {
		return ar_zero
	}

	number := &big.Int{}
	number.SetBytes(numberOrig.Bytes())

	result := ""
	groupLevel := 0

	for number.Cmp(big_1) >= 0 {
		// separate number into groups
		groupNumberBig := &big.Int{}
		number.DivMod(number, big_1000, groupNumberBig)
		groupNumber := int(groupNumberBig.Int64())

		// convert group into its text
		groupDescription := processGroup(
			groupNumber,
			groupLevel,
			feminine,
		)

		// here we add the new converted group to the previous concatenated text
		if groupDescription != "" {
			if groupLevel > 0 {
				if len(result) > 0 {
					// FIXME: huh??
					result = "و " + result
				}
				if groupNumber != 2 && groupNumber%100 != 1 {
					if groupNumber >= 3 && groupNumber <= 10 {
						// for numbers between 3 and 9 we use plural name
						result = group_words[groupLevel].Plural + " " + result
					} else {
						if len(result) > 0 {
							// use appending case
							result = group_words[groupLevel].Appended + " " + result
						} else {
							// use normal case
							result = group_words[groupLevel].Normal + " " + result
						}
					}
				}
			}
			result = groupDescription + " " + result
		}

		groupLevel++
	}

	return strings.TrimSpace(result)
}

func getDigitWord(digit int, groupLevel int, feminine bool) string {
	if feminine && (groupLevel == -1 || groupLevel == 0) {
		return small_words[digit][1]
	}
	return small_words[digit][0]
}

func processTens(tens int, hundreds int, groupLevel int, feminine bool) string {
	if tens < 20 {
		// if we are processing under 20 numbers
		if tens == 2 && hundreds == 0 && groupLevel > 0 {
			// This is special case for number 2 when it comes alone in the group
			// In the case of individuals
			return group_words[groupLevel].Genitive + "ن"
		}
		// General case
		if tens == 1 && groupLevel > 0 {
			return group_words[groupLevel].Normal
		}
		// Get Feminine status for this digit
		return getDigitWord(tens, groupLevel, feminine)
	}
	ones := tens % 10
	if ones == 0 {
		return small_words[tens][0]
	}

	return getDigitWord(ones, groupLevel, feminine) + ar_and + small_words[tens/10*10][0]
}

func processGroup(groupNumber int, groupLevel int, feminine bool) string {
	tens := groupNumber % 100
	hundreds := groupNumber / 100 * 100

	if tens == 0 {
		if hundreds == 200 && groupLevel > 0 {
			// Genitive case: حالة المضاف
			return group_words[0].Genitive
		}
		return small_words[hundreds][0]
	}

	result := ""
	if hundreds > 0 {
		// الحالة العادية
		result = small_words[hundreds][0]
	}

	tmp_result := processTens(tens, hundreds, groupLevel, feminine)
	if result != "" {
		result += ar_and
	}
	result += tmp_result

	return result
}
