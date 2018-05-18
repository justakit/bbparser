package main

import (
	"fmt"

	"github.com/KitlerUA/bbparser"
)

func main() {
	tests := []string{}
	parser := bbparser.NewDefault()
	for _, str := range tests {
		fmt.Println(parser.Parse(str))
	}

}
