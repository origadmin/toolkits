package toolkits

//go:generate protoc -I. -I./third_party --go_out=paths=source_relative:. ./third_party/errors/rpcerr/*.proto

//go:generate protoc -I. -I./third_party --go_out=paths=source_relative:. ./middlewares/*.proto
//go:generate protoc -I. -I./third_party --validate_out=paths=source_relative,lang=go:. ./middlewares/*.proto

//go:generate protoc -I. -I./third_party --go_out=paths=source_relative:. ./runtime/config/*.proto
//go:generate protoc -I. -I./third_party --validate_out=paths=source_relative,lang=go:. ./runtime/config/*.proto
