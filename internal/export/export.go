package export

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fedy2/cwq/internal/cloudwatch"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

func Export() {
	client := cloudwatch.NewClient()

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
