package main

import "C"
import "sync"

var (
	change sync.RWMutex
	store        = map[C.int]any{}
	last   C.int = 0
)

func CGORef[T any](to *T) C.int {
	change.Lock()
	defer change.Unlock()

	id := last
	last++
	store[id] = to
	return id
}

func CGoDeref[T any](id C.int) (*T, bool) {
	if value, has := store[id]; has {
		out, can := value.(*T)
		return out, can
	} else {
		return nil, false
	}
}

func CGoUnref(id C.int) {
	change.Lock()
	defer change.Unlock()

	delete(store, id)
}
