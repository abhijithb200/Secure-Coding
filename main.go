package main

import (
	"bytes"
	"context"
	"fmt"
	"lsp/server"
	"os"

	"github.com/sourcegraph/jsonrpc2"
)

func main() {
	server := &server.Server{}
	defer server.Stop()

	ctx := context.Background()
	rwc := x{}
	handler := jsonrpc2.HandlerWithError(server.Handle)
	stream := jsonrpc2.NewBufferedStream(rwc, jsonrpc2.VSCodeObjectCodec{})
	conn := jsonrpc2.NewConn(ctx, stream, handler)
	<-conn.DisconnectNotify() // channel which will close until a connection is diconnected
}

type x struct{}

func (x) Read(p []byte) (n int, err error) {
	return os.Stdin.Read(p)
}

func (x) Write(p []byte) (n int, err error) {
	return os.Stdout.Write(p)
}

func (x) Close() error {
	var errs listErr
	if err := os.Stdin.Close(); err != nil {
		errs = append(errs, err)
	}
	if err := os.Stdout.Close(); err != nil {
		errs = append(errs, err)
	}
	if len(errs) == 1 {
		return errs[0]
	}
	return errs
}

type listErr []error

func (l listErr) Error() string {
	var buf bytes.Buffer

	buf.Write([]byte("Multiple Errors Found: \n"))
	for _, e := range l {
		fmt.Fprintf(&buf, "%v\n", e) // format print and write to buf
	}
	return buf.String()
}
