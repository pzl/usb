package usb

type Interface struct {
	ID        int // interface number
	Alternate int
	Endpoints []Endpoint

	d *Device
	//@todo: isKernelDriverActive -- should it be a `Driver string` property? method? bool?
}

// Kernel interface release handled automatically
func (i *Interface) Claim() error { return backingUsbfs{}.claim(*i) }

// Kernel interface re-claim handled automatically
func (i *Interface) Release() error { return backingUsbfs{}.release(*i) }

func (i *Interface) SetAlt() error {
	return nil //@todo
}

func (i *Interface) GetDriver() (string, error) {
	return i.d.dataSource.getDriver(*i.d, i.ID)
}
