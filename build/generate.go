package main

import (
	"fmt"
	"github.com/deweysasser/cumulus/code_generation"
	"github.com/rs/zerolog"
	"os"
)

func main() {

	zerolog.SetGlobalLevel(zerolog.WarnLevel)

	pkg := os.Args[1]

	for _, arg := range os.Args[2:] {
		fmt.Println("reading", arg)
		if err := code_generation.GenerateFieldCode(pkg, arg); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(2)
		}
	}
}
