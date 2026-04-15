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

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_latest_commits",
		Description: "Get latest commits of all the projects inside working directory",
	}, functions.GetLatestCommit)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "search_todo",
		Description: "Search for todos in all the projects inside working directory",
	}, functions.SearchTodo)

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}