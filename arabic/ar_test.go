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

type TestCase struct {
	String string
	BigInt *big.Int
	Words  string
}

func loadTestData() []TestCase {
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
	data := []TestCase{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "\t", 2)
		if len(parts) != 2 {
			log.Fatalf("bad line: %v", line)
		}
		num_str := parts[0]
		words := parts[1]
		bn := &big.Int{}
		bn.SetString(num_str, 10)
		data = append(data, TestCase{
			String: num_str,
			BigInt: bn,
			Words:  words,
		})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func TestConvertString(t *testing.T) {
	is := is.New(t).Lax()
	for _, tc := range testData {
		is.Equal(arabic.ConvertString(tc.String), tc.Words)
	}
}

func TestConvertString2(t *testing.T) {
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
	for _, tc := range testData {
		is.Msg("number=%v", tc.String).Equal(arabic.ConvertBigInt(tc.BigInt), tc.Words)
	}
}

func TestConvertBigIntTiny(t *testing.T) {
	is := is.New(t)
	is.Equal(arabic.ConvertBigInt(big.NewInt(1)), "واحد")
	is.Equal(arabic.ConvertBigInt(big.NewInt(2)), "اثنان")
}
