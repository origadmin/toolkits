package toolkits

//go:generate protoc -I./third_party --go_out=paths=source_relative:. ./third_party/errors/rpcerr/*.proto

//go:generate protoc -I./third_party --go_out=paths=source_relative:. ./third_party/middlewares/*.proto

//go:generate protoc -I./third_party --validate_out=paths=source_relative,lang=go:. ./third_party/middlewares/*.proto
