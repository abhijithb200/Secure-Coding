package server

import (
	"context"
	"fmt"
	"net/url"

	"github.com/sourcegraph/jsonrpc2"
)

type handler struct {
	conn *jsonrpc2.Conn
}

func NewHandler() jsonrpc2.Handler {
	handler := &handler{}
	return jsonrpc2.HandlerWithError(handler.handle)
}

func (h *handler) handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (result interface{}, err error) {
	switch req.Method {
	case "initialize":
		return h.handleInitialize(ctx, conn, req)
	case "textDocument/didOpen":
		// return h.handleTextDocumentDidOpen(ctx, conn, req)
	case "textDocument/didSave":
		return h.handleTextDocumentDidSave(ctx, conn, req)
		
	case "textDocument/definition":
		return nil,nil
	case "textDocument/completion":
		return nil,nil
	case "textDocument/hover":
		return nil,nil

	case "initialized":
		return

	}
	return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeMethodNotFound, Message: fmt.Sprintf("method not supported: %s", req.Method)}
}

func uriToDocumentURI(uri string) DocumentURI {
	return DocumentURI(fmt.Sprintf("file://%s", uri))
}

func documentURIToURI(duri DocumentURI) string {
	s := string(duri)[len("file://"):][1:]
	decodedValue, _ := url.QueryUnescape(s) 
	return decodedValue
}
