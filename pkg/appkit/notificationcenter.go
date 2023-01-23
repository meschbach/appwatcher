package appkit

import "C"
import (
	"fmt"
	"github.com/meschbach/appwatcher/pkg/bridge"
	"runtime"
	"sync"
	"unsafe"
)

/*
#cgo CFLAGS: -x objective-c -Wincompatible-pointer-types
#cgo LDFLAGS: -framework Foundation -framework Appkit
#import <Foundation/Foundation.h>
#import <Appkit/Appkit.h>

typedef struct wrappedNotification {
	NSNotification* notification;
	NSAutoreleasePool *pool;
} notification;

typedef NSNotificationCenter* notificationCenter;
void NotificationTrampoline(uint, uint, notification);

static void* trampoline(NSNotificationCenter* center, uint dispatcher, uint target){
	return [center addObserverForName:NSWorkspaceDidActivateApplicationNotification object:nil
    		queue:nil usingBlock:^(NSNotification *note) {
		NSAutoreleasePool *pool = [[NSAutoreleasePool alloc] init];
		notification n ;
		n.pool = pool;
		n.notification = [note retain];
		NotificationTrampoline(dispatcher, target, n);
    }];
}

static void trampolineCleanup(notification n){
	[n.pool release];
}

typedef NSDictionary* objcDictionary;

static objcDictionary notification_userInfo(notification notice){
	return [notice.notification userInfo];
}

typedef NSRunningApplication* runningApplication;
static runningApplication dictionary_valueAsRunningApplication(objcDictionary bundle){
//https://developer.apple.com/documentation/appkit/nsworkspacedidactivateapplicationnotification
	return bundle[NSWorkspaceApplicationKey];
}
*/
import "C"

//export NotificationTrampoline
func NotificationTrampoline(dispatcher C.uint, target C.uint, notice C.notification) {
	nc, ok := bridge.CGoDeref[NotificationCenter](bridge.Ref(dispatcher))
	if !ok {
		//TODO: better errors?
		panic("Expected notification center, got none")
	}
	nc.dispatch(target, notice)
}

type NotificationCenter struct {
	object    *C.struct_NSNotificationCenter
	lock      sync.RWMutex
	consumers map[C.uint]chan *RunningApplication
	nextID    C.uint
}

func (n *NotificationCenter) WorkspaceDidActivateApplication() (chan *RunningApplication, func()) {
	n.lock.Lock()
	defer n.lock.Unlock()

	id := n.nextID
	n.nextID++

	out := make(chan *RunningApplication)
	n.consumers[id] = out
	ref := bridge.CGORef(n)
	fmt.Printf("Registering\n")
	C.trampoline(n.object, C.uint(ref), id)
	return out, func() {
		n.lock.Lock()
		defer n.lock.Unlock()
		delete(n.consumers, id)

		close(out)
	}
}

func (n *NotificationCenter) dispatch(target C.uint, notice C.notification) {
	n.lock.RLock()
	defer n.lock.RUnlock()

	if out, has := n.consumers[target]; has {
		userInfo := C.notification_userInfo(notice)
		running := C.dictionary_valueAsRunningApplication(userInfo)
		finalizeNotification := func(app *RunningApplication) {
			C.trampolineCleanup(notice)
		}
		app := ImportRunningApplication(unsafe.Pointer(running))
		runtime.SetFinalizer(app, finalizeNotification)
		out <- app
	}
}

func ImportNotificationCenter(obj unsafe.Pointer) *NotificationCenter {
	return &NotificationCenter{
		object:    (*C.struct_NSNotificationCenter)(obj),
		lock:      sync.RWMutex{},
		consumers: make(map[C.uint]chan *RunningApplication),
		nextID:    0,
	}
}
