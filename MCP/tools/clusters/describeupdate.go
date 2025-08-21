package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/amazon-elastic-kubernetes-service/mcp-server/config"
	"github.com/amazon-elastic-kubernetes-service/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func DescribeupdateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		updateIdVal, ok := args["updateId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: updateId"), nil
		}
		updateId, ok := updateIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: updateId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["nodegroupName"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("nodegroupName=%v", val))
		}
		if val, ok := args["addonName"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("addonName=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/clusters/%s/updates/%s%s", cfg.BaseURL, name, updateId, queryString)
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
		var result models.DescribeUpdateResponse
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

func CreateDescribeupdateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_clusters_name_updates_updateId",
		mcp.WithDescription("<p>Returns descriptive information about an update against your Amazon EKS cluster or associated managed node group or Amazon EKS add-on.</p> <p>When the status of the update is <code>Succeeded</code>, the update is complete. If an update fails, the status is <code>Failed</code>, and an error detail explains the reason for the failure.</p>"),
		mcp.WithString("name", mcp.Required(), mcp.Description("The name of the Amazon EKS cluster associated with the update.")),
		mcp.WithString("updateId", mcp.Required(), mcp.Description("The ID of the update to describe.")),
		mcp.WithString("nodegroupName", mcp.Description("The name of the Amazon EKS node group associated with the update. This parameter is required if the update is a node group update.")),
		mcp.WithString("addonName", mcp.Description("The name of the add-on. The name must match one of the names returned by <a href=\"https://docs.aws.amazon.com/eks/latest/APIReference/API_ListAddons.html\"> <code>ListAddons</code> </a>. This parameter is required if the update is an add-on update.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DescribeupdateHandler(cfg),
	}
}
