package fetch

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/fatih/color"
	"github.com/mmcquillan/lawsg/config"
	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

func Logs(options config.Options) {

	// initialize
	width, _ := terminal.Width()
	count := int64(0)
	loop := true
	nextToken := ""
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("ERROR: Cannot create an AWS session ", err)
		os.Exit(1)
	}
	svc := cloudwatchlogs.New(sess)

	// get stream length
	streamLen := 0
	if !options.NoStream {
		streamLen = StreamLength(options)
	}

	// log pull loop
	for loop {

		// debug
		if options.Debug {
			fmt.Println("==> Starting Request Loop")
		}

		// build basic params
		params := &cloudwatchlogs.FilterLogEventsInput{
			StartTime:    aws.Int64(options.StartTime),
			EndTime:      aws.Int64(options.EndTime),
			Interleaved:  aws.Bool(true),
			LogGroupName: aws.String(options.Group),
		}

		// set the chunk to request
		if options.Number > 0 && options.Number-count < options.Chunk {
			params.Limit = aws.Int64(options.Number - count)
		} else {
			params.Limit = aws.Int64(options.Chunk)
		}

		// set the log filter
		if options.Filter != "" {
			params.FilterPattern = aws.String(options.Filter)
		}

		// set the stream(s) filter
		if options.Stream != "" {
			streams := strings.Split(options.Stream, ",")
			params.LogStreamNames = aws.StringSlice(streams)
		}

		// set the token marker
		if nextToken != "" {
			params.NextToken = aws.String(nextToken)
		}

		// debug
		if options.Debug {
			fmt.Printf("==> Requesting %d with token %s\n", *params.Limit, nextToken)
		}

		// pull logs
		resp, err := svc.FilterLogEvents(params)
		if err != nil {
			fmt.Println("ERROR: Cannot make AWS request ", err)
			os.Exit(1)
		}

		// loop over events
		for _, event := range resp.Events {

			// init event
			msg := ""

			// handle no group
			if !options.NoGroup {
				msg += color.GreenString(options.Group) + " "
			}

			// handle no stream
			if !options.NoStream {
				s := len(*event.LogStreamName)
				msg += color.CyanString(*event.LogStreamName) + strings.Repeat(" ", streamLen-s+1)
			}

			// handle date format
			dateFormat := time.RFC3339
			if options.DateFormat != "" {
				dateFormat = options.DateFormat
			}

			// handle no time and tz
			if !options.NoTime {
				t := time.Unix(*event.Timestamp/1000, 0)
				if options.TimeZone {
					msg += fmt.Sprintf("%s ", t.Local().Format(dateFormat))
				} else {
					msg += fmt.Sprintf("%s ", t.UTC().Format(dateFormat))
				}
			}

			// add message
			if options.TrimLeft > 0 {
				m := *event.Message
				if len(m) > options.TrimLeft {
					msg += m[options.TrimLeft:]
				}
			} else {
				msg += *event.Message
			}

			// handle no wrap
			if options.NoWrap {
				if uint(len(msg)) > width {
					msg = msg[0:(width-4)] + "..."
				}
			}

			// handle green
			if options.Green != "" {
				reColor := color.New(color.BgGreen).SprintFunc()
				for _, s := range strings.Split(options.Green, ",") {
					msg = strings.Replace(msg, s, reColor(s), -1)
				}
			}

			// handle yellow
			if options.Yellow != "" {
				reColor := color.New(color.BgYellow).SprintFunc()
				for _, s := range strings.Split(options.Yellow, ",") {
					msg = strings.Replace(msg, s, reColor(s), -1)
				}
			}

			// handle red
			if options.Red != "" {
				reColor := color.New(color.BgRed).SprintFunc()
				for _, s := range strings.Split(options.Red, ",") {
					msg = strings.Replace(msg, s, reColor(s), -1)
				}
			}

			// handle no color
			if options.NoColor {
				regx := regexp.MustCompile("\\033\\[[0-9;]*m")
				msg = regx.ReplaceAllString(msg, "")
			}

			fmt.Printf("%s\n", msg)
			color.Unset()
			count++

		}

		// debug
		if options.Debug {
			fmt.Println("==> Completed Event Output")
		}

		// evaluate next steps
		if count >= options.Number {
			loop = false
		}

		// refresh next token
		if resp.NextToken != nil {
			if *resp.NextToken != nextToken {
				nextToken = *resp.NextToken
			} else {
				loop = false
			}
		}

		// handle tailing
		if options.Tail {
			if options.Debug {
				fmt.Println("==> Sleeping...")
			}
			time.Sleep(3 * 1000 * time.Millisecond)
			options.Chunk = 200
			loop = true
		}

		// debug
		if options.Debug {
			fmt.Printf("==> Completed Loop token %s - continue %t\n", nextToken, loop)
		}

	}
}
