package main

import (
	"fmt"

	metacredsaws "github.com/4armed/metacreds/pkg/aws"
)

func main() {

	c, p := metacredsaws.New()
	creds, err := metacredsaws.Retrieve(c, p)
	if err != nil {
		fmt.Printf("Retrieve creds error: %v", err)
	}

	region, _ := c.Region()

	profile := metacredsaws.GenerateProfile("pentest", region, creds)
	fmt.Printf("\n%v\n", profile)
}
