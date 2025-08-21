package main

import (
	"github.com/amazon-elastic-kubernetes-service/mcp-server/config"
	"github.com/amazon-elastic-kubernetes-service/mcp-server/models"
	tools_clusters "github.com/amazon-elastic-kubernetes-service/mcp-server/tools/clusters"
	tools_addons "github.com/amazon-elastic-kubernetes-service/mcp-server/tools/addons"
	tools_cluster_registrations "github.com/amazon-elastic-kubernetes-service/mcp-server/tools/cluster_registrations"
	tools_tags "github.com/amazon-elastic-kubernetes-service/mcp-server/tools/tags"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_clusters.CreateListnodegroupsTool(cfg),
		tools_clusters.CreateCreatenodegroupTool(cfg),
		tools_addons.CreateDescribeaddonversionsTool(cfg),
		tools_cluster_registrations.CreateRegisterclusterTool(cfg),
		tools_clusters.CreateDeletefargateprofileTool(cfg),
		tools_clusters.CreateDescribefargateprofileTool(cfg),
		tools_clusters.CreateListaddonsTool(cfg),
		tools_clusters.CreateCreateaddonTool(cfg),
		tools_clusters.CreateUpdateaddonTool(cfg),
		tools_tags.CreateUntagresourceTool(cfg),
		tools_clusters.CreateUpdatenodegroupversionTool(cfg),
		tools_clusters.CreateDeleteclusterTool(cfg),
		tools_clusters.CreateDescribeclusterTool(cfg),
		tools_clusters.CreateDeleteaddonTool(cfg),
		tools_clusters.CreateDescribeaddonTool(cfg),
		tools_addons.CreateDescribeaddonconfigurationTool(cfg),
		tools_clusters.CreateAssociateencryptionconfigTool(cfg),
		tools_clusters.CreateListfargateprofilesTool(cfg),
		tools_clusters.CreateCreatefargateprofileTool(cfg),
		tools_clusters.CreateUpdatenodegroupconfigTool(cfg),
		tools_cluster_registrations.CreateDeregisterclusterTool(cfg),
		tools_clusters.CreateDeletenodegroupTool(cfg),
		tools_clusters.CreateDescribenodegroupTool(cfg),
		tools_clusters.CreateListupdatesTool(cfg),
		tools_clusters.CreateUpdateclusterversionTool(cfg),
		tools_clusters.CreateListclustersTool(cfg),
		tools_clusters.CreateCreateclusterTool(cfg),
		tools_tags.CreateListtagsforresourceTool(cfg),
		tools_tags.CreateTagresourceTool(cfg),
		tools_clusters.CreateListidentityproviderconfigsTool(cfg),
		tools_clusters.CreateDescribeidentityproviderconfigTool(cfg),
		tools_clusters.CreateUpdateclusterconfigTool(cfg),
		tools_clusters.CreateDescribeupdateTool(cfg),
		tools_clusters.CreateAssociateidentityproviderconfigTool(cfg),
		tools_clusters.CreateDisassociateidentityproviderconfigTool(cfg),
	}
}
