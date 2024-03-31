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

type SmallWord struct {
	Male   string
	Female string
}

var small_words = map[uint16]SmallWord{
	1: {
		Male:   "واحد",
		Female: "إحدى",
	},
	2: {
		Male:   "اثنان",
		Female: "اثنتان",
	},
	3: {
		Male:   "ثلاثة",
		Female: "ثلاث",
	},
	4: {
		Male:   "أربعة",
		Female: "أربع",
	},
	5: {
		Male:   "خمسة",
		Female: "خمس",
	},
	6: {
		Male:   "ستة",
		Female: "ست",
	},
	7: {
		Male:   "سبعة",
		Female: "سبع",
	},
	8: {
		Male:   "ثمانية",
		Female: "ثمان",
	},
	9: {
		Male:   "تسعة",
		Female: "تسع",
	},
	10: {
		Male:   "عشرة",
		Female: "عشر",
	},
	11: {
		Male:   "أحد عشر",
		Female: "إحدى عشرة",
	},
	12: {
		Male:   "اثنا عشر",
		Female: "اثنتا عشرة",
	},
	13: {
		Male:   "ثلاثة عشر",
		Female: "ثلاث عشرة",
	},
	14: {
		Male:   "أربعة عشر",
		Female: "أربع عشرة",
	},
	15: {
		Male:   "خمسة عشر",
		Female: "خمس عشرة",
	},
	16: {
		Male:   "ستة عشر",
		Female: "ست عشرة",
	},
	17: {
		Male:   "سبعة عشر",
		Female: "سبع عشرة",
	},
	18: {
		Male:   "ثمانية عشر",
		Female: "ثماني عشرة",
	},
	19: {
		Male:   "تسعة عشر",
		Female: "تسع عشرة",
	},
	20: {
		Male:   "عشرون",
		Female: "عشرون",
	},
	30: {
		Male:   "ثلاثون",
		Female: "ثلاثون",
	},
	40: {
		Male:   "أربعون",
		Female: "أربعون",
	},
	50: {
		Male:   "خمسون",
		Female: "خمسون",
	},
	60: {
		Male:   "ستون",
		Female: "ستون",
	},
	70: {
		Male:   "سبعون",
		Female: "سبعون",
	},
	80: {
		Male:   "ثمانون",
		Female: "ثمانون",
	},
	90: {
		Male:   "تسعون",
		Female: "تسعون",
	},
	100: {
		Male:   "مائة",
		Female: "مائة",
	},
	200: {
		Male:   "مئتان",
		Female: "مئتان",
	},
	300: {
		Male:   "ثلاثمائة",
		Female: "ثلاثمائة",
	},
	400: {
		Male:   "أربعمائة",
		Female: "أربعمائة",
	},
	500: {
		Male:   "خمسمائة",
		Female: "خمسمائة",
	},
	600: {
		Male:   "ستمائة",
		Female: "ستمائة",
	},
	700: {
		Male:   "سبعمائة",
		Female: "سبعمائة",
	},
	800: {
		Male:   "ثمانمائة",
		Female: "ثمانمائة",
	},
	900: {
		Male:   "تسعمائة",
		Female: "تسعمائة",
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

type Group struct {
	level  uint64
	number uint16
}

// groupNumber < 1000
func convertGroup(group Group, feminine bool, appending bool) string {
	// convert group into its text
	groupDescription := processGroup(group, feminine)
	if groupDescription == "" {
		return ""
	}
	if group.level == 0 {
		return groupDescription
	}
	if group.number != 2 && group.number%100 != 1 {
		if group.number >= 3 && group.number <= 10 {
			// for numbers between 3 and 9 we use plural name
			return groupDescription + " " + group_words[group.level].Plural
		}
		if appending {
			// use appending case
			return groupDescription + " " + group_words[group.level].Appended
		}
		// use normal case
		return groupDescription + " " + group_words[group.level].Normal

	}
	return groupDescription
}

func extractGroups(numberBytes []byte) []Group {
	number := &big.Int{}
	number.SetBytes(numberBytes)
	groups := []Group{}
	for number.Cmp(big_1) >= 0 {
		groupNumberBig := &big.Int{}
		number.DivMod(number, big_1000, groupNumberBig)
		groups = append(groups, Group{
			number: uint16(groupNumberBig.Int64()),
			level:  uint64(len(groups)),
		})
	}
	return groups
}

func convertBigInt(numberOrig *big.Int, feminine bool) string {
	if numberOrig.Cmp(big_0) == 0 {
		return ar_zero
	}
	// separate number into groups
	groups := extractGroups(numberOrig.Bytes())
	result := []string{}
	for _, group := range groups {
		groupResult := convertGroup(group, feminine, len(result) > 0)
		if groupResult == "" {
			continue
		}
		result = append([]string{groupResult}, result...)
	}
	return strings.Join(result, ar_and)
}

func getDigitWord(digit uint16, groupLevel uint64, feminine bool) string {
	if feminine && groupLevel == 0 {
		return small_words[digit].Female
	}
	return small_words[digit].Male
}

func processTens(tens uint16, hundreds uint16, groupLevel uint64, feminine bool) string {
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
		return small_words[tens].Male
	}
	return getDigitWord(ones, groupLevel, feminine) + ar_and + small_words[tens/10*10].Male
}

func processGroup(group Group, feminine bool) string {
	tens := group.number % 100
	hundreds := group.number / 100 * 100
	if hundreds == 0 {
		return processTens(tens, hundreds, group.level, feminine)
	}
	if tens == 0 {
		if hundreds == 200 && group.level > 0 {
			// genitive case - حالة المضاف
			return group_words[0].Genitive
		}
		return small_words[hundreds].Male
	}
	// normal case - الحالة العادية
	return small_words[hundreds].Male + ar_and + processTens(tens, hundreds, group.level, feminine)
}
