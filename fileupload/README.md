# This is an upload intermediary service

## The features to be done now are as follows

1. The client uploads the file and receives it from the server
2. The client bridges the http upload and the server receives it.
3. Both the client and the server should support HTTP and grpc.
4. Add gRPC upload function
5. Unified HTTP and gRPC upload interfaces
6. Support multi-file upload (by calling uploader/receiver multiple times)
7. Provides better error handling and resource management
8. Added folder upload support
9. Implemented the function of resumable transmission
10. Buffer pools are used to optimize memory usage
11. Maintain interface consistency
12. Better error handling is provided
13. Provides HTTP to gRPC bridge upload
14. The bridge from HTTP to gRPC supports resumable upload
15. HTTP to gRPC bridging optimizes memory usage using buffer pools
16. Contextual support is provided for the HTTP to gRPC bridge upload

## Implementation of the underlying logic

1. The user clicks the upload button to upload the file (HTTP).
2. The server receives the uploader that bridges HTTP and gRPC) and forwards it to the grpc server.
3. The grpc server receives the data and saves it to the file/OSS.

Examples:
```
func main() {

 // Create a gRPC Builder

  grpcBuilder := NewBuilder(
    WithServiceType(ServiceTypeGRPC),
    WithURI("grpc-server-address:port"),
  )

  // Create a bridge uploader
  bridgeUploader := NewBridgeUploader(grpcBuilder, &uploadBuilder{
    bufPool: &sync.Pool{
      New: func() interface{} {
        return make([]byte, 32*1024)
      },
    },
  })

  // Register the HTTP processor
  http.HandleFunc("/upload", bridgeUploader.ServeHTTP)
  http.HandleFunc("/upload/resume", bridgeUploader.ServeHTTPWithResume)
  // Start the HTTP server
  http.ListenAndServe(":8080", nil)

}
```

### Workflow:

1. The client uploads the file via HTTP
2. BridgeUploader receives HTTP requests
3. Use HTTPReceiver to read the contents of the file
4. Create a gRPC uploader and forward the file
5. Return the gRPC response to the HTTP client