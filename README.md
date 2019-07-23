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

Conversion from libusb is noted here

### Device Handling

| Libusb function | `usb` equiv | Done? |
|-----------------|-------------|-------|
| `libusb_device` | `usb.Device` | ✔|
| `libusb_get_device_list` | `usb.List()` | ✔|
| `libusb_get_bus_number` | `usb.Device.Bus` | ✔|
| `libusb_get_port_number` | `usb.Device.Port` | ✔|
| `libusb_get_port_numbers` | `usb.Device.Ports` | ✔|
| `libusb_get_parent` | `usb.Device.Parent` | ✔|
| `libusb_get_device_address` | `usb.Device.Device` | ✔|
| `libusb_get_device_speed` | `usb.Device.Speed` | ✔|
| `libusb_get_max_packet_size` | `usb.Endpoint.MaxPacketSize` | ✔|
| `libusb_get_max_iso_packet_size` | `usb.Endpoint.MaxISOPacketSize` | ✔|
| `libusb_open` | `usb.Open()` | ✔|
| `libusb_open_device_with_vid_pid` | `usb.VidPid().Open()` | ✔|
| `libusb_close` | `usb.Device.Close()` | ✔|
| `libusb_get_configuration` | `usb.Device.ActiveConfig` | ✔|
| `libusb_set_configuration` | `usb.Device.SetConfiguration()` | ✔|
| `libusb_claim_interface` | `usb.Interface.Claim()` | ✔|
| `libusb_release_interface` | `usb.Interface.Release()` | ✔|
| `libusb_set_interface_alt_setting` | `usb.Interface.SetAlt()` | ✔|
| `libusb_clear_halt` | - | |
| `libusb_reset_device` | `usb.Device.Reset()` | ✔|
| `libusb_kernel_driver_active` | - | |
| `libusb_speed` | `usb.Speed` | ✔|
| `libusb_supported_speed` | (reuse usb.Speed?) | |
| `libusb_2_0_extension_attributes` | - ||
| `libusb_ss_usb_device_capability_attributes` | - ||
| `libusb_bos_type` | - ||



### Descriptors
libusb_*_descriptor -> see gusb library

### Hotplug
libusb_hotplug_* -> TODO

### Async IO
TODO

### Polling, Timing
TODO

### Synchronous IO
libusb_control_transfer -> TODO
libusb_bulk_transfer -> TODO
libusb_interrupt_transfer -> TODO

**Things that are not needed with this library:**

- libusb_version
- libusb_context
- libusb_option
- libusb_set_option
- libusb_init
- libusb_exit
- libusb_device_handle
- libusb_get_device
- libusb_set_auto_detach_kernel_driver - always on
- libusb_detach_kernel_driver - automatically detached with claim
- libusb_attach_kernel_driver - automatically reattached with release
- libusb_free_device_list
- libusb_ref_device
- libusb_unref_device

**Deprecated calls:**

- libusb_get_port_path
- libusb_set_debug



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