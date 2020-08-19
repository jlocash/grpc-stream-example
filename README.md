# grpc-stream-example
This repository serves as an example implementation of server-side streaming using gRPC

## Build
To compile, you must first have the protobuf compiler extensions for golang and gRPC installed.
Once you have those installed, simply run ```make compile``` and server/client binaries will be compiled for darwin and Linux platforms. These binaries can be located in the ```build``` directory

## Run
To run the server, simply execute: 
```
$ ./build/grpc-server-<platform>-amd64
```

With that running, in a separate terminal execute:
```
$ ./build/grpc-client-<platform>-amd64
```

The messages printed to stdout are sent over the stream from the server using a goroutine and channel. You could remove the goroutine entirely and instead communicate through the client channels from event-driven operations.
