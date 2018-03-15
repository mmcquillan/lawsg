# Build steps

1. Ensure you have updated the help docs.

```
go run lawsgs.go docs > HELP.md
```


2. Prepare Binaries for a Release
```
go install -ldflags "-X main.version=XXX" github.com/mmcquillan/lawsg
cd $GOBIN
tar -cvzf lawsg-XXX.tar.gz ./lawsg
openssl sha -sha256 lawsg-XXX.tar.gz
```


3. Create a github release.

https://github.com/mmcquillan/lawsg/releases


4. Update Homebrew Forumla if needed.

https://github.com/mmcquillan/homebrew-tools/blob/master/lawsg.rb
