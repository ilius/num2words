package persian

import (
	"strconv"
	"strings"
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
// var faBigNumEU = append(faBigNumFirst, "میلیارد", "بیلیون", "بیلیارد", "تریلیون", "تریلیارد")

// American
// var faBigNumUS = append(faBigNumFirst,
// 	"بیلیون",
// 	"تریلیون",
// 	"کوآدریلیون",
// 	"کوینتیلیون",
// 	"سکستیلیون",
// )

// Common in Iran (the rest are uncommon or mistaken)
var faBigNumIran = append(faBigNumFirst, "میلیارد", "تریلیون")

var faBigNum = faBigNumIran

func split3(st string) []int {
	n := len(st)
	d := n / 3
	m := n % 3
	parts := make([]int, d)
	for i := range d {
		p_int, err := strconv.ParseInt(st[n-3*i-3:n-3*i], 10, 64)
		if err != nil {
			panic(err)
		}
		parts[i] = int(p_int)
	}
	if m > 0 {
		p_int, err := strconv.ParseInt(st[:m], 10, 64)
		if err != nil {
			panic(err)
		}
		parts = append(parts, int(p_int))
	}
	return parts
}

func join_reversed(parts []string, sep string) string {
	r_parts := make([]string, len(parts))
	n := len(parts)
	for i := range n {
		r_parts[n-i-1] = parts[i]
	}
	return strings.Join(r_parts, sep)
}

func convert_int(num int) string {
	return ConvertString(strconv.FormatUint(uint64(num), 10))
}

func ConvertString(st string) string {
	if len(st) > 3 {
		parts := split3(st)
		k := len(parts)
		wparts := []string{}
		for i := range k {
			faOrder := ""
			p := parts[i]
			if p == 0 {
				continue
			}
			var wpart string
			if i == 0 {
				wpart = convert_int(p)
			} else {
				if i < len(faBigNum) {
					faOrder = faBigNum[i]
				} else {
					faOrder = ""
					d := i / 3
					m := i % 3
					t9 := faBigNum[3]
					for j := range d {
						if j > 0 {
							faOrder += "\u200c"
						}
						faOrder += t9
					}
					if m != 0 {
						if faOrder != "" {
							faOrder = "\u200c" + faOrder
						}
						faOrder = faBigNum[m] + faOrder
					}
				}
				if i == 1 && p == 1 {
					wpart = faOrder
				} else {
					wpart = convert_int(p) + " " + faOrder
				}
			}
			wparts = append(wparts, wpart)
		}
		return join_reversed(wparts, " و ")
	}
	// now assume that n <= 999
	n_i64, err := strconv.ParseInt(st, 10, 64)
	if err != nil {
		panic(err)
	}
	n := int(n_i64)
	if _, ok := faBaseNum[n]; ok {
		return faBaseNum[n]
	}
	y := n % 10
	d := int((n % 100) / 10)
	s := int(n / 100)
	// print s, d, y
	dy := 10*d + y
	fa := ""
	if s != 0 {
		if _, ok := faBaseNum[s*100]; ok {
			fa += faBaseNum[s*100]
		} else {
			fa += faBaseNum[s] + faBaseNum[100]
		}
		if d != 0 || y != 0 {
			fa += " و "
		}
	}
	if d != 0 {
		if _, ok := faBaseNum[dy]; ok {
			fa += faBaseNum[dy]
			return fa
		}
		fa += faBaseNum[d*10]
		if y != 0 {
			fa += " و "
		}
	}
	if y != 0 {
		fa += faBaseNum[y]
	}
	return fa
}

func ConvertOrdinalString(str string) string {
	if str == "1" {
		return "اول" // or "یکم"
	}
	if str == "10" {
		return "دهم"
	}
	norm_fa := ConvertString(str)
	if norm_fa == "" {
		return ""
	}
	if strings.HasSuffix(norm_fa, "ی") {
		norm_fa += "\u200cام"
	} else if strings.HasSuffix(norm_fa, "سه") {
		norm_fa = norm_fa[:len(norm_fa)-1] + "وم"
	} else {
		norm_fa += "م"
	}
	return norm_fa
}
