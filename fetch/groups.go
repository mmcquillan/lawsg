package fetch

import (
	"fmt"
	"os"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/fatih/color"
	"github.com/mmcquillan/lawsg/config"
)

func Groups(options config.Options) {
	block := 50
	count := block
	nextToken := ""
	var groups []string
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("ERROR: Cannot create an AWS session ", err)
		os.Exit(1)
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
			fmt.Println("ERROR: Cannot make AWS request ", err)
			os.Exit(1)
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
		if options.NoColor {
			fmt.Printf("%s\n", group)
		} else {
			color.Green("%s\n", group)
		}
	}
}
