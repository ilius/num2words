package arabic_test

import (
	"bufio"
	"compress/gzip"
	"log"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/ilius/is/v2"
	"github.com/ilius/num2words/arabic"
)

var testData = loadTestData()

func loadTestData() map[string]string {
	file, err := os.Open("test-data.gz")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	zfile, err := gzip.NewReader(file)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(zfile)
	data := map[string]string{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "\t", 2)
		if len(parts) != 2 {
			log.Fatalf("bad line: %v", line)
		}
		num_str := parts[0]
		words_expected := parts[1]
		data[num_str] = words_expected
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

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

func TestConvertBigInt(t *testing.T) {
	is := is.New(t).Lax()
	for num_str, words_expected := range testData {
		bn := &big.Int{}
		bn.SetString(num_str, 10)
		is.Equal(arabic.ConvertBigInt(bn), words_expected)
	}
}

func TestConvertBigIntTiny(t *testing.T) {
	is := is.New(t)
	is.Equal(arabic.ConvertBigInt(big.NewInt(1)), "واحد")
	is.Equal(arabic.ConvertBigInt(big.NewInt(2)), "اثنان")
}
