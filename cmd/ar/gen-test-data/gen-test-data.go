package main

import (
	"compress/gzip"
	"math/big"
	"os"

	"github.com/ilius/num2words/arabic"
)

// my select of prime numbers: 7, 71, 719, 7121, 71171, 711121, 7113221

func main() {
	file, err := os.Create("test-data.gz")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	gzw := gzip.NewWriter(file)
	defer gzw.Close()

	add := func(bn *big.Int) {
		_, err := gzw.Write([]byte(bn.String() + "\t" + arabic.ConvertBigInt(bn) + "\n"))
		if err != nil {
			panic(err)
		}
	}

	for n := range 100 {
		add(big.NewInt(int64(n)))
	}
	for n := int64(100); n < 1000; n += 7 {
		add(big.NewInt(n))
	}
	for n := int64(1000); n < 10_000; n += 71 {
		add(big.NewInt(n))
	}
	for n := int64(10_000); n < 100_000; n += 719 {
		add(big.NewInt(n))
	}
	for n := int64(100_000); n < 1_000_00; n += 7_121 {
		add(big.NewInt(n))
	}
	for n := int64(1_000_00); n < 10_000_000; n += 71_171 {
		add(big.NewInt(n))
	}
	for n := int64(10_000_000); n < 100_000_000; n += 711_121 {
		add(big.NewInt(n))
	}
	for n := int64(100_000_00); n < 1_000_000_000; n += 7_113_221 {
		add(big.NewInt(n))
	}
	for n := int64(10_000_000); n < 10_100_000; n += 71 {
		add(big.NewInt(n))
	}
	for n := int64(1000); n < 10_000; n += 71 {
		bn := big.NewInt(n)
		bn.Mul(bn, big.NewInt(1_001_001))
		add(bn)
	}
	for n := int64(1000); n < 10_000; n += 71 {
		bn := big.NewInt(n)
		bn.Mul(bn, big.NewInt(1_001_001_001))
		add(bn)
	}
	for n := int64(1000); n < 10_000; n += 71 {
		bn := big.NewInt(n)
		bn.Mul(bn, big.NewInt(1_001_001_001_001))
		add(bn)
	}
	for n := int64(1000); n < 10_000; n += 71 {
		bn := big.NewInt(n)
		bn.Mul(bn, big.NewInt(1_001_001_001_001_001))
		add(bn)
	}
}
