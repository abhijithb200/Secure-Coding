package server

import (
	"context"
	"fmt"

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
		return h.handleTextDocumentDidOpen(ctx, conn, req)
	case "initialized":
		return

	}
	return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeMethodNotFound, Message: fmt.Sprintf("method not supported: %s", req.Method)}
}
