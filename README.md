# appwatcher
An OSX program to journal which applications are currently at the forefront of the focus stack.

## Context
Found focus would randomly change for a few second occasionally.  I found a Python script on a StackOverflow post
however it polled every second instead of using events.  Thought writing a similar program in Golang would be an
interesting way to learn about `cgo` and it sure was.

## Development
Running `go build -o appwatcher ./cmd` will generate the binary.  Placed most of the correct code in `pkg` under the
OSX framework it is related to.  You will need to compile it on OSX to build successfully.

`NotificationCenter` is actually in `Foundation` however in this case it is strongly bound to the notification of
activation changes for applications.  So in my case it didn't make sense to further factor out.

## Contributing
Welcoming contributions which improve either the `appwatch` application or the bindings against OSX specific frameworks.

## Future
Need to journal the data somewhere useful, such as Postgres or a SQLite type database.  This will allow review of the
applications and filtering.  Might make sense to continue building out appkit and foundation library.
