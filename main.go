package main

import (
	"context"
	"encoding/json"

	"errors"

	"github.com/sourcegraph/jsonrpc2"
)

const version = "UNTAGGED"
const (
	Full TextDocumentSyncKind = iota
	Incremental
)

type Server struct {
	isInitialized bool
}

func (s *Server) Handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) ([]byte, error) {

	switch req.Method {
	case "initialize":
		if req.Params == nil {
			return nil, errors.Errorf("Cannot have nil params for intilize")
		}
		var params initializeParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			return nil, errors.Wrapf(err, "Unable to unmarshal intializeParams")
		}

		res := InitializeResult{
			ServerInfo: struct {
				Name    string "json:\"name\""
				Version string "json:\"version\""
			}{Name: "LSP", Version: version},
			Capabilities: ServerCapabilities{
				TextDocumentSync: Full,
				CompletionProvider: {
					WorkDoneProgressOption:{WorkDoneProgressOption:{
						WorkDoneProgress:false
					}},
				},
			},
		}

	case "textDocument/didOpen":
	case "textDocument/didSave":

	}
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

type CompletionOption struct {
	WorkDoneProgressOption
	ResolveProvider   bool     `json:"resolveProvider"`
	TriggerCharacters []string `json:"triggerCharacters"`
}

type WorkDoneProgressOption struct {
	WorkDoneProgress bool `json:"workDoneProgress"`
}

type TextDocumentSyncKind int
