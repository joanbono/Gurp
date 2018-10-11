![](img/Gurp_banner.png)

## Requirements

+ BurpSuite Professional version `v2.0.0beta` or greater from [PortSwigger](https://portswigger.net/burp)

> Enable the API under *User Options* > *Misc* > *REST API*. 



***

## Dependencies

```bash
go get -u -v github.com/joanbono/Gurp
```

Add `rsrc` to the `$PATH` to build Windows binaries using the icon.

Run `make all` to build binaries for Windows, Linux and Darwin.

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

# Build all
make all
```

***

## Usage

Refer to [`USAGE`](USAGE.md)
