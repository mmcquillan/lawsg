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
      --stream=STREAM        Stream Filter (can be a comma seperated list)
  -s, --starttime=STARTTIME  Start Time
  -e, --endtime=ENDTIME      End Time
  -n, --number=NUMBER        Number of Rows
      --chunk=CHUNK          Chunk Size
      --ns                   No Streams
      --nw                   No Wrapping Lines (will be truncated)
      --nc                   No Color
      --nt                   No Time
  -z, --tz                   Convert my Time Zone
  -t, --tail                 Tail of Log
      --trimleft=TRIMLEFT    Trim Left of Event Message
      --green=GREEN          Green Highlight
      --yellow=YELLOW        Yellow Highlight
      --red=RED              Red Highlight
      --debug                Debug

Args:
  <group>  Log Group or 'list' to show groups
```
