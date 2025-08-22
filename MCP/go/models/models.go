package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ResourceTagMapping represents the ResourceTagMapping schema from the OpenAPI specification
type ResourceTagMapping struct {
	Resourcearn interface{} `json:"ResourceARN,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Compliancedetails interface{} `json:"ComplianceDetails,omitempty"`
}

// TagFilter represents the TagFilter schema from the OpenAPI specification
type TagFilter struct {
	Key interface{} `json:"Key,omitempty"`
	Values interface{} `json:"Values,omitempty"`
}

// FailedResourcesMap represents the FailedResourcesMap schema from the OpenAPI specification
type FailedResourcesMap struct {
}

// TagMap represents the TagMap schema from the OpenAPI specification
type TagMap struct {
}

// UntagResourcesInput represents the UntagResourcesInput schema from the OpenAPI specification
type UntagResourcesInput struct {
	Resourcearnlist interface{} `json:"ResourceARNList"`
	Tagkeys interface{} `json:"TagKeys"`
}

// Summary represents the Summary schema from the OpenAPI specification
type Summary struct {
	Targetidtype interface{} `json:"TargetIdType,omitempty"`
	Lastupdated interface{} `json:"LastUpdated,omitempty"`
	Noncompliantresources interface{} `json:"NonCompliantResources,omitempty"`
	Region interface{} `json:"Region,omitempty"`
	Resourcetype interface{} `json:"ResourceType,omitempty"`
	Targetid interface{} `json:"TargetId,omitempty"`
}

// StartReportCreationInput represents the StartReportCreationInput schema from the OpenAPI specification
type StartReportCreationInput struct {
	S3bucket interface{} `json:"S3Bucket"`
}

// FailureInfo represents the FailureInfo schema from the OpenAPI specification
type FailureInfo struct {
	Statuscode interface{} `json:"StatusCode,omitempty"`
	Errorcode interface{} `json:"ErrorCode,omitempty"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
}

// TagResourcesOutput represents the TagResourcesOutput schema from the OpenAPI specification
type TagResourcesOutput struct {
	Failedresourcesmap interface{} `json:"FailedResourcesMap,omitempty"`
}

// GetTagValuesInput represents the GetTagValuesInput schema from the OpenAPI specification
type GetTagValuesInput struct {
	Key interface{} `json:"Key"`
	Paginationtoken interface{} `json:"PaginationToken,omitempty"`
}

// GetTagKeysInput represents the GetTagKeysInput schema from the OpenAPI specification
type GetTagKeysInput struct {
	Paginationtoken interface{} `json:"PaginationToken,omitempty"`
}

// GetResourcesOutput represents the GetResourcesOutput schema from the OpenAPI specification
type GetResourcesOutput struct {
	Paginationtoken interface{} `json:"PaginationToken,omitempty"`
	Resourcetagmappinglist interface{} `json:"ResourceTagMappingList,omitempty"`
}

// GetTagKeysOutput represents the GetTagKeysOutput schema from the OpenAPI specification
type GetTagKeysOutput struct {
	Tagkeys interface{} `json:"TagKeys,omitempty"`
	Paginationtoken interface{} `json:"PaginationToken,omitempty"`
}

// ComplianceDetails represents the ComplianceDetails schema from the OpenAPI specification
type ComplianceDetails struct {
	Compliancestatus interface{} `json:"ComplianceStatus,omitempty"`
	Keyswithnoncompliantvalues interface{} `json:"KeysWithNoncompliantValues,omitempty"`
	Noncompliantkeys interface{} `json:"NoncompliantKeys,omitempty"`
}

// GetResourcesInput represents the GetResourcesInput schema from the OpenAPI specification
type GetResourcesInput struct {
	Paginationtoken interface{} `json:"PaginationToken,omitempty"`
	Resourcearnlist interface{} `json:"ResourceARNList,omitempty"`
	Resourcetypefilters interface{} `json:"ResourceTypeFilters,omitempty"`
	Resourcesperpage interface{} `json:"ResourcesPerPage,omitempty"`
	Tagfilters interface{} `json:"TagFilters,omitempty"`
	Tagsperpage interface{} `json:"TagsPerPage,omitempty"`
	Excludecompliantresources interface{} `json:"ExcludeCompliantResources,omitempty"`
	Includecompliancedetails interface{} `json:"IncludeComplianceDetails,omitempty"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// StartReportCreationOutput represents the StartReportCreationOutput schema from the OpenAPI specification
type StartReportCreationOutput struct {
}

// GetTagValuesOutput represents the GetTagValuesOutput schema from the OpenAPI specification
type GetTagValuesOutput struct {
	Paginationtoken interface{} `json:"PaginationToken,omitempty"`
	Tagvalues interface{} `json:"TagValues,omitempty"`
}

// DescribeReportCreationOutput represents the DescribeReportCreationOutput schema from the OpenAPI specification
type DescribeReportCreationOutput struct {
	Status interface{} `json:"Status,omitempty"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
	S3location interface{} `json:"S3Location,omitempty"`
}

// TagResourcesInput represents the TagResourcesInput schema from the OpenAPI specification
type TagResourcesInput struct {
	Resourcearnlist interface{} `json:"ResourceARNList"`
	Tags interface{} `json:"Tags"`
}

// UntagResourcesOutput represents the UntagResourcesOutput schema from the OpenAPI specification
type UntagResourcesOutput struct {
	Failedresourcesmap interface{} `json:"FailedResourcesMap,omitempty"`
}

// DescribeReportCreationInput represents the DescribeReportCreationInput schema from the OpenAPI specification
type DescribeReportCreationInput struct {
}

// GetComplianceSummaryOutput represents the GetComplianceSummaryOutput schema from the OpenAPI specification
type GetComplianceSummaryOutput struct {
	Summarylist interface{} `json:"SummaryList,omitempty"`
	Paginationtoken interface{} `json:"PaginationToken,omitempty"`
}

// GetComplianceSummaryInput represents the GetComplianceSummaryInput schema from the OpenAPI specification
type GetComplianceSummaryInput struct {
	Groupby interface{} `json:"GroupBy,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Paginationtoken interface{} `json:"PaginationToken,omitempty"`
	Regionfilters interface{} `json:"RegionFilters,omitempty"`
	Resourcetypefilters interface{} `json:"ResourceTypeFilters,omitempty"`
	Tagkeyfilters interface{} `json:"TagKeyFilters,omitempty"`
	Targetidfilters interface{} `json:"TargetIdFilters,omitempty"`
}
