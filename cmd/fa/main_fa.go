package main

import (
	"fmt"
	"os"

	"github.com/ilius/num2words/persian"
)

func main() {
	for _, arg := range os.Args[1:] {
		// num, err := strconv.ParseInt(arg)
		// if err != nil {
		// 	panic(err)
		// }
		words := persian.Convert(arg)
		fmt.Printf("%s\t%s\n", arg, words)
	}
}
