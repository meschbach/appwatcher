package bridge

import "C"
import "sync"

// Ref wraps uint to provide access to a specific object
// ASSUMPTION: uint corresponds to uint within C
type Ref uint

var (
	change sync.RWMutex
	store      = map[Ref]any{}
	last   Ref = 0
)

func CGORef[T any](to *T) Ref {
	change.Lock()
	defer change.Unlock()

	id := last
	last++
	store[id] = to
	return id
}

func CGoDeref[T any](id Ref) (*T, bool) {
	if value, has := store[id]; has {
		out, can := value.(*T)
		return out, can
	} else {
		return nil, false
	}
}

func CGoUnref(id Ref) {
	change.Lock()
	defer change.Unlock()

	delete(store, id)
}
