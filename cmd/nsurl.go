package main

/*
#cgo CFLAGS: -x objective-c -Wincompatible-pointer-types
#cgo LDFLAGS: -framework Foundation -framework Appkit
#import <Cocoa/Cocoa.h>
#import <Foundation/Foundation.h>
#import <Appkit/Appkit.h>

const char* nsurl_absoluteString(NSURL* url) {
	return [[url absoluteString] UTF8String];
}

const char* nsurl_fileSystemRepresentation(NSURL* url) {
	return [url fileSystemRepresentation];
}
*/
import "C"

type NSURL struct {
	object *C.NSURL
}

func (n *NSURL) Absolute() string {
	//return ""
	result := C.nsurl_absoluteString(n.object)
	return C.GoString(result)
}

func (n *NSURL) FileSystemPath() string {
	//return ""
	result := C.nsurl_fileSystemRepresentation(n.object)
	return C.GoString(result)
}
