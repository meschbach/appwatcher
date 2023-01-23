package foundation

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation
#import <Foundation/Foundation.h>

typedef NSString* objcString;

const char* nsstring_utf8(objcString str){
	return [str UTF8String];
}
*/
import "C"
import "unsafe"

type ObjcString C.objcString

type String struct {
	object C.objcString
}

// WrapString wraps a pointer to an String* to be operated upon in Go.
// TODO: Revisit for a better way.  Effectively I need to figure out how to correctly express to cgo the return type is
// the same despite being in other packages.
func WrapString(obj unsafe.Pointer) *String {
	return &String{object: C.objcString(obj)}
}

// Internalize extracts the given string into a native Go string.
func (n *String) Internalize() string {
	cString := C.nsstring_utf8(n.object)
	return C.GoString(cString)
}
