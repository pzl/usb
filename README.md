usb
====

> Pure Go USB stack

[![GoDoc][godoc-badge]][godoc]
[![Go Report Card][goreport-badge]][goreport]
[![MIT license][MIT-badge]][license]
![Tag][tag-badge]


![Logo][logo]

`usb` is a pure Go implementation of the USB stack (no CGo). It is a `libusb` alternative for situations you need a go-native library.

`usb` currently only supports linux, but with enough interest or requests, we can add more platforms.



Getting Started
-----------------

```go
package main

import  "github.com/pzl/usb"

func main() {
    dev, err := usb.VidPid(0x0c45, 0x6300) // get device by IDs
    if err != nil {
        // handle...
    }
}

```


Usage
-------

The top-level of this project can be used as a high-level library. See the docs at [godoc][godoc].

The `gusb` sub-directory can be used as a more low-level library, if that suits your needs. Documentation also at [godoc][godoc].


Status
-------

Conversion from libusb is underway. Almost all enumeration and device descriptor function is implemented. IO (both sync and async) is in progress.

See [API](API.md) for a direct mapping of libusb calls to `usb` methods.


Release History
----------------
- 0.0.1
    + Work in Progress

License
--------

MIT

Copyright 2019 Dan Panzarella

See [`LICENSE`][license] file for full License details


<!-- biblio -->
[godoc-badge]: https://godoc.org/github.com/pzl/usb?status.svg
[goreport-badge]: https://goreportcard.com/badge/github.com/pzl/usb?style=flat
[MIT-badge]: https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat
[tag-badge]: https://img.shields.io/github/tag/pzl/usb.svg?style=flat
[release-badge]: https://img.shields.io/github/release/pzl/usb/all.svg?style=flat
[logo]: gusb.svg
[godoc]: https://godoc.org/github.com/pzl/usb
[goreport]: https://goreportcard.com/report/github.com/pzl/usb
[license]: LICENSE