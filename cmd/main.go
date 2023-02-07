package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	cfg := Config{}

	replay := &cobra.Command{
		Use: "replay",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Replay(cmd.Context(), cfg)
		},
	}

	root := &cobra.Command{
		Use: "appwatcher",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(cfg)
		},
	}
	root.Flags().BoolVarP(&cfg.NewOnly, "new-only", "n", false, "will filter to only show new applications")
	root.PersistentFlags().StringVar(&cfg.UseTStoragePath, "tstorage-path", "", "stores time series data using tstorage format.  must be a directory.")
	root.AddCommand(replay)

	if err := root.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
