package aws_test

import (
	"testing"

	metacredsaws "github.com/4armed/metacreds/pkg/aws"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/awstesting/unit"
)

func TestGenerateProfile(t *testing.T) {

	server := initTestServer("2014-12-16T01:51:37Z")
	defer server.Close()

	c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
	p := &ec2rolecreds.EC2RoleProvider{
		Client: c,
	}
	creds, err := metacredsaws.Retrieve(c, p)
	if err != nil {
		t.Errorf("%v", err)
	}

	profile := metacredsaws.GenerateProfile("testing", "eu-west-1", creds)
	t.Logf("profile: %v", profile)
	// TODO: Add an actual test here
}
