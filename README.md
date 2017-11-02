# lawsg - The AWS Cloudwatch Logs Viewer


Install by running:

`go get github.com/mmcquillan/lawsg`



Ensure you set the Environment Variables:

`AWS_ACCESS_KEY_ID`

`AWS_SECRET_ACCESS_KEY`



Command Line Arguments
```
usage: lawsg [<flags>] <group>

Flags:
      --help                 Show context-sensitive help (also try --help-long and --help-man).
  -f, --filter=FILTER        Cloudwatch Filter Pattern
      --stream=STREAM        Stream Filter (comma seperated)
  -s, --starttime=STARTTIME  Start Time
  -e, --endtime=ENDTIME      End Time
  -l, --last=LAST            Last X minutes of logs
  -n, --number=NUMBER        Number of Rows
      --chunk=CHUNK          Chunk Size
      --ns                   No Streams
      --nw                   No Wrapping Lines (will be truncated)
      --nc                   No Color
      --nt                   No Time
  -z, --tz                   Convert my Time Zone
  -t, --tail                 Tail of Log
      --trimleft=TRIMLEFT    Trim Left of Event Message
      --green=GREEN          Green Highlight (comma seperated)
      --yellow=YELLOW        Yellow Highlight (comma seperated)
      --red=RED              Red Highlight (comma seperated)
      --debug                Debug

Args:
  <group>  Log Group or 'list' to show groups
```


Environment Variables
```
LAWSG_FILTER
LAWSG_STREAM
LAWSG_STARTTIME
LAWSG_ENDTIME
LAWSG_LAST
LAWSG_NUMBER
LAWSG_CHUNK
LAWSG_TAIL
LAWSG_TIMEZONE
LAWSG_NO_TIME
LAWSG_NO_STREAM
LAWSG_NO_COLOR
LAWSG_NOWRAP
LAWSG_TRIM_LEFT
LAWSG_GREEN
LAWSG_YELLOW
LAWSG_RED
LAWSG_DEBUG
```
