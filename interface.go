package usb

import (
	"github.com/pzl/usb/gusb"
)

type Interface struct {
	ID        int // interface number
	Alternate int
	Endpoints []Endpoint

	d *Device
	//@todo: isKernelDriverActive -- should it be a `Driver string` property? method? bool?
}

// Kernel interface release handled automatically
func (i *Interface) Claim() error {
	//SYSFS method: @todo
	//	write interface basename to SYSFS_PATH/drivers/DRIVERNAME/unbind
	//	write interface basename to SYSFS_PATH/drivers/usbfs/bind

	// USBFS method - IOCTL
	return gusb.Claim(i.d.f, int32(i.ID))
}

// Kernel interface re-claim handled automatically
func (i *Interface) Release() error {
	//SYSFS method: @todo
	//	write interface basename to SYSFS_PATH/drivers/usbfs/unbind
	//	... not sure we can tell kernel to rebind to the appropriate driver by ourself? perhaps the uevent file?

	// USBFS method - IOCTL
	return gusb.Release(i.d.f, int32(i.ID))
}

func (i *Interface) SetAlt() error {
	return nil //@todo
}

func (i *Interface) GetDriver() (string, error) {
	return i.d.dataSource.getDriver(*i.d, i.ID)
}
