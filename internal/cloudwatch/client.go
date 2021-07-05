package cloudwatch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

type Client struct {
	CwClient cloudwatchlogs.Client
}

type QueryDefinition = types.QueryDefinition

func NewClient() (*Client, error) {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	client := cloudwatchlogs.NewFromConfig(cfg)

	return &Client{
		CwClient: *client,
	}, nil
}

func (client *Client) DescribeQueryDefinitions(queryDefinitionNamePrefix string) ([]QueryDefinition, error) {
	describeInput := &cloudwatchlogs.DescribeQueryDefinitionsInput{}
	if queryDefinitionNamePrefix != "" {
		describeInput.QueryDefinitionNamePrefix = &queryDefinitionNamePrefix
	}

	queryDefinitions := make([]QueryDefinition, 0)

	for {
		output, err := client.CwClient.DescribeQueryDefinitions(context.TODO(), describeInput)
		if err != nil {
			return nil, err
		}

		queryDefinitions = append(queryDefinitions, output.QueryDefinitions...)
		if output.NextToken == nil {
			return queryDefinitions, nil
		}

		describeInput.NextToken = output.NextToken
	}
}

func (client *Client) PutQueryDefinitions(queryDefinitions []QueryDefinition) error {

	for _, queryDefinition := range queryDefinitions {

		_, err := client.CwClient.PutQueryDefinition(context.TODO(), &cloudwatchlogs.PutQueryDefinitionInput{
			Name:              queryDefinition.Name,
			QueryString:       queryDefinition.QueryString,
			LogGroupNames:     queryDefinition.LogGroupNames,
			QueryDefinitionId: queryDefinition.QueryDefinitionId,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
