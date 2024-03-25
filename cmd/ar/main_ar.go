package main

import (
	"fmt"
	"os"

	"github.com/ilius/num2words/arabic"
)

func main() {
	for _, arg := range os.Args[1:] {
		words := arabic.ConvertString(arg)
		// if err != nil {
		// 	panic(err)
		// }
		fmt.Println(arg)
		fmt.Println(words)
		// fmt.Println(arabic.ConvertOrdinalString(arg))
		fmt.Println()
	}
}
