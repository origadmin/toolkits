package toolkits

//go:generate protoc -I./third_party --go_out=paths=source_relative:. ./third_party/pagination/*.proto

//go:generate protoc -I./third_party --go_out=paths=source_relative:. ./third_party/errors/rpcerr/*.proto
