package main

import (
	"fmt"
	"github.com/deweysasser/cumulus/program"
	"github.com/pkg/profile"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {

	var options program.Options

	context, err := options.Parse(os.Args[1:])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch {
	case options.Profile.CPU:
		defer profile.Start(profile.CPUProfile).Stop()
	case options.Profile.Memory:
		defer profile.Start(profile.MemProfile).Stop()
	}

	// This ends up calling options.Run()
	if err := context.Run(options); err != nil {
		log.Err(err).Msg("Program failed")
		os.Exit(1)
	}
}
