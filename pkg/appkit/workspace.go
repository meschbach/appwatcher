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

// Workspace wraps the NSWorkspace object.
// See https://developer.apple.com/documentation/appkit/nsworkspace?language=objc
type Workspace struct {
	object *C.NSWorkspace
}

// SharedWorkspace Locates the shared desktop instance.
// See https://developer.apple.com/documentation/appkit/nsworkspace/1530344-sharedworkspace?language=objc
func SharedWorkspace() *Workspace {
	obj := C.nsworkspace_shared()
	return &Workspace{object: obj}
}

func (w *Workspace) NotificationCenter() *NotificationCenter {
	obj := C.nsworkspace_notificationCenter(w.object)
	return ImportNotificationCenter(unsafe.Pointer(obj))
}
