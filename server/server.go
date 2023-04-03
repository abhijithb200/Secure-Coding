package server

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/sourcegraph/jsonrpc2"
)

const version = "UNTAGGED"
const (
	Full TextDocumentSyncKind = iota
	Incremental
)

type Server struct {
	isInitialized bool

	Endpoint string
}

func (s *Server) Stop() {}

func (s *Server) Handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (interface{}, error) {

	switch req.Method {
	case "initialize":
		if s.isInitialized {
			return nil, NoOpErr{}
		}
		if req.Params == nil {
			return nil, errors.Errorf("Cannot have nil params for intilize")
		}
		var params initializeParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			return nil, errors.Wrapf(err, "Unable to unmarshal intializeParams")
		}

		retVal := InitializeResult{
			ServerInfo: struct {
				Name    string "json:\"name\""
				Version string "json:\"version\""
			}{Name: "LSP", Version: version},
			Capabilities: ServerCapabilities{
				TextDocumentSync: Full,
				CompletionProvider: CompletionOption{
					WorkDoneProgressOption: WorkDoneProgressOption{
						WorkDoneProgress: true,
					},
					ResolveProvider:   false,
					TriggerCharacters: []string{"{", ":"},
				},
			},
		}
		return retVal, nil

	case "shutdown":
		return nil, nil
	case "exit":
		return nil, nil
	case "textDocument/didOpen":
		if req.Params == nil {
			return nil, errors.Errorf("Cannot have nil parameter for textDocument/didOpen")
		}
	case "textDocument/didSave":
	case "textDocument/hover":
	case "textDocument/completion":

	default:
		return nil, errors.Errorf("Method %q not supported", req.Method)

	}

	return nil, nil
}

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
