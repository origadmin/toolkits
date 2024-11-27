module github.com/origadmin/toolkits/contrib/security/jwt

go 1.23.1

toolchain go1.23.2

replace (
	github.com/origadmin/toolkits => ../../../
	github.com/origadmin/toolkits/idgen => ../../../idgen
)

require (
	github.com/goexts/generic v0.1.1
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/origadmin/toolkits v0.0.0-00010101000000-000000000000
	github.com/origadmin/toolkits/idgen v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/coocood/freecache v1.2.4 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)