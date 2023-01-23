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

typedef NSString* objcString;

objcString nsrunningapplication_bundleIdentifier(runningApplication app) {
	return [app bundleIdentifier];
}
*/
import "C"

type RunningApplication struct {
	object C.runningApplication
}

func FrontmostApplication() *RunningApplication {
	return &RunningApplication{C.nsworkspace_frontmostApplication()}
}

func (n *RunningApplication) BundleURL() *NSURL {
	return &NSURL{object: C.nsrunningapplication_bundleurl(n.object)}
}

func (n *RunningApplication) BundleIdentifier() *NSString {
	str := C.nsrunningapplication_bundleIdentifier(n.object)
	internal := NSStringFromC(str)
	return internal
}
