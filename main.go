package main

import (
	"context"
	"log"
	functions_shell_history "shell-history-mcp-server/internal/mcp/functions"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{Name: "shell_history", Version: "v1.0.0"}, nil)
	mcp.AddTool(server, &mcp.Tool{Name: "shell_history", Description: "Get shell history commands"}, functions_shell_history.GetShellHistory)

	// Run the server over stdin/stdout, until the client disconnects.
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
