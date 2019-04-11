/*

Package usb is the high-level interface to working with USB devices in pure Go.



Getting Started

The first task is finding your device. You can optionally set a configuration if the device has multiple, but this is rare. Before communicating with the device, you need to claim an interface. Then you may send to the endpoints in that interface.

Libusb Compatibility Table

Most of the things you could achieve with libusb are available here.
	// Device handling
	libusb_device                              -> usb.Device
	libusb_get_device_list                     -> usb.List()
	libusb_get_bus_number                      -> usb.Device.Bus
	libusb_get_port_number                     -> usb.Device.Port
	libusb_get_port_numbers                    -> usb.Device.Ports
	libusb_get_parent                          -> usb.Device.Parent
	libusb_get_device_address                  -> usb.Device.Device
	libusb_get_device_speed                    -> usb.Device.Speed
	libusb_get_max_packet_size                 -> usb.Endpoint.MaxPacketSize
	libusb_get_max_iso_packet_size             -> usb.Endpoint.MaxISOPacketSize
	libusb_open                                -> usb.Open()
	libusb_open_device_with_vid_pid            -> usb.VidPid().Open()
	libusb_close                               -> usb.Device.Close()
	libusb_get_configuration                   -> usb.Device.ActiveConfig
	libusb_set_configuration                   -> usb.Device.SetConfiguration()
	libusb_claim_interface                     -> usb.Interface.Claim()
	libusb_release_interface                   -> usb.Interface.Release()
	libusb_set_interface_alt_setting           -> usb.Interface.SetAlt()
	libusb_clear_halt                          -> TODO
	libusb_reset_device                        -> usb.Device.Reset()
	libusb_kernel_driver_active                -> TODO
	libusb_speed                               -> usb.Speed
	libusb_supported_speed                     -> TODO (reuse usb.Speed?)
	libusb_2_0_extension_attributes            -> TODO
	libusb_ss_usb_device_capability_attributes -> TODO
	libusb_bos_type                            -> TODO

	// Descriptors
	libusb_*_descriptor -> see gusb library

	// Hotplug
	libusb_hotplug_* -> TODO

	// Async IO
	TODO

	// Polling, Timing
	TODO

	// Synchronous IO
	libusb_control_transfer -> TODO
	libusb_bulk_transfer -> TODO
	libusb_interrupt_transfer -> TODO

Things that are not needed with this library:
	libusb_version
	libusb_context
	libusb_option
	libusb_set_option
	libusb_init
	libusb_exit
	libusb_device_handle
	libusb_get_device
	libusb_set_auto_detach_kernel_driver - always on
	libusb_detach_kernel_driver - automatically detached with claim
	libusb_attach_kernel_driver - automatically reattached with release
	libusb_free_device_list
	libusb_ref_device
	libusb_unref_device

Deprecated calls:
	libusb_get_port_path
	libusb_set_debug


*/
package usb

// check http://libusb.sourceforge.net/api-1.0/libusb_caveats.html
