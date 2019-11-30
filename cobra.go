// Cobra is a very popular package for writing cmd line utilities in golang. Here is a simple example
// illustrating how it can be done

package main

import (
	"os"
	"time"

	"github.com/spf13/cobra"
)

func printCommand() *cobra.Command { // this pretty much the subcommand
	return &cobra.Command{
		Use: "curtime",
		RunE: func(cmd *cobra.Command, args []string) error { // RunE essentially is the function that is going to be executed when this subCommand is invoked, E because in case of an error, we can return the error and cobra would know how to percolate this error upto the command
			now := time.Now()
			prettyTime := now.Format(time.RubyDate)
			cmd.Println("helloworld ", prettyTime)
			return nil
		},
	}
}

func main() {
	cmd := &cobra.Command{
		Use:          "Hellloworld",
		Short:        "Hello world",
		SilenceUsage: true,
	}

	// AddCommand is basically like adding a subcommand, printCommand here is basically a subCommand
	cmd.AddCommand(printCommand())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

}

// in order to run this simply do a go build cobra.go and run the executable ./cobra
