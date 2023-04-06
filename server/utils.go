package server

const (
	TDSKNone        TextDocumentSyncKind = 0
	TDSKFull        TextDocumentSyncKind = 1
	TDSKIncremental TextDocumentSyncKind = 2
)

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities,omitempty"`
}

type ServerCapabilities struct {
	TextDocumentSync                int                `json:"textDocumentSync,omitempty"`
	HoverProvider                   bool               `json:"hoverProvider,omitempty"`
	CompletionProvider              *CompletionOptions `json:"completionProvider,omitempty"`
	DefinitionProvider              bool               `json:"definitionProvider,omitempty"`
	TypeDefinitionProvider          bool               `json:"typeDefinitionProvider,omitempty"`
	ReferencesProvider              bool               `json:"referencesProvider,omitempty"`
	DocumentHighlightProvider       bool               `json:"documentHighlightProvider,omitempty"`
	DocumentSymbolProvider          bool               `json:"documentSymbolProvider,omitempty"`
	WorkspaceSymbolProvider         bool               `json:"workspaceSymbolProvider,omitempty"`
	ImplementationProvider          bool               `json:"implementationProvider,omitempty"`
	CodeActionProvider              bool               `json:"codeActionProvider,omitempty"`
	DocumentFormattingProvider      bool               `json:"documentFormattingProvider,omitempty"`
	DocumentRangeFormattingProvider bool               `json:"documentRangeFormattingProvider,omitempty"`
	RenameProvider                  bool               `json:"renameProvider,omitempty"`
}

type CompletionOptions struct {
	ResolveProvider   bool     `json:"resolveProvider,omitempty"`
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`
}

type TextDocumentSyncOptionsOrKind struct {
	Kind    *TextDocumentSyncKind
	Options *TextDocumentSyncOptions
}

type TextDocumentSyncKind int

type TextDocumentSyncOptions struct {
	OpenClose         bool                 `json:"openClose,omitempty"`
	Change            TextDocumentSyncKind `json:"change"`
	WillSave          bool                 `json:"willSave,omitempty"`
	WillSaveWaitUntil bool                 `json:"willSaveWaitUntil,omitempty"`
	Save              *SaveOptions         `json:"save,omitempty"`
}

type SaveOptions struct {
	IncludeText bool `json:"includeText"`
}
