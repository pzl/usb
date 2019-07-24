/*

Package usb is the high-level interface to working with USB devices in pure Go.



Getting Started

The first task is finding your device. You can optionally set a configuration if the device has multiple, but this is rare. Before communicating with the device, you need to claim an interface. Then you may send to the endpoints in that interface.


*/
package usb

// check http://libusb.sourceforge.net/api-1.0/libusb_caveats.html
