# Build steps

```
go install -ldflags "-X main.version=XXX" github.com/mmcquillan/lawsg
cd $GOBIN
tar -cvzf lawsg-XXX.tar.gz ./lawsg
openssl sha -sha256 lawsg-XXX.tar.gz
```
