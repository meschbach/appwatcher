package main

import (
	"context"
	"fmt"
	"github.com/meschbach/appwatcher/pkg/appkit"
	"github.com/meschbach/appwatcher/pkg/foundation"
	"github.com/meschbach/appwatcher/pkg/junk"
	"github.com/meschbach/appwatcher/pkg/storage"
	"golang.org/x/exp/slices"
)

type Config struct {
	NewOnly bool
	//UseTStoragePath will store the data capture using the tstorage backend backed by the given directory if not empty
	UseTStoragePath string
}

// Run watches applications for changes and outputs the results to the console
// TODO: Would really love to push in a context and exit when it closes
func Run(c Config) error {
	//handle OS signals
	processContext, processDone := context.WithCancel(context.Background())
	defer processDone()

	parallelSinks := junk.NewTee[*appkit.RunningApplication]()
	teeContext, shutdownTee := context.WithCancel(processContext)
	go func() {
		<-processContext.Done()
		shutdownTee()
	}()

	if len(c.UseTStoragePath) > 0 {
		storageMechanism, err := storage.NewTStorage(teeContext, c.UseTStoragePath)
		if err != nil {
			return err
		}

		storageSink := make(chan *appkit.RunningApplication, 16)
		defer close(storageSink)
		parallelSinks.Add(storageSink)
		go storageMechanism.Consume(teeContext, storageSink)
	}

	stdoutSink := make(chan *appkit.RunningApplication)
	go func() {
		<-processContext.Done()
		close(stdoutSink)
	}()
	parallelSinks.Add(stdoutSink)
	go func() {
		for msg := range stdoutSink {
			fmt.Printf("App bundle: %s (%s)\n", msg.BundleURL().FileSystemPath(), msg.BundleIdentifier().Internalize())
		}
	}()

	absoluteURL := appkit.FrontmostApplication().BundleURL().FileSystemPath()
	fmt.Printf("App URL %s\n", absoluteURL)

	message, done := appkit.SharedWorkspace().NotificationCenter().WorkspaceDidActivateApplication()
	go func() {
		<-processContext.Done()
		fmt.Printf("Done processing events.")
		done()
	}()

	if c.NewOnly {
		fmt.Printf("[II] Filtering for new applications only.\n")
		message = newOnlyFilter(message)
	}

	//pump the NS system
	go parallelSinks.PumpNonblocking(teeContext, message)
	foundation.RunMainLoop()
	fmt.Printf("NSMainLoop returned.  Shutting down.\n")
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
