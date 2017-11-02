package fetch

import (
	"fmt"
	"log"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func List() {
	block := 50
	count := block
	nextToken := ""
	var groups []string
	sess, err := session.NewSession()
	if err != nil {
		log.Fatal("ERROR: Cannot create an AWS session ", err)
	}
	svc := cloudwatchlogs.New(sess)
	for count >= block {
		params := &cloudwatchlogs.DescribeLogGroupsInput{
			Limit: aws.Int64(int64(block)),
		}
		if nextToken != "" {
			params.NextToken = aws.String(nextToken)
		}
		resp, err := svc.DescribeLogGroups(params)
		if err != nil {
			log.Fatal("ERROR: Cannot make AWS request ", err)
		}
		count = len(resp.LogGroups)
		groupList := make([]string, count, count)
		for i, logGroup := range resp.LogGroups {
			groupList[i] = *logGroup.LogGroupName
		}
		groups = append(groups, groupList...)
		if count >= block {
			nextToken = *resp.NextToken
		}
	}
	sort.Strings(groups)
	for _, group := range groups {
		fmt.Printf("%s\n", group)
	}
}
