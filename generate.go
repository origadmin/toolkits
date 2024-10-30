package toolkits

//go:generate protoc -I. -I./third_party --go_out=paths=source_relative:. ./errors/rpcerr/*.proto

//go:generate protoc -I. -I./third_party --go_out=paths=source_relative:. ./runtime/middlewares/*.proto
//go:generate protoc -I. -I./third_party --validate_out=paths=source_relative,lang=go:. ./runtime/middlewares/*.proto

//go:generate protoc -I. -I./third_party --go_out=paths=source_relative:. ./runtime/config/*.proto
//go:generate protoc -I. -I./third_party --validate_out=paths=source_relative,lang=go:. ./runtime/config/*.proto
