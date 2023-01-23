package main

/*
#cgo CFLAGS: -x objective-c -Wincompatible-pointer-types
#cgo LDFLAGS: -framework Foundation -framework Appkit
#import <Cocoa/Cocoa.h>
#import <Foundation/Foundation.h>
#import <Appkit/Appkit.h>

typedef NSRunningApplication* runningApplication;
runningApplication nsworkspace_frontmostApplication() {
	return [[NSWorkspace sharedWorkspace] frontmostApplication];
}

NSURL* nsrunningapplication_bundleurl(runningApplication app) {
	return [app bundleURL];
}

NSString* nsrunningapplication_bundleIdentifier(runningApplication app) {
	return [app bundleIdentifier];
}
*/
import "C"
import (
	"github.com/meschbach/appwatcher/pkg/foundation"
	"unsafe"
)

type RunningApplication struct {
	object C.runningApplication
}

func FrontmostApplication() *RunningApplication {
	return &RunningApplication{C.nsworkspace_frontmostApplication()}
}

func (n *RunningApplication) BundleURL() *foundation.NSURL {
	return foundation.ImportURL(unsafe.Pointer(C.nsrunningapplication_bundleurl(n.object)))
}

func (n *RunningApplication) BundleIdentifier() *foundation.String {
	str := C.nsrunningapplication_bundleIdentifier(n.object)
	internal := foundation.WrapString(unsafe.Pointer(str))
	return internal
}
