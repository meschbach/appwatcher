package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework Appkit
#import <Cocoa/Cocoa.h>
#import <Foundation/Foundation.h>
#import <Appkit/Appkit.h>

NSWorkspace* nsworkspace_shared() {
	return [NSWorkspace sharedWorkspace];
}

NSNotificationCenter* nsworkspace_notificationCenter(NSWorkspace* workspace) {
	return [workspace notificationCenter];
}
*/
import "C"
import (
	"github.com/meschbach/appwatcher/pkg/appkit"
	"sync"
)

type Workspace struct {
	object *C.NSWorkspace
}

func SharedWorkspace() *Workspace {
	obj := C.nsworkspace_shared()
	return &Workspace{object: obj}
}

func (w *Workspace) NotificationCenter() *NotificationCenter {
	obj := C.nsworkspace_notificationCenter(w.object)
	return &NotificationCenter{
		object:    obj,
		lock:      sync.RWMutex{},
		consumers: make(map[C.int]chan *appkit.RunningApplication),
		nextID:    0,
	}
}
