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

var small_words_masc = map[int]string{
	1:  "واحد",
	2:  "اثنان",
	3:  "ثلاثة",
	4:  "أربعة",
	5:  "خمسة",
	6:  "ستة",
	7:  "سبعة",
	8:  "ثمانية",
	9:  "تسعة",
	10: "عشرة",
	11: "أحد عشر",
	12: "اثنا عشر",
	13: "ثلاثة عشر",
	14: "أربعة عشر",
	15: "خمسة عشر",
	16: "ستة عشر",
	17: "سبعة عشر",
	18: "ثمانية عشر",
	19: "تسعة عشر",
}

var small_words_feminine = map[int]string{
	1:  "إحدى",
	2:  "اثنتان",
	3:  "ثلاث",
	4:  "أربع",
	5:  "خمس",
	6:  "ست",
	7:  "سبع",
	8:  "ثمان",
	9:  "تسع",
	10: "عشر",
	11: "إحدى عشرة",
	12: "اثنتا عشرة",
	13: "ثلاث عشرة",
	14: "أربع عشرة",
	15: "خمس عشرة",
	16: "ست عشرة",
	17: "سبع عشرة",
	18: "ثماني عشرة",
	19: "تسع عشرة",
}

var ten_words = map[int]string{
	20: "عشرون",
	30: "ثلاثون",
	40: "أربعون",
	50: "خمسون",
	60: "ستون",
	70: "سبعون",
	80: "ثمانون",
	90: "تسعون",
}

var arabicHundreds = []string{
	"",
	"مائة",
	"مئتان",
	"ثلاثمائة",
	"أربعمائة",
	"خمسمائة",
	"ستمائة",
	"سبعمائة",
	"ثمانمائة",
	"تسعمائة",
}

var arabicTwos = []string{
	"مئتان",
	"ألفان",
	"مليونان",
	"ملياران",
	"تريليونان",
	"كوادريليونان",
	"كوينتليونان",
	"سكستيليونان",
}

var arabicAppendedTwos = []string{
	"مئتا",
	"ألفا",
	"مليونا",
	"مليارا",
	"تريليونا",
	"كوادريليونا",
	"كوينتليونا",
	"سكستيليونا",
}

var arabicGroup = []string{
	"مائة",
	"ألف",
	"مليون",
	"مليار",
	"تريليون",
	"كوادريليون",
	"كوينتليون",
	"سكستيليون",
}

var arabicAppendedGroup = []string{
	"",
	"ألفاً",
	"مليوناً",
	"ملياراً",
	"تريليوناً",
	"كوادريليوناً",
	"كوينتليوناً",
	"سكستيليوناً",
}

var arabicPluralGroups = []string{
	"",
	"آلاف",
	"ملايين",
	"مليارات",
	"تريليونات",
	"كوادريليونات",
	"كوينتليونات",
	"سكستيليونات",
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
		groupDescription := processArabicGroup(
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
						result = arabicPluralGroups[groupLevel] + " " + result
					} else {
						if len(result) > 0 {
							// use appending case
							result = arabicAppendedGroup[groupLevel] + " " + result
						} else {
							// use normal case
							result = arabicGroup[groupLevel] + " " + result
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

func getDigitFeminineStatus(digit int, groupLevel int, feminine bool) string {
	if feminine && (groupLevel == -1 || groupLevel == 0) {
		return small_words_feminine[digit]
	}
	return small_words_masc[digit]
}

func processArabicGroupTens(tens int, hundreds int, groupLevel int, feminine bool) string {
	if tens < 20 {
		// if we are processing under 20 numbers
		if tens == 2 && hundreds == 0 && groupLevel > 0 {
			// This is special case for number 2 when it comes alone in the group
			return arabicTwos[groupLevel] // في حالة الافراد
		}
		// General case
		if tens == 1 && groupLevel > 0 {
			return arabicGroup[groupLevel]
		}
		// Get Feminine status for this digit
		return getDigitFeminineStatus(tens, groupLevel, feminine)
	}
	ones := tens % 10
	if ones == 0 {
		return ten_words[tens]
	}

	return getDigitFeminineStatus(ones, groupLevel, feminine) + ar_and + ten_words[tens/10*10]
}

func processArabicGroup(groupNumber int, groupLevel int, feminine bool) string {
	tens := groupNumber % 100
	hundreds := groupNumber / 100
	result := ""

	if hundreds > 0 {
		if tens == 0 && hundreds == 2 { // حالة المضاف
			if groupLevel == 0 {
				result = arabicHundreds[hundreds]
			} else {
				result = arabicAppendedTwos[0]
			}
		} else { // الحالة العادية
			result = arabicHundreds[hundreds]
		}
	}

	if tens > 0 {
		tmp_result := processArabicGroupTens(tens, hundreds, groupLevel, feminine)
		if result != "" {
			result += ar_and
		}
		result += tmp_result
	}

	return result
}
