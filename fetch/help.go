package fetch

import (
	"fmt"
	"time"

	"github.com/mmcquillan/lawsg/config"
	"github.com/mmcquillan/lawsg/util"
)

// Help - display the help
func Help(options config.Options) {

	if options.Group == "" {
		options.Group = "less"
	}

	if util.Member(options.Group, "docs") {
		fmt.Println("# LAWSG HELP")
		fmt.Println("```")
	}

	fmt.Println("")
	fmt.Println("LAWSG - The AWS Cloudwatch Logs Viewer")
	fmt.Println("")

	if util.Member(options.Group, "less") {
		fmt.Println("  lawsg <command> <group> [options]")
		fmt.Println("")
	}

	if util.Member(options.Group, "usage,more,docs") {
		fmt.Println("USAGE:")
		fmt.Println("  lawsg help [ more | <topic name> ]")
		fmt.Println("  lawsg groups [options]")
		fmt.Println("  lawsg streams <group name> [options]")
		fmt.Println("  lawsg get <group name> [options]")
		fmt.Println("")
	}

	if util.Member(options.Group, "filter,less,more,docs") {
		fmt.Println("FILTER OPTIONS:")
		fmt.Println("  -f --filter         Cloudwatch Filter for Event Logs")
		fmt.Println("  -m --stream         Comma delimited list of Streams")
		fmt.Println("  -s --starttime      Start Time for the Event Logs [default: 10 min before now]")
		fmt.Println("  -e --endtime        End Time for the Event Logs [default: now]")
		fmt.Println("  -n --number         Number of Log Events to show")
		fmt.Println("  -t --tail           Active tailing of Event Logs (experimental)")
		fmt.Println("")
	}

	if util.Member(options.Group, "display,less,more,docs") {
		fmt.Println("DISPLAY OPTIONS:")
		fmt.Println("     --tz             Convert Event Log display to local time")
		fmt.Println("     --spacing        Adds spacing between Log Events")
		fmt.Println("     --ng             Display No Group column")
		fmt.Println("     --ns             Display No Stream column")
		fmt.Println("     --nt             Display No Time column")
		fmt.Println("     --nc             Display No Color")
		fmt.Println("     --nw             Display No Wrapping of lines (truncates)")
		fmt.Println("     --stream-ltrim   Trims Left side of Stream Name")
		fmt.Println("     --stream-rtrim   Trims Right side of Stream Name")
		fmt.Println("     --message-ltrim  Trims Left side of Event Message")
		fmt.Println("     --message-rtrim  Trims Right side of Event Message")
		fmt.Println("     --green          Comma delimited Words to highlight Green")
		fmt.Println("     --yellow         Comma delimited Words to highlight Yellow")
		fmt.Println("     --red            Comma delimited Words to highlight Red")
		fmt.Println("")
	}

	if util.Member(options.Group, "less") {
		fmt.Println("MORE HELP: lawsg help more")
	}

	if util.Member(options.Group, "advanced,more,docs") {
		fmt.Println("ADVANCED OPTIONS:")
		fmt.Println("  -c --command        Command to run groups, streams, get, help (or first argument)")
		fmt.Println("  -g --group          Group for the command (or second argument)")
		fmt.Println("     --chunk          Chunk size for retrieving Event Logs [default: 10000]")
		fmt.Println("     --refresh        Tail Refresh interval in seconds [default: 5]")
		fmt.Println("     --cache          Enable local cache")
		fmt.Println("     --cachedir       Directory for the local cache")
		fmt.Println("     --stats          Display Stats from request")
		fmt.Println("     --debug          Debug of Output")
		fmt.Println("")
	}

	if util.Member(options.Group, "cloudwatch,filter,more,docs") {
		fmt.Println("CLOUDWATCH FILTER EXAMPLES:")
		fmt.Println("  '\"toy\"'           Matches toy, toys or toyota")
		fmt.Println("  '\"toy\" -\"fun\"'    Matches toy but not fun")
		fmt.Println("  '\"eggs ham\"'      Matches multiple words 'green eggs and ham'")
		fmt.Println("                    More: https://goo.gl/eHztRU")
		fmt.Println("")
	}

	if util.Member(options.Group, "saved,config,more,docs") {
		fmt.Println("SAVED CONFIG:")
		fmt.Println("  Display options can be saved in a config file for each group (or global)")
		fmt.Println("  Save to: ~/.lawsg/config.json")
		fmt.Println("  Example:")
		fmt.Println("    {")
		fmt.Println("      \"global\": {\"no_group\": true, , \"date_format\": \"yyyyMMdd-H:mm:ss\"},")
		fmt.Println("      \"/var/log/myapp.log\": {\"yellow\": \"[WARN]\", \"red\": \"[ERR]\"}")
		fmt.Println("    }")
		fmt.Println("  Fields:")
		fmt.Println("    time_zone         bool")
		fmt.Println("    spacing           bool")
		fmt.Println("    no_group          bool")
		fmt.Println("    no_strean         bool")
		fmt.Println("    no_time           bool")
		fmt.Println("    no_color          bool")
		fmt.Println("    no_wrap           bool")
		fmt.Println("    stream_ltrim      int")
		fmt.Println("    stream_rtrim      int")
		fmt.Println("    message_ltrim     int")
		fmt.Println("    message_rtrim     int")
		fmt.Println("    date_format       string")
		fmt.Println("    green             string")
		fmt.Println("    yellow            string")
		fmt.Println("    red               string")
		fmt.Println("")
	}

	if util.Member(options.Group, "env,environment,more,docs") {
		fmt.Println("ENVIRONMENT VARIABLES:")
		fmt.Println("  LAWSG_FILTER         Cloudwatch Filter for Event Logs")
		fmt.Println("  LAWSG_STREAM         Comma delimited list of Streams")
		fmt.Println("  LAWSG_START_TIME     Start Time for the Event Logs [default: 10 min before now]")
		fmt.Println("  LAWSG_END_TIME       End Time for the Event Logs [default: now]")
		fmt.Println("  LAWSG_NUMBER         Number of Log Events to show")
		fmt.Println("  LAWSG_TAIL           Active tailing of Event Logs (experimental)")
		fmt.Println("  LAWSG_TIMEZONE       Convert Event Log display to local time")
		fmt.Println("  LAWSG_SPACING        Adds spacing between Log Events")
		fmt.Println("  LAWSG_NO_GROUP       Display No Group column")
		fmt.Println("  LAWSG_NO_STREAM      Display No Stream column")
		fmt.Println("  LAWSG_NO_TIME        Display No Time column")
		fmt.Println("  LAWSG_NO_COLOR       Display No Color")
		fmt.Println("  LAWSG_NO_WRAP        Display No Wrapping of lines (truncates)")
		fmt.Println("  LAWSG_STREAM_LTRIM   Trims Left side of Stream Name")
		fmt.Println("  LAWSG_STREAM_RTRIM   Trims Right side of Stream Name")
		fmt.Println("  LAWSG_MESSAGE_LTRIM  Trims Left side of Event Message")
		fmt.Println("  LAWSG_MESSAGE_RTRIM  Trims Right side of Event Message")
		fmt.Println("  LAWSG_TRIM_LEFT      Trims Left side of Event Message")
		fmt.Println("  LAWSG_DATE_FORMAT    The Date Format to use for display")
		fmt.Println("  LAWSG_GREEN          Comma delimited Words to highlight Green")
		fmt.Println("  LAWSG_YELLOW         Comma delimited Words to highlight Yellow")
		fmt.Println("  LAWSG_RED            Comma delimited Words to highlight Red")
		fmt.Println("  LAWSG_COMMAND        Command to run groups, streams, get, help (or first argument)")
		fmt.Println("  LAWSG_GROUP          Group for the command (or second argument)")
		fmt.Println("  LAWSG_CHUNK          Chunk size for retrieving Event Logs [default: 10000]")
		fmt.Println("  LAWSG_REFRESH        Tail Refresh interval in seconds [default: 5]")
		fmt.Println("  LAWSG_CACHE          Enable local cache")
		fmt.Println("  LAWSG_CACHE_DIR      Directory for the local cache")
		fmt.Println("  LAWSG_STATS          Display Stats from request")
		fmt.Println("  LAWSG_DEBUG          Debug of Output")
		fmt.Println("")
	}

	if util.Member(options.Group, "date,datetime,time,more,docs") {
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

	if util.Member(options.Group, "dateformat,more,docs") {
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

	if util.Member(options.Group, "aws,more,docs") {
		fmt.Println("AWS CREDENTIALS:")
		fmt.Println("  Normal AWS Env Vars are used")
		fmt.Println("  AWS_REGION")
		fmt.Println("  AWS_ACCESS_KEY_ID")
		fmt.Println("  AWS_SECRET_ACCESS_KEY")
		fmt.Println("")
	}

	if util.Member(options.Group, "about,more,docs") {
		fmt.Println("ABOUT:")
		fmt.Println("  Free to Use and Contribute via the MIT License")
		fmt.Println("  Maintained by Matt McQuillan")
		fmt.Println("  https://github.com/mmcquillan/lawsg")
		fmt.Println("")
	}

	if util.Member(options.Group, "docs") {
		fmt.Println("```")
	}

	if options.Stats {
		stat := fmt.Sprintf("\n[ STATS: %dms exec ]\n", (time.Now().UnixNano()-options.Timer)/1000000)
		fmt.Print(stat)
	}

}
