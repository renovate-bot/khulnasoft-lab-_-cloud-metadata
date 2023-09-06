package storage

import "github.com/khulnasoft-lab/cloud-metadata/pkg/metadata"

var EnableUbla = metadata.Metadata{
	ID:          "AVD-GCP-0064",
	Title:       "Ensure that Cloud Storage buckets have uniform bucket-level access enabled",
	Description: "When you enable uniform bucket-level access on a bucket, Access Control Lists (ACLs) are disabled, and only bucket-level Identity and Access Management (IAM) permissions grant access to that bucket and the objects it contains. You revoke all access granted by object ACLs and the ability to administrate permissions using bucket ACLs.",
	Impact:      "ACLs are difficult to manage and often lead to incorrect/unintended configurations.",
	Severity:    "MEDIUM",
	Links:       []string {
		"https://cloud.google.com/storage/docs/uniform-bucket-level-access", "https://jbrojbrojbro.medium.com/you-make-the-rules-with-authentication-controls-for-cloud-storage-53c32543747b", 
	},
}

