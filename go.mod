module github.com/origadmin/toolkits

go 1.24.0

require (
	github.com/bits-and-blooms/bloom/v3 v3.7.1
	github.com/goexts/generic v0.8.0
	github.com/origadmin/toolkits/errors v1.0.0
	github.com/stretchr/testify v1.11.1
)

require (
	github.com/bits-and-blooms/bitset v1.24.2 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	golang.org/x/exp v0.0.0-20241009180824-f66d83c29e7c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v0.0.90
	v0.0.89
	v0.0.88
	v0.0.87
)
