# OrigAdmin Go Toolkits

This repository contains a collection of Go-specific packages designed to provide idiomatic solutions for common tasks in Go-based software development within the [OrigAdmin](https://github.com/origadmin) organization. It aims to streamline your workflow, enhance code quality, and promote consistency across Go projects.

## Introduction

The `toolkits` module is a collection of Go packages providing idiomatic solutions for common tasks encountered in Go-based software development. It aims to streamline your workflow, enhance code quality, and promote consistency across Go projects.

### Key Features

- **Idiomatic Go**: Designed with best practices and conventions in mind, ensuring seamless integration with your existing Go codebase.
- **Modular Structure**: Organized into distinct packages, each addressing a specific aspect of Go development, allowing you to pick and choose the components you need.
- **Well-Documented APIs**: Comprehensive documentation, including package-level documentation, function comments, and usage examples, to facilitate quick understanding and adoption.
- **Robust Testing**: Rigorous unit tests and integration tests accompany each package, ensuring stability, reliability, and maintainability.
- **Open Source & Community-Driven**: The Toolkit is open-source under a permissive license, encouraging community contributions, bug reports, and feature requests.

## Available Packages

The `toolkits` module currently offers the following packages:

1.  [**codec**](./codec): Utilities for encoding and decoding data in various formats.
2.  [**crypto**](./crypto): A unified and extensible framework for cryptographic operations, primarily password hashing and verification.
3.  [**decode**](./decode): Generic utilities for decoding various data structures.
4.  [**errors**](./errors): Enhanced error handling utilities, such as error wrapping, context propagation, and error inspection.
5.  [**helpers**](./helpers): A collection of small, reusable utility functions.
6.  [**identifier**](./identifier): Tools for generating and managing unique identifiers.
7.  [**io**](./io): Extended I/O utilities.
8.  [**mail**](./mail): Utilities for sending and managing emails.
9.  [**metrics**](./metrics): Tools and interfaces for collecting and exposing application metrics.
10. [**net**](./net): Network-related utilities.
11. [**slogx**](./slogx): Extensions and helpers for structured logging.
12. [**version**](./version): Utilities for managing and exposing application version information.

## Getting Started

### To incorporate packages from `toolkits` into your project, follow these steps:

1. **Add the dependency**:

- Add the `toolkits` module as a dependency in your `go.mod` file, specifying the latest version:

```bash
go get github.com/origadmin/toolkits@vX.Y.Z
```

- Replace `vX.Y.Z` with the desired version or `latest` to fetch the most recent release.

2. **Import required packages**:

- In your Go source files, import the necessary packages from `toolkits`:

```go
import (
    "github.com/origadmin/toolkits/crypto/hash"
    "github.com/origadmin/toolkits/errors"
    // ... and other packages as needed
)
```

3. **Use the toolkit components**:

- Refer to the package-specific documentation (e.g., `toolkits/crypto/README.md`) and examples to learn how to utilize the components in your code.
- You can access the documentation by running `godoc` locally or visiting the package documentation hosted on godoc.org.

## Contributing

We welcome contributions from the community to improve and expand the Toolkit. To contribute, please follow these guidelines:

1. **Familiarize yourself with the project**: Read the [CONTRIBUTING.md] file for details on the contribution process, code style, and Pull Request requirements.
2. **Submit an issue or proposal**: If you encounter any bugs, have feature suggestions, or want to discuss potential changes, create an issue in the [GitHub repository](https://github.com/origadmin/toolkit).
3. **Create a Pull Request**: After implementing your changes, submit a Pull Request following the guidelines outlined in [CONTRIBUTING.md].

## Contributors

### Code of Conduct

All contributors and participants are expected to abide by the [Contributor Covenant][ContributorHomepage], version [2.1][v2.1]. This document outlines the expected behavior when interacting with the Toolkit community.

## State

![Alt](https://repobeats.axiom.co/api/embed/bc9ad4ec869e9769ecbf84bb4a37c365a0cad47f.svg "Repobeats analytics image")

## License

The Toolkit is distributed under the terms of the [MIT License][MIT]. This permissive license allows for free use, modification, and distribution of the toolkit in both commercial and non-commercial contexts.

[CONTRIBUTING.md]: CONTRIBUTING.md
[ContributorHomepage]: https://www.contributor-covenant.org
[v2.1]: https://www.contributor-covenant.org/version/2/1/code_of_conduct.html
[MIT]: LICENSE
