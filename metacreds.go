package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	p := &ec2rolecreds.EC2RoleProvider{
		Client: ec2metadata.New(session.New()),
	}

	creds, err := p.Retrieve()
	if err != nil {
		fmt.Printf("Retrieve creds error: %v", err)
	}

	fmt.Printf("export AWS_ACCESS_KEY_ID=%v", creds.AccessKeyID)
	fmt.Printf("export AWS_SECRET_ACCESS_KEY=%v", creds.SecretAccessKey)
	fmt.Printf("export AWS_SESSION_TOKEN=%v", creds.SessionToken)
}
