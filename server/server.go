package server

import (
	"context"
	"encoding/json"
	"regexp"

	"github.com/pkg/errors"

	"github.com/sourcegraph/jsonrpc2"
	"github.com/vektah/gqlparser/ast"
	"github.com/vektah/gqlparser/parser"
)

const version = "UNTAGGED"
const (
	Full TextDocumentSyncKind = iota
	Incremental
)

type Server struct {
	isInitialized bool

	Endpoint string // "global" where the schema sits
	m        map[string]fileMetadata
}

func New() *Server {
	return &Server{
		m: make(map[string]fileMetadata),
	}
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

		var params DidOpenTextDocumentParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			return nil, errors.Wrap(err, "Unable to unmarshall DidOpenTextDocumentParams")
		}

		m := fileMetadata{
			TextDocumentItem: params.TextDocument,
		}
		s.m[m.URI] = m

		parsed, err := parse(m.URI, m.TextDocumentItem.Text)
		if err != nil {
			return nil, errors.Wrapf(err, "Cannot parse input file")
		}
		s.m[m.URI].parsed = parsed

		switch parsed.(type) {
		case *ast.QueryDocument:
		case *ast.SchemaDocument:
			s.m[m.URI].schemaURI = parsed
		default:
		}

	case "textDocument/didChange":
		if req.Params == nil {
			return nil, errors.Errorf("Cannot have nil parameter for textDocument/didChange")
		}

		var params DidChangeTextDocumentParams

	// when the file is saved - parse the text
	case "textDocument/didSave":
		if req.Params == nil {
			return nil, errors.Errorf("Cannot have nil parameter for textDocument/didOpen")
		}
		var params DidSaveTextDocumentParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			return nil, errors.Wrap(err, "Unable to unmarshall DidSaveTextDocumentParams")
		}

		id := params.TextDocument.URL
		if params.Text != "" {
			s.m[id].TextDocument.Text = params.Text
		}
		return nil, nil

	case "textDocument/hover":
	case "textDocument/completion":
	case "textDocument/codeAction":

	default:
		return nil, errors.Errorf("Method %q not supported", req.Method)

	}

	return nil, nil
}

var isSchema = regexp.MustCompile(`type (.+?) \{`)

func parse(uri, a string) (interface{}, error) {

	found := isSchema.FindAllString(a, -1)

	if len(found) == 0 {
		// its query => Name : filename , Input : File content
		return parser.ParseQuery(&ast.Source{Name: uri, Input: a})
	}
	// parse it as a schema

	return parser.ParseSchema(&ast.Source{Name: uri, Input: a})

}
