package main

import (
	"context"
	"os"

	"lsp/server"

	"github.com/sourcegraph/jsonrpc2"
)

func main() {
	handler := server.NewHandler()
	<-jsonrpc2.NewConn(context.Background(), jsonrpc2.NewBufferedStream(stdrwc{}, jsonrpc2.VSCodeObjectCodec{}), handler).DisconnectNotify()
}

type stdrwc struct{}

func (stdrwc) Read(p []byte) (int, error) {
	f, _ := os.OpenFile("C:/Users/abhij/OneDrive/Desktop/Test/input.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	f.Write(p)
	return os.Stdin.Read(p)
}

func (c stdrwc) Write(p []byte) (int, error) {
	f, _ := os.OpenFile("C:/Users/abhij/OneDrive/Desktop/Test/output.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	f.Write(p)
	return os.Stdout.Write(p)
}

func (c stdrwc) Close() error {
	if err := os.Stdin.Close(); err != nil {
		return err
	}
	return os.Stdout.Close()
}
