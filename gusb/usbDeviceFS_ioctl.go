package gusb

import "unsafe"

// these ones are common across 32, 64- bit

// https://git.kernel.org/pub/scm/linux/kernel/git/stable/linux.git/plain/Documentation/ioctl/ioctl-decoding.txt?id=HEAD
// LSByte is command ID (0, 2, 3, 9...)
// 0x55 is 'U' for USB
// MSBit is W mode, next bit is R mode.
//   0xc = _IOWR
//   0x8 = _IOW
//   0x4 = _IOR
//   0x0 = _IO
// remaining bits (14bit) are size

type IoctlRequest uint32

// Universal IOCTL numbers
const (
	USBDEVFS_SETINTERFACE     IoctlRequest = 0x80085504
	USBDEVFS_SETCONFIGURATION IoctlRequest = 0x80045505
	USBDEVFS_GETDRIVER        IoctlRequest = 0x41045508
	USBDEVFS_CONNECTINFO      IoctlRequest = 0x40085511
	USBDEVFS_CLAIMINTERFACE   IoctlRequest = 0x8004550f
	USBDEVFS_RELEASEINTERFACE IoctlRequest = 0x80045510
	USBDEVFS_DISCONNECT_CLAIM IoctlRequest = 0x8108551b
	USBDEVFS_RESETEP          IoctlRequest = 0x80045503
	USBDEVFS_CLEAR_HALT       IoctlRequest = 0x80045515
	USBDEVFS_CLAIM_PORT       IoctlRequest = 0x80045518
	USBDEVFS_RELEASE_PORT     IoctlRequest = 0x80045519
	USBDEVFS_HUB_PORTINFO     IoctlRequest = 0x80805513
	USBDEVFS_GET_CAPABILITIES IoctlRequest = 0x8004551a
	USBDEVFS_ALLOC_STREAMS    IoctlRequest = 0x8008551c
	USBDEVFS_FREE_STREAMS     IoctlRequest = 0x8008551d
	USBDEVFS_DROP_PRIVILEGES  IoctlRequest = 0x4004551e
	USBDEVFS_DISCARDURB       IoctlRequest = 0x0000550b
	USBDEVFS_RESET            IoctlRequest = 0x00005514
	USBDEVFS_DISCONNECT       IoctlRequest = 0x00005516
	USBDEVFS_CONNECT          IoctlRequest = 0x00005517
	USBDEVFS_GET_SPEED        IoctlRequest = 0x0000551f

/*
	USBDEVFS_CONTROL32       = 0xc0105500
	USBDEVFS_BULK32          = 0xc0105502
	USBDEVFS_DISCSIGNAL32    = 0x8000550e
	USBDEVFS_SUBMITURB32     = 0x8000550a
	USBDEVFS_IOCTL32         = 0xc0005512
	USBDEVFS_REAPURB32       = 0x4000550c
	USBDEVFS_REAPURBNDELAY32 = 0x4000550d
*/
)

func slicePtr(b []byte) VoidPtr {
	return VoidPtr(uintptr(unsafe.Pointer(&b[0])))
}
