package aws

import (
	"fmt"
)

type Creds struct {
	// AWS Access key ID
	AccessKeyID string `json:"AccessKeyID"`

	// AWS Secret Access Key
	SecretAccessKey string `json:"SecretAccessKey"`

	// AWS Session Token
	SessionToken string `json:"Token"`
}

// GenerateProfile creates an AWS CLI profile based on the supplied credentials
func GenerateProfile(profileName string, region string, creds Creds) string {
	profile := `[profile %v]
region = %v
aws_access_key_id = %v
aws_secret_access_key = %v
aws_session_token = %v`

	return fmt.Sprintf(profile, profileName, region, creds.AccessKeyID, creds.SecretAccessKey, creds.SessionToken)
}
