package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/meschbach/appwatcher/pkg/appkit"
	"github.com/nakabonne/tstorage"
	"time"
)

type TStorageEngine struct {
	store tstorage.Storage
}

func NewTStorage(ctx context.Context, filePath string) (*TStorageEngine, error) {
	storage, err := tstorage.NewStorage(
		tstorage.WithTimestampPrecision(tstorage.Milliseconds),
		tstorage.WithDataPath(filePath),
		tstorage.WithWALBufferedSize(0),
		tstorage.WithWriteTimeout(1*time.Second),
	)
	if err != nil {
		return nil, err
	}
	return &TStorageEngine{store: storage}, nil
}

func (t *TStorageEngine) Close() {
	//TODO: Complain somewhere if there is a problem
	t.store.Close()
}

func (t *TStorageEngine) Consume(ctx context.Context, activeApplication <-chan *appkit.RunningApplication) {
	for {
		select {
		case msg := <-activeApplication:
			if err := t.storeRecord(ctx, msg); err != nil {
				fmt.Printf("[WWW] tstorage: error while writing data point: %e\n", err)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (t *TStorageEngine) storeRecord(ctx context.Context, msg *appkit.RunningApplication) error {
	labels := []tstorage.Label{
		{Name: "bundleIdentifier", Value: msg.BundleIdentifier().Internalize()},
		{Name: "bundleURL", Value: msg.BundleURL().FileSystemPath()},
	}

	return t.store.InsertRows([]tstorage.Row{
		{
			Metric: "active",
			Labels: labels,
			DataPoint: tstorage.DataPoint{
				Timestamp: time.Now().Unix(),
				Value:     1,
			},
		},
	})
}

func (t *TStorageEngine) ReplayAll(ctx context.Context) ([]int, error) {
	_, err := t.store.Select("active", nil, 0, time.Now().Unix())
	if err != nil {
		if errors.Is(err, tstorage.ErrNoDataPoints) {
			return nil, nil
		}
		return nil, err
	}
	return nil, nil
}
