package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pzl/usb"
)

/*
 * Locate a device using VID and PID
 * then print its parent tree
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

	device, err := usb.VidPid(uint16(vid), uint16(pid))
	if err == usb.ErrDeviceNotFound {
		fmt.Println("Device Not found")
		return
	} else if err != nil {
		panic(err)
	}

	fmt.Printf("Device: %s\n", printDevice(device))
	for p, i := device.Parent, 1; p != nil; p, i = p.Parent, i+1 {
		fmt.Printf("%s⮡ %s\n", strings.Repeat(" ", i), printDevice(p))
	}
}

func printDevice(d *usb.Device) string {
	return fmt.Sprintf("Bus %03d Device %03d: ID %04x:%04x %s %s", d.Bus, d.Device, d.Vendor.ID, d.Product.ID, d.Vendor.Name(), d.Product.Name())
}
