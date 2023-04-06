package server

import (
	"context"
	"fmt"

	"github.com/sourcegraph/jsonrpc2"
)

type handler struct {
}

func NewHandler() jsonrpc2.Handler {
	handler := &handler{}
	return jsonrpc2.HandlerWithError(handler.handle)
}

func (h *handler) handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (result interface{}, err error) {
	switch req.Method {
	case "initialize":
		return InitializeResult{
			Capabilities: ServerCapabilities{
				TextDocumentSync:           1,
				DocumentFormattingProvider: true,
				DefinitionProvider:         true,
				HoverProvider:              true,
				ReferencesProvider:         true,
				CompletionProvider: &CompletionOptions{
					TriggerCharacters: []string{"*", "."},
					ResolveProvider:   true,
				},
			},
		}, nil
	case "initialized":
		return

	}
	return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeMethodNotFound, Message: fmt.Sprintf("method not supported: %s", req.Method)}
}
