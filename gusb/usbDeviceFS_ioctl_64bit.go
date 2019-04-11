// +build amd64 arm64 arm64be ppc64 ppc64le mips64 mips64le s390x sparc64

package gusb

// 64-bit specific IOCTL numbers
const (
	USBDEVFS_CONTROL       IoctlRequest = 0xc0185500 // struct CtrlTransfer
	USBDEVFS_BULK          IoctlRequest = 0xc0185502 // struct BulkTransfer
	USBDEVFS_SUBMITURB     IoctlRequest = 0x8038550a // struct Urb
	USBDEVFS_REAPURB       IoctlRequest = 0x4008550c // void *
	USBDEVFS_REAPURBNDELAY IoctlRequest = 0x4008550d // void *
	USBDEVFS_DISCSIGNAL    IoctlRequest = 0x8010550e // struct DisconnectSignal
	USBDEVFS_IOCTL         IoctlRequest = 0xc0105512 // struct IoctlPacket
)

// Matches width of Kernel void * on each platform.
type VoidPtr uint64 // 64-bit width
type pad4 uint32
