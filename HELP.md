# LAWSG HELP
```

LAWSG - The AWS Cloudwatch Logs Viewer

USAGE:
  lawsg help [ more | <topic name> ]
  lawsg groups [options]
  lawsg streams <group name> [options]
  lawsg get <group name> [options]

FILTER OPTIONS:
  -f --filter         Cloudwatch Filter for Event Logs
  -m --stream         Comma delimited list of Streams
  -s --starttime      Start Time for the Event Logs [default: 10 min before now]
  -e --endtime        End Time for the Event Logs [default: now]
  -n --number         Number of Log Events to show
  -t --tail           Active tailing of Event Logs (experimental)

DISPLAY OPTIONS:
     --tz             Convert Event Log display to local time
     --spacing        Adds spacing between Log Events
     --ng             Display No Group column
     --ns             Display No Stream column
     --nt             Display No Time column
     --nc             Display No Color
     --nw             Display No Wrapping of lines (truncates)
     --stream-ltrim   Trims Left side of Stream Name
     --stream-rtrim   Trims Right side of Stream Name
     --message-ltrim  Trims Left side of Event Message
     --message-rtrim  Trims Right side of Event Message
     --green          Comma delimited Words to highlight Green
     --yellow         Comma delimited Words to highlight Yellow
     --red            Comma delimited Words to highlight Red

ADVANCED OPTIONS:
  -c --command        Command to run groups, streams, get, help (or first argument)
  -g --group          Group for the command (or second argument)
     --chunk          Chunk size for retrieving Event Logs [default: 10000]
     --refresh        Tail Refresh interval in seconds [default: 5]
     --cache          Enable local cache
     --cachedir       Directory for the local cache
     --stats          Display Stats from request
     --debug          Debug of Output

CLOUDWATCH FILTER EXAMPLES:
  '"toy"'           Matches toy, toys or toyota
  '"toy" -"fun"'    Matches toy but not fun
  '"eggs ham"'      Matches multiple words 'green eggs and ham'
                    More: https://goo.gl/eHztRU

SAVED CONFIG:
  Display options can be saved in a config file for each group (or global)
  Save to: ~/.lawsg/config.json
  Example:
    {
      "global": {"no_group": true, , "date_format": "yyyyMMdd-H:mm:ss"},
      "/var/log/myapp.log": {"yellow": "[WARN]", "red": "[ERR]"}
    }
  Fields:
    time_zone         bool
    spacing           bool
    no_group          bool
    no_strean         bool
    no_time           bool
    no_color          bool
    no_wrap           bool
    stream_ltrim      int
    stream_rtrim      int
    message_ltrim     int
    message_rtrim     int
    date_format       string
    green             string
    yellow            string
    red               string

ENVIRONMENT VARIABLES:
  LAWSG_FILTER         Cloudwatch Filter for Event Logs
  LAWSG_STREAM         Comma delimited list of Streams
  LAWSG_START_TIME     Start Time for the Event Logs [default: 10 min before now]
  LAWSG_END_TIME       End Time for the Event Logs [default: now]
  LAWSG_NUMBER         Number of Log Events to show
  LAWSG_TAIL           Active tailing of Event Logs (experimental)
  LAWSG_TIMEZONE       Convert Event Log display to local time
  LAWSG_SPACING        Adds spacing between Log Events
  LAWSG_NO_GROUP       Display No Group column
  LAWSG_NO_STREAM      Display No Stream column
  LAWSG_NO_TIME        Display No Time column
  LAWSG_NO_COLOR       Display No Color
  LAWSG_NO_WRAP        Display No Wrapping of lines (truncates)
  LAWSG_STREAM_LTRIM   Trims Left side of Stream Name
  LAWSG_STREAM_RTRIM   Trims Right side of Stream Name
  LAWSG_MESSAGE_LTRIM  Trims Left side of Event Message
  LAWSG_MESSAGE_RTRIM  Trims Right side of Event Message
  LAWSG_TRIM_LEFT      Trims Left side of Event Message
  LAWSG_DATE_FORMAT    The Date Format to use for display
  LAWSG_GREEN          Comma delimited Words to highlight Green
  LAWSG_YELLOW         Comma delimited Words to highlight Yellow
  LAWSG_RED            Comma delimited Words to highlight Red
  LAWSG_COMMAND        Command to run groups, streams, get, help (or first argument)
  LAWSG_GROUP          Group for the command (or second argument)
  LAWSG_CHUNK          Chunk size for retrieving Event Logs [default: 10000]
  LAWSG_REFRESH        Tail Refresh interval in seconds [default: 5]
  LAWSG_CACHE          Enable local cache
  LAWSG_CACHE_DIR      Directory for the local cache
  LAWSG_STATS          Display Stats from request
  LAWSG_DEBUG          Debug of Output

DATETIME EXAMPLES:
  'now' 'n' current time
  '12s ago' twelve seconds ago
  '12m ago' twelve minutes ago
  '12h ago' twelve hours ago
  '12d ago' twelve days ago
  '12w ago' twelve weeks ago
  '12y ago' twelve years ago
  ...and most every other parsable date/time

DATEFORMAT EXAMPLES:
  Two digit year              y       06
  Four digit year             Y       2006
  Long month name             MMMM    January
  Short month name            MMM     Jan
  Two digit month             MM      01
  Single digit month          M       1
  Two digit day of month      dd      02
  Single digit day of month   d       2
  Long day of week            EE      Monday
  Short day of week           E       Mon
  Two digit 12 hours          hh      03
  Single digit 12 hour        h       3
  Two digit 24 hours          H       15
  Two digit minutes           mm      04
  Single digit minutes        m       4
  Two digit seconds           ss      05
  Single digit seconds        s       5
  Milliseconds                S       .0
  AM/PM                       a       pm
  Time Zone                   z       MST
  Time Zone offset            Z       -0700

AWS CREDENTIALS:
  Normal AWS Env Vars are used
  AWS_REGION
  AWS_ACCESS_KEY_ID
  AWS_SECRET_ACCESS_KEY

ABOUT:
  Free to Use and Contribute via the MIT License
  Maintained by Matt McQuillan
  https://github.com/mmcquillan/lawsg

```
