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

// text document

type DidOpenTextDocumentParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}

type TextDocumentItem struct {
	/**
	 * The text document's URI.
	 */
	URI DocumentURI `json:"uri"`

	/**
	 * The text document's language identifier.
	 */
	LanguageID string `json:"languageId"`

	/**
	 * The version number of this document (it will strictly increase after each
	 * change, including undo/redo).
	 */
	Version int `json:"version"`

	/**
	 * The content of the opened text document.
	 */
	Text string `json:"text"`
}

type DocumentURI string

type PublishDiagnosticsParams struct {
	URI         DocumentURI  `json:"uri"`
	Diagnostics []Diagnostic `json:"diagnostics"`
}

type Diagnostic struct {
	/**
	 * The range at which the message applies.
	 */
	Range Range `json:"range"`

	/**
	 * The diagnostic's severity. Can be omitted. If omitted it is up to the
	 * client to interpret diagnostics as error, warning, info or hint.
	 */
	Severity DiagnosticSeverity `json:"severity,omitempty"`

	/**
	 * The diagnostic's code. Can be omitted.
	 */
	Code string `json:"code,omitempty"`

	/**
	 * A human-readable string describing the source of this
	 * diagnostic, e.g. 'typescript' or 'super lint'.
	 */
	Source string `json:"source,omitempty"`

	/**
	 * The diagnostic's message.
	 */
	Message string `json:"message"`
}

type DiagnosticSeverity int

const (
	Error       DiagnosticSeverity = 1
	Warning                        = 2
	Information                    = 3
	Hint                           = 4
)

type Range struct {
	/**
	 * The range's start position.
	 */
	Start Position `json:"start"`

	/**
	 * The range's end position.
	 */
	End Position `json:"end"`
}

type Position struct {
	/**
	 * Line position in a document (zero-based).
	 */
	Line int `json:"line"`

	/**
	 * Character offset on a line in a document (zero-based).
	 */
	Character int `json:"character"`
}
