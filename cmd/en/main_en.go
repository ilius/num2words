package main

import (
	"fmt"
	"os"

	"github.com/ilius/num2words/english"
)

func main() {
	for _, arg := range os.Args[1:] {
		words, err := english.ConvertString(arg)
		if err != nil {
			panic(err)
		}
		fmt.Println(arg)
		fmt.Println(words)
		fmt.Println()
	}
}
