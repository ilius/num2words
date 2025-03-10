package tajik

// Copyright @ 2024 Saeed Rasooli <saeed.gnu@gmail.com> (ilius)
//
// This library is free software; you can redistribute it and/or
// modify it under the terms of the GNU Lesser General Public
// License as published by the Free Software Foundation; either
// version 2.1 of the License, or (at your option) any later version.
//
// This library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
// Lesser General Public License for more details.

import (
	"math/big"
	"strconv"
	"strings"
)

var (
	big_zero     = big.NewInt(0)
	big_one      = big.NewInt(1)
	big_ten      = big.NewInt(10)
	big_thausand = big.NewInt(1000)
)

const (
	zwnj     = " "
	tg_and   = "у "
	tg_zero  = "Нол"
	tg_first = "Аввал" // or "Як??"
	tg_tenth = "Даҳум"
)

var small_words = map[uint16]string{
	0:   tg_zero,
	1:   "Як",
	2:   "ду",
	3:   "Се",
	4:   "Чаҳор",
	5:   "Пань",
	6:   "Шаш",
	7:   "Ҳафт",
	8:   "Ҳашт",
	9:   "Ну",
	10:  "даҳ",
	11:  "Ёздаҳ",
	12:  "дувоздаҳ",
	13:  "Сенздаҳ",
	14:  "Чаҳорум",
	15:  "Понздаҳ",
	16:  "Шонздаҳ",
	17:  "Ҳабдаҳ",
	18:  "Ҳаждаҳ",
	19:  "Нуздаҳ",
	20:  "Бист",
	30:  "Сӣ",
	40:  "Чил",
	50:  "Панҷоҳ",
	60:  "Шаст",
	70:  "Ҳафтод",
	80:  "Ҳаштод",
	90:  "Навад",
	100: "Сад",
	200: "Дусад",
	300: "Сесад",
	500: "Панҷсад",
}

var big_words_first = []string{"Як", "Ҳазор", "Миллион"}

// European
// var big_words_europe = append(
// 	big_words_first,
// 	"Миллиард", // Milliard
// 	"بیلیون",
// 	"بیلیارد",
// 	"Триллион", // Trillion
// 	"تریلیارد",
// )

// American
// var big_words_US = append(
// 	big_words_first,
// 	"بیلیون",
// 	"Триллион", // Trillion
// 	"کوآدریلیون",
// 	"کوینتیلیون",
// 	"سکستیلیون",
// )

// Common in Iran (the rest are uncommon or mistaken)
var big_words = append(
	big_words_first,
	"Миллиард", // Milliard
	"Триллион", // Trillion
)

func extractGroupsByString(numStr string) ([]uint16, error) {
	digitCount := len(numStr)
	groupCount := digitCount / 3
	groups := make([]uint16, groupCount)
	for i := range groupCount {
		p_int, err := strconv.ParseUint(numStr[digitCount-3*i-3:digitCount-3*i], 10, 64)
		if err != nil {
			return nil, err
		}
		groups[i] = uint16(p_int)
	}
	m := digitCount % 3
	if m > 0 {
		p_int, err := strconv.ParseUint(numStr[:m], 10, 64)
		if err != nil {
			return nil, err
		}
		groups = append(groups, uint16(p_int))
	}
	return groups, nil
}

func bigIntCountDigits(bnBytes []byte) int {
	bn := &big.Int{}
	bn.SetBytes(bnBytes)
	if bn.Cmp(big_zero) == 0 {
		return 1
	}
	count := 0
	for bn.Cmp(big_zero) != 0 {
		bn.Div(bn, big_ten)
		count++
	}
	return count
}

func extractGroupsByBigInt(bn *big.Int, digitCount int) []uint16 {
	groupCount := digitCount / 3
	groups := make([]uint16, groupCount)
	for i := range groupCount {
		mod := &big.Int{}
		div := &big.Int{}
		div.DivMod(bn, big_thausand, mod)
		groups[i] = uint16(mod.Uint64())
		bn = div
	}
	m := digitCount % 3
	if m > 0 {
		groups = append(groups, uint16(bn.Uint64()))
	}
	return groups
}

func joinReversed(groups []string, sep string) string {
	r_groups := make([]string, len(groups))
	n := len(groups)
	for i := range n {
		r_groups[n-i-1] = groups[i]
	}
	return strings.Join(r_groups, sep)
}

// n >= 1000
func convertLarge(groups []uint16) string {
	k := len(groups)
	w_groups := []string{}
	for i := range k {
		p := groups[i]
		if p == 0 {
			continue
		}
		if i == 0 {
			w_groups = append(w_groups, convertSmall(p))
			continue
		}
		order := ""
		if i < len(big_words) {
			order = big_words[i]
		} else {
			order = ""
			d := i / 3
			m := i % 3
			t9 := big_words[3]
			for j := range d {
				if j > 0 {
					order += zwnj
				}
				order += t9
			}
			if m != 0 {
				if order != "" {
					order = zwnj + order
				}
				order = big_words[m] + order
			}
		}
		var w_group string
		if i == 1 && p == 1 {
			w_group = order
		} else {
			w_group = convertSmall(p) + " " + order
		}
		w_groups = append(w_groups, w_group)
	}
	return joinReversed(w_groups, tg_and)
}

// num < 1000
func convertSmall(num uint16) string {
	{
		word, ok := small_words[num]
		if ok {
			return word
		}
	}
	ones := num % 10
	tens := (num % 100) / 10
	hundreds := num / 100
	result := ""
	if hundreds != 0 {
		word, ok := small_words[hundreds*100]
		if ok {
			result += word
		} else {
			result += small_words[hundreds] + small_words[100]
		}
		if tens != 0 || ones != 0 {
			result += tg_and
		}
	}
	if tens != 0 {
		word, ok := small_words[num%100]
		if ok {
			result += word
			return result // OK, Done
		}
		result += small_words[tens*10]
		if ones != 0 {
			result += tg_and
		}
	}
	if ones != 0 {
		result += small_words[ones]
	}
	return result
}

// ConvertString: only for non-negative integers
func ConvertString(str string) (string, error) {
	if len(str) <= 3 { // n <= 999
		n_i64, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return "", err
		}
		return convertSmall(uint16(n_i64)), nil
	}
	// n >= 1000
	groups, err := extractGroupsByString(str)
	if err != nil {
		return "", err
	}
	return convertLarge(groups), nil
}

// ConvertBigInt: only for non-negative integers
func ConvertBigInt(bn *big.Int) string {
	digitCount := bigIntCountDigits(bn.Bytes())
	if digitCount <= 3 { // n <= 999
		return convertSmall(uint16(bn.Uint64()))
	}
	// n >= 1000
	return convertLarge(extractGroupsByBigInt(bn, digitCount))
}

func ConvertBigIntSigned(bn *big.Int) string {
	if bn.Cmp(big_zero) < 0 {
		return "Манфӣ " + ConvertBigInt(bn.Abs(bn))
	}
	return ConvertBigInt(bn)
}

func addOrdinalSuffix(result string) string {
	if strings.HasSuffix(result, "ӣ") {
		return result[:len(result)-1] + "юм"
	}
	if strings.HasSuffix(result, "се") || strings.HasSuffix(result, "Се") {
		resultRunes := []rune(result)
		return string(resultRunes[:len(resultRunes)-1]) + "вум"
	}
	return result + "юм"
}

func ConvertOrdinalString(str string) (string, error) {
	if str == "1" {
		return tg_first, nil
	}
	if str == "10" {
		return tg_tenth, nil
	}
	result, err := ConvertString(str)
	if err != nil {
		return "", err
	}
	return addOrdinalSuffix(result), nil
}

func ConvertOrdinalBigInt(bn *big.Int) string {
	if bn.Cmp(big_one) == 0 {
		return tg_first
	}
	if bn.Cmp(big_ten) == 0 {
		return tg_tenth
	}
	result := ConvertBigInt(bn)
	return addOrdinalSuffix(result)
}
