# go-memo

Simple terminal UI app to store and recall frequently used commands, grouped by context

Commands can be quickly browsed and copied to the clipboard

![](img/demo.gif)


## Build and run

Give it a try with

```bash
go build -o memo
./memo
```

## For dev

Debug information can be written to `$TMPDIR/go-memo-debug.log` by uncommenting these lines in `main.go`

```go
	debug.Start()
	defer debug.Stop()
```

