## Language Server Protocol(LSP)

- https://microsoft.github.io/language-server-protocol/

The Language Server Protocol (LSP) defines the protocol used between an editor or IDE and a language server that provides language features like auto complete, go to definition, find all references etc. The goal of the Language Server Index Format (LSIF, pronounced like "else if") is to support rich code navigation in development tools or a Web UI without needing a local copy of the source code.

- It standardizes the communication between tooling and code editor
- Language Server can be implemented in any language tooling and code editor



## Language Server Extension

- https://code.visualstudio.com/api/language-extensions/language-server-extension-guide

With Language Servers, you can implement autocomplete, error-checking (diagnostics), jump-to-definition. In VS code, language server has two parts:

1. Language Client : A normal VS Code extension written in JS/TS
2. Language Server : A language analysis tool running in a separate process. It can be implemented in any langugages, as long as it can communicate with Language Client following Language Server Protocol. As it is running in a separate process, CPU and Memory usage decreases and VS code will not be slow.