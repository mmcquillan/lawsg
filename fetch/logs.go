package fetch

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/fatih/color"
	"github.com/mmcquillan/joda"
	"github.com/mmcquillan/lawsg/config"
	"github.com/mmcquillan/lawsg/util"
	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

var wg sync.WaitGroup

// Logs - Pulls Event Logs based on option filters
func Logs(options config.Options) {
	if options.Group == "*" {
		groupList := ListGroups(options)
		for _, group := range groupList {
			wg.Add(1)
			options.Group = group
			go getLogs(options)
		}
	} else if strings.Contains(options.Group, ",") {
		groupList := strings.Split(options.Group, ",")
		for _, group := range groupList {
			wg.Add(1)
			group = strings.TrimSpace(group)
			options.Group = group
			go getLogs(options)
		}
	} else {
		wg.Add(1)
		getLogs(options)
	}
	wg.Wait()
}

func getLogs(options config.Options) {

	// initialize
	w, _ := terminal.Width()
	width := int(w)
	count := int64(0)
	tdiff := int64(0)
	callcnt := 0
	streamLen := 0
	streamCnt := 0
	loop := true
	nextToken := ""
	lastTimestamp := options.EndTime

	// aws session
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("ERROR: Cannot create an AWS session ", err)
		os.Exit(1)
	}
	svc := cloudwatchlogs.New(sess)

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
			if !strings.Contains(options.Filter, "\"") {
				options.Filter = fmt.Sprintf("\"%s\"", options.Filter)
			}
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
			MatchGroups(options)
			os.Exit(1)
		}

		// get stream length
		streamCnt = len(resp.SearchedLogStreams)
		for _, stream := range resp.SearchedLogStreams {
			if len(*stream.LogStreamName) > streamLen {
				streamLen = len(*stream.LogStreamName)
			}
		}

		// loop over events
		for _, event := range resp.Events {

			// handle multi-line
			var messages []string
			if options.MultiLine {
				messages = strings.Split(*event.Message, "\n")
			} else {
				messages = make([]string, 1, 1)
				messages[0] = *event.Message
			}

			// loop over messages
			for _, message := range messages {

				// init event
				msg := ""

				// handle sort key
				if options.SortKey {
					msg += fmt.Sprintf("%d ", *event.Timestamp)
				}

				// handle no group
				if !options.NoGroup {
					msg += color.GreenString(options.Group) + " "
				}

				// handle no stream
				if !options.NoStream {
					s := streamLen - len(*event.LogStreamName) + 1
					if s < 1 {
						s = 1
					}
					sm := *event.LogStreamName + strings.Repeat(" ", s)
					if options.StreamLTrim > 0 || options.StreamRTrim > 0 {
						if options.StreamLTrim > 0 && options.StreamLTrim < len(sm) {
							sm = sm[options.StreamLTrim:]
						}
						if options.StreamRTrim > 0 && options.StreamRTrim < len(sm) {
							sm = sm[0 : len(sm)-options.StreamRTrim]
						}
					}
					msg += color.CyanString(sm)
				}

				// handle date format
				dateFormat := time.RFC3339
				if options.DateFormat != "" {
					dateFormat = joda.Format(options.DateFormat)
				}

				// handle no time and tz
				if !options.NoTime {
					t := time.Unix(*event.Timestamp/1000, 0)
					if options.TimeZone {
						msg += color.MagentaString(t.Local().Format(dateFormat)) + " "
					} else {
						msg += color.MagentaString(t.UTC().Format(dateFormat)) + " "
					}
				}

				// add message
				if options.MessageLTrim > 0 || options.MessageRTrim > 0 {
					if options.MessageLTrim > 0 && options.MessageLTrim < len(message) {
						message = message[options.MessageLTrim:]
					}
					if options.MessageRTrim > 0 && options.MessageRTrim < len(message) {
						message = message[0 : len(message)-options.MessageRTrim]
					}
				}
				msg += message

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
					msg = util.Unformat(msg)
				}

				// handle no wrap
				if options.NoWrap {
					formatWidth := 4
					cntFormat := util.CountFormat(msg) * formatWidth
					if len(msg)-(cntFormat) > width {
						offset := width - 4
						cf := util.CountFormat(msg[0:offset]) * formatWidth
						offset = offset + cf
						msg = msg[0:offset] + "..."
					}
					if options.Debug {
						msg += "\n" + strings.Repeat("---------|", int(width/10))
					}
				}

				// output and reset
				fmt.Printf("%s\n", msg)
				color.Unset()

				// handle spacing
				if options.Spacing {
					fmt.Printf("\n")
				}

				// keep track
				lastTimestamp = *event.Timestamp
				tdiff = tdiff + (*event.IngestionTime - *event.Timestamp)
				count++

			}
		}

		// debug
		if options.Debug {
			fmt.Println("==> Completed Event Output")
		}

		// evaluate next steps
		if count >= options.Number && options.Number != 0 {
			loop = false
		}

		// refresh next token
		if resp.NextToken != nil {
			if string(*resp.NextToken) != string(nextToken) {
				if options.Debug {
					fmt.Println("==> Diff Tokens")
				}
				nextToken = *resp.NextToken
			} else {
				if options.Debug {
					fmt.Println("==> Same Tokens")
				}
				loop = false
			}
		} else {
			loop = false
		}

		// handle tailing
		if options.Tail {
			if options.Debug {
				fmt.Println("==> Sleeping...")
			}
			for i := 0; i < options.Refresh; i++ {
				umsg := "Updating in " + strconv.Itoa(options.Refresh-i) + " second"
				if options.Refresh-i == 1 {
					umsg += "..."
				} else {
					umsg += "s..."
				}
				fmt.Print(umsg)
				time.Sleep(time.Second * 1)
				fmt.Print(strings.Repeat("\b", len(umsg)))
			}
			options.StartTime = lastTimestamp + 1
			options.EndTime = time.Now().Unix() * 1000
			loop = true
		}

		// debug
		if options.Debug {
			fmt.Printf("==> Completed Loop token %s - continue %t\n", nextToken, loop)
		}

		callcnt++

	}

	if options.Stats {
		stat := "\n[ STATS:"
		stat += fmt.Sprintf(" %dms exec", (time.Now().UnixNano()-options.Timer)/1000000)
		if callcnt > 1 {
			stat += fmt.Sprintf(" | %d aws calls", callcnt)
		} else {
			stat += fmt.Sprintf(" | %d aws call", callcnt)
		}
		stat += fmt.Sprintf(" | %d events", count)
		if streamCnt > 1 {
			stat += fmt.Sprintf(" | %d streams", streamCnt)
		} else {
			stat += fmt.Sprintf(" | %d stream", streamCnt)
		}
		if count > 0 {
			stat += fmt.Sprintf(" | %dms avg ingestion", (tdiff / count))
		}
		stat += " ]\n"
		fmt.Print(stat)
	}

	wg.Done()

}
