# Toolkit

## Introduction

The Toolkit is a collection of Go-specific packages providing idiomatic solutions for common tasks encountered in
Go-based software development. It aims to streamline your workflow, enhance code quality, and promote consistency across
Go projects. This README focuses on the Toolkit, detailing its contents, usage, and contribution guidelines.

### Key Features

- **Idiomatic Go**: Designed with best practices and conventions in mind, ensuring seamless integration with your
  existing Go codebase.
- **Modular Structure**: Organized into distinct packages, each addressing a specific aspect of Go development, allowing
  you to pick and choose the components you need.
- **Well-Documented APIs**: Comprehensive documentation, including package-level documentation, function comments, and
  usage examples, to facilitate quick understanding and adoption.
- **Robust Testing**: Rigorous unit tests and integration tests accompany each package, ensuring stability, reliability,
  and maintainability.
- **Open Source & Community-Driven**: The Toolkit is open-source under a permissive license, encouraging community
  contributions, bug reports, and feature requests.

## Available Packages

The Toolkit main package is `toolkit` and is the interface package.:

1. [**registry**](registry): The registry package is register implementation for the main package definition.

The Toolkit currently offers the following packages:

1. [**codec**](codec): A set of utilities for serializing and deserializing data in a variety of formats.
2. [**errors**](errors): Enhanced error handling utilities, such as error wrapping, context propagation, and error
   inspection.
3. [**idgen**](idgen): A package for generating unique identifiers.
4. [**crypto**](crypto): A powerful library designed to provide cryptographic functionalities, including hashing, encryption, and decryption. This toolkit is built with performance and security in mind, making it an essential component for any application that requires secure data handling.

## Getting Started

### To incorporate the Toolkit into your project, follow these steps:

1. **Add the dependency**:

- Add the Toolkit as a dependency in your `go.mod` file, specifying the latest version:

```bash
go get github.com/origadmin/toolkit@vX.Y.Z
```

- Replace `vX.Y.Z` with the desired version or `latest` to fetch the most recent release.

2. **Import required packages**:

- In your Go source files, import the necessary packages from the Toolkit:

```go
import (
"github.com/origadmin/toolkits/errors"
"github.com/origadmin/runtime"
//    "github.com/origadmin/runtime/xxx"
"github.com/origadmin/runtime/config"
"github.com/origadmin/contrib"
//    "github.com/origadmin/contrib/xxx"    
"github.com/origadmin/contrib/config"
)
```

3. **Use the toolkit components**:

- Refer to the package documentation and examples to learn how to utilize the toolkit components in your code.
- You can access the documentation by running `godoc` locally or visiting the package documentation hosted on godoc.org.

## Call relationships intro

In `OrigAdmin`, the call relationship between packages is as follows:

- [Contrib](https://github.com/origadmin/contrib) : in view of the toolkits interface, the realization of the runtime and kratos.
    - [contrib/consul/config ](https://github.com/origadmin/contrib/consul/config) : the runtime config to encapsulate an implementation, encapsulates the consul of the client
    - [contrib/consul/registry](https://github.com/origadmin/contrib/consul/registry) : An implementation of kratos config's encapsulation that encapsulates consul's client
    - [contrib/config](https://github.com/origadmin/contrib/config) : kratos config encapsulation of implementation

- [Runtime](https://github.com/origadmin/runtime) : encapsulates the kratos runtime required interfaces, including basic functions, initialize the application, as well as the service registry, service discovery, etc.
    - [runtime/registry](https://github.com/origadmin/runtime/registry) : provide basic service registry found the function definition.
    - [runtime/config ](https://github.com/origadmin/runtime/config) : provide configuration file is read and parse the definition of the configuration of the analytic function.
- [Toolkit](https://github.com/origadmin/toolkits) : toolkits provide some of the basic function of the general interface definition or implementation, such as the serialization and deserialization, error handling, a unique identifier generated, etc.
    -  [toolkits/codec](https://github.com/origadmin/toolkits/codec) : provide a variety of formats data serialization and deserialization.
    -  [toolkit/errors](https://github.com/origadmin/toolkits/errors) : provide enhanced error handling utility, such as wrong packaging, context propagation and error checking.

```shell
your_project --> contrib --> contrib/config --> runtime --> toolkits
             --> contrib/config
             --> runtime --> rumtime/registry
             --> rumtime/registry
             --> toolkits
             --> toolkits/codec
             --> toolkits/errors
             --> other_package
```

```mermaid
graph LR
A(your_project) 
A --> B(contrib)
A --> C(runtime)
A --> D(toolkits)
B --> E(contrib/config)
C --> F(runtime/config)
C --> G(runtime/registry)
D --> H(toolkits/codec)
D --> I(toolkits/errors)
```

## Contributing

We welcome contributions from the community to improve and expand the Toolkit. To contribute, please follow these
guidelines:

1. **Familiarize yourself with the project**: Read the [CONTRIBUTING] file for details on the contribution process, code
   style, and Pull Request requirements.
2. **Submit an issue or proposal**: If you encounter any bugs, have feature suggestions, or want to discuss potential
   changes, create an issue in the [GitHub repository](https://github.com/origadmin/toolkit).
3. **Create a Pull Request**: After implementing your changes, submit a Pull Request following the guidelines outlined
   in [CONTRIBUTING].

## Contributors

### Code of Conduct

All contributors and participants are expected to abide by the [Contributor Covenant][ContributorHomepage],
version [2.1][v2.1]. This document outlines the expected behavior when interacting with the Toolkit community.

## State

![Alt](https://repobeats.axiom.co/api/embed/bc9ad4ec869e9769ecbf84bb4a37c365a0cad47f.svg "Repobeats analytics image")

## License

The Toolkit is distributed under the terms of the [MIT]. This permissive license allows for free use, modification, and
distribution of the toolkit in both commercial and non-commercial contexts.

[CONTRIBUTING]: CONTRIBUTING.md

[ContributorHomepage]: https://www.contributor-covenant.org

[v2.1]: https://www.contributor-covenant.org/version/2/1/code_of_conduct.html

[MIT]: LICENSE
