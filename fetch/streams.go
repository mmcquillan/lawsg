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
)

// Streams - Pulls back the list of streams for a given group
func Streams(options config.Options) {
	streams := getStreams(options)
	sort.Strings(streams)
	for _, s := range streams {
		if options.NoColor {
			fmt.Printf("%s\n", s)
		} else {
			color.Cyan("%s\n", s)
		}
	}
	if options.Stats {
		stat := fmt.Sprintf("\n[ STATS: %dms exec ]\n", (time.Now().UnixNano()-options.Timer)/1000000)
		fmt.Print(stat)
	}
}

func getStreams(options config.Options) (streams []string) {
	block := 50
	count := block
	nextToken := ""
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("ERROR: Cannot create an AWS session ", err)
		os.Exit(1)
	}
	svc := cloudwatchlogs.New(sess)
	for count >= block {
		params := &cloudwatchlogs.DescribeLogStreamsInput{
			Limit:        aws.Int64(int64(block)),
			LogGroupName: aws.String(options.Group),
		}
		if options.Stream == "" {
			params.OrderBy = aws.String("LastEventTime")
		} else {
			params.LogStreamNamePrefix = aws.String(options.Stream)
		}
		if nextToken != "" {
			params.NextToken = aws.String(nextToken)
		}
		resp, err := svc.DescribeLogStreams(params)
		if err != nil {
			fmt.Println("ERROR: Cannot make AWS request ", err)
			MatchGroups(options)
			os.Exit(1)
		}
		for _, s := range resp.LogStreams {
			if *s.LastEventTimestamp > options.StartTime-(1000*60*60) {
				streams = append(streams, *s.LogStreamName)
			}
		}
		count = len(resp.LogStreams)
		if count >= block {
			nextToken = *resp.NextToken
		}
	}
	return streams
}
