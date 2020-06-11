package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	c := ec2metadata.New(session.New())
	p := &ec2rolecreds.EC2RoleProvider{
		Client: c,
	}

	creds, err := p.Retrieve()
	if err != nil {
		fmt.Printf("Retrieve creds error: %v", err)
	}

	region, _ := c.Region()

	fmt.Printf("export AWS_REGION=%v", region)
	fmt.Printf("export AWS_ACCESS_KEY_ID=%v", creds.AccessKeyID)
	fmt.Printf("export AWS_SECRET_ACCESS_KEY=%v", creds.SecretAccessKey)
	fmt.Printf("export AWS_SESSION_TOKEN=%v", creds.SessionToken)
}
