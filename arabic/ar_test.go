package arabic_test

import (
	"strings"
	"testing"

	"github.com/ilius/is/v2"
	"github.com/ilius/num2words/arabic"
)

func TestConvertString(t *testing.T) {
	is := is.New(t).Lax()
	test := func(num_str string, words []string) {
		actual_words := strings.Split(arabic.ConvertString(num_str), " ")
		is.Equal(actual_words, words)
	}
	words := strings.Split(strings.ReplaceAll(`تسعة سكستيليونات 
و ثمانمائة و اثنان و سبعون كوينتليوناً 
و ستمائة و سبعة و سبعون كوادريليوناً 
و ثمانمائة و تسعة و عشرون تريليوناً و ستمائة 
و أربعة و خمسون ملياراً 
و سبعمائة و أربعة و سبعون مليوناً و خمسمائة 
و خمسة و ثمانون ألفاً و مئتان و تسعة و ستون`, "\n", ""), " ")
	test(
		"9872677829654774585269",
		words,
	)
}
