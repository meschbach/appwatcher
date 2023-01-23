package foundation

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation
#import <Foundation/Foundation.h>

const char* nsurl_absoluteString(NSURL* url) {
	return [[url absoluteString] UTF8String];
}

const char* nsurl_fileSystemRepresentation(NSURL* url) {
	return [url fileSystemRepresentation];
}
*/
import "C"
import "unsafe"

type NSURL struct {
	object *C.NSURL
}

func ImportURL(pointer unsafe.Pointer) *NSURL {
	return &NSURL{object: (*C.NSURL)(pointer)}
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
