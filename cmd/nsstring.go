package main

/*
#cgo CFLAGS: -x objective-c -Wincompatible-pointer-types
#cgo LDFLAGS: -framework Foundation
#import <Foundation/Foundation.h>

typedef NSString* objcString;

const char* nsstring_utf8(objcString str){
	return [str UTF8String];
}
*/
import "C"

type NSString struct {
	object C.objcString
}

func NSStringFromC(obj C.objcString) *NSString {
	return &NSString{object: obj}
}

func (n *NSString) Internalize() string {
	cString := C.nsstring_utf8(n.object)
	return C.GoString(cString)
}
