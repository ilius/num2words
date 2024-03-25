package persian_test

import (
	"bufio"
	"compress/gzip"
	"log"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/ilius/is/v2"
	"github.com/ilius/num2words/persian"
)

var testData = loadTestData()

func loadTestData() map[string]string {
	file, err := os.Open("test-data.gz")
	if err != nil {
		panic(err)
	}
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
	for num_str, words_expected := range testData {
		words, err := persian.ConvertString(num_str)
		if err != nil {
			log.Fatal(err)
		}
		if words != words_expected {
			log.Fatalf(
				"test failed: num=%v, expected: %#v, actual: %#v",
				num_str, words_expected, words,
			)
		}
	}
}

func TestConvertBigInt(t *testing.T) {
	is := is.New(t)
	for num_str, words_expected := range testData {
		bn := &big.Int{}
		bn.SetString(num_str, 10)
		is.Equal(persian.ConvertBigInt(bn), words_expected)

		bn_neg := &big.Int{}
		bn_neg.SetString(num_str, 10)
		bn_neg.Neg(bn_neg)
		if num_str == "0" {
			is.Equal(persian.ConvertBigIntSigned(bn_neg), "صفر")
		} else {
			is.Equal(persian.ConvertBigIntSigned(bn_neg), "منفی "+words_expected)
		}
	}
}

func Benchmark_convert_string_bigInt(b *testing.B) {
	b.Run("string", func(b *testing.B) {
		for num_str := range testData {
			persian.ConvertString(num_str)
		}
	})
	b.Run("big.Int", func(b *testing.B) {
		for num_str := range testData {
			bn := &big.Int{}
			bn.SetString(num_str, 10)
			persian.ConvertBigInt(bn)
		}
	})
}

/*
go test -bench=. -benchtime=1000x -count=5
	Benchmark_convert_string_bigInt/string-8         	    1000	      3500 ns/op
	Benchmark_convert_string_bigInt/string-8         	    1000	      3388 ns/op
	Benchmark_convert_string_bigInt/string-8         	    1000	      3383 ns/op
	Benchmark_convert_string_bigInt/string-8         	    1000	      3769 ns/op
	Benchmark_convert_string_bigInt/string-8         	    1000	      3614 ns/op
	Benchmark_convert_string_bigInt/big.Int-8        	    1000	      6050 ns/op
	Benchmark_convert_string_bigInt/big.Int-8        	    1000	      6301 ns/op
	Benchmark_convert_string_bigInt/big.Int-8        	    1000	      7827 ns/op
	Benchmark_convert_string_bigInt/big.Int-8        	    1000	      6571 ns/op
	Benchmark_convert_string_bigInt/big.Int-8        	    1000	      6183 ns/op
*/
