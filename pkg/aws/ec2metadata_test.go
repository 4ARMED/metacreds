package aws_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	metacredsaws "github.com/4armed/metacreds/pkg/aws"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/awstesting/unit"
)

type CredsResp struct {
	Code            string
	Type            string
	AccessKeyId     string
	SecretAccessKey string
	Token           string
	Expiration      string
	LastUpdated     string
}

var credsResp = CredsResp{
	Code:            "Success",
	Type:            "AWS-HMAC",
	AccessKeyId:     "accessKey",
	SecretAccessKey: "secret",
	Token:           "token",
	LastUpdated:     "2009-11-23T0:00:00Z",
}

const instanceIdentityDocument = `{
  "devpayProductCodes" : null,
  "marketplaceProductCodes" : [ "1abc2defghijklm3nopqrs4tu" ], 
  "availabilityZone" : "us-east-1d",
  "privateIp" : "10.158.112.84",
  "version" : "2010-08-31",
  "region" : "us-east-1",
  "instanceId" : "i-1234567890abcdef0",
  "billingProducts" : null,
  "instanceType" : "t1.micro",
  "accountId" : "123456789012",
  "pendingTime" : "2015-11-19T16:32:11Z",
  "imageId" : "ami-5fb8c835",
  "kernelId" : "aki-919dcaf8",
  "ramdiskId" : null,
  "architecture" : "x86_64"
}`

func initTestServer(expireOn string) *httptest.Server {
	credsResp.Expiration = expireOn
	credsRespJson, err := json.Marshal(credsResp)
	if err != nil {
		fmt.Printf("[!] that went wrong: %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/latest/meta-data/iam/security-credentials/":
			fmt.Fprintln(w, "RoleName")
		case "/latest/meta-data/iam/security-credentials/RoleName":
			fmt.Fprint(w, string(credsRespJson))
		case "/latest/dynamic/instance-identity/document":
			fmt.Fprint(w, instanceIdentityDocument)
		default:
			http.Error(w, "Not found", http.StatusNotFound)
		}
	}))

	return server
}

func TestRetrieve(t *testing.T) {
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

	if creds.AccessKeyID != credsResp.AccessKeyId {
		t.Errorf("Retrieve: expected %v, actual %v", credsResp.AccessKeyId, creds.AccessKeyID)
	}

	if creds.SecretAccessKey != credsResp.SecretAccessKey {
		t.Errorf("Retrieve: expected %v, actual %v", credsResp.SecretAccessKey, creds.SecretAccessKey)
	}

	if creds.SessionToken != credsResp.Token {
		t.Errorf("Retrieve: expected %v, actual %v", credsResp.Token, creds.SessionToken)
	}

}

func TestRegion(t *testing.T) {
	server := initTestServer("2014-12-16T01:51:37Z")
	defer server.Close()

	c := ec2metadata.New(unit.Session, &aws.Config{Endpoint: aws.String(server.URL + "/latest")})
	region, err := c.Region()
	if err != nil {
		t.Errorf("%v", err)
	}

	if region != "us-east-1" {
		t.Errorf("Region: expected %v, actual %v", "us-east-1", region)
	}

}
