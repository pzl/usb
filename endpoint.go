package usb

type Endpoint struct {
	Address          int
	TransferType     int
	MaxPacketSize    int
	MaxISOPacketSize int

	i *Interface
}

/* ---- Synchronous Sending ---- */

func (e *Endpoint) CtrlTransfer() {

}

func (e *Endpoint) Bulk() {

}

func (e *Endpoint) Interrupt() {

}
