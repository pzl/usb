package main

import (
	"fmt"

	"github.com/pzl/usb"
)

func main() {
	devs, err := usb.List()
	if err != nil {
		panic(err)
	}
	for _, d := range devs {
		fmt.Printf("Bus %03d Device %03d: ID %04x:%04x %s %s\n", d.Bus, d.Device, d.Vendor.ID, d.Product.ID, d.Vendor.Name(), d.Product.Name())
	}
}
