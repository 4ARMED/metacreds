package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
)

func main() {
	creds := credentials.NewCredentials(&ec2rolecreds.EC2RoleProvider{})

	// Retrieve the credentials value
	credValue, err := creds.Get()
	if err != nil {
		fmt.Printf(`Could not retrieve creds: %v`, err)
		return
	}

	fmt.Printf(`credValue: %v`, credValue)
}
