// +build 386 amd64p32 arm armbe mips mipsle mips64p32 mips64p32le ppc s390 sparc

package gusb

// 32-bit-specific IOCTL numbers
const (
	USBDEVFS_CONTROL       IoctlRequest = 0xc0105500 // struct CtrlTransfer
	USBDEVFS_BULK          IoctlRequest = 0xc0105502 // struct BulkTransfer
	USBDEVFS_SUBMITURB     IoctlRequest = 0x802c550a // struct Urb
	USBDEVFS_REAPURB       IoctlRequest = 0x4004550c // void *
	USBDEVFS_REAPURBNDELAY IoctlRequest = 0x4004550d // void *
	USBDEVFS_DISCSIGNAL    IoctlRequest = 0x8008550e // struct DisconnectSignal
	USBDEVFS_IOCTL         IoctlRequest = 0xc00c5512 // struct IoctlPacket
)

// Matches width of Kernel void * on each platform.
type VoidPtr uint32 // 32-bit def
type pad4 struct{}  // supposedly takes up 0 bytes
