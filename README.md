usb
====

This is a pure Go implementation of the USB stack (no Cgo). It is a `libusb` alternative for situations you need a go-native library.

`usb` currently only supports linux, but with enough interest or requests, we can add more platforms.


Usage
-------

The top-level of this project can be used as a high-level library. See the docs at [godoc](https://godoc.org/github.com/pzl/usb).

The `gusb` sub-directory can be used as a more low-level library, if that suits your needs. Documentation also at [godoc](https://github.com/pzl/usb/gusb)


License
--------

MIT

Copyright 2019 Dan Panzarella

See [`LICENSE`](LICENSE) file for full License details