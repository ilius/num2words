package english

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

const (
	en_and     = ", "
	en_zero    = "Zero"
	en_hundred = "Hundred"
)

var (
	big_zero     = big.NewInt(0)
	big_ten      = big.NewInt(10)
	big_thausand = big.NewInt(1000)
)

var small_words = map[uint16]string{
	0:  en_zero,
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
	11: "Eleven",
	12: "Twelve",
	13: "Thirteen",
	14: "Fourteen",
	15: "Fifteen",
	16: "Sixteen",
	17: "Seventeen",
	18: "Eighteen",
	19: "Nineteen",
	20: "Twenty",
	30: "Thirty",
	40: "Forty",
	50: "Fifty",
	60: "Sixty",
	70: "Seventy",
	80: "Eighty",
	90: "Ninety",
}

var big_words = []string{
	"One",
	"Thousand",
	"Million",
	"Billion",
}

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
					order += " "
				}
				order += t9
			}
			if m != 0 {
				if order != "" {
					order = " " + order
				}
				order = big_words[m] + order
			}
		}
		w_group := convertSmall(p) + " " + order
		w_groups = append(w_groups, w_group)
	}
	return joinReversed(w_groups, en_and)
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
		result += small_words[hundreds] + " " + en_hundred
		if tens != 0 || ones != 0 {
			result += " "
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
			result += " "
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
		return "Negative " + ConvertBigInt(bn.Abs(bn))
	}
	return ConvertBigInt(bn)
}
