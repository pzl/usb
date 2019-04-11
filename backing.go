package usb

import "errors"

var ErrNotImplemented = errors.New("not implemented")

// something that can fetch these fields
type dataBacking interface {
	// at startup
	getDevNum(Device) (int, error) // usbfs can't determine bus number alone
	getVendorName(Device) (string, error)
	getProductName(Device) (string, error)
	getPort(Device) (int, error)
	getActiveConfig(Device) (int, error)
	getSpeed(Device) (Speed, error)

	// dynamic calls
	getDriver(d Device, intf int) (string, error)
	setConfiguration(Device, int) error
}
