package usb_test

import (
	"fmt"

	"github.com/pzl/usb"
)

func ExampleList() {
	devices, err := usb.List()
	if err != nil {
		//handle
	}

	for _, d := range devices {
		fmt.Printf("%04x:%04x - %s, %s\n", d.Vendor.ID, d.Product.ID, d.Vendor.Name(), d.Product.Name())
	}
}

func ExampleOpen() {
	dev, err := usb.Open(1, 3)
	if err != nil {
		//handle error
	}
	defer dev.Close()

	//do something
}

func Example() {
	dev, err := usb.VidPid(0x0c45, 0x6300) // opens the first device it finds with these
	if err == usb.ErrDeviceNotFound {
		fmt.Println("Not found")
		return
	}
	err = dev.Open()
	if err != nil {
		fmt.Printf("Error opening device: %v\n", err)
		return
	}
	defer dev.Close()

	err = dev.ClaimInterface(1)
	if err != nil {
		fmt.Printf("Error claiming interface: %v\n", err)
		return
	}
	defer dev.ReleaseInterface(1)

	// @todo this is super ugly
	dev.ActiveConfig.Interfaces[1].Endpoints[1].CtrlTransfer( /*...*/ )
}
