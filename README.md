# go-memo

Simple terminal UI app to store and recall frequently used commands, grouped by context

Commands can be quickly browsed and copied to the clipboard

![](img/demo.gif)


## Design

This project follows the **Elm architecture**: a **model** contains the application's state, which is **updated** based on user's messages and **rendered** to stdout.

**Views** are created following the **strategy** design pattern (the interface is in **internal/app/views.go**), which provides very good scalability for this case.

TODO: database

## Get started

First off, you'll need to create a database containing groups of commands. The easiest way to get started is to create one from the SQL instructions in `examples/basic.sql`.

`go-memo` works with **SQLite** and looks for databases in `$HOME/.config/memo/commands.db`

```bash
mkdir -p $HOME/.config/memo
sqlite3 $HOME/.config/memo/commands.db < examples/basic.sql
```

## Build and run

Give it a try with

```bash
go build -o memo
./memo
```

## For dev


Run unit tests with

```bash
go test -v ./...
```

Additionally, this project has a small (and colourful!) debug package. Printing to stdout is not a good idea, since it is **blocked by the TUI**.
Debug information can be written to `$TMPDIR/go-memo-debug.log` by uncommenting these lines in `main.go`

```go
  debug.Start()
  defer debug.Stop()
```

