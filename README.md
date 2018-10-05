![](img/Gurp_banner.png)

## Requirements

+ BurpSuite Professional `v2.0.0beta` or greater from [PortSwigger](https://portswigger.net/burp)

***

## Dependencies

```bash
go get -u -v github.com/fatih/color
go get -u -v github.com/integrii/flaggy
go get -u -v github.com/tidwall/gjson
go get -u -v github.com/grokify/html-strip-tags-go
go get -u -v github.com/akavel/rsrc
go get -u -v github.com/tomsteele/go-nmap
```

Add `rsrc` to the `$PATH` to build Windows binaries using the icon.

This can be automated by running `make deps` and `make all`.
To install the binary in your `$GOPATH`, run `make install`.

***

## Binaries

Latest version available [here](https://github.com/joanbono/Gurp/releases/latest).

***

## Building

```bash
# macOS binary
make darwin

# Linux binary
make linux

# Windows binary
make windows

# Build releases
make all
```

***

## Usage

Refer to [`USAGE`](USAGE.md)
