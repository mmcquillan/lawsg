package fetch

import (
	"fmt"

	"github.com/mmcquillan/lawsg/config"
	"github.com/mmcquillan/lawsg/util"
)

// Help - display the help
func Help(options config.Options) {

	if options.Group == "" {
		options.Group = "less"
	}

	fmt.Println("")
	fmt.Println("LAWSG - The AWS Cloudwatch Logs Viewer")
	fmt.Println("")

	if util.Member(options.Group, "less") {
		//fmt.Println("USAGE:")
		fmt.Println("  lawsg <command> <group> [options]")
		fmt.Println("")
	}

	if util.Member(options.Group, "usage,more") {
		fmt.Println("USAGE:")
		fmt.Println("  lawsg help [ more | <topic name> ]")
		fmt.Println("  lawsg groups")
		fmt.Println("  lawsg streams <group name>")
		fmt.Println("  lawsg get <group name> [options]")
		fmt.Println("")
	}

	if util.Member(options.Group, "filter,less,more") {
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

	if util.Member(options.Group, "display,less,more") {
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

	if util.Member(options.Group, "less") {
		fmt.Println("MORE HELP: lawsg help more")
	}

	if util.Member(options.Group, "advanced,more") {
		fmt.Println("ADVANCED OPTIONS:")
		fmt.Println("     --chunk      Chunk size for retreiving Event Logs [default: 10000]")
		fmt.Println("     --refresh    Tail Refresh interval in seconds [default: 5]")
		fmt.Println("     --debug      Debug of Output")
		fmt.Println("")
	}

	if util.Member(options.Group, "env,environment,more") {
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
		fmt.Println("  LAWSG_REFRESH      Tail Refresh interval in seconds [default: 5]")
		fmt.Println("  LAWSG_DEBUG        Debug of Output")
		fmt.Println("")
	}

	if util.Member(options.Group, "date,datetime,time,more") {
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

	if util.Member(options.Group, "dateformat,more") {
		fmt.Println("DATEFORMAT EXAMPLES:")
		fmt.Println("  Two digit year              y       06")
		fmt.Println("  Four digit year             Y       2006")
		fmt.Println("  Long month name             MMMM    January")
		fmt.Println("  Short month name            MMM     Jan")
		fmt.Println("  Two digit month             MM      01")
		fmt.Println("  Single digit month          M       1")
		fmt.Println("  Two digit day of month      dd      02")
		fmt.Println("  Single digit day of month   d       2")
		fmt.Println("  Long day of week            EE      Monday")
		fmt.Println("  Short day of week           E       Mon")
		fmt.Println("  Two digit 12 hours          hh      03")
		fmt.Println("  Single digit 12 hour        h       3")
		fmt.Println("  Two digit 24 hours          H       15")
		fmt.Println("  Two digit minutes           mm      04")
		fmt.Println("  Single digit minutes        m       4")
		fmt.Println("  Two digit seconds           ss      05")
		fmt.Println("  Single digit seconds        s       5")
		fmt.Println("  Milliseconds                S       .0")
		fmt.Println("  AM/PM                       a       pm")
		fmt.Println("  Time Zone                   z       MST")
		fmt.Println("  Time Zone offset            Z       -0700")
		fmt.Println("")
	}

	if util.Member(options.Group, "aws,more") {
		fmt.Println("AWS CREDENTIALS:")
		fmt.Println("  Normal AWS Env Vars are used")
		fmt.Println("  AWS_REGION")
		fmt.Println("  AWS_ACCESS_KEY_ID")
		fmt.Println("  AWS_SECRET_ACCESS_KEY")
		fmt.Println("")
	}

	if util.Member(options.Group, "about,more") {
		fmt.Println("ABOUT:")
		fmt.Println("  Free to Use and Contribute via the MIT License")
		fmt.Println("  Maintained by Matt McQuillan")
		fmt.Println("  https://github.com/mmcquillan/lawsg")
		fmt.Println("")
	}

}
