# lawsg - The AWS Cloudwatch Logs Viewer

### Why another cloudwatch CLI?
- Query one or more Cloudwatch groups
- Advanced formatting to help you find what you need
- Highlight keywords in Green, Yellow or Red to capture your attention
- Display preferences via command line flags, env variables or a config file
- Help for when you can't quite remember the group name


### Features
- [Detailed Help](HELP.md)


### Get Started
1. Install on Mac via brew: `brew install mmcquillan/tools/lawsg`
2. Alternatively compile via go: `go get github.com/mmcquillan/lawsg`
3. Ensure you set your AWS Environment Variables:
    - `AWS_REGION`
    - `AWS_ACCESS_KEY_ID`
    - `AWS_SECRET_ACCESS_KEY`
4. Get started with: `lawsg help more`
5. Enjoy all your wonderful logs!


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
