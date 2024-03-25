package persian

import (
	"math/big"
	"strconv"
	"strings"
)

var (
	big_zero     = big.NewInt(0)
	big_ten      = big.NewInt(10)
	big_thausand = big.NewInt(1000)
)

const (
	fa_and = " و "
	zwnj   = "\u200c"
)

var faBaseNum = map[int]string{
	1:   "یک",
	2:   "دو",
	3:   "سه",
	4:   "چهار",
	5:   "پنج",
	6:   "شش",
	7:   "هفت",
	8:   "هشت",
	9:   "نه",
	10:  "ده",
	11:  "یازده",
	12:  "دوازده",
	13:  "سیزده",
	14:  "چهارده",
	15:  "پانزده",
	16:  "شانزده",
	17:  "هفده",
	18:  "هجده",
	19:  "نوزده",
	20:  "بیست",
	30:  "سی",
	40:  "چهل",
	50:  "پنجاه",
	60:  "شصت",
	70:  "هفتاد",
	80:  "هشتاد",
	90:  "نود",
	100: "صد",
	200: "دویست",
	300: "سیصد",
	500: "پانصد",
}

var faBigNumFirst = []string{"یک", "هزار", "میلیون"}

// European
// var faBigNumEU = append(
// 	faBigNumFirst,
// 	"میلیارد",
// 	"بیلیون",
// 	"بیلیارد",
// 	"تریلیون",
// 	"تریلیارد",
// )

// American
// var faBigNumUS = append(
// 	faBigNumFirst,
// 	"بیلیون",
// 	"تریلیون",
// 	"کوآدریلیون",
// 	"کوینتیلیون",
// 	"سکستیلیون",
// )

// Common in Iran (the rest are uncommon or mistaken)
var faBigNumIran = append(
	faBigNumFirst,
	"میلیارد",
	"تریلیون",
)

var faBigNum = faBigNumIran

func split3(st string) ([]uint16, error) {
	digitCount := len(st)
	partCount := digitCount / 3
	parts := make([]uint16, partCount)
	for i := range partCount {
		p_int, err := strconv.ParseUint(st[digitCount-3*i-3:digitCount-3*i], 10, 64)
		if err != nil {
			return nil, err
		}
		parts[i] = uint16(p_int)
	}
	m := digitCount % 3
	if m > 0 {
		p_int, err := strconv.ParseUint(st[:m], 10, 64)
		if err != nil {
			return nil, err
		}
		parts = append(parts, uint16(p_int))
	}
	return parts, nil
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

func split3BigInt(bn *big.Int, digitCount int) ([]uint16, error) {
	partCount := digitCount / 3
	parts := make([]uint16, partCount)
	for i := range partCount {
		mod := &big.Int{}
		div := &big.Int{}
		div.DivMod(bn, big_thausand, mod)
		parts[i] = uint16(mod.Uint64())
		bn = div
	}
	m := digitCount % 3
	if m > 0 {
		parts = append(parts, uint16(bn.Uint64()))
	}
	return parts, nil
}

func join_reversed(parts []string, sep string) string {
	r_parts := make([]string, len(parts))
	n := len(parts)
	for i := range n {
		r_parts[n-i-1] = parts[i]
	}
	return strings.Join(r_parts, sep)
}

func convert_int(num uint64) (string, error) {
	return ConvertString(strconv.FormatUint(num, 10))
}

func convertStringLarge(parts []uint16) (string, error) {
	k := len(parts)
	w_parts := []string{}
	for i := range k {
		p := parts[i]
		if p == 0 {
			continue
		}
		if i == 0 {
			w_part, err := convert_int(uint64(p))
			if err != nil {
				return "", err
			}
			w_parts = append(w_parts, w_part)
			continue
		}
		faOrder := ""
		if i < len(faBigNum) {
			faOrder = faBigNum[i]
		} else {
			faOrder = ""
			d := i / 3
			m := i % 3
			t9 := faBigNum[3]
			for j := range d {
				if j > 0 {
					faOrder += zwnj
				}
				faOrder += t9
			}
			if m != 0 {
				if faOrder != "" {
					faOrder = zwnj + faOrder
				}
				faOrder = faBigNum[m] + faOrder
			}
		}
		var w_part string
		if i == 1 && p == 1 {
			w_part = faOrder
		} else {
			w_part_tmp, err := convert_int(uint64(p))
			if err != nil {
				return "", err
			}
			w_part = w_part_tmp + " " + faOrder
		}
		w_parts = append(w_parts, w_part)
	}
	return join_reversed(w_parts, fa_and), nil
}

// n < 1000
func convertStringSmall(n int) (string, error) {
	if _, ok := faBaseNum[n]; ok {
		return faBaseNum[n], nil
	}
	yekan := n % 10
	dahgan := int((n % 100) / 10)
	sadgan := int(n / 100)
	dahgan_yekan := 10*dahgan + yekan
	result := ""
	if sadgan != 0 {
		if _, ok := faBaseNum[sadgan*100]; ok {
			result += faBaseNum[sadgan*100]
		} else {
			result += faBaseNum[sadgan] + faBaseNum[100]
		}
		if dahgan != 0 || yekan != 0 {
			result += fa_and
		}
	}
	if dahgan != 0 {
		if _, ok := faBaseNum[dahgan_yekan]; ok {
			result += faBaseNum[dahgan_yekan]
			return result, nil
		}
		result += faBaseNum[dahgan*10]
		if yekan != 0 {
			result += fa_and
		}
	}
	if yekan != 0 {
		result += faBaseNum[yekan]
	}
	return result, nil
}

func ConvertString(str string) (string, error) {
	if len(str) > 3 {
		parts, err := split3(str)
		if err != nil {
			return "", err
		}
		return convertStringLarge(parts)
	}
	// now assume that n <= 999
	n_i64, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return convertStringSmall(int(n_i64))
}

func ConvertOrdinalString(str string) (string, error) {
	if str == "1" {
		return "اول", nil // or "یکم"
	}
	if str == "10" {
		return "دهم", nil
	}
	norm_fa, err := ConvertString(str)
	if err != nil {
		return "", err
	}
	if strings.HasSuffix(norm_fa, "ی") {
		norm_fa += "\u200cام"
	} else if strings.HasSuffix(norm_fa, "سه") {
		norm_fa = norm_fa[:len(norm_fa)-1] + "وم"
	} else {
		norm_fa += "م"
	}
	return norm_fa, nil
}

func ConvertBigInt(bn *big.Int) (string, error) {
	digitCount := bigIntCountDigits(bn.Bytes())
	if digitCount > 3 {
		parts, err := split3BigInt(bn, digitCount)
		if err != nil {
			return "", err
		}
		return convertStringLarge(parts)
	}
	// now assume that n <= 999
	return convertStringSmall(int(bn.Int64()))
}
