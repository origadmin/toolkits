package third_party

//go:generate protoc --proto_path=. --proto_path=. --go_out=paths=source_relative:. --go-gin_out=paths=source_relative:. ./pagination/pagination.proto
//go:generate protoc --proto_path=. --proto_path=. --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. ./pagination/pagination.proto
//go:generate protoc --proto_path=. --proto_path=. --go_out=paths=source_relative:. --go-http_out=paths=source_relative:. ./pagination/pagination.proto
//go:generate protoc --proto_path=. --proto_path=. --go_out=paths=source_relative:. --go-errors_out=paths=source_relative:. ./pagination/pagination.proto
//go:generate protoc --proto_path=. --proto_path=. --go_out=paths=source_relative:. ./pagination/pagination.proto
//go:generate protoc --proto_path=. --proto_path=. --openapiv2_out . --openapiv2_opt logtostderr=true ./pagination/pagination.proto
