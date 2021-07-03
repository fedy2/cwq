package cloudwatch

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

type Client struct {
	CwClient cloudwatchlogs.Client
}

type QueryDefinition = types.QueryDefinition

func NewClient() *Client {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println(err)
	}

	client := cloudwatchlogs.NewFromConfig(cfg)

	return &Client{
		CwClient: *client,
	}
}

func (client *Client) DescribeQueryDefinitions() ([]QueryDefinition, error) {
	// TODO: handle multi page
	// TODO: enable passing QueryDefinitionNamePrefix option
	output, err := client.CwClient.DescribeQueryDefinitions(context.TODO(), &cloudwatchlogs.DescribeQueryDefinitionsInput{})
	return output.QueryDefinitions, err
}
