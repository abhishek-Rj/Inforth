package functions

import (
	"context"
	"os/exec"

	"github.com/abhishek-Rj/Inforth/config"
	"github.com/abhishek-Rj/Inforth/schema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func GetLatestCommit(ctx context.Context, req *mcp.CallToolRequest, input schema.GetLatestCommitRequest) (*mcp.CallToolResult, any, error) {
	cmd := exec.Command("git", "log", "-n", "10")
	cmd.Dir = config.App.MainDir + "/" + input.ProjectName
	
	opt, err := cmd.CombinedOutput()
	if err != nil {
		return nil, nil, err
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(opt)},
		},
	}, nil, nil
}