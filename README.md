# lawsg - The AWS Cloudwatch Logs Viewer

### Why another cloudwatch CLI?
- Query one or more Cloudwatch groups
- Advanced formatting to help you find what you need
- Highlight keywords in Green, Yellow or Red to capture your attention
- Display preferences via command line flags, env variables or a config file
- Help for when you can't quite remember the group name


### Get Started
1. Install on Mac via brew: `brew install mmcquillan/tools/lawsg`
2. Alternatively compile via go: `go get github.com/mmcquillan/lawsg`
3. Ensure you set your AWS Environment Variables:
    - `AWS_REGION`
    - `AWS_ACCESS_KEY_ID`
    - `AWS_SECRET_ACCESS_KEY`
4. Get started with: `lawsg help more`
5. Enjoy all your wonderful logs!


### Features
```
USAGE:
  lawsg help [ more | <topic name> ]
  lawsg groups [options]
  lawsg streams <group name> [options]
  lawsg get <group name> [options]
  lawsg version

FILTER OPTIONS:
  -f --filter         Cloudwatch Filter for Event Logs
  -m --stream         Comma delimited list of Streams
  -s --starttime      Start Time for the Event Logs [default: 10 min before now]
  -e --endtime        End Time for the Event Logs [default: now]
  -n --number         Number of Log Events to show
  -t --tail           Active tailing of Event Logs

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
     --multi-line     Handles multiple logs entries in one Event Message
     --green          Comma delimited Words to highlight Green
     --yellow         Comma delimited Words to highlight Yellow
     --red            Comma delimited Words to highlight Red

ADVANCED OPTIONS:
  -c --command        Command to run groups, streams, get, help (or first argument)
  -g --group          Group for the command (or second argument)
     --chunk          Chunk size for retrieving Event Logs [default: 10000]
     --sortkey        Add each line with a sortable time based key [default: false]
     --refresh        Tail Refresh interval in seconds [default: 5]
     --region         Override or set the AWS Region
     --env            Environment prefix multiple AWS ENV Vars (ex: STAGING_AWS_REGION)
     --stats          Display Stats from request
     --debug          Debug of Output
```
[More Detailed Help...](HELP.md)


### Releases

v0.3.0
- Added `--env` var for multiple AWS ENV vars
- Added `--region` var for overriding or setting the AWS region
- Added messaging if Start Time is parsed incorrectly
- Added `--sortkey` to prefix each line with the epoch timestamp for sorting

v0.2.0
- Added get for multiple or all log groups
- Added a countdown before update on the tail option

v0.1.1
- Added the multi-line formatting option to handle multiple log entries per Event Message
- Fixed a help issue

v0.1.0
- Initial release
