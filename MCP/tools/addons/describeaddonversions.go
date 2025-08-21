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

func DescribeaddonversionsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["kubernetesVersion"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("kubernetesVersion=%v", val))
		}
		if val, ok := args["maxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxResults=%v", val))
		}
		if val, ok := args["nextToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("nextToken=%v", val))
		}
		if val, ok := args["addonName"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("addonName=%v", val))
		}
		if val, ok := args["types"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("types=%v", val))
		}
		if val, ok := args["publishers"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("publishers=%v", val))
		}
		if val, ok := args["owners"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("owners=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/addons/supported-versions%s", cfg.BaseURL, queryString)
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
		var result models.DescribeAddonVersionsResponse
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

func CreateDescribeaddonversionsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_addons_supported-versions",
		mcp.WithDescription("Describes the versions for an add-on. Information such as the Kubernetes versions that you can use the add-on with, the <code>owner</code>, <code>publisher</code>, and the <code>type</code> of the add-on are returned. "),
		mcp.WithString("kubernetesVersion", mcp.Description("The Kubernetes versions that you can use the add-on with.")),
		mcp.WithNumber("maxResults", mcp.Description("The maximum number of results to return.")),
		mcp.WithString("nextToken", mcp.Description("<p>The <code>nextToken</code> value returned from a previous paginated <code>DescribeAddonVersionsRequest</code> where <code>maxResults</code> was used and the results exceeded the value of that parameter. Pagination continues from the end of the previous results that returned the <code>nextToken</code> value.</p> <note> <p>This token should be treated as an opaque identifier that is used only to retrieve the next items in a list and not for other programmatic purposes.</p> </note>")),
		mcp.WithString("addonName", mcp.Description("The name of the add-on. The name must match one of the names returned by <a href=\"https://docs.aws.amazon.com/eks/latest/APIReference/API_ListAddons.html\"> <code>ListAddons</code> </a>.")),
		mcp.WithArray("types", mcp.Description("The type of the add-on. For valid <code>types</code>, don't specify a value for this property.")),
		mcp.WithArray("publishers", mcp.Description("The publisher of the add-on. For valid <code>publishers</code>, don't specify a value for this property.")),
		mcp.WithArray("owners", mcp.Description("The owner of the add-on. For valid <code>owners</code>, don't specify a value for this property.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DescribeaddonversionsHandler(cfg),
	}
}
