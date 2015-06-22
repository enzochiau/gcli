package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
	"github.com/tcnksm/gcli/command"
)

// Run execute RunCustom() with color and output to Stdout/Stderr.
// It returns exit code.
func Run(args []string) int {

	// Meta-option for executables.
	// It defines output color and its stdout/stderr stream.
	meta := &command.Meta{
		UI: &cli.ColoredUi{
			InfoColor:  cli.UiColorBlue,
			ErrorColor: cli.UiColorRed,
			Ui: &cli.BasicUi{
				Writer:      os.Stdout,
				ErrorWriter: os.Stderr,
				Reader:      os.Stdin,
			},
		}}

	return RunCustom(args, Commands(meta))
}

// RunCustom execute mitchellh/cli and return its exit code.
func RunCustom(args []string, commands map[string]cli.CommandFactory) int {

	cli := &cli.CLI{
		Args:       args[1:],
		Commands:   commands,
		Version:    Version,
		HelpFunc:   cli.BasicHelpFunc(Name),
		HelpWriter: os.Stdout,
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute: %s\n", err.Error())
	}

	return exitCode
}
