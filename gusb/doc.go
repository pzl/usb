/*
Package gusb contains lower-level USB interaction. If you want to parse descriptors, send Ioctls manually, or other inspections, use this.

Descriptors

This is generally what you see when running lsusb, especially with the verbose flag. It reads the classification information the USB device provides to describe itself. The second byte of a descriptor header is a type constant. In C those IDs are USB_DT_* constants. In Go, they begin with DT. The struct types themselves are:

	DeviceDescriptor
	ConfigDescriptor
	InterfaceDescriptor
	EndpointDescriptor
	StringDescriptor
	DevQualifierDescriptor


*/
package gusb
