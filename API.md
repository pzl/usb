API Mapping
============

below includes a direct mapping of Libusb functions to `usb` concepts.

## [Device Handling](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html)

| Libusb | `usb` equiv | Done? |
|--------|-------------|-------|
| [`libusb_device`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga77eedd00d01eb7569b880e861a971c2b) | [`usb.Device`][usb-device] | ✔|
| [`libusb_get_device_list`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gac0fe4b65914c5ed036e6cbec61cb0b97) | [`usb.List()`](https://godoc.org/github.com/pzl/usb#List) | ✔|
| [`libusb_get_bus_number`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gaf2718609d50c8ded2704e4051b3d2925) | [`usb.Device.Bus`][usb-device] | ✔|
| [`libusb_get_port_number`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga14879a0ea7daccdcddb68852d86c00c4) | [`usb.Device.Port`][usb-device] | ✔|
| [`libusb_get_port_numbers`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gaa4b7b2b50a9ce2aa396b0af2b979544d) | [`usb.Device.Ports`][usb-device] | ✔|
| [`libusb_get_parent`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga97bb4dfff6bbb897ed9dfd6fa1a1deed) | [`usb.Device.Parent`][usb-device] | ✔|
| [`libusb_get_device_address`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gab6d4e39ac483ebaeb108f2954715305d) | [`usb.Device.Device`][usb-device] | ✔|
| [`libusb_get_device_speed`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga58c4e448ecd5cd4782f2b896ec40b22b) | [`usb.Device.Speed`][usb-device] | ✔|
| [`libusb_get_max_packet_size`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gac81968047e262409e78f3fe24321b604) | [`usb.Endpoint.MaxPacketSize`][usb-ep] | ✔|
| [`libusb_get_max_iso_packet_size`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gaec10b71c7209760db55ee0f8768bb4f0) | [`usb.Endpoint.MaxISOPacketSize`][usb-ep] | ✔|
| [`libusb_open`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga3f184a8be4488a767b2e0ae07e76d1b0) | [`usb.Open()`](https://godoc.org/github.com/pzl/usb#Open) | ✔|
| [`libusb_open_device_with_vid_pid`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga11ba48adb896b1492bbd3d0bf7e0f665) | [`usb.VidPid().Open()`](https://godoc.org/github.com/pzl/usb#VidPid) | ✔|
| [`libusb_close`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga779bc4f1316bdb0ac383bddbd538620e) | [`usb.Device.Close()`](https://godoc.org/github.com/pzl/usb#Device.Close) | ✔|
| [`libusb_get_configuration`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gae921014b888b105471a31d54c77c1c4d) | [`usb.Device.ActiveConfig`][usb-device] | ✔|
| [`libusb_set_configuration`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga785ddea63a2b9bcb879a614ca4867bed) | [`usb.Device.SetConfiguration()`](https://godoc.org/github.com/pzl/usb#Device.SetConfiguration) | ✔|
| [`libusb_claim_interface`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gaee5076addf5de77c7962138397fd5b1a) | [`usb.Interface.Claim()`](https://godoc.org/github.com/pzl/usb#Interface.Claim) | ✔|
| [`libusb_release_interface`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga49b5cb0d894f6807cd1693ef29aecbfa) | [`usb.Interface.Release()`](https://godoc.org/github.com/pzl/usb#Interface.Release) | ✔|
| [`libusb_set_interface_alt_setting`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga4858ad4f0f58fd1dc0afaead1fe6479a) | [`usb.Interface.SetAlt()`](https://godoc.org/github.com/pzl/usb#Interface.SetAlt) | ✔|
| [`libusb_clear_halt`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gab794bbc0b055d140f186f5a4d39c0891) | - | |
| [`libusb_reset_device`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gafee9c4638f1713ca5faa867948878111) | [`usb.Device.Reset()`](https://godoc.org/github.com/pzl/usb#Device.Reset) | ✔|
| [`libusb_kernel_driver_active`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga1cabd4660a274f715eeb82de112e0779) | - | |
| [`libusb_speed`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga2959abf1184f87b2ce06fe90db6ce614) | [`usb.Speed`](https://godoc.org/github.com/pzl/usb#Speed) | ✔|
| [`libusb_supported_speed`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga1454797ecc0de4d084c1619c420014f6) | (reuse usb.Speed?) | |
| [`libusb_2_0_extension_attributes`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gacb8cfa928bffdd0066a3dd2e6aba0558) | - ||
| [`libusb_ss_usb_device_capability_attributes`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gaad5a5399176a35a64164dafad7fe4fcd) | - ||
| [`libusb_bos_type`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga6ccabbf3b3728ae69608ba83bba4e64c) | - ||



## [Descriptors](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html)

These are implemented in the `gusb` subpackage

| Libusb | `gusb` equiv | Done? |
|--------|--------------|-------|


## [Hotplug](http://libusb.sourceforge.net/api-1.0/group__libusb__hotplug.html)

| Libusb | `usb` equiv | Done? |
|--------|-------------|-------|
| [`libusb_hotplug_register_callback`](http://libusb.sourceforge.net/api-1.0/group__libusb__hotplug.html#ga00e0c69ddf1fb1b6774dc918192e8dc7) | - ||
| [`libusb_hotplug_deregister_callback`](http://libusb.sourceforge.net/api-1.0/group__libusb__hotplug.html#ga8110f57eab2064375934f1449b2602bc) | - ||


## [Async IO](http://libusb.sourceforge.net/api-1.0/group__libusb__asyncio.html)
TODO

## [Polling, Timing](http://libusb.sourceforge.net/api-1.0/group__libusb__poll.html)
TODO

## [Synchronous IO](http://libusb.sourceforge.net/api-1.0/group__libusb__syncio.html)

| Libusb | `usb` equiv | Done? |
|--------|-------------|-------|
| [`libusb_control_transfer`](http://libusb.sourceforge.net/api-1.0/group__libusb__syncio.html#gadb11f7a761bd12fc77a07f4568d56f38) | - ||
| [`libusb_bulk_transfer`](http://libusb.sourceforge.net/api-1.0/group__libusb__syncio.html#gab8ae853ab492c22d707241dc26c8a805) | - ||
| [`libusb_interrupt_transfer`](http://libusb.sourceforge.net/api-1.0/group__libusb__syncio.html#gac412bda21b7ecf57e4c76877d78e6486) | - ||



## Not Included

These pieces are not being ported to `usb`

| Libusb | reason |
|--------|--------|
| [`libusb_version`](http://libusb.sourceforge.net/api-1.0/structlibusb__version.html) | 
| [`libusb_context`](http://libusb.sourceforge.net/api-1.0/group__libusb__lib.html#ga4ec088aa7b79c4a9599e39bf36a72833) | this is not a shared library
| [`libusb_option`](http://libusb.sourceforge.net/api-1.0/group__libusb__lib.html#ga07d4ec54cf575d672ba94c72b3c0de7c) | not planned
| [`libusb_set_option`](http://libusb.sourceforge.net/api-1.0/group__libusb__lib.html#gaf6ce5db28dac96b1680877a123da4fa8) | not planned
| [`libusb_init`](http://libusb.sourceforge.net/api-1.0/group__libusb__lib.html#ga9517c37281bba0b51cc62eba728be48b) |
| [`libusb_exit`](http://libusb.sourceforge.net/api-1.0/group__libusb__lib.html#ga86532f222d4f1332a5f8f5eef9a92da9) |
| [`libusb_device_handle`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga7df95821d20d27b5597f1d783749d6a4) | handles not used
| [`libusb_get_device`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gafe70b8a797893d4d16985980acec956a) | handles not used
| [`libusb_set_auto_detach_kernel_driver`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gac35b26fef01271eba65c60b2b3ce1cbf) | this feature is always enabled
| [`libusb_detach_kernel_driver`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga5e0cc1d666097e915748593effdc634a) | automatically detached with claim
| [`libusb_attach_kernel_driver`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gadeba36e900db663c0b7cf1b164a20d02) | automatically reattached with release
| [`libusb_free_device_list`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gad3b8561d064bb3e1b8851ddeed3cd7d6) | not manually memory managed
| [`libusb_ref_device`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gaabaa4193adcabba1789cc1165ac41a03) | not reference counted
| [`libusb_unref_device`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga3cc62e6a191b7a9f213e62b81ec30f4d) | not reference counted
| [`libusb_get_port_path`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga9d392b8dff7abf5e475c72fd071c3c34) | _deprecated in libusb_
| [`libusb_set_debug`](http://libusb.sourceforge.net/api-1.0/group__libusb__lib.html#ga5f8376b7a863a5a8d5b8824feb8a427a) | _deprecated in libusb_



<!-- Biblio -->

[usb-device]: https://godoc.org/github.com/pzl/usb#Device
[usb-ep]: https://godoc.org/github.com/pzl/usb#Endpoint