package fetch

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/fatih/color"
	"github.com/mmcquillan/lawsg/config"
	"github.com/schollz/closestmatch"
)

func Groups(options config.Options) {
	groups := getGroups(options)
	for _, g := range groups {
		if options.NoColor {
			fmt.Printf("%s\n", g)
		} else {
			color.Green("%s\n", g)
		}
	}
	if options.Stats {
		stat := fmt.Sprintf("\n[ STATS: %dms exec ]\n", (time.Now().UnixNano()-options.Timer)/1000000)
		fmt.Print(stat)
	}
}

func MatchGroups(options config.Options) {
	groups := getGroups(options)
	bagSizes := []int{2, 3, 4}
	cm := closestmatch.New(groups, bagSizes)
	matches := cm.ClosestN(options.Group, 3)
	if len(matches) > 0 {
		fmt.Println("\nDid you mean...")
		for _, group := range matches {
			fmt.Printf("- %s\n", group)
		}
		fmt.Println("")
	}
}

func getGroups(options config.Options) []string {
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
	return groups
}
