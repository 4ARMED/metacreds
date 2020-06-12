package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/credentials"
)

// GenerateProfile creates an AWS CLI profile based on the supplied credentials
func GenerateProfile(profileName, region string, creds credentials.Value) string {
	profile := `[profile %v]
region = %v
aws_access_key_id = %v
aws_secret_access_key = %v
aws_session_token = %v`

	return fmt.Sprintf(profile, profileName, region, creds.AccessKeyID, creds.SecretAccessKey, creds.SessionToken)
}
