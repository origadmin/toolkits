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

## Call relationships between packages

In OrigAdmin, the call relationship between packages is as follows:

- [**contrib**](https://github.com/origadmin/contrib): defines configuration structs and parsing functions for reading and parsing configuration files.
    - [**contrib/config**](https://github.com/origadmin/contrib/config): provides configuration parsing functions for reading and parsing configuration files.
  
- [**runtime**](https://github.com/origadmin/runtime): Provides the basic functions of the runtime environment, such as logging and configuration management.
    - [**runtime/registry**](https://github.com/origadmin/runtime/registry): provides basic registry functions, such as service registration and service discovery.
    - [**runtime/config**](https://github.com/origadmin/runtime/config): provides configuration parsing functions for reading and parsing configuration files.

- [**toolkits**](https://github.com/origadmin/toolkits): provides basic functions for toolkit, such as error handling, context management, and logging.
    - [**toolkits/codec**](https://github.com/origadmin/toolkits/codec): provides functions for serializing and deserializing data in a variety of formats.
    - [**toolkits/errors**](https://github.com/origadmin/toolkits/errors): provides enhanced error handling utilities, such as error wrapping, context propagation, and error inspection.

```shell
your_project -> contrib -> contrib/config -> runtime -> toolkits
			 -> runtime -> rumtime/registry
			 -> other_package
```

```mermaid
graph LR
A(your_project) --> B(contrib)
A --> C(runtime)
A --> D(toolkits)
B --> E(contrib/config)
B --> C(runtime)
B --> D(toolkits)
C --> F(runtime/config)
C --> G(runtime/registry)
D --> H(toolkits/codec)
D --> I(toolkits/errors)
I --> J(other_package)
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
