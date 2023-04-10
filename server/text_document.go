package server

import (
	"context"
	"encoding/json"

	"github.com/sourcegraph/jsonrpc2"
)

func (h *handler) handleTextDocumentDidOpen(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (result interface{}, err error) {
	if req.Params == nil {
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
	}

	var params DidOpenTextDocumentParams
	if err := json.Unmarshal(*req.Params, &params); err != nil {
		return nil, err
	}

	msg := PublishDiagnosticsParams{
		URI: params.TextDocument.URI,
		Diagnostics: []Diagnostic{
			Diagnostic{
				Message:  "I am abhi",
				Severity: 4,
				Range: Range{
					Start: Position{
						Line:      0,
						Character: 0,
					},
					End: Position{
						Line:      0,
						Character: 0,
					},
				},
			},
		},
	}

	conn.Notify(ctx, "textDocument/publishDiagnostics", msg)
	return nil, nil
}
