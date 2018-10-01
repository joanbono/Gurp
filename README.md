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
```

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