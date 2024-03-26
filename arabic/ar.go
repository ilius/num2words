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

// Ones
var arabicOnes = []string{
	"",
	"واحد",
	"اثنان",
	"ثلاثة",
	"أربعة",
	"خمسة",
	"ستة",
	"سبعة",
	"ثمانية",
	"تسعة",
	"عشرة",
	"أحد عشر",
	"اثنا عشر",
	"ثلاثة عشر",
	"أربعة عشر",
	"خمسة عشر",
	"ستة عشر",
	"سبعة عشر",
	"ثمانية عشر",
	"تسعة عشر",
}

var arabicFeminineOnes = []string{
	"",
	"إحدى",
	"اثنتان",
	"ثلاث",
	"أربع",
	"خمس",
	"ست",
	"سبع",
	"ثمان",
	"تسع",
	"عشر",
	"إحدى عشرة",
	"اثنتا عشرة",
	"ثلاث عشرة",
	"أربع عشرة",
	"خمس عشرة",
	"ست عشرة",
	"سبع عشرة",
	"ثماني عشرة",
	"تسع عشرة",
}

// Tens
var arabicTens = []string{
	"عشرون",
	"ثلاثون",
	"أربعون",
	"خمسون",
	"ستون",
	"سبعون",
	"ثمانون",
	"تسعون",
}

// Hundreds
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

// Twos
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

// Appended
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

// Twos

// Group
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

// Appended

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

// Group

// Plural groups
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
		// tempNumber -> remaining
		groupDescription := processArabicGroup(
			groupNumber,
			groupLevel,
			int(number.Int64()),
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
		return arabicFeminineOnes[digit]
	}
	return arabicOnes[digit]
}

func processArabicGroup(groupNumber int, groupLevel int, remaining int, feminine bool) string {
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
		if tens < 20 { // if we are processing under 20 numbers
			if tens == 2 && hundreds == 0 && groupLevel > 0 { // This is special case for number 2 when it comes alone in the group
				result = arabicTwos[groupLevel] // في حالة الافراد
			} else { // General case
				if result != "" {
					result += ar_and
				}

				if tens == 1 && groupLevel > 0 {
					result += arabicGroup[groupLevel]
				} else if (tens == 1 || tens == 2) && (groupLevel == 0 || groupLevel == -1) && hundreds == 0 && remaining == 0 {
					// Special case for 1 and 2 numbers like: ليرة سورية و ليرتان سوريتان
				} else {
					// Get Feminine status for this digit
					result += getDigitFeminineStatus(tens, groupLevel, feminine)
				}
			}
		} else {
			ones := tens % 10
			tens = tens/10 - 2 // 20's offset

			if ones > 0 {
				if result != "" {
					result += ar_and
				}

				// Get Feminine status for this digit
				result += getDigitFeminineStatus(ones, groupLevel, feminine)
			}

			if result != "" {
				result += ar_and
			}

			// Get Tens text
			result += arabicTens[tens]
		}
	}

	return result
}
