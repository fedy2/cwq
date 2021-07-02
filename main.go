package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
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

	for _, object := range output.QueryDefinitions {
		fmt.Printf("id=%s name=%s query:%s\n", aws.ToString(object.QueryDefinitionId), aws.ToString(object.Name), aws.ToString(object.QueryString))
	}

}
