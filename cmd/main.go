package main

import (
	"fmt"
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
	run()
}
