package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/shiimaxx/image-convert/lib"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (c *CLI) Run(args []string) int {
	var (
		version bool
		srcExt  string
		destExt string
	)

	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(c.outStream)

	flags.StringVar(&srcExt, "src", "", "source extension")
	flags.StringVar(&srcExt, "s", "", "source extension(Short)")
	flags.StringVar(&destExt, "dest", "", "destination extension")
	flags.StringVar(&destExt, "d", "", "destination extension(Short)")
	flags.BoolVar(&version, "version", false, "print version information")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	if version {
		fmt.Fprintf(c.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	if len(flags.Args()) < 1 {
		fmt.Fprintln(c.errStream, "Missing arguments")
		return ExitCodeError
	}

	filePath := flags.Args()[0]
	finfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Fprintf(c.errStream, "%s: No such file or directory\n", filePath)
		return ExitCodeError
	}
	if !finfo.IsDir() {
		fmt.Fprintf(c.errStream, "%s: Is a not directory\n", filePath)
		return ExitCodeError
	}

	imageFiles, err := converter.MakeImageFiles(filePath, srcExt)
	for _, f := range imageFiles {
		err := converter.Convert(f, destExt)
		if err != nil {
			fmt.Fprintln(c.errStream, err)
			return ExitCodeError
		}
	}

	return ExitCodeOK
}
