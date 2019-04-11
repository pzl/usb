package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pzl/usb"
)

/*
 * Locate a device using it's VID and PID
 *
 */
func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Arguments required: <vid> <pid>")
		os.Exit(1)
	}
	vid, err := strconv.ParseUint(strings.TrimPrefix(os.Args[1], "0x"), 16, 16)
	if err != nil {
		panic(err)
	}
	pid, err := strconv.ParseUint(strings.TrimPrefix(os.Args[2], "0x"), 16, 16)
	if err != nil {
		panic(err)
	}

	fmt.Printf("looking for %04x:%04x\n", vid, pid)

	device, err := usb.VidPid(uint16(vid), uint16(pid))
	if err == usb.ErrDeviceNotFound {
		fmt.Println("Device Not found")
		return
	} else if err != nil {
		panic(err)
	}

	fmt.Printf("Device: %s %s\n", device.Vendor.Name(), device.Product.Name())
}
