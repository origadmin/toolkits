module github.com/origadmin/toolkits

go 1.22.2

require (
	github.com/origadmin/toolkits/errors v0.0.4
	github.com/stretchr/testify v1.9.0
	golang.org/x/crypto v0.22.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v0.0.3
	v0.0.2
	v0.0.1
)
