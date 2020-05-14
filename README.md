# urlencode

Converts a single argument (a raw url string without a query) to an escaped version.

## example

```
% urlencode "https://store.xkcd.com/products/citation needed sticker pack"
https://store.xkcd.com/products/citation%20needed%20sticker%20pack
%
```

## building

You will need a [Go](https://golang.org/) tool suite installed.  There are a variety of ways to build a `Go` command.

You can build the command directly from this directory:
```
% go build urlencode.go
```

You can even fetch it from the source code repository, build, and install it in one step:
```
% go get github.com/ranger6/urlencode
```
