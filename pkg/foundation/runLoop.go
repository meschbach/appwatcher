package foundation

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation
#import <Foundation/Foundation.h>

void runMainLoop(){
	[[NSRunLoop mainRunLoop] run];
}
*/
import "C"

func RunMainLoop() {
	C.runMainLoop()
}
