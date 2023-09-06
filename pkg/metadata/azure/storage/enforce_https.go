package storage

import "github.com/khulnasoft-lab/cloud-metadata/pkg/metadata"

var EnforceHttps = metadata.Metadata{
	ID:          "AVD-AZU-0049",
	Title:       "Storage accounts should be configured to only accept transfers that are over secure connections",
	Description: "You can configure your storage account to accept requests from secure connections only by setting the Secure transfer required property for the storage account. 

When you require secure transfer, any requests originating from an insecure connection are rejected. 

Microsoft recommends that you always require secure transfer for all of your storage accounts.",
	Impact:      "Insecure transfer of data into secure accounts could be read if intercepted",
	Severity:    "HIGH",
	Links:       []string {
		"https://docs.microsoft.com/en-us/azure/storage/common/storage-require-secure-transfer", 
	},
}

