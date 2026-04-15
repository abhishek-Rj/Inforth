package functions

import (
	"context"
	"os/exec"

	"github.com/abhishek-Rj/Inforth/config"
	"github.com/abhishek-Rj/Inforth/schema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func ListProjects(ctx context.Context, req *mcp.CallToolRequest, input schema.ListProjectsRequest) (*mcp.CallToolResult, any, error) {
	list, err := exec.Command("ls", "-l", config.App.MainDir).Output()
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
			Content: []mcp.Content{
					&mcp.TextContent{Text: string(list)},
			},
	}, nil, nil
}