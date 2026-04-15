package main

import (
	"context"
	"log"

	"github.com/abhishek-Rj/Inforth/config"
	"github.com/abhishek-Rj/Inforth/functions"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	config.Load()
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "get_project_updates",
		Version: "1.0.0",
	}, nil)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_projects",
		Description: "List all the projects inside working directory",
	}, functions.ListProjects)

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}