package server

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/sourcegraph/jsonrpc2"
)


func (h *handler) handleTextDocumentDidSave(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (result interface{}, err error) {
	if req.Params == nil {
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
	}
	
	var params DidOpenTextDocumentParams
	if err := json.Unmarshal(*req.Params, &params); err != nil {
		return nil, err
	}

	
	src, _ := ioutil.ReadFile(documentURIToURI(params.TextDocument.URI))
	s := string(src)
	p := Post{
		Contents: s,
	}
	v,_ := json.Marshal(p)

	foo2 := FinalReport{}
	
	getJson("http://localhost:3000/analyze", &foo2,v)

	msg := PublishDiagnosticsParams{
		URI: params.TextDocument.URI,
		Diagnostics: []Diagnostic{
			Diagnostic{
				Message:  documentURIToURI(params.TextDocument.URI),
				Code: "This is a huge problem",
				Severity: 1,
				Range: Range{
					Start: Position{
						Line:      10,
						Character: 0,
					},
					End: Position{
						Line:      11,
						Character: 0,
					},
				},
			},
		},
	}

	conn.Notify(ctx, "textDocument/publishDiagnostics", msg)
	return nil, nil
}

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
				Message:  "I am pari",
				Severity: 1,
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
