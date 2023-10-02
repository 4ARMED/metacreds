package aws

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
)

// New sets up a new AWS session and provider
func New() (*ec2metadata.EC2Metadata, *ec2rolecreds.EC2RoleProvider, error) {
	s, err := session.NewSession()
	if err != nil {
		return nil, nil, err
	}
	c := ec2metadata.New(s)
	p := &ec2rolecreds.EC2RoleProvider{
		Client: c,
	}

	return c, p, nil
}

// Retrieve wraps aws-sdk-go Retrieve() credentials
func Retrieve(c *ec2metadata.EC2Metadata, p *ec2rolecreds.EC2RoleProvider) (credentials.Value, error) {
	creds, err := p.Retrieve()
	return creds, err
}

// Region wraps Region() so we can test
func Region(c *ec2metadata.EC2Metadata) (string, error) {
	region, err := c.Region()
	return region, err
}
