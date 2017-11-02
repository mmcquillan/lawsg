package fetch

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/mmcquillan/lawsg/config"
)

func Streams(options config.Options) (streams []string) {
	block := 50
	count := block
	nextToken := ""
	sess, err := session.NewSession()
	if err != nil {
		log.Fatal("ERROR: Cannot create an AWS session ", err)
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
			log.Fatal("ERROR: Cannot make AWS request ", err)
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

func StreamLength(options config.Options) (length int) {
	streams := Streams(options)
	length = 0
	for _, s := range streams {
		l := len(s)
		if l > length {
			length = l
		}
	}
	return length
}
