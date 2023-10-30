package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	metacredsaws "github.com/4armed/metacreds/pkg/aws"
)

func main() {

	var creds metacredsaws.Creds

	stdin := flag.Bool("s", false, "process STDIN instead of metadata service")
	region := flag.String("region", "eu-west-1", "AWS region")

	flag.Parse()

	if *stdin {
		err := json.NewDecoder(os.Stdin).Decode(&creds)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		c, p, err := metacredsaws.New()
		if err != nil {
			fmt.Printf("New session error: %v", err)
			return
		}
		metadataCreds, err := metacredsaws.Retrieve(c, p)
		if err != nil {
			fmt.Printf("Retrieve creds error: %v", err)
			return
		}

		awsRegion, _ := c.Region()
		region = &awsRegion

		creds = metacredsaws.Creds{
			AccessKeyID:     metadataCreds.AccessKeyID,
			SecretAccessKey: metadataCreds.SecretAccessKey,
			SessionToken:    metadataCreds.SessionToken,
		}
	}

	profile := metacredsaws.GenerateProfile("pentest", *region, creds)
	fmt.Printf("%v\n", profile)
}
