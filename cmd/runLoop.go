package main

/*
#cgo CFLAGS: -x objective-c -Wincompatible-pointer-types
#cgo LDFLAGS: -framework Foundation -framework Appkit
#import <Foundation/Foundation.h>
#import <Appkit/Appkit.h>

void runMainLoop(){
	[[NSRunLoop mainRunLoop] run];
}
*/
import "C"

func run() {
	C.runMainLoop()
}
