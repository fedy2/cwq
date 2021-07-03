package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

func main() {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println(err)
	}

	client := cloudwatchlogs.NewFromConfig(cfg)

	output, err := client.DescribeQueryDefinitions(context.TODO(), &cloudwatchlogs.DescribeQueryDefinitionsInput{})

	if err != nil {
		fmt.Println(err)
	}

	queryDefinitionJson, err := json.MarshalIndent(output.QueryDefinitions, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(queryDefinitionJson))
}
