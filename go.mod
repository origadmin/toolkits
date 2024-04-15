module github.com/origadmin/toolkits

go 1.22.0

require (
	github.com/BurntSushi/toml v1.3.2
	github.com/hashicorp/go-multierror v1.1.1
	github.com/json-iterator/go v1.1.12
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.9.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
)

retract (
	v0.0.3
	v0.0.2
	v0.0.1
)
