package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/pzl/usb"
)

/*
 * Open a device by bus and device number.
 * Then, for the given interface number, print the
 * current driver assigned by the kernel. Claim that interface, and print
 * the driver now assigned. Then release the interface and print what should
 * be the kernel default driver again.
 */
func main() {
	if len(os.Args) < 4 {
		fmt.Fprintln(os.Stderr, "Arguments required: <bus> <dev> <interface>")
		os.Exit(1)
	}
	bus := mustInt(os.Args[1])
	dev := mustInt(os.Args[2])
	intf := mustInt(os.Args[3])

	device, err := usb.Open(bus, dev)
	if err != nil {
		panic(err)
	}
	defer device.Close()

	drv, _ := device.GetDriver(intf)
	fmt.Printf("kernel-bound driver:          %s\n", drv)

	err = device.ClaimInterface(intf)
	if err != nil {
		panic(err)
	}
	drv, _ = device.GetDriver(intf)
	fmt.Printf("successfully attached driver: %s\n", drv)
	time.Sleep(2000 * time.Millisecond) // hold on to it for 2s

	err = device.ReleaseInterface(intf)
	if err != nil {
		panic(err)
	}
	drv, _ = device.GetDriver(intf)
	fmt.Printf("Released to kernel driver:    %s\n", drv)
}

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
