API Mapping
============

below includes a direct mapping of Libusb functions to `usb` concepts.

## [Device Handling](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html)

| Libusb | `usb` equiv | Done? |
|--------|-------------|-------|
| `struct` [`libusb_device`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga77eedd00d01eb7569b880e861a971c2b) | [`usb.Device`][usb-device] | ✔|
| [`libusb_get_device_list()`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gac0fe4b65914c5ed036e6cbec61cb0b97) | [`usb.List()`](https://godoc.org/github.com/pzl/usb#List) | ✔|
| [`libusb_get_bus_number()`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gaf2718609d50c8ded2704e4051b3d2925) | [`usb.Device.Bus`][usb-device] | ✔|
| [`libusb_get_port_number()`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga14879a0ea7daccdcddb68852d86c00c4) | [`usb.Device.Port`][usb-device] | ✔|
| [`libusb_get_port_numbers()`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gaa4b7b2b50a9ce2aa396b0af2b979544d) | [`usb.Device.Ports`][usb-device] | ✔|
| [`libusb_get_parent()`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga97bb4dfff6bbb897ed9dfd6fa1a1deed) | [`usb.Device.Parent`][usb-device] | ✔|
| [`libusb_get_device_address()`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gab6d4e39ac483ebaeb108f2954715305d) | [`usb.Device.Device`][usb-device] | ✔|
| [`libusb_get_device_speed()`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga58c4e448ecd5cd4782f2b896ec40b22b) | [`usb.Device.Speed`][usb-device] | ✔|
| [`libusb_get_max_packet_size()`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gac81968047e262409e78f3fe24321b604) | [`usb.Endpoint.MaxPacketSize`][usb-ep] | ✔|
| [`libusb_get_max_iso_packet_size()`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gaec10b71c7209760db55ee0f8768bb4f0) | [`usb.Endpoint.MaxISOPacketSize`][usb-ep] | ✔|
| [`libusb_open()`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga3f184a8be4488a767b2e0ae07e76d1b0) | [`usb.Open()`](https://godoc.org/github.com/pzl/usb#Open) | ✔|
| [`libusb_open_device_with_vid_pid()`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga11ba48adb896b1492bbd3d0bf7e0f665) | [`usb.VidPid().Open()`](https://godoc.org/github.com/pzl/usb#VidPid) | ✔|
| [`libusb_close()`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga779bc4f1316bdb0ac383bddbd538620e) | [`usb.Device.Close()`](https://godoc.org/github.com/pzl/usb#Device.Close) | ✔|
| [`libusb_get_configuration()`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gae921014b888b105471a31d54c77c1c4d) | [`usb.Device.ActiveConfig`][usb-device] | ✔|
| [`libusb_set_configuration()`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga785ddea63a2b9bcb879a614ca4867bed) | [`usb.Device.SetConfiguration()`](https://godoc.org/github.com/pzl/usb#Device.SetConfiguration) | ✔|
| [`libusb_claim_interface()`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gaee5076addf5de77c7962138397fd5b1a) | [`usb.Interface.Claim()`](https://godoc.org/github.com/pzl/usb#Interface.Claim) | ✔|
| [`libusb_release_interface()`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga49b5cb0d894f6807cd1693ef29aecbfa) | [`usb.Interface.Release()`](https://godoc.org/github.com/pzl/usb#Interface.Release) | ✔|
| [`libusb_set_interface_alt_setting()`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga4858ad4f0f58fd1dc0afaead1fe6479a) | [`usb.Interface.SetAlt()`](https://godoc.org/github.com/pzl/usb#Interface.SetAlt) | ✔|
| [`libusb_clear_halt()`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gab794bbc0b055d140f186f5a4d39c0891) | - | |
| [`libusb_reset_device()`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gafee9c4638f1713ca5faa867948878111) | [`usb.Device.Reset()`](https://godoc.org/github.com/pzl/usb#Device.Reset) | ✔|
| [`libusb_kernel_driver_active()`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga1cabd4660a274f715eeb82de112e0779) | - | |
| `enum` [`libusb_speed`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga2959abf1184f87b2ce06fe90db6ce614) | [`usb.Speed`](https://godoc.org/github.com/pzl/usb#Speed) | ✔|
| `enum` [`libusb_supported_speed`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga1454797ecc0de4d084c1619c420014f6) | (reuse usb.Speed?) | |
| `enum` [`libusb_2_0_extension_attributes`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gacb8cfa928bffdd0066a3dd2e6aba0558) | - ||
| `enum` [`libusb_ss_usb_device_capability_attributes`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gaad5a5399176a35a64164dafad7fe4fcd) | - ||
| `enum` [`libusb_bos_type`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga6ccabbf3b3728ae69608ba83bba4e64c) | - ||


#### [Details](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#details)

This is the equivalent pattern for enumerating devices, choosing your device, and opening it

```go
var found *usb.Device

devs, err := usb.List()
if err != nil {
    panic(err)
}

for _, d := range devs {
    if isMyDevice(d) {
        found = d
        break
    }
}

if found == nil {
    panic("device not found")
}

err = found.Open()
if err != nil {
    panic(err)
}
defer found.Close()
```


The concept of needing a handle separately from a device object is not carried over to this library. The device is the object from which you may obtain descriptor data, and it is the device you use for I/O operations.

Similarly, the list obtained from `usb.List()` does not need to be freed. And devices are not reference counted.


## [Descriptors](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html)

These are implemented in the `gusb` subpackage

| Libusb | `gusb` equiv | Done? |
|--------|--------------|-------|
| `struct` [`libusb_device_descriptor`](http://libusb.sourceforge.net/api-1.0/structlibusb__device__descriptor.html) | [`gusb.DeviceDescriptor`](https://godoc.org/github.com/pzl/usb/gusb#DeviceDescriptor) | |
| `struct` [`libusb_endpoint_descriptor`](http://libusb.sourceforge.net/api-1.0/structlibusb__endpoint__descriptor.html) | [`gusb.EndpointDescriptor`](https://godoc.org/github.com/pzl/usb/gusb#EndpointDescriptor) | |
| `struct` [`libusb_interface_descriptor`](http://libusb.sourceforge.net/api-1.0/structlibusb__interface__descriptor.html) | [`gusb.InterfaceDescriptor`](https://godoc.org/github.com/pzl/usb/gusb#InterfaceDescriptor) | |
| `struct` [`libusb_interface`](http://libusb.sourceforge.net/api-1.0/structlibusb__interface.html) |  | |
| `struct` [`libusb_config_descriptor`](http://libusb.sourceforge.net/api-1.0/structlibusb__config__descriptor.html) | [`gusb.ConfigDescriptor`](https://godoc.org/github.com/pzl/usb/gusb#ConfigDescriptor) | |
| `struct` [`libusb_ss_endpoint_companion_descriptor`](http://libusb.sourceforge.net/api-1.0/structlibusb__ss__endpoint__companion__descriptor.html) |  | |
| `struct` [`libusb_bos_dev_capability_descriptor`](http://libusb.sourceforge.net/api-1.0/structlibusb__bos__dev__capability__descriptor.html) |  | |
| `struct` [`libusb_bos_descriptor`](http://libusb.sourceforge.net/api-1.0/structlibusb__bos__descriptor.html) |  | |
| `struct` [`libusb_2_0_extension_descriptor`](http://libusb.sourceforge.net/api-1.0/structlibusb__usb__2__0__extension__descriptor.html) |  | |
| `struct` [`libusb_ss_usb_device_capability_descriptor`](http://libusb.sourceforge.net/api-1.0/structlibusb__ss__usb__device__capability__descriptor.html) | | |
| `struct` [`libusb_container_id_descriptor`](http://libusb.sourceforge.net/api-1.0/structlibusb__container__id__descriptor.html) | | |
| `enum` [`libusb_class_code`](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html#gac529888effbefef8af4f9d17ebc903a1) | [`gusb.USBClass`](https://godoc.org/github.com/pzl/usb/gusb#USBClass) | |
| `enum` [`libusb_descriptor_type`](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html#ga0a2e8a25dfdebf29fdd4764dcdbc1a9c) | [`gusb.DT`](https://godoc.org/github.com/pzl/usb/gusb#DT) | |
| `enum` [`libusb_endpoint_direction`](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html#ga86c880af878493aa8f805c2aba654b8b) | [`gusb.EndpointDirection`](https://godoc.org/github.com/pzl/usb/gusb#EndpointDirection) | |
| `enum` [`libusb_transfer_type`](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html#gacb52027036a07de6ecc6c2bf07d07c71) | [`gusb.TransferType`](https://godoc.org/github.com/pzl/usb/gusb#TransferType) | |
| `enum` [`libusb_iso_sync_type`](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html#gaa2231202dfe12eddca49b1193d44a441) | [`gusb.ISOSyncType`](https://godoc.org/github.com/pzl/usb/gusb#ISOSyncType) | |
| `enum` [`libusb_iso_usage_type`](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html#gab72474551b0eb965401e6febb856007c) | [`gusb.ISOSyncMode`](https://godoc.org/github.com/pzl/usb/gusb#ISOSyncMode) | |
| [`libusb_get_device_descriptor()`](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html#ga5e9ab08d490a7704cf3a9b0439f16f00) | [``]() |
| [`libusb_get_active_config_descriptor()`](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html#ga425885149172b53b3975a07629c8dab3) | [``]() |
| [`libusb_get_config_descriptor()`](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html#gaa635d9aec77de4895dd0896ccf001532) | [``]() |
| [`libusb_get_config_descriptor_by_value()`](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html#ga3e7bffc5d08404c4d6491e73b967bf67) | [``]() |
| [`libusb_get_ss_endpoint_companion_descriptor()`](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html#gade95f9708956c3d45d9969e860fc7241) | [``]() |
| [`libusb_get_bos_descriptor()`](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html#ga68fd9576677c12aa397192916dc49a0b) | [``]() |
| [`libusb_get_usb_2_0_extension_descriptor()`](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html#gac76954b54b97d90c760716fc048b6555) | [``]() |
| [`libusb_get_ss_usb_device_capability_descriptor()`](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html#ga6f9118c2c4c5a42f9e4040e78af63f78) | [``]() |
| [`libusb_get_container_id_descriptor()`](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html#gabe419223add0c8190a940cc4fae19c7f) | [``]() |
| [`libusb_get_string_descriptor_ascii()`](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html#ga240aac96d92cb9f51e3aea79a4adbf42) | [``]() |
| [`libusb_get_descriptor()`](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html#ga9e34f7ecf3817e9bfe77ed09238940df) | [``]() |
| [`libusb_get_string_descriptor()`](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html#ga09103309f98471387561bce0719d88ad) | [``]() |


## [Hotplug](http://libusb.sourceforge.net/api-1.0/group__libusb__hotplug.html)

| Libusb | `usb` equiv | Done? |
|--------|-------------|-------|
| [`libusb_hotplug_register_callback()`](http://libusb.sourceforge.net/api-1.0/group__libusb__hotplug.html#ga00e0c69ddf1fb1b6774dc918192e8dc7) | - ||
| [`libusb_hotplug_deregister_callback()`](http://libusb.sourceforge.net/api-1.0/group__libusb__hotplug.html#ga8110f57eab2064375934f1449b2602bc) | - ||


## [Async IO](http://libusb.sourceforge.net/api-1.0/group__libusb__asyncio.html)
TODO

## [Polling, Timing](http://libusb.sourceforge.net/api-1.0/group__libusb__poll.html)
TODO

## [Synchronous IO](http://libusb.sourceforge.net/api-1.0/group__libusb__syncio.html)

| Libusb | `usb` equiv | Done? |
|--------|-------------|-------|
| [`libusb_control_transfer()`](http://libusb.sourceforge.net/api-1.0/group__libusb__syncio.html#gadb11f7a761bd12fc77a07f4568d56f38) | - ||
| [`libusb_bulk_transfer()`](http://libusb.sourceforge.net/api-1.0/group__libusb__syncio.html#gab8ae853ab492c22d707241dc26c8a805) | - ||
| [`libusb_interrupt_transfer()`](http://libusb.sourceforge.net/api-1.0/group__libusb__syncio.html#gac412bda21b7ecf57e4c76877d78e6486) | - ||



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
| [`libusb_free_config_descriptor`](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html#ga18d2b08a065857ff7ae4f3f719c115cc) | not manually memory managed
| [`libusb_free_ss_endpoint_companion_descriptor`](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html#ga52738e36d2c0e6a0607405d783a2c894) | not manually memory managed
| [`libusb_free_bos_descriptor`](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html#gac9bff9809a7b8663190dda0455998767) | not manually memory managed
| [`libusb_free_usb_2_0_extension_descriptor`](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html#ga5efd69bf490c4359356f5e5bea8dd48c) | not manually memory managed
| [`libusb_free_ss_usb_device_capability_descriptor`](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html#ga8fbf4e49c3ae0c6b7bc2d540ad9a75fb) | not manually memory managed
| [`libusb_free_container_id_descriptor`](http://libusb.sourceforge.net/api-1.0/group__libusb__desc.html#gafeb7c49aa9d3a9c4964f397220d8eab4) | not manually memory managed
| [`libusb_ref_device`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#gaabaa4193adcabba1789cc1165ac41a03) | not reference counted
| [`libusb_unref_device`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga3cc62e6a191b7a9f213e62b81ec30f4d) | not reference counted
| [`libusb_get_port_path`](http://libusb.sourceforge.net/api-1.0/group__libusb__dev.html#ga9d392b8dff7abf5e475c72fd071c3c34) | _deprecated in libusb_
| [`libusb_set_debug`](http://libusb.sourceforge.net/api-1.0/group__libusb__lib.html#ga5f8376b7a863a5a8d5b8824feb8a427a) | _deprecated in libusb_



<!-- Biblio -->

[usb-device]: https://godoc.org/github.com/pzl/usb#Device
[usb-ep]: https://godoc.org/github.com/pzl/usb#Endpoint
