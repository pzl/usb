package gusb

// from /usr/include/linux/usbdevice_fs.h

// Below are the structs & data types sent through Ioctl for each request type

type CtrlTransfer struct {
	RequestType uint8
	Request     uint8
	Value       uint16
	Index       uint16
	Length      uint16
	Timeout     uint32  // milliseconds
	_           pad4    // 4 bytes on 64 bit, 0b on 32
	Data        VoidPtr // void *
}

type BulkTransfer struct {
	Ep      uint32
	Len     uint32
	Timeout uint32 //ms
	_       pad4   //4b padding on 64bit, 0 on 32bit
	Data    VoidPtr
}

type SetInterface struct {
	Interface  uint32
	AltSetting uint32
}

// 8 on pi (4 + 4), 16 on desktop (4 + _4 + 8)
type DisconnectSignal struct {
	Signr uint32 // unsigned int
	_     pad4
	Data  VoidPtr
}

const MAXDRIVERNAME = 255

type GetDriverS struct {
	Interface uint32
	Driver    [MAXDRIVERNAME + 1]byte
}

type ConnectInfo struct {
	Devnum uint32
	Slow   uint8 // unsigned char
}

// this is super not correct
// 44,56
type URB struct {
	Type         uint8
	Endpoint     uint8
	Status       int32
	Flags        uint32
	Buffer       VoidPtr
	BufferLength int32
	ActualLength int32
	StartFrame   int32
	//fucking, a union?!
	ErrorCount   int32
	Signr        uint32
	UserContext  VoidPtr
	IsoFrameDesc struct { // 12,12
		Length       uint32
		ActualLength uint32
		Status       uint32
	}
}

type IoctlPacket struct { //usbdevfs_ioctl
	IfNo      int32 //interface number
	IoctlCode int32
	Data      VoidPtr
}

type HubPortinfo struct {
	NPorts int8
	Port   [127]int8
}

type DisconnectClaim struct {
	Interface uint32
	Flags     uint32
	Driver    [MAXDRIVERNAME + 1]byte
}

type Streams struct {
	NumStreams uint32 //not used by FREE_STREAMS
	NumEps     uint32
	Eps        []uint8
}

type DeviceSpeed int

const (
	SpeedUnknown DeviceSpeed = iota
	SpeedLow
	SpeedFull
	SpeedHigh
	SpeedWireless
	SpeedSuper
	SpeedSuperPlus
)

// use uint32 as argument for
// ClaimInterface ReleaseInterface ResetEp SetConfiguration
// ClearHalt ClaimPort ReleasePort GetCapabilites DropPrivileges
