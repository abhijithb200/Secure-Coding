package server

type InitializeResult struct {
	ServerInfo struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}
	Capabilities ServerCapabilities `json:"capabilities"`
}

type initializeParams struct {
	ProcessId    int                `json:"processid"`
	RootURI      string             `json:"rootUri"`
	Capabilities ClientCapabilities `json:"capabilities"`
}

type ClientCapabilities struct{}

type ServerCapabilities struct {
	TextDocumentSync          TextDocumentSyncKind `json:"textDocumentSync"`
	CompletionProvider        CompletionOption     `json:"completionProvider"`
	HoverProvider             bool                 `json:"hoverProvider"`
	TypeDefinitionProvider    bool                 `json:"typeDefinitionProvider"`
	ImplementationProvider    bool                 `json:"implementationProvider"`
	DocumentHighlightProvider bool                 `json:"documentHighlightProvider"`
	DocumentSymbolProvider    bool                 `json:"documentSymbolProvider"`
}
type TextDocumentSyncKind int

type CompletionOption struct {
	WorkDoneProgressOption
	ResolveProvider   bool     `json:"resolveProvider"`
	TriggerCharacters []string `json:"triggerCharacters"`
}

type WorkDoneProgressOption struct {
	WorkDoneProgress bool `json:"workDoneProgress"`
}

type NoOpErr struct{}

func (NoOpErr) Error() string {
	return "NoOp"
}

// it is the parameter that is sent when a textDocument/didOpenFile command is send

type DidOpenTextDocumentParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}

// it is the document itself
type TextDocumentItem struct {
	URI        string `json:"uri"`
	LanguageID string `json:"languageId"`
	Version    int    `json:"version"`
	Text       string `json:"text"` // contents of the file
}

type fileMetadata struct {
	TextDocumentItem
	schemaURI string

	isSchema bool
	parsed   interface{}
}

type DidSaveTextDocumentParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Text         string                 `json:"text`
}

type TextDocumentIdentifier struct {
	URL string `json:"uri"`
}

type DidChangeTextDocumentParams struct {
	TextDocument   VersionedTextDocumentIdentifier  `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:contentChanges`
}
