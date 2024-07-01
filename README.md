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

The Toolkit currently offers the following packages:

1. **`errors`**: Enhanced error handling utilities, such as error wrapping, context propagation, and error inspection.
2. **`httpclient`**: A configurable and extensible HTTP client with support for timeouts, retries, and middleware
   chaining.
3. **`jsonutil`**: Convenient functions for JSON encoding, decoding, and manipulation, including custom JSON
   marshalers/unmarshalers and error handling improvements.
4. **`concurrency`**: Tools for managing concurrent tasks, synchronization primitives, and thread-safe data structures.
5. **`config`**: Simplified configuration management, supporting various file formats and environment variable
   integration.
6. **`logging`**: A flexible logging framework with customizable log levels, formatters, and integrations with popular
   logging backends.

## Getting Started

To incorporate the Toolkit into your project, follow these steps:

1. **Add the dependency**: Add the Toolkit as a dependency in your `go.mod` file, specifying the latest version:

```bash
go get github.com/origadmin/toolkit@vX.Y.Z
```

Replace `vX.Y.Z` with the desired version or `latest` to fetch the most recent release.

2. **Import required packages**: In your Go source files, import the necessary packages from the Toolkit:

```go
import (
"github.com/origadmin/toolkit/errors"
"github.com/origadmin/toolkit/httpclient"
"github.com/origadmin/toolkit/jsonutil"
"github.com/origadmin/toolkit/concurrency"
"github.com/origadmin/toolkit/config"
"github.com/origadmin/toolkit/logging"
)
```

3. **Use the toolkit components**: Refer to the package documentation and examples to learn how to utilize the toolkit
   components in your code. You can access the documentation by running `godoc` locally or visiting the package
   documentation hosted on godoc.org.

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
