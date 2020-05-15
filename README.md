# urlencode

Converts a single argument (a raw url string without a query) to an escaped version.

## example

```
$ urlencode "https://store.xkcd.com/products/citation needed sticker pack"
https://store.xkcd.com/products/citation%20needed%20sticker%20pack
$
```

## building

You will need a [Go](https://golang.org/) tool suite installed. Go modules are used; but are hardly necessary for such a simple tool.  So, a version of Go >= 1.11 should be used.  However, the tool should build just fine with an older version.

There are a variety of ways to build a `Go` command.

You can build the command directly from this directory:
```
$ go build
```

You can even fetch it from the source code repository, build, and install it in `$GOPATH/bin` in one step. From any directory:
```
$ go get github.com/ranger6/urlencode
```
