package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	cfg := Config{}
	root := &cobra.Command{
		Use: "appwatcher",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(cfg)
		},
	}
	root.Flags().BoolVarP(&cfg.NewOnly, "new-only", "n", false, "will filter to only show new applications")
	if err := root.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
