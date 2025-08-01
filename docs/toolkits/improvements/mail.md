# Mail Module Improvements & Todo List

This document outlines identified areas for improvement and pending tasks within the `mail` module.

---

## 1. Lack of Implementation for `Sender` Interface

### Problem Description

The `mail` module currently defines a `Sender` interface but provides no concrete implementations. This means the
package is not directly usable for sending emails without external code providing an implementation.

### Current State Analysis

- **Relevant File**: `mail/mail.go`
- **Observation**: Only the `Sender` interface and its methods (`Send`, `SendTo`) are defined. No structs implement this
  interface within the package.

### Proposed Solution(s)

Provide at least one concrete implementation of the `Sender` interface. A common approach would be to implement an SMTP
client using Go's standard library (`net/smtp`) or integrate with a popular email service provider's SDK.

### Expected Impact

- **Usability**: The `mail` module will become directly functional for sending emails.
- **Completeness**: Provide a ready-to-use solution for email sending.

### Verification Plan

- **Unit Tests**: Develop comprehensive unit tests for the implemented `Sender` to verify email sending functionality,
  including recipient handling, subject, body, and attachments.

---

## 2. Missing `README.md` Documentation

### Problem Description

The `mail` module lacks a `README.md` file, which is crucial for providing an overview of the package's purpose, its
functionalities, and how to use it.

### Current State Analysis

- **Observation**: No `README.md` file was found within the `mail` directory.

### Proposed Solution(s)

Create a `README.md` file that includes:

- A brief overview of the `mail` package as an email sending abstraction.
- Explanation of the `Sender` interface and its methods.
- Instructions on how to implement the `Sender` interface.
- If a concrete implementation is provided, details on its configuration and usage.
- Clarification on the `file ...string` parameter (e.g., local file paths for attachments).
- Clarification on the `body string` parameter's expected content type (e.g., plain text, HTML).

### Expected Impact

- **Improved Usability**: Developers can quickly understand and use the `mail` package.
- **Better Maintainability**: Clear documentation reduces the learning curve for new contributors.

### Verification Plan

- **Documentation Review**: Ensure the `README.md` is clear, concise, and accurate.

---

## 3. Lack of Comprehensive Unit Tests

### Problem Description

Since there are no concrete implementations of the `Sender` interface, there are no unit tests for the `mail` module.
This means that once an implementation is added, its correctness and robustness will not be verified.

### Current State Analysis

- **Observation**: No `_test.go` files were found within the `mail` directory.

### Proposed Solution(s)

Once a concrete implementation of the `Sender` interface is provided, develop comprehensive unit tests to verify its
functionality. Tests should cover:

- Successful email sending.
- Error handling (e.g., invalid recipients, connection issues).
- Correct handling of CC, BCC, subject, body, and attachments.
- Concurrent sending scenarios.

### Expected Impact

- **Increased Reliability**: Ensure the correctness and robustness of the email sending functionality.
- **Easier Maintenance**: Future changes can be made with confidence, knowing that existing functionality is protected
  by tests.

### Verification Plan

- **Unit Tests**: Develop and execute unit tests using Go's testing framework.

---

## 4. Limited Attachment Handling

### Problem Description

The `Send` and `SendTo` methods currently accept `file ...string` for attachments, implying local file paths. This might
be restrictive if there's a need to send attachments from memory (e.g., generated content) or handle more complex
attachment properties (e.g., MIME types, inline attachments).

### Current State Analysis

- **Relevant File**: `mail/mail.go`
- **Relevant Code Snippet**:
    ```go
    type Sender interface {
        Send(ctx context.Context, to []string, cc []string, bcc []string, subject string, body string, file ...string) error
        SendTo(ctx context.Context, to []string, subject string, body string, file ...string) error
    }
    ```

### Proposed Solution(s)

Consider introducing a more flexible `Attachment` type or struct that can encapsulate various attachment sources (file
path, byte slice, `io.Reader`) and properties (filename, MIME type, Content-ID for inline attachments). This
`Attachment` type would then be used in the `Send` methods.

### Expected Impact

- **Increased Flexibility**: Support a wider range of attachment scenarios.
- **Improved API Design**: Provide a more robust and extensible way to handle attachments.

### Verification Plan

- **API Design Review**: Review the new `Attachment` type and its integration into the `Sender` interface.
- **Unit Tests**: Develop tests for various attachment types and sources.

---

## 5. Unspecified Email Body Content Type

### Problem Description

The `body string` parameter in the `Send` and `SendTo` methods does not explicitly specify the content type (e.g., plain
text, HTML). This can lead to issues with email clients rendering the message correctly.

### Current State Analysis

- **Relevant File**: `mail/mail.go`
- **Relevant Code Snippet**:
    ```go
    type Sender interface {
        Send(ctx context.Context, to []string, cc []string, bcc []string, subject string, body string, file ...string) error
        SendTo(ctx context.Context, to []string, subject string, body string, file ...string) error
    }
    ```

### Proposed Solution(s)

Add a parameter or a field in a `Message` struct (if introduced) to specify the `Content-Type` of the email body (e.g.,
`text/plain`, `text/html`). This will ensure email clients render the message as intended.

### Expected Impact

- **Improved Email Rendering**: Ensure emails are displayed correctly across different clients.
- **Enhanced API Clarity**: Make the expected body format explicit.

### Verification Plan

- **API Design Review**: Review the updated `Send` method signatures or `Message` struct.
- **Unit Tests**: Test sending emails with different content types.
