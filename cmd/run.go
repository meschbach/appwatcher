package main

import (
	"fmt"
	"github.com/meschbach/appwatcher/pkg/appkit"
	"github.com/meschbach/appwatcher/pkg/foundation"
)

// Run watches applications for changes and outputs the results to the console
// TODO: Would really love to push in a context and exit when it closes
func Run() error {
	absoluteURL := appkit.FrontmostApplication().BundleURL().FileSystemPath()
	fmt.Printf("App URL %s\n", absoluteURL)

	message, done := appkit.SharedWorkspace().NotificationCenter().WorkspaceDidActivateApplication()
	defer done()
	go func() {
		for msg := range message {
			fmt.Printf("App bundle: %s (%s)\n", msg.BundleURL().FileSystemPath(), msg.BundleIdentifier().Internalize())
		}
	}()
	foundation.RunMainLoop()
	return nil
}
