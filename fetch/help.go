package fetch

import (
	"fmt"
	"strings"

	"github.com/mmcquillan/lawsg/config"
)

// Help - display the help
func Help(options config.Options) {

	if options.Group == "" {
		options.Group = "less"
	}

	fmt.Println("")
	fmt.Println("LAWSG - The AWS Cloudwatch Logs Viewer")
	fmt.Println("")

	if oneOf(options.Group, "less") {
		fmt.Println("USAGE:")
		fmt.Println("  lawsg <command> <group> [options]")
		fmt.Println("")
	}

	if oneOf(options.Group, "usage,more") {
		fmt.Println("USAGE:")
		fmt.Println("  lawsg help [ more | <topic name> ]")
		fmt.Println("  lawsg groups")
		fmt.Println("  lawsg streams <group name>")
		fmt.Println("  lawsg get <group name> [options]")
		fmt.Println("")
	}

	if oneOf(options.Group, "filter,less,more") {
		fmt.Println("FILTER OPTIONS:")
		fmt.Println("  -c --command    Command to run groups, streams, get, help (or first argument)")
		fmt.Println("  -g --group      Group for the command (or second argument)")
		fmt.Println("  -f --filter     Coudwatch Filter for Event Logs")
		fmt.Println("  -m --stream     Comman delimited list of Streams")
		fmt.Println("  -s --starttime  Start Time for the Event Logs [default: 10 min before now]")
		fmt.Println("  -e --endtime    End Time for the Event Logs [default: now]")
		fmt.Println("  -n --number     Number of Log Events to show")
		fmt.Println("  -t --tail       Active tailing of Event Logs (experimental)")
		fmt.Println("")
	}

	if oneOf(options.Group, "display,less,more") {
		fmt.Println("DISPLAY OPTIONS:")
		fmt.Println("     --tz         Convert Event Log display to local time")
		fmt.Println("     --spacing    Adds spacing between Log Events")
		fmt.Println("     --ng         Display No Group column")
		fmt.Println("     --ns         Display No Stream column")
		fmt.Println("     --nt         Display No Time column")
		fmt.Println("     --nc         Display No Color")
		fmt.Println("     --nw         Display No Wrapping of lines (truncates)")
		fmt.Println("     --trimleft   Trims Left side of Event Message")
		fmt.Println("     --dateformat The Date Format to use for display")
		fmt.Println("     --green      Comma delimited Words to highlight Green")
		fmt.Println("     --yellow     Comma delimited Words to highlight Yellow")
		fmt.Println("     --red        Comma delimited Words to highlight Red")
		fmt.Println("")
	}

	if oneOf(options.Group, "less") {
		fmt.Println("lawsg help more - to see more help")
	}

	if oneOf(options.Group, "advanced,more") {
		fmt.Println("ADVANCED OPTIONS:")
		fmt.Println("     --chunk      Chunk size for retreiving Event Logs [default: 10000]")
		fmt.Println("     --debug      Debug of Output")
		fmt.Println("")
	}

	if oneOf(options.Group, "env,environment,more") {
		fmt.Println("ENVIRONMENT VARIABLES:")
		fmt.Println("  LAWSG_COMMAND      Command to run groups, streams, get, help (or first argument)")
		fmt.Println("  LAWSG_GROUP        Group for the command (or second argument)")
		fmt.Println("  LAWSG_FILTER       Coudwatch Filter for Event Logs")
		fmt.Println("  LAWSG_STREAM       Comman delimited list of Streams")
		fmt.Println("  LAWSG_START_TIME   Start Time for the Event Logs [default: 10 min before now]")
		fmt.Println("  LAWSG_END_TIME     End Time for the Event Logs [default: now]")
		fmt.Println("  LAWSG_NUMBER       Number of Log Events to show")
		fmt.Println("  LAWSG_TAIL         Active tailing of Event Logs (experimental)")
		fmt.Println("  LAWSG_TIMEZONE     Convert Event Log display to local time")
		fmt.Println("  LAWSG_SPACING      Adds spacing between Log Events")
		fmt.Println("  LAWSG_NO_GROUP     Display No Group column")
		fmt.Println("  LAWSG_NO_STREAM    Display No Stream column")
		fmt.Println("  LAWSG_NO_TIME      Display No Time column")
		fmt.Println("  LAWSG_NO_COLOR     Display No Color")
		fmt.Println("  LAWSG_NO_WRAP      Display No Wrapping of lines (truncates)")
		fmt.Println("  LAWSG_TRIM_LEFT    Trims Left side of Event Message")
		fmt.Println("  LAWSG_DATE_FORMAT  The Date Format to use for display")
		fmt.Println("  LAWSG_GREEN        Comma delimited Words to highlight Green")
		fmt.Println("  LAWSG_YELLOW       Comma delimited Words to highlight Yellow")
		fmt.Println("  LAWSG_RED          Comma delimited Words to highlight Red")
		fmt.Println("  LAWSG_CHUNK        Chunk size for retreiving Event Logs [default: 10000]")
		fmt.Println("  LAWSG_DEBUG        Debug of Output")
		fmt.Println("")
	}

	if oneOf(options.Group, "date,datetime,time,more") {
		fmt.Println("DATETIME EXAMPLES:")
		fmt.Println("  'now' 'n' current time")
		fmt.Println("  '12s ago' twelve seconds ago")
		fmt.Println("  '12m ago' twelve minutes ago")
		fmt.Println("  '12h ago' twelve hours ago")
		fmt.Println("  '12d ago' twelve days ago")
		fmt.Println("  '12w ago' twelve weeks ago")
		fmt.Println("  '12y ago' twelve years ago")
		fmt.Println("  ...and most every other parsable date/time")
		fmt.Println("")
	}

	if oneOf(options.Group, "aws,more") {
		fmt.Println("AWS CREDENTIALS:")
		fmt.Println("  Normal AWS Env Vars are used")
		fmt.Println("  AWS_REGION")
		fmt.Println("  AWS_ACCESS_KEY_ID")
		fmt.Println("  AWS_SECRET_ACCESS_KEY")
		fmt.Println("")
	}

	if oneOf(options.Group, "about,more") {
		fmt.Println("ABOUT:")
		fmt.Println("  Free to Use and Contribute via the MIT License")
		fmt.Println("  Maintained by Matt McQuillan")
		fmt.Println("  https://github.com/mmcquillan/lawsg")
		fmt.Println("")
	}

}

func oneOf(needle string, haystack string) bool {
	for _, h := range strings.Split(haystack, ",") {
		if strings.ToUpper(needle) == strings.ToUpper(h) {
			return true
		}
	}
	return false
}
