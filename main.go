package main

import (
	"github.com/spf13/cobra"

	"github.com/vseinstrumentiru/combiner/cmd"
)

func main() {
	root := &cobra.Command{}

	root.AddCommand(
		cmd.Combine(),
	)

	if err := root.Execute(); err != nil {
		println(err.Error())
	}
}
