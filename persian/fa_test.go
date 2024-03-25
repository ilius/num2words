package persian_test

import (
	"bufio"
	"compress/gzip"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/ilius/num2words/persian"
)

func TestGeneratedData(t *testing.T) {
	file, err := os.Open("test-data.gz")
	if err != nil {
		panic(err)
	}
	zfile, err := gzip.NewReader(file)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(zfile)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "\t", 2)
		if len(parts) != 2 {
			log.Fatalf("bad line: %v", line)
		}
		num_str := parts[0]
		words_expected := parts[1]
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
		// log.Println(num_str)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
