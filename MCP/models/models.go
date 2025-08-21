package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// CreateClusterResponse represents the CreateClusterResponse schema from the OpenAPI specification
type CreateClusterResponse struct {
	Cluster interface{} `json:"cluster,omitempty"`
}

// AssociateEncryptionConfigResponse represents the AssociateEncryptionConfigResponse schema from the OpenAPI specification
type AssociateEncryptionConfigResponse struct {
	Update Update `json:"update,omitempty"` // An object representing an asynchronous update.
}

// UpdateClusterVersionResponse represents the UpdateClusterVersionResponse schema from the OpenAPI specification
type UpdateClusterVersionResponse struct {
	Update interface{} `json:"update,omitempty"`
}

// OIDC represents the OIDC schema from the OpenAPI specification
type OIDC struct {
	Issuer interface{} `json:"issuer,omitempty"`
}

// ListAddonsResponse represents the ListAddonsResponse schema from the OpenAPI specification
type ListAddonsResponse struct {
	Addons interface{} `json:"addons,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// RegisterClusterResponse represents the RegisterClusterResponse schema from the OpenAPI specification
type RegisterClusterResponse struct {
	Cluster Cluster `json:"cluster,omitempty"` // An object representing an Amazon EKS cluster.
}

// DescribeNodegroupRequest represents the DescribeNodegroupRequest schema from the OpenAPI specification
type DescribeNodegroupRequest struct {
}

// NodegroupHealth represents the NodegroupHealth schema from the OpenAPI specification
type NodegroupHealth struct {
	Issues interface{} `json:"issues,omitempty"`
}

// UpdateClusterConfigResponse represents the UpdateClusterConfigResponse schema from the OpenAPI specification
type UpdateClusterConfigResponse struct {
	Update Update `json:"update,omitempty"` // An object representing an asynchronous update.
}

// AddonIssue represents the AddonIssue schema from the OpenAPI specification
type AddonIssue struct {
	Code interface{} `json:"code,omitempty"`
	Message interface{} `json:"message,omitempty"`
	Resourceids interface{} `json:"resourceIds,omitempty"`
}

// ListIdentityProviderConfigsRequest represents the ListIdentityProviderConfigsRequest schema from the OpenAPI specification
type ListIdentityProviderConfigsRequest struct {
}

// DeleteNodegroupRequest represents the DeleteNodegroupRequest schema from the OpenAPI specification
type DeleteNodegroupRequest struct {
}

// ControlPlanePlacementRequest represents the ControlPlanePlacementRequest schema from the OpenAPI specification
type ControlPlanePlacementRequest struct {
	Groupname interface{} `json:"groupName,omitempty"`
}

// DescribeUpdateRequest represents the DescribeUpdateRequest schema from the OpenAPI specification
type DescribeUpdateRequest struct {
}

// DescribeAddonVersionsRequest represents the DescribeAddonVersionsRequest schema from the OpenAPI specification
type DescribeAddonVersionsRequest struct {
}

// RegisterClusterRequest represents the RegisterClusterRequest schema from the OpenAPI specification
type RegisterClusterRequest struct {
	Connectorconfig interface{} `json:"connectorConfig"`
	Name interface{} `json:"name"`
	Tags interface{} `json:"tags,omitempty"`
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
}

// ListUpdatesRequest represents the ListUpdatesRequest schema from the OpenAPI specification
type ListUpdatesRequest struct {
}

// FargateProfile represents the FargateProfile schema from the OpenAPI specification
type FargateProfile struct {
	Tags interface{} `json:"tags,omitempty"`
	Podexecutionrolearn interface{} `json:"podExecutionRoleArn,omitempty"`
	Subnets interface{} `json:"subnets,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Fargateprofilearn interface{} `json:"fargateProfileArn,omitempty"`
	Fargateprofilename interface{} `json:"fargateProfileName,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Clustername interface{} `json:"clusterName,omitempty"`
	Selectors interface{} `json:"selectors,omitempty"`
}

// CreateClusterRequest represents the CreateClusterRequest schema from the OpenAPI specification
type CreateClusterRequest struct {
	Version interface{} `json:"version,omitempty"`
	Name interface{} `json:"name"`
	Resourcesvpcconfig interface{} `json:"resourcesVpcConfig"`
	Logging interface{} `json:"logging,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Outpostconfig interface{} `json:"outpostConfig,omitempty"`
	Rolearn interface{} `json:"roleArn"`
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Encryptionconfig interface{} `json:"encryptionConfig,omitempty"`
	Kubernetesnetworkconfig interface{} `json:"kubernetesNetworkConfig,omitempty"`
}

// DescribeClusterRequest represents the DescribeClusterRequest schema from the OpenAPI specification
type DescribeClusterRequest struct {
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Tags interface{} `json:"tags"`
}

// DescribeFargateProfileRequest represents the DescribeFargateProfileRequest schema from the OpenAPI specification
type DescribeFargateProfileRequest struct {
}

// UpdateNodegroupConfigResponse represents the UpdateNodegroupConfigResponse schema from the OpenAPI specification
type UpdateNodegroupConfigResponse struct {
	Update Update `json:"update,omitempty"` // An object representing an asynchronous update.
}

// IdentityProviderConfigResponse represents the IdentityProviderConfigResponse schema from the OpenAPI specification
type IdentityProviderConfigResponse struct {
	Oidc interface{} `json:"oidc,omitempty"`
}

// DescribeAddonRequest represents the DescribeAddonRequest schema from the OpenAPI specification
type DescribeAddonRequest struct {
}

// UpdateNodegroupConfigRequest represents the UpdateNodegroupConfigRequest schema from the OpenAPI specification
type UpdateNodegroupConfigRequest struct {
	Scalingconfig interface{} `json:"scalingConfig,omitempty"`
	Taints interface{} `json:"taints,omitempty"`
	Updateconfig interface{} `json:"updateConfig,omitempty"`
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Labels interface{} `json:"labels,omitempty"`
}

// EncryptionConfig represents the EncryptionConfig schema from the OpenAPI specification
type EncryptionConfig struct {
	Provider interface{} `json:"provider,omitempty"`
	Resources interface{} `json:"resources,omitempty"`
}

// ConnectorConfigResponse represents the ConnectorConfigResponse schema from the OpenAPI specification
type ConnectorConfigResponse struct {
	Activationcode interface{} `json:"activationCode,omitempty"`
	Activationexpiry interface{} `json:"activationExpiry,omitempty"`
	Activationid interface{} `json:"activationId,omitempty"`
	Provider interface{} `json:"provider,omitempty"`
	Rolearn interface{} `json:"roleArn,omitempty"`
}

// AddonVersionInfo represents the AddonVersionInfo schema from the OpenAPI specification
type AddonVersionInfo struct {
	Addonversion interface{} `json:"addonVersion,omitempty"`
	Architecture interface{} `json:"architecture,omitempty"`
	Compatibilities interface{} `json:"compatibilities,omitempty"`
	Requiresconfiguration interface{} `json:"requiresConfiguration,omitempty"`
}

// NodegroupUpdateConfig represents the NodegroupUpdateConfig schema from the OpenAPI specification
type NodegroupUpdateConfig struct {
	Maxunavailable interface{} `json:"maxUnavailable,omitempty"`
	Maxunavailablepercentage interface{} `json:"maxUnavailablePercentage,omitempty"`
}

// ListNodegroupsResponse represents the ListNodegroupsResponse schema from the OpenAPI specification
type ListNodegroupsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Nodegroups interface{} `json:"nodegroups,omitempty"`
}

// OidcIdentityProviderConfigRequest represents the OidcIdentityProviderConfigRequest schema from the OpenAPI specification
type OidcIdentityProviderConfigRequest struct {
	Groupsclaim interface{} `json:"groupsClaim,omitempty"`
	Groupsprefix interface{} `json:"groupsPrefix,omitempty"`
	Identityproviderconfigname interface{} `json:"identityProviderConfigName"`
	Issuerurl interface{} `json:"issuerUrl"`
	Requiredclaims interface{} `json:"requiredClaims,omitempty"`
	Usernameclaim interface{} `json:"usernameClaim,omitempty"`
	Usernameprefix interface{} `json:"usernamePrefix,omitempty"`
	Clientid interface{} `json:"clientId"`
}

// Certificate represents the Certificate schema from the OpenAPI specification
type Certificate struct {
	Data interface{} `json:"data,omitempty"`
}

// AssociateEncryptionConfigRequest represents the AssociateEncryptionConfigRequest schema from the OpenAPI specification
type AssociateEncryptionConfigRequest struct {
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Encryptionconfig interface{} `json:"encryptionConfig"`
}

// KubernetesNetworkConfigRequest represents the KubernetesNetworkConfigRequest schema from the OpenAPI specification
type KubernetesNetworkConfigRequest struct {
	Ipfamily interface{} `json:"ipFamily,omitempty"`
	Serviceipv4cidr interface{} `json:"serviceIpv4Cidr,omitempty"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// Nodegroup represents the Nodegroup schema from the OpenAPI specification
type Nodegroup struct {
	Createdat interface{} `json:"createdAt,omitempty"`
	Instancetypes interface{} `json:"instanceTypes,omitempty"`
	Remoteaccess interface{} `json:"remoteAccess,omitempty"`
	Disksize interface{} `json:"diskSize,omitempty"`
	Capacitytype interface{} `json:"capacityType,omitempty"`
	Health interface{} `json:"health,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Nodegrouparn interface{} `json:"nodegroupArn,omitempty"`
	Modifiedat interface{} `json:"modifiedAt,omitempty"`
	Scalingconfig interface{} `json:"scalingConfig,omitempty"`
	Nodegroupname interface{} `json:"nodegroupName,omitempty"`
	Releaseversion interface{} `json:"releaseVersion,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Launchtemplate interface{} `json:"launchTemplate,omitempty"`
	Amitype interface{} `json:"amiType,omitempty"`
	Noderole interface{} `json:"nodeRole,omitempty"`
	Taints interface{} `json:"taints,omitempty"`
	Resources interface{} `json:"resources,omitempty"`
	Subnets interface{} `json:"subnets,omitempty"`
	Labels interface{} `json:"labels,omitempty"`
	Version interface{} `json:"version,omitempty"`
	Updateconfig interface{} `json:"updateConfig,omitempty"`
	Clustername interface{} `json:"clusterName,omitempty"`
}

// DeleteNodegroupResponse represents the DeleteNodegroupResponse schema from the OpenAPI specification
type DeleteNodegroupResponse struct {
	Nodegroup interface{} `json:"nodegroup,omitempty"`
}

// CreateAddonRequest represents the CreateAddonRequest schema from the OpenAPI specification
type CreateAddonRequest struct {
	Resolveconflicts interface{} `json:"resolveConflicts,omitempty"`
	Serviceaccountrolearn interface{} `json:"serviceAccountRoleArn,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Addonname interface{} `json:"addonName"`
	Addonversion interface{} `json:"addonVersion,omitempty"`
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Configurationvalues interface{} `json:"configurationValues,omitempty"`
}

// DescribeIdentityProviderConfigRequest represents the DescribeIdentityProviderConfigRequest schema from the OpenAPI specification
type DescribeIdentityProviderConfigRequest struct {
	Identityproviderconfig interface{} `json:"identityProviderConfig"`
}

// VpcConfigResponse represents the VpcConfigResponse schema from the OpenAPI specification
type VpcConfigResponse struct {
	Endpointprivateaccess interface{} `json:"endpointPrivateAccess,omitempty"`
	Endpointpublicaccess interface{} `json:"endpointPublicAccess,omitempty"`
	Publicaccesscidrs interface{} `json:"publicAccessCidrs,omitempty"`
	Securitygroupids interface{} `json:"securityGroupIds,omitempty"`
	Subnetids interface{} `json:"subnetIds,omitempty"`
	Vpcid interface{} `json:"vpcId,omitempty"`
	Clustersecuritygroupid interface{} `json:"clusterSecurityGroupId,omitempty"`
}

// ControlPlanePlacementResponse represents the ControlPlanePlacementResponse schema from the OpenAPI specification
type ControlPlanePlacementResponse struct {
	Groupname interface{} `json:"groupName,omitempty"`
}

// Taint represents the Taint schema from the OpenAPI specification
type Taint struct {
	Effect interface{} `json:"effect,omitempty"`
	Key interface{} `json:"key,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// CreateNodegroupRequest represents the CreateNodegroupRequest schema from the OpenAPI specification
type CreateNodegroupRequest struct {
	Labels interface{} `json:"labels,omitempty"`
	Remoteaccess interface{} `json:"remoteAccess,omitempty"`
	Releaseversion interface{} `json:"releaseVersion,omitempty"`
	Amitype interface{} `json:"amiType,omitempty"`
	Version interface{} `json:"version,omitempty"`
	Nodegroupname interface{} `json:"nodegroupName"`
	Subnets interface{} `json:"subnets"`
	Tags interface{} `json:"tags,omitempty"`
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Noderole interface{} `json:"nodeRole"`
	Scalingconfig interface{} `json:"scalingConfig,omitempty"`
	Taints interface{} `json:"taints,omitempty"`
	Disksize interface{} `json:"diskSize,omitempty"`
	Capacitytype interface{} `json:"capacityType,omitempty"`
	Instancetypes interface{} `json:"instanceTypes,omitempty"`
	Updateconfig interface{} `json:"updateConfig,omitempty"`
	Launchtemplate interface{} `json:"launchTemplate,omitempty"`
}

// DisassociateIdentityProviderConfigResponse represents the DisassociateIdentityProviderConfigResponse schema from the OpenAPI specification
type DisassociateIdentityProviderConfigResponse struct {
	Update Update `json:"update,omitempty"` // An object representing an asynchronous update.
}

// DescribeClusterResponse represents the DescribeClusterResponse schema from the OpenAPI specification
type DescribeClusterResponse struct {
	Cluster interface{} `json:"cluster,omitempty"`
}

// DescribeAddonResponse represents the DescribeAddonResponse schema from the OpenAPI specification
type DescribeAddonResponse struct {
	Addon Addon `json:"addon,omitempty"` // An Amazon EKS add-on. For more information, see <a href="https://docs.aws.amazon.com/eks/latest/userguide/eks-add-ons.html">Amazon EKS add-ons</a> in the <i>Amazon EKS User Guide</i>.
}

// AssociateIdentityProviderConfigResponse represents the AssociateIdentityProviderConfigResponse schema from the OpenAPI specification
type AssociateIdentityProviderConfigResponse struct {
	Update Update `json:"update,omitempty"` // An object representing an asynchronous update.
	Tags interface{} `json:"tags,omitempty"`
}

// Provider represents the Provider schema from the OpenAPI specification
type Provider struct {
	Keyarn interface{} `json:"keyArn,omitempty"`
}

// OutpostConfigRequest represents the OutpostConfigRequest schema from the OpenAPI specification
type OutpostConfigRequest struct {
	Controlplaneinstancetype interface{} `json:"controlPlaneInstanceType"`
	Controlplaneplacement interface{} `json:"controlPlanePlacement,omitempty"`
	Outpostarns interface{} `json:"outpostArns"`
}

// IdentityProviderConfig represents the IdentityProviderConfig schema from the OpenAPI specification
type IdentityProviderConfig struct {
	Name interface{} `json:"name"`
	TypeField interface{} `json:"type"`
}

// VpcConfigRequest represents the VpcConfigRequest schema from the OpenAPI specification
type VpcConfigRequest struct {
	Endpointprivateaccess interface{} `json:"endpointPrivateAccess,omitempty"`
	Endpointpublicaccess interface{} `json:"endpointPublicAccess,omitempty"`
	Publicaccesscidrs interface{} `json:"publicAccessCidrs,omitempty"`
	Securitygroupids interface{} `json:"securityGroupIds,omitempty"`
	Subnetids interface{} `json:"subnetIds,omitempty"`
}

// Update represents the Update schema from the OpenAPI specification
type Update struct {
	Params interface{} `json:"params,omitempty"`
	Status interface{} `json:"status,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
	Id interface{} `json:"id,omitempty"`
}

// DeleteFargateProfileResponse represents the DeleteFargateProfileResponse schema from the OpenAPI specification
type DeleteFargateProfileResponse struct {
	Fargateprofile interface{} `json:"fargateProfile,omitempty"`
}

// ListFargateProfilesResponse represents the ListFargateProfilesResponse schema from the OpenAPI specification
type ListFargateProfilesResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Fargateprofilenames interface{} `json:"fargateProfileNames,omitempty"`
}

// AssociateIdentityProviderConfigRequest represents the AssociateIdentityProviderConfigRequest schema from the OpenAPI specification
type AssociateIdentityProviderConfigRequest struct {
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Oidc interface{} `json:"oidc"`
	Tags interface{} `json:"tags,omitempty"`
}

// RemoteAccessConfig represents the RemoteAccessConfig schema from the OpenAPI specification
type RemoteAccessConfig struct {
	Sourcesecuritygroups interface{} `json:"sourceSecurityGroups,omitempty"`
	Ec2sshkey interface{} `json:"ec2SshKey,omitempty"`
}

// ClusterHealth represents the ClusterHealth schema from the OpenAPI specification
type ClusterHealth struct {
	Issues interface{} `json:"issues,omitempty"`
}

// ListNodegroupsRequest represents the ListNodegroupsRequest schema from the OpenAPI specification
type ListNodegroupsRequest struct {
}

// ListClustersResponse represents the ListClustersResponse schema from the OpenAPI specification
type ListClustersResponse struct {
	Clusters interface{} `json:"clusters,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// ClusterIssue represents the ClusterIssue schema from the OpenAPI specification
type ClusterIssue struct {
	Code interface{} `json:"code,omitempty"`
	Message interface{} `json:"message,omitempty"`
	Resourceids interface{} `json:"resourceIds,omitempty"`
}

// UpdateTaintsPayload represents the UpdateTaintsPayload schema from the OpenAPI specification
type UpdateTaintsPayload struct {
	Removetaints interface{} `json:"removeTaints,omitempty"`
	Addorupdatetaints interface{} `json:"addOrUpdateTaints,omitempty"`
}

// DeleteFargateProfileRequest represents the DeleteFargateProfileRequest schema from the OpenAPI specification
type DeleteFargateProfileRequest struct {
}

// OidcIdentityProviderConfig represents the OidcIdentityProviderConfig schema from the OpenAPI specification
type OidcIdentityProviderConfig struct {
	Usernameprefix interface{} `json:"usernamePrefix,omitempty"`
	Clientid interface{} `json:"clientId,omitempty"`
	Groupsprefix interface{} `json:"groupsPrefix,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Usernameclaim interface{} `json:"usernameClaim,omitempty"`
	Clustername interface{} `json:"clusterName,omitempty"`
	Groupsclaim interface{} `json:"groupsClaim,omitempty"`
	Identityproviderconfigarn interface{} `json:"identityProviderConfigArn,omitempty"`
	Issuerurl interface{} `json:"issuerUrl,omitempty"`
	Identityproviderconfigname interface{} `json:"identityProviderConfigName,omitempty"`
	Requiredclaims interface{} `json:"requiredClaims,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// ListIdentityProviderConfigsResponse represents the ListIdentityProviderConfigsResponse schema from the OpenAPI specification
type ListIdentityProviderConfigsResponse struct {
	Identityproviderconfigs interface{} `json:"identityProviderConfigs,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// OutpostConfigResponse represents the OutpostConfigResponse schema from the OpenAPI specification
type OutpostConfigResponse struct {
	Controlplaneinstancetype interface{} `json:"controlPlaneInstanceType"`
	Controlplaneplacement interface{} `json:"controlPlanePlacement,omitempty"`
	Outpostarns interface{} `json:"outpostArns"`
}

// AddonHealth represents the AddonHealth schema from the OpenAPI specification
type AddonHealth struct {
	Issues interface{} `json:"issues,omitempty"`
}

// UpdateAddonResponse represents the UpdateAddonResponse schema from the OpenAPI specification
type UpdateAddonResponse struct {
	Update Update `json:"update,omitempty"` // An object representing an asynchronous update.
}

// DescribeAddonConfigurationRequest represents the DescribeAddonConfigurationRequest schema from the OpenAPI specification
type DescribeAddonConfigurationRequest struct {
}

// DescribeFargateProfileResponse represents the DescribeFargateProfileResponse schema from the OpenAPI specification
type DescribeFargateProfileResponse struct {
	Fargateprofile interface{} `json:"fargateProfile,omitempty"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
}

// LogSetup represents the LogSetup schema from the OpenAPI specification
type LogSetup struct {
	Enabled interface{} `json:"enabled,omitempty"`
	Types interface{} `json:"types,omitempty"`
}

// ListUpdatesResponse represents the ListUpdatesResponse schema from the OpenAPI specification
type ListUpdatesResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Updateids interface{} `json:"updateIds,omitempty"`
}

// Cluster represents the Cluster schema from the OpenAPI specification
type Cluster struct {
	Status interface{} `json:"status,omitempty"`
	Certificateauthority interface{} `json:"certificateAuthority,omitempty"`
	Logging interface{} `json:"logging,omitempty"`
	Encryptionconfig interface{} `json:"encryptionConfig,omitempty"`
	Rolearn interface{} `json:"roleArn,omitempty"`
	Version interface{} `json:"version,omitempty"`
	Resourcesvpcconfig interface{} `json:"resourcesVpcConfig,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Endpoint interface{} `json:"endpoint,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Kubernetesnetworkconfig interface{} `json:"kubernetesNetworkConfig,omitempty"`
	Identity interface{} `json:"identity,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Connectorconfig interface{} `json:"connectorConfig,omitempty"`
	Health interface{} `json:"health,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Platformversion interface{} `json:"platformVersion,omitempty"`
	Outpostconfig interface{} `json:"outpostConfig,omitempty"`
}

// ListClustersRequest represents the ListClustersRequest schema from the OpenAPI specification
type ListClustersRequest struct {
}

// DeregisterClusterRequest represents the DeregisterClusterRequest schema from the OpenAPI specification
type DeregisterClusterRequest struct {
}

// NodegroupScalingConfig represents the NodegroupScalingConfig schema from the OpenAPI specification
type NodegroupScalingConfig struct {
	Maxsize interface{} `json:"maxSize,omitempty"`
	Minsize interface{} `json:"minSize,omitempty"`
	Desiredsize interface{} `json:"desiredSize,omitempty"`
}

// ListAddonsRequest represents the ListAddonsRequest schema from the OpenAPI specification
type ListAddonsRequest struct {
}

// DeleteClusterRequest represents the DeleteClusterRequest schema from the OpenAPI specification
type DeleteClusterRequest struct {
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
}

// ListFargateProfilesRequest represents the ListFargateProfilesRequest schema from the OpenAPI specification
type ListFargateProfilesRequest struct {
}

// ErrorDetail represents the ErrorDetail schema from the OpenAPI specification
type ErrorDetail struct {
	Errorcode interface{} `json:"errorCode,omitempty"`
	Errormessage interface{} `json:"errorMessage,omitempty"`
	Resourceids interface{} `json:"resourceIds,omitempty"`
}

// RequiredClaimsMap represents the RequiredClaimsMap schema from the OpenAPI specification
type RequiredClaimsMap struct {
}

// UpdateNodegroupVersionRequest represents the UpdateNodegroupVersionRequest schema from the OpenAPI specification
type UpdateNodegroupVersionRequest struct {
	Releaseversion interface{} `json:"releaseVersion,omitempty"`
	Version interface{} `json:"version,omitempty"`
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Force interface{} `json:"force,omitempty"`
	Launchtemplate interface{} `json:"launchTemplate,omitempty"`
}

// AutoScalingGroup represents the AutoScalingGroup schema from the OpenAPI specification
type AutoScalingGroup struct {
	Name interface{} `json:"name,omitempty"`
}

// DescribeIdentityProviderConfigResponse represents the DescribeIdentityProviderConfigResponse schema from the OpenAPI specification
type DescribeIdentityProviderConfigResponse struct {
	Identityproviderconfig interface{} `json:"identityProviderConfig,omitempty"`
}

// LabelsMap represents the LabelsMap schema from the OpenAPI specification
type LabelsMap struct {
}

// NodegroupResources represents the NodegroupResources schema from the OpenAPI specification
type NodegroupResources struct {
	Autoscalinggroups interface{} `json:"autoScalingGroups,omitempty"`
	Remoteaccesssecuritygroup interface{} `json:"remoteAccessSecurityGroup,omitempty"`
}

// DescribeAddonConfigurationResponse represents the DescribeAddonConfigurationResponse schema from the OpenAPI specification
type DescribeAddonConfigurationResponse struct {
	Addonname interface{} `json:"addonName,omitempty"`
	Addonversion interface{} `json:"addonVersion,omitempty"`
	Configurationschema interface{} `json:"configurationSchema,omitempty"`
}

// UpdateClusterVersionRequest represents the UpdateClusterVersionRequest schema from the OpenAPI specification
type UpdateClusterVersionRequest struct {
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Version interface{} `json:"version"`
}

// MarketplaceInformation represents the MarketplaceInformation schema from the OpenAPI specification
type MarketplaceInformation struct {
	Productid interface{} `json:"productId,omitempty"`
	Producturl interface{} `json:"productUrl,omitempty"`
}

// DeleteClusterResponse represents the DeleteClusterResponse schema from the OpenAPI specification
type DeleteClusterResponse struct {
	Cluster interface{} `json:"cluster,omitempty"`
}

// Issue represents the Issue schema from the OpenAPI specification
type Issue struct {
	Resourceids interface{} `json:"resourceIds,omitempty"`
	Code interface{} `json:"code,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

// Logging represents the Logging schema from the OpenAPI specification
type Logging struct {
	Clusterlogging interface{} `json:"clusterLogging,omitempty"`
}

// Compatibility represents the Compatibility schema from the OpenAPI specification
type Compatibility struct {
	Clusterversion interface{} `json:"clusterVersion,omitempty"`
	Defaultversion interface{} `json:"defaultVersion,omitempty"`
	Platformversions interface{} `json:"platformVersions,omitempty"`
}

// DeleteAddonRequest represents the DeleteAddonRequest schema from the OpenAPI specification
type DeleteAddonRequest struct {
}

// UpdateLabelsPayload represents the UpdateLabelsPayload schema from the OpenAPI specification
type UpdateLabelsPayload struct {
	Addorupdatelabels interface{} `json:"addOrUpdateLabels,omitempty"`
	Removelabels interface{} `json:"removeLabels,omitempty"`
}

// Identity represents the Identity schema from the OpenAPI specification
type Identity struct {
	Oidc interface{} `json:"oidc,omitempty"`
}

// UpdateNodegroupVersionResponse represents the UpdateNodegroupVersionResponse schema from the OpenAPI specification
type UpdateNodegroupVersionResponse struct {
	Update Update `json:"update,omitempty"` // An object representing an asynchronous update.
}

// CreateAddonResponse represents the CreateAddonResponse schema from the OpenAPI specification
type CreateAddonResponse struct {
	Addon Addon `json:"addon,omitempty"` // An Amazon EKS add-on. For more information, see <a href="https://docs.aws.amazon.com/eks/latest/userguide/eks-add-ons.html">Amazon EKS add-ons</a> in the <i>Amazon EKS User Guide</i>.
}

// CreateNodegroupResponse represents the CreateNodegroupResponse schema from the OpenAPI specification
type CreateNodegroupResponse struct {
	Nodegroup interface{} `json:"nodegroup,omitempty"`
}

// FargateProfileSelector represents the FargateProfileSelector schema from the OpenAPI specification
type FargateProfileSelector struct {
	Labels interface{} `json:"labels,omitempty"`
	Namespace interface{} `json:"namespace,omitempty"`
}

// UpdateParam represents the UpdateParam schema from the OpenAPI specification
type UpdateParam struct {
	Value interface{} `json:"value,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
}

// KubernetesNetworkConfigResponse represents the KubernetesNetworkConfigResponse schema from the OpenAPI specification
type KubernetesNetworkConfigResponse struct {
	Ipfamily interface{} `json:"ipFamily,omitempty"`
	Serviceipv4cidr interface{} `json:"serviceIpv4Cidr,omitempty"`
	Serviceipv6cidr interface{} `json:"serviceIpv6Cidr,omitempty"`
}

// CreateFargateProfileRequest represents the CreateFargateProfileRequest schema from the OpenAPI specification
type CreateFargateProfileRequest struct {
	Tags interface{} `json:"tags,omitempty"`
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Fargateprofilename interface{} `json:"fargateProfileName"`
	Podexecutionrolearn interface{} `json:"podExecutionRoleArn"`
	Selectors interface{} `json:"selectors,omitempty"`
	Subnets interface{} `json:"subnets,omitempty"`
}

// FargateProfileLabel represents the FargateProfileLabel schema from the OpenAPI specification
type FargateProfileLabel struct {
}

// DeregisterClusterResponse represents the DeregisterClusterResponse schema from the OpenAPI specification
type DeregisterClusterResponse struct {
	Cluster Cluster `json:"cluster,omitempty"` // An object representing an Amazon EKS cluster.
}

// DeleteAddonResponse represents the DeleteAddonResponse schema from the OpenAPI specification
type DeleteAddonResponse struct {
	Addon Addon `json:"addon,omitempty"` // An Amazon EKS add-on. For more information, see <a href="https://docs.aws.amazon.com/eks/latest/userguide/eks-add-ons.html">Amazon EKS add-ons</a> in the <i>Amazon EKS User Guide</i>.
}

// AddonInfo represents the AddonInfo schema from the OpenAPI specification
type AddonInfo struct {
	TypeField interface{} `json:"type,omitempty"`
	Addonname interface{} `json:"addonName,omitempty"`
	Addonversions interface{} `json:"addonVersions,omitempty"`
	Marketplaceinformation interface{} `json:"marketplaceInformation,omitempty"`
	Owner interface{} `json:"owner,omitempty"`
	Publisher interface{} `json:"publisher,omitempty"`
}

// LaunchTemplateSpecification represents the LaunchTemplateSpecification schema from the OpenAPI specification
type LaunchTemplateSpecification struct {
	Name interface{} `json:"name,omitempty"`
	Version interface{} `json:"version,omitempty"`
	Id interface{} `json:"id,omitempty"`
}

// ConnectorConfigRequest represents the ConnectorConfigRequest schema from the OpenAPI specification
type ConnectorConfigRequest struct {
	Provider interface{} `json:"provider"`
	Rolearn interface{} `json:"roleArn"`
}

// DescribeUpdateResponse represents the DescribeUpdateResponse schema from the OpenAPI specification
type DescribeUpdateResponse struct {
	Update interface{} `json:"update,omitempty"`
}

// DisassociateIdentityProviderConfigRequest represents the DisassociateIdentityProviderConfigRequest schema from the OpenAPI specification
type DisassociateIdentityProviderConfigRequest struct {
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Identityproviderconfig interface{} `json:"identityProviderConfig"`
}

// CreateFargateProfileResponse represents the CreateFargateProfileResponse schema from the OpenAPI specification
type CreateFargateProfileResponse struct {
	Fargateprofile interface{} `json:"fargateProfile,omitempty"`
}

// UpdateAddonRequest represents the UpdateAddonRequest schema from the OpenAPI specification
type UpdateAddonRequest struct {
	Addonversion interface{} `json:"addonVersion,omitempty"`
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Configurationvalues interface{} `json:"configurationValues,omitempty"`
	Resolveconflicts interface{} `json:"resolveConflicts,omitempty"`
	Serviceaccountrolearn interface{} `json:"serviceAccountRoleArn,omitempty"`
}

// UpdateClusterConfigRequest represents the UpdateClusterConfigRequest schema from the OpenAPI specification
type UpdateClusterConfigRequest struct {
	Clientrequesttoken interface{} `json:"clientRequestToken,omitempty"`
	Logging interface{} `json:"logging,omitempty"`
	Resourcesvpcconfig VpcConfigRequest `json:"resourcesVpcConfig,omitempty"` // An object representing the VPC configuration to use for an Amazon EKS cluster.
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"tags,omitempty"`
}

// DescribeNodegroupResponse represents the DescribeNodegroupResponse schema from the OpenAPI specification
type DescribeNodegroupResponse struct {
	Nodegroup interface{} `json:"nodegroup,omitempty"`
}

// DescribeAddonVersionsResponse represents the DescribeAddonVersionsResponse schema from the OpenAPI specification
type DescribeAddonVersionsResponse struct {
	Addons interface{} `json:"addons,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// TagMap represents the TagMap schema from the OpenAPI specification
type TagMap struct {
}

// Addon represents the Addon schema from the OpenAPI specification
type Addon struct {
	Owner interface{} `json:"owner,omitempty"`
	Publisher interface{} `json:"publisher,omitempty"`
	Serviceaccountrolearn interface{} `json:"serviceAccountRoleArn,omitempty"`
	Addonname interface{} `json:"addonName,omitempty"`
	Addonversion interface{} `json:"addonVersion,omitempty"`
	Configurationvalues interface{} `json:"configurationValues,omitempty"`
	Health interface{} `json:"health,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Addonarn interface{} `json:"addonArn,omitempty"`
	Clustername interface{} `json:"clusterName,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Modifiedat interface{} `json:"modifiedAt,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Marketplaceinformation interface{} `json:"marketplaceInformation,omitempty"`
}
