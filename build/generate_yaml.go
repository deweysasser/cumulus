package main

import (
	"github.com/deweysasser/cumulus/code_generation"
	"github.com/rs/zerolog"
)

func main() {

	zerolog.SetGlobalLevel(zerolog.WarnLevel)

	code_generation.GenerateAllYaml()
}
