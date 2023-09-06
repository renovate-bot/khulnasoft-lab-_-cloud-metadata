package iam

import "github.com/khulnasoft-lab/cloud-metadata/pkg/metadata"

var NoFolderLevelDefaultServiceAccountAssignment = metadata.Metadata{
	ID:          "AVD-GCP-0041",
	Title:       "Roles should not be assigned to default service accounts",
	Description: "Default service accounts should not be used - consider creating specialised service accounts for individual purposes.",
	Impact:      "Violation of principal of least privilege",
	Severity:    "MEDIUM",
	Links:       []string {
		"", 
	},
}

