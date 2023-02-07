package junk

import (
	"context"
)

// Tee duplicates an incoming message to a set of outgoing channels.
type Tee[T any] struct {
	dispatchTo []chan<- T
}

func NewTee[T any]() *Tee[T] {
	return &Tee[T]{}
}

// Add the given target to the set of channels to duplicate a message to.
func (t *Tee[T]) Add(target chan<- T) {
	t.dispatchTo = append(t.dispatchTo, target)
}

// PumpNonblocking will dispatch a message from into all target channels.  If a channel is full or otherwise not receiving
// then the message will be dropped.  Will continue pumping until the given ctx is Done.  No mutations should occur
// while pumping.
func (t *Tee[T]) PumpNonblocking(ctx context.Context, from <-chan T) {
	for {
		select {
		case msg := <-from:
			for _, to := range t.dispatchTo {
				select {
				case to <- msg:
				default:
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
