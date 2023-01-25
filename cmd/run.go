package main

import (
	"fmt"
	"github.com/meschbach/appwatcher/pkg/appkit"
	"github.com/meschbach/appwatcher/pkg/foundation"
	"golang.org/x/exp/slices"
)

type Config struct {
	NewOnly bool
}

// Run watches applications for changes and outputs the results to the console
// TODO: Would really love to push in a context and exit when it closes
func Run(c Config) error {
	absoluteURL := appkit.FrontmostApplication().BundleURL().FileSystemPath()
	fmt.Printf("App URL %s\n", absoluteURL)

	message, done := appkit.SharedWorkspace().NotificationCenter().WorkspaceDidActivateApplication()
	defer done()

	if c.NewOnly {
		fmt.Printf("[II] Filtering for new applications only.\n")
		message = newOnlyFilter(message)
	}
	go func() {
		for msg := range message {
			fmt.Printf("App bundle: %s (%s)\n", msg.BundleURL().FileSystemPath(), msg.BundleIdentifier().Internalize())
		}
	}()
	foundation.RunMainLoop()
	return nil
}

func newOnlyFilter(input chan *appkit.RunningApplication) chan *appkit.RunningApplication {
	output := make(chan *appkit.RunningApplication)
	go func() {
		var seen []string
		for msg := range input {
			path := msg.BundleURL().FileSystemPath()
			if slices.Contains(seen, path) {
				continue
			}
			seen = append(seen, path)
			output <- msg
		}
	}()
	return output
}
