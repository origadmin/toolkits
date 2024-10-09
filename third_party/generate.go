package third_party

//go:generate protoc -I. --go_out=paths=source_relative:. ./pagination/pagination.proto
//go:generate protoc -I. --openapiv2_out . --openapiv2_opt logtostderr=true ./pagination/pagination.proto
