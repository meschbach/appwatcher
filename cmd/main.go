package main

import (
	"fmt"
	"github.com/meschbach/appwatcher/pkg/foundation"
)

func main() {
	absoluteURL := FrontmostApplication().BundleURL().FileSystemPath()
	fmt.Printf("App URL %s\n", absoluteURL)

	message, done := SharedWorkspace().NotificationCenter().workspaceDidActivateApplication()
	defer done()
	go func() {
		for msg := range message {
			fmt.Printf("App bundle: %s (%s)\n", msg.BundleURL().FileSystemPath(), msg.BundleIdentifier().Internalize())
		}
	}()
	foundation.RunMainLoop()
}
