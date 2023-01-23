package appkit

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Appkit
#import <Appkit/Appkit.h>

NSWorkspace* nsworkspace_shared() {
	return [NSWorkspace sharedWorkspace];
}

NSNotificationCenter* nsworkspace_notificationCenter(NSWorkspace* workspace) {
	return [workspace notificationCenter];
}
*/
import "C"
import "unsafe"

type Workspace struct {
	object *C.NSWorkspace
}

func SharedWorkspace() *Workspace {
	obj := C.nsworkspace_shared()
	return &Workspace{object: obj}
}

func (w *Workspace) NotificationCenter() unsafe.Pointer {
	obj := C.nsworkspace_notificationCenter(w.object)
	return unsafe.Pointer(obj)
}
