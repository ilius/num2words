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
	for _, tc := range testData {
		words, err := persian.ConvertString(tc.String)
		if err != nil {
			log.Fatal(err)
		}
		if words != tc.Words {
			log.Fatalf(
				"test failed: num=%v, expected: %#v, actual: %#v",
				tc.String, tc.Words, words,
			)
		}
	}
}

func TestConvertBigInt(t *testing.T) {
	is := is.New(t)
	for _, tc := range testData {
		is.Equal(persian.ConvertBigInt(tc.BigInt), tc.Words)
	}
}

var big_zero = big.NewInt(0)

func TestConvertBigIntSigned(t *testing.T) {
	is := is.New(t)
	for _, tc := range testData {
		bn_neg := &big.Int{}
		bn_neg.SetBytes(tc.BigInt.Bytes())
		bn_neg.Neg(bn_neg)
		if tc.BigInt.Cmp(big_zero) == 0 {
			is.Equal(persian.ConvertBigIntSigned(bn_neg), "صفر")
		} else {
			is.Equal(persian.ConvertBigIntSigned(bn_neg), "منفی "+tc.Words)
		}
	}
}

func Benchmark_convert_string_bigInt(b *testing.B) {
	b.Run("string", func(b *testing.B) {
		for _, tc := range testData {
			_, _ = persian.ConvertString(tc.String)
		}
	})
	b.Run("big.Int", func(b *testing.B) {
		for _, tc := range testData {
			_ = persian.ConvertBigInt(tc.BigInt)
		}
	})
}

/*
go test -bench=. -benchtime=1000x -count=5
	Benchmark_convert_string_bigInt/string-8         	    1000	      3152 ns/op
	Benchmark_convert_string_bigInt/string-8         	    1000	      3158 ns/op
	Benchmark_convert_string_bigInt/string-8         	    1000	      3205 ns/op
	Benchmark_convert_string_bigInt/string-8         	    1000	      3135 ns/op
	Benchmark_convert_string_bigInt/string-8         	    1000	      3954 ns/op
	Benchmark_convert_string_bigInt/big.Int-8        	    1000	      5158 ns/op
	Benchmark_convert_string_bigInt/big.Int-8        	    1000	      5272 ns/op
	Benchmark_convert_string_bigInt/big.Int-8        	    1000	      5864 ns/op
	Benchmark_convert_string_bigInt/big.Int-8        	    1000	      5045 ns/op
	Benchmark_convert_string_bigInt/big.Int-8        	    1000	      5496 ns/op
*/
