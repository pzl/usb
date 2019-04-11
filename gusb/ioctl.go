package gusb

import (
	"bytes"
	"encoding/binary"
	"os"
	"unsafe"

	"golang.org/x/sys/unix"
)

// Hand-craft an IOCTL to send to an open file descriptor.
// data must be a pointer.
func Ioctl(f *os.File, ioctl IoctlRequest, data interface{}) (int, error) {
	// USB explicitly uses LE byte order. Serialize to pass to kernel
	b := new(bytes.Buffer)
	if err := binary.Write(b, binary.LittleEndian, data); err != nil {
		return -1, err
	}
	// the conversion from unsafe.Pointer to uintptr MUST
	// occur in the call expression. For compiler to recognize
	// this pattern, and have the GC not muck with things
	//nolint:unconvert
	r, _, err := unix.Syscall(
		unix.SYS_IOCTL,                           // ioctl
		uintptr(f.Fd()),                          // file
		uintptr(uint32(ioctl)),                   // request
		uintptr(unsafe.Pointer(&(b.Bytes()[0]))), // argument
	)
	if err != 0 {
		//return -1, os.NewSyscallError("ioctl", err)
		return int(r), err
	}
	// read back the (possibly) kernel-modified bytes into the original struct given
	if err := binary.Read(b, binary.LittleEndian, data); err != nil {
		return int(r), err // @todo: more user-friendly error what's going on here. Ioctl may have succeeded, but parsing failed
	}
	return int(r), nil
}

/*
Can be used to calculate an IOCTL number dynamically. Here's an example translation from the C def for USBDEVFS_CONTROL
	#define USBDEVFS_CONTROL     _IOWR('U', 0, struct usbdevfs_ctrltransfer)
WR means read and write, so both should be set to true. char is 'U', which is true for all the usbfs subsystem. num is 0 (second param). size is: sizeof that struct (0x18 on 64 bit).

so call Ioctlnum(true, true, 0, 0x18) will give you 0xC0185500 = USBDEVFS_CONTROL.
*/
func Ioctlnum(read bool, write bool, num uint8, size uint16) uint32 {
	const IOCTL_USB_CHAR = 'U' // 0x55
	r, w := 0, 0
	if read {
		r++
	}
	if write {
		w++
	}
	return uint32(num) |
		uint32(IOCTL_USB_CHAR)<<8 |
		uint32(size)<<16 |
		uint32(w)<<30 |
		uint32(r)<<31
}
