package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/amazon-elastic-kubernetes-service/mcp-server/config"
	"github.com/amazon-elastic-kubernetes-service/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CreateclusterHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/clusters", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.CreateClusterResponse
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

func CreateCreateclusterTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_clusters",
		mcp.WithDescription("<p>Creates an Amazon EKS control plane. </p> <p>The Amazon EKS control plane consists of control plane instances that run the Kubernetes software, such as <code>etcd</code> and the API server. The control plane runs in an account managed by Amazon Web Services, and the Kubernetes API is exposed by the Amazon EKS API server endpoint. Each Amazon EKS cluster control plane is single tenant and unique. It runs on its own set of Amazon EC2 instances.</p> <p>The cluster control plane is provisioned across multiple Availability Zones and fronted by an Elastic Load Balancing Network Load Balancer. Amazon EKS also provisions elastic network interfaces in your VPC subnets to provide connectivity from the control plane instances to the nodes (for example, to support <code>kubectl exec</code>, <code>logs</code>, and <code>proxy</code> data flows).</p> <p>Amazon EKS nodes run in your Amazon Web Services account and connect to your cluster's control plane over the Kubernetes API server endpoint and a certificate file that is created for your cluster.</p> <p>In most cases, it takes several minutes to create a cluster. After you create an Amazon EKS cluster, you must configure your Kubernetes tooling to communicate with the API server and launch nodes into your cluster. For more information, see <a href="https://docs.aws.amazon.com/eks/latest/userguide/managing-auth.html">Managing Cluster Authentication</a> and <a href="https://docs.aws.amazon.com/eks/latest/userguide/launch-workers.html">Launching Amazon EKS nodes</a> in the <i>Amazon EKS User Guide</i>.</p>"),
		mcp.WithObject("outpostConfig", mcp.Description("Input parameter: The configuration of your local Amazon EKS cluster on an Amazon Web Services Outpost. Before creating a cluster on an Outpost, review <a href=\"https://docs.aws.amazon.com/eks/latest/userguide/eks-outposts-local-cluster-create.html\">Creating a local cluster on an Outpost</a> in the <i>Amazon EKS User Guide</i>. This API isn't available for Amazon EKS clusters on the Amazon Web Services cloud.")),
		mcp.WithObject("resourcesVpcConfig", mcp.Required(), mcp.Description("Input parameter: An object representing the VPC configuration to use for an Amazon EKS cluster.")),
		mcp.WithObject("tags", mcp.Description("Input parameter: The metadata to apply to the cluster to assist with categorization and organization. Each tag consists of a key and an optional value. You define both.")),
		mcp.WithString("clientRequestToken", mcp.Description("Input parameter: Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")),
		mcp.WithArray("encryptionConfig", mcp.Description("Input parameter: The encryption configuration for the cluster.")),
		mcp.WithString("version", mcp.Description("Input parameter: <p>The desired Kubernetes version for your cluster. If you don't specify a value here, the default version available in Amazon EKS is used.</p> <note> <p>The default version might not be the latest version available.</p> </note>")),
		mcp.WithObject("logging", mcp.Description("Input parameter: An object representing the logging configuration for resources in your cluster.")),
		mcp.WithString("roleArn", mcp.Required(), mcp.Description("Input parameter: The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to Amazon Web Services API operations on your behalf. For more information, see <a href=\"https://docs.aws.amazon.com/eks/latest/userguide/service_IAM_role.html\">Amazon EKS Service IAM Role</a> in the <i> <i>Amazon EKS User Guide</i> </i>.")),
		mcp.WithObject("kubernetesNetworkConfig", mcp.Description("Input parameter: The Kubernetes network configuration for the cluster.")),
		mcp.WithString("name", mcp.Required(), mcp.Description("Input parameter: The unique name to give to your cluster.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreateclusterHandler(cfg),
	}
}
