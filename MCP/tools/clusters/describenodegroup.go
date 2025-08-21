package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/amazon-elastic-kubernetes-service/mcp-server/config"
	"github.com/amazon-elastic-kubernetes-service/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func DescribenodegroupHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		nameVal, ok := args["name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: name"), nil
		}
		name, ok := nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: name"), nil
		}
		nodegroupNameVal, ok := args["nodegroupName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: nodegroupName"), nil
		}
		nodegroupName, ok := nodegroupNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: nodegroupName"), nil
		}
		url := fmt.Sprintf("%s/clusters/%s/node-groups/%s", cfg.BaseURL, name, nodegroupName)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.DescribeNodegroupResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateDescribenodegroupTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_clusters_name_node-groups_nodegroupName",
		mcp.WithDescription("Returns descriptive information about an Amazon EKS node group."),
		mcp.WithString("name", mcp.Required(), mcp.Description("The name of the Amazon EKS cluster associated with the node group.")),
		mcp.WithString("nodegroupName", mcp.Required(), mcp.Description("The name of the node group to describe.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DescribenodegroupHandler(cfg),
	}
}
