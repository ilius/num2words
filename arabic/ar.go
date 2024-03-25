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

func ConvertBigInt(number *big.Int) string {
	return convertBigInt(number, false)
}

func convertBigInt(number *big.Int, feminine bool) string {
	return strings.TrimSpace(convertToArabic(number, feminine))
}

func ConvertString(number string) string {
	num_big := &big.Int{}
	num_big.SetString(number, 10)
	return convertBigInt(num_big, false)
}

func getDigitFeminineStatus(digit int, groupLevel int, feminine bool) string {
	if groupLevel == -1 || groupLevel == 0 {
		if !feminine {
			return arabicOnes[digit]
		}
		return arabicFeminineOnes[digit]
	}
	return arabicOnes[digit]
}

func convertToArabic(number *big.Int, feminine bool) string {
	if number.Cmp(big_0) == 0 {
		return ar_zero
	}

	tempNumber := &big.Int{}
	tempNumber.SetBytes(number.Bytes())

	result := ""
	group := 0

	for tempNumber.Cmp(big_1) >= 0 {

		// separate number into groups
		mod := &big.Int{}
		tempNumber.DivMod(tempNumber, big_1000, mod)
		numberToProcess := int(mod.Int64())

		// convert group into its text
		tempValue := int(tempNumber.Int64())
		groupDescription := processArabicGroup(numberToProcess, group, tempValue, feminine)

		// here we add the new converted group to the previous concatenated text
		if groupDescription != "" {
			if group > 0 {
				if len(result) > 0 {
					// FIXME: huh??
					result = "و " + result
				}
				if numberToProcess != 2 && numberToProcess%100 != 1 {
					if numberToProcess >= 3 && numberToProcess <= 10 {
						// for numbers between 3 and 9 we use plural name
						result = arabicPluralGroups[group] + " " + result
					} else {
						if len(result) > 0 {
							// use appending case
							result = arabicAppendedGroup[group] + " " + result
						} else {
							// use normal case
							result = arabicGroup[group] + " " + result
						}
					}
				}
			}
			result = groupDescription + " " + result
		}

		group++
	}

	return result
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
