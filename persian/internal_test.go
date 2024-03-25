package persian

import (
	"log"
	"math/big"
	"testing"

	"github.com/ilius/is/v2"
)

func Test_split3(t *testing.T) {
	is := is.New(t).Lax()

	test := func(str string, parts []uint16) {
		actual_parts, err := splitGroups(str)
		if err != nil {
			log.Fatal(err)
		}
		is.Equal(actual_parts, parts)
	}

	test("1", []uint16{1})
	test("12", []uint16{12})
	test("123", []uint16{123})
	test("1234", []uint16{234, 1})
	test("12345", []uint16{345, 12})
	test("123456", []uint16{456, 123})
	test("1234567", []uint16{567, 234, 1})
}

func Test_bigIntCountDigits(t *testing.T) {
	is := is.New(t).Lax()
	test := func(str string, count int) {
		bn := &big.Int{}
		_, ok := bn.SetString(str, 10)
		if !ok {
			log.Fatalf("failed to parse %v as big int", str)
		}
		actual := bigIntCountDigits(bn.Bytes())
		is.Equal(actual, count)
		is.Equal(bn.String(), str)
	}

	test("1", 1)
	test("123", 3)
	test("1234", 4)
	test("12345", 5)
	test("123456", 6)
	test("1234567", 7)
}

func Test_split3BigInt(t *testing.T) {
	is := is.New(t).Lax()

	test := func(str string, parts []uint16) {
		bn := &big.Int{}
		_, ok := bn.SetString(str, 10)
		if !ok {
			log.Fatalf("failed to parse %v as big int", str)
		}
		// log.Println(bn.Int64())
		actual_parts := splitGroupsBigInt(bn, bigIntCountDigits(bn.Bytes()))
		is.Equal(actual_parts, parts)
	}

	// test("1", []uint16{1})
	// test("123", []uint16{123})
	// test("1234", []uint16{234, 1})
	// test("12345", []uint16{345, 12})
	// test("123456", []uint16{456, 123})
	test("1234567", []uint16{567, 234, 1})
}
