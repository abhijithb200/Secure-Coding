package server

import (
	"context"

	"github.com/sourcegraph/jsonrpc2"
)

func (h *handler) handleInitialize(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (result interface{}, err error) {

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
}
