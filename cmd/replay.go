package main

import (
	"context"
	"fmt"
	"github.com/meschbach/appwatcher/pkg/storage"
)

func Replay(ctx context.Context, cfg Config) error {
	dataStorage, err := storage.NewTStorage(ctx, cfg.UseTStoragePath)
	if err != nil {
		return err
	}

	messages, err := dataStorage.ReplayAll(ctx)
	if err != nil {
		return err
	}

	for _, m := range messages {
		fmt.Printf("%#v\n", m)
	}

	return nil
}
