package launcher

import (
	"bf/internal/config"
	"bf/pkg/brainfuck"
	"fmt"
	"os"
)

func Launch() {
	conf, argErr := config.FromArgs()
	if argErr != nil {
		fmt.Println(argErr)
		os.Exit(1)
	}

	code, fileErr := os.ReadFile(conf.SourceFilename)
	if fileErr != nil {
		fmt.Println(fileErr)
		os.Exit(1)
	}

	bf := brainfuck.Default()
	runErr := bf.Run(string(code))
	if runErr != nil {
		fmt.Println(runErr)
		os.Exit(1)
	}
}
