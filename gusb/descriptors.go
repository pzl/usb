package gusb

// From:
//  -/usr/include/usb.h (libusb-compat, i.e. v0.1)
//  -/usr/include/linux/usb/ch9.h

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
)

/*
 * --------------------------- CONSTANTS/DEFINES ------------------------------
 */

type DescriptorRange uint8

// USB Descriptor classification mask.
// refers to bDescriptorType
const (
	DescRangeGlobal   DescriptorRange = 0x00
	DescRangeClass    DescriptorRange = 0x20
	DescRangeVendor   DescriptorRange = 0x40
	DescRangeReserved DescriptorRange = 0x60
)

type DT uint8

// Equivalent to USB_DT_* constants.
// determines type via bDescriptorType (aka <descriptor>.Descriptor)
const (
	DTDevice              DT = 0x01 // usb 1.0+
	DTConfig              DT = 0x02
	DTString              DT = 0x03
	DTInterface           DT = 0x04
	DTEndpoint            DT = 0x05
	DTDeviceQualifier     DT = 0x06 // usb 2.0+
	DTOtherSpeed          DT = 0x07
	DTInterfacePower      DT = 0x08
	DTOTG                 DT = 0x09
	DTDebug               DT = 0x0a
	DTInterfaceAssoc      DT = 0x0b
	DTSecurity            DT = 0x0c // wireless USB
	DTKey                 DT = 0x0d
	DTEncType             DT = 0x0e
	DTBOS                 DT = 0x0f
	DTDeviceCapability    DT = 0x10
	DTWirelessEPComp      DT = 0x11
	DTWireAdapter         DT = 0x21 // should these be class-specific?
	DTRPipe               DT = 0x22
	DTCsRadioControl      DT = 0x23
	DTPipeUsage           DT = 0x24 // T10 UAS Spec
	DTSSEndpointComp      DT = 0x30 // usb 3.0+
	DTSSPISOCEndpointComp DT = 0x31 // usb 3.1+
)

func (d DT) String() string {
	global := []string{"invalid", "Device", "Config", "String", "Interface", "Endpoint", "Device Qualifier",
		"Other Speed", "Interface Power", "OTG", "Debug", "Interface Assoc", "Security", "Encryption Type", "BOS", "Device Capability", "Wireless Endpoint Comp", "Wire Adapter"}
	if int(d) < len(global) {
		return global[d]
	}

	switch d {
	/* for now we will ignore these: perhaps they need to be moved to class-specific parsing, but many C headers have these defined with the globals
	case DTWireAdapter:
		return "Wire Adapter"
	case DTRPipe:
		return "RPipe"
	case DTCsRadioControl:
		return "CS Radio Control"
	case DTPipeUsage:
		return "Pipe Usage"
	*/
	case DTSSEndpointComp:
		return "USB 3.0 SS Endpoint Comp"
	case DTSSPISOCEndpointComp:
		return "USB 3.1 SSP ISOC Endpoint Comp"
	}

	var rng string
	switch {
	case d&DT(DescRangeClass) != 0:
		rng = "Class Specific"
	case d&DT(DescRangeVendor) != 0:
		rng = "Vendor Specific"
	case d&DT(DescRangeReserved) != 0:
		rng = "Reserved Class Type"
	default:
		rng = "unknown"
	}
	return fmt.Sprintf("%s: 0x%02x", rng, uint8(d))
}

type USBClass uint8

// USB class IDs https://www.usb.org/defined-class-codes .
// refers to C: bDeviceClass and bInterfaceClass or Go: DeviceDescriptor.Class and InterfaceDescriptor.Class.
const (
	USBClassSeeInterface       USBClass = 0x00 // device, use class information from the interface descriptors
	USBClassAudio              USBClass = 0x01 // interface
	USBClassComm               USBClass = 0x02 // both
	USBClassHID                USBClass = 0x03 // interface
	USBClassPhysical           USBClass = 0x05 // interface
	USBClassStillImage         USBClass = 0x06 // interface
	USBClassPrinter            USBClass = 0x07 // interface
	USBClassMassStorage        USBClass = 0x08 // interface
	USBClassHub                USBClass = 0x09 // device
	USBClassCDCData            USBClass = 0x0a // interface
	USBClassCSCId              USBClass = 0x0b // interface, chip+ smart card
	USBClassContentSec         USBClass = 0x0d // interface, content security
	USBClassVideo              USBClass = 0x0e // interface
	USBClassHealth             USBClass = 0x0f // interface
	USBClassAV                 USBClass = 0x10 // interface
	USBClassBillboard          USBClass = 0x11 // device
	USBClassCBridge            USBClass = 0x12 // interface
	USBClassDiagnostic         USBClass = 0xdc // both
	USBClassWirelessController USBClass = 0xe0 // interface
	USBClassMisc               USBClass = 0xef // both
	USBClassAppSpecific        USBClass = 0xfe // interface
	USBClassVendorSpecific     USBClass = 0xff // both
)

func (c USBClass) String() string {
	s := []string{"Use Interface Class", "Audio", "Comm & CDC Control", "HID", "invalid", "Physical", "Image", "Printer", "Mass Storage", "Hub", "CDC-Data", "Smart Card", "undefined", "Content Security", "Video", "Personal Healthcare", "Audio/Video Devices", "Billboard", "USB Type-C Bridge"}
	if int(c) < len(s) {
		return s[c]
	}

	switch c {
	case USBClassDiagnostic:
		return "Diagnostic"
	case USBClassWirelessController:
		return "Wireless Controller"
	case USBClassMisc:
		return "Miscellaneous"
	case USBClassAppSpecific:
		return "Application Specific"
	case USBClassVendorSpecific:
		return "Vendor Specific"
	}

	return fmt.Sprintf("unknown: 0x%02x", uint8(c))
}

type USBSubClass uint8

// USB Subclassifications.
// refers to C: bDeviceSubClass, bInterfaceSubClass or Go: DeviceDescriptor.SubClass, InterfaceDescriptor.SubClass
const (
	// for USBClassAV
	AVSubclassControlIntf USBSubClass = 0x01 // interface
	AVSubclassVideoStream USBSubClass = 0x02 // interface
	AVSubclassAudioStream USBSubClass = 0x03 // interface

	// for HID, it is used as a boot interface support flag
	HIDBootSupportFalse USBSubClass = 0
	HIDBootSupportTrue  USBSubClass = 1 // 2-255 are reserved

	SubclassVendorSpecific USBSubClass = 0xff
)

//@todo: print subclass info, need class (and protocol?) as context

type USBProtocolDesc uint8

// class/vendor specific protocol values.
// refers to C: bDeviceProtocol, bInterfaceProtocol or Go: DeviceDescriptor.Protocol, InterfaceDescriptor.Protocol
const (
	// for USBClassHub
	HubProtocolFullSpeed     USBProtocolDesc = 0x00
	HubProtocolHiSpeed       USBProtocolDesc = 0x01
	HubProtocolHiSpeedManyTT USBProtocolDesc = 0x02

	// for HID, only defines role for boot interface (if supported)
	HIDBootAsNone     USBProtocolDesc = 0x0
	HIDBootAsKeyboard USBProtocolDesc = 0x1
	HIDBootAsMouse    USBProtocolDesc = 0x2 // 3-255 are reserved
)

//@todo: print protocol info, need class & subclass as context

//@todo: what are these defining?
const (
	USBDescTypeHID      = 0x21
	USBDescTypeReport   = 0x22
	USBDescTypePhysical = 0x23
	USBDescTypeHub      = 0x29
)

/*
 * -------------------------- STRUCTS ---------------------------
 */

//  struct usb_descriptor_header
type DescHeader struct {
	Length     uint8 // bLength
	Descriptor DT    // bDescriptorType, uint8
}

// defines the class,subclass,protocol triplet found in Device and Interface Descriptors
type DescClasses struct {
	Class    USBClass        // b<thing>Class    uint8
	SubClass USBSubClass     // b<thing>SubClass uint8
	Protocol USBProtocolDesc // b<thing>Protocol uint8
}

func (dc DescClasses) String() string {
	//@todo: audio.h, cdc.h, video.h
	sub := strconv.Itoa(int(dc.SubClass))
	prot := strconv.Itoa(int(dc.Protocol))

	switch dc.Class {
	case USBClassAV:
		switch dc.SubClass {
		case AVSubclassControlIntf:
			sub = "Control Interface"
		case AVSubclassVideoStream:
			sub = "Video Stream"
		case AVSubclassAudioStream:
			sub = "Audio Stream"
		}
	case USBClassHID:
		switch dc.SubClass {
		case HIDBootSupportFalse:
			sub = "not boot device"
		case HIDBootSupportTrue:
			sub = "Boot device"
			switch dc.Protocol {
			case HIDBootAsNone:
				prot = "Boot as none"
			case HIDBootAsKeyboard:
				prot = "keyboard"
			case HIDBootAsMouse:
				prot = "mouse"
			default:
				prot = "invalid HID boot device: " + strconv.Itoa(int(dc.Protocol))
			}
		default:
			sub = "invalid HID subclass: " + strconv.Itoa(int(dc.SubClass))
		}
	}

	return fmt.Sprintf("Class: %s, SubClass: %s, Protocol: %s", dc.Class, sub, prot)
}

//  DEVICE

// USB_DT_DEVICE, aka DTDevice. (struct usb_device_descriptor)
type DeviceDescriptor struct {
	DescHeader
	USBVer        USBVer // bcdUSB, uint16
	DescClasses          // 3 * uint8. Class, Subclass, Protocol
	MaxPacketSize uint8  // for endpoint 0. One of 8,16,32,64
	Vendor        USBID  // idVendor,  uint16
	Product       USBID  // idProduct, uint16
	Version       USBVer // bcdDevice, uint16
	ManufStr      uint8  // iManufacturer
	ProductStr    uint8  // iProduct
	SerialStr     uint8  // iSerial
	NumConfigs    uint8  // bNumConfigurations

	Configs   []ConfigDescriptor
	extradata []byte // @todo: parse and fill

	// internal use, not part of Descriptor spec
	PathInfo DevicePath
}

func NewDevice(b []byte) (DeviceDescriptor, error) {
	const DFSize = 18
	if len(b) < DFSize {
		return DeviceDescriptor{}, errors.New("not enough bytes to construct device descriptor")
	}

	dev := DeviceDescriptor{
		DescHeader: DescHeader{
			Length:     b[0],
			Descriptor: DT(b[1]),
		},
		USBVer: USBVer(binary.LittleEndian.Uint16(b[2:])),
		DescClasses: DescClasses{
			Class:    USBClass(b[4]),
			SubClass: USBSubClass(b[5]),
			Protocol: USBProtocolDesc(b[6]),
		},
		MaxPacketSize: b[7],
		Vendor:        USBID(binary.LittleEndian.Uint16(b[8:])),
		Product:       USBID(binary.LittleEndian.Uint16(b[10:])),
		Version:       USBVer(binary.LittleEndian.Uint16(b[12:])),
		ManufStr:      b[14],
		ProductStr:    b[15],
		SerialStr:     b[16],
		NumConfigs:    b[17],
		Configs:       make([]ConfigDescriptor, b[17]),
	}
	if len(b) > DFSize {
		dev.extradata = b[DFSize:]
	}
	return dev, nil
}

func (d DeviceDescriptor) String() string {
	return fmt.Sprintf("%s %s:%s, USB Ver %s, Class: %s, SubClass: %d, Protocol: %d, Release: %s, Serial: %d, Manuf: %d, Product: %d, EP 0 Max Packet: %db, NumConfigs: %d", d.Descriptor, d.Vendor, d.Product, d.USBVer, d.Class, d.SubClass, d.Protocol, d.Version, d.SerialStr, d.ManufStr, d.ProductStr, d.MaxPacketSize, d.NumConfigs)
}

/*
 * String Descriptor
 */

//  struct usb_string_descriptor
// bDescriptorType, C: USB_DT_STRING, Go: DescString
type StringDescriptor struct {
	DescHeader
	S string
}

func NewString(b []byte) (StringDescriptor, error) {
	if len(b) < 2 {
		return StringDescriptor{}, errors.New("not enough bytes to create String Descriptor")
	}
	return StringDescriptor{
		DescHeader: DescHeader{
			Length:     b[0],
			Descriptor: DT(b[1]),
		},
		S: string(b[2:]),
	}, nil
}
func (s StringDescriptor) String() string { return s.S }

/*
 * Endpoint Descriptor
 */

//  struct usb_endpoint_descriptor
// bDescriptorType: C: USB_DT_ENDPOINT, Go: DescEndpoint
type EndpointDescriptor struct { // leftovers & interpreted
	//@todo: bRefresh && bSynchAddress provided via audio endpoints. See ch9.h, line 410
	DescHeader
	Address       EndpointAddress //uint8
	Attributes    uint8
	MaxPacketSize uint16
	Interval      uint8
	TransferType  TransferType // parsed from Attributes
	ISOSyncType   ISOSyncType  // parsed from Attributes
	ISOSyncMode   ISOSyncMode  // parsed from Attributes
	extradata     []byte
}

func NewEndpoint(b []byte) (EndpointDescriptor, error) {
	const (
		EFSize           = 7
		EndpointTypeMask = 0x3      // Attributes->TransferType
		ISOSyncMask      = 0x3 << 2 // Attributes->IsoSyncType
		ISOModeMask      = 0x3 << 4 // Attributes->IsoSyncMode
	)
	if len(b) < EFSize {
		return EndpointDescriptor{}, errors.New("not enough bytes to create Endpoint Descriptor")
	}
	e := EndpointDescriptor{
		DescHeader: DescHeader{
			Length:     b[0],
			Descriptor: DT(b[1]),
		},
		Address:       EndpointAddress(b[2]),
		Attributes:    b[3],
		MaxPacketSize: binary.LittleEndian.Uint16(b[4:]),
		Interval:      b[6],
		TransferType:  TransferType(b[3] & EndpointTypeMask),
	}

	if e.TransferType == EndpointTypeIsochronous {
		e.ISOSyncType = ISOSyncType(e.Attributes & ISOSyncMask)
		e.ISOSyncMode = ISOSyncMode(e.Attributes & ISOModeMask)
	}
	if len(b) > EFSize {
		e.extradata = b[EFSize:]
	}
	return e, nil
}

func (e EndpointDescriptor) String() string {
	return fmt.Sprintf("%s %s (0x%02x), Type: %s. Max Packet: %db. [%s]", e.Descriptor, e.Address, uint8(e.Address), e.TransferType, e.MaxPacketSize, e.extradata)
}

type TransferType int

const (
	EndpointTypeControl TransferType = iota
	EndpointTypeIsochronous
	EndpointTypeBulk
	EndpointTypeInterrupt
	EndpointTypeBulkStream
)

func (t TransferType) String() string {
	return []string{"Ctrl", "Isochronous", "Bulk", "Interrupt", "Bulk Stream"}[t]
}

type ISOSyncType int

const (
	ISOSyncTypeNone ISOSyncType = iota
	ISOSyncTypeAsync
	ISOSyncTypeAdaptive
	ISOSyncTypeSync
)

type ISOSyncMode int

const (
	ISOUsageData ISOSyncMode = iota
	ISOUsageFeedback
	ISOUsageImplicit
)

type EndpointDirection uint8

const (
	EndpointDirOUT EndpointDirection = iota
	EndpointDirIN
)

func (ed EndpointDirection) String() string { return []string{"OUT", "IN"}[ed] }

const endpointAddrMask = 0x0F // EndpointAddress
type EndpointAddress uint8

func (e EndpointAddress) Dir() EndpointDirection { return EndpointDirection(e >> 7) }
func (e EndpointAddress) Num() int               { return int(e & endpointAddrMask) }
func (e EndpointAddress) String() string         { return fmt.Sprintf("%d %s", e.Num(), e.Dir()) }

/*
 * Interface Descriptors
 */

const (
	mAXINTERFACES = 32
	mAXALTERNATES = 128
)

// struct usb_interface_descriptor
type InterfaceDescriptor struct {
	DescHeader
	InterfaceNumber  uint8
	AlternateSetting uint8
	NumEndpoints     uint8
	DescClasses      // 3 * uint8. Class,Subclass,Protocol
	StrIndex         uint8
	Endpoints        []EndpointDescriptor
	extradata        []byte
}

func NewInterface(b []byte) (InterfaceDescriptor, error) {
	const IFSize = 9
	if len(b) < IFSize {
		return InterfaceDescriptor{}, errors.New("not enough bytes to create Interface Descriptor")
	}
	interf := InterfaceDescriptor{
		DescHeader: DescHeader{
			Length:     b[0],
			Descriptor: DT(b[1]),
		},
		InterfaceNumber:  b[2],
		AlternateSetting: b[3],
		NumEndpoints:     b[4],
		DescClasses: DescClasses{
			Class:    USBClass(b[5]),
			SubClass: USBSubClass(b[6]),
			Protocol: USBProtocolDesc(b[7]),
		},
		StrIndex:  b[8],
		Endpoints: make([]EndpointDescriptor, b[4]),
	}
	if len(b) > IFSize {
		interf.extradata = b[IFSize:]
	}
	return interf, nil
}

func (i InterfaceDescriptor) String() string {
	return fmt.Sprintf("%s %d, Alternate: %d. Endpoints: %d. %s. Str Index: %d. Extra: [%v]", i.Descriptor, i.InterfaceNumber, i.AlternateSetting, i.NumEndpoints, i.DescClasses, i.StrIndex, i.extradata)
}

/*
 * Configuration Descriptor
 */

const mAXCONFIGS = 8

// struct usb_config_descriptor
type ConfigDescriptor struct {
	DescHeader
	TotalLength    uint16 // wTotalLength
	NumInterfaces  uint8  // bNumInterfaces
	Value          uint8  // bConfigurationValue
	StrIndex       uint8  // iConfiguration
	Attributes     uint8  // bmAttributes
	MaxPower       uint8  // MaxPower
	SelfPowered    bool   // Attributes https://www.beyondlogic.org/usbnutshell/usb5.shtml#ConfigurationDescriptors
	RemoteWakeup   bool   // Attributes
	BatteryPowered bool   // Attributes (ch9.h)
	Interfaces     []InterfaceDescriptor
	extradata      []byte
}

func NewConfig(b []byte) (ConfigDescriptor, error) {
	const (
		CFSize        = 9
		BattPowerMask = (1 << 4)
		WakeupMask    = (1 << 5)
		SelfPowerMask = (1 << 6)
	)
	if len(b) < CFSize {
		return ConfigDescriptor{}, errors.New("not enough bytes to create Config Descriptor")
	}

	config := ConfigDescriptor{
		DescHeader: DescHeader{
			Length:     b[0],
			Descriptor: DT(b[1]),
		},
		TotalLength:    binary.LittleEndian.Uint16(b[2:]),
		NumInterfaces:  b[4],
		Value:          b[5],
		StrIndex:       b[6],
		Attributes:     b[7],
		MaxPower:       b[8],
		Interfaces:     make([]InterfaceDescriptor, b[4]),
		RemoteWakeup:   b[7]&WakeupMask != 0,
		SelfPowered:    b[7]&SelfPowerMask != 0,
		BatteryPowered: b[7]&BattPowerMask != 0,
	}
	if len(b) > CFSize {
		config.extradata = b[CFSize:]
	}
	return config, nil
}

func (cf ConfigDescriptor) String() string {
	return fmt.Sprintf("%s %d, Interfaces: %d. StrIndex: %d. Max Power: %dmA. Battery Powered: %t. Self Powered: %t. Remote Wakeup: %t", cf.Descriptor, cf.Value, cf.NumInterfaces, cf.StrIndex, cf.MaxPower*2, cf.BatteryPowered, cf.SelfPowered, cf.RemoteWakeup)
}

//	struct usb_qualifer_descriptor
type DevQualifierDescriptor struct {
	DescHeader
	Version       USBVer
	DescClasses   // 3 * uint8. Class, SubClass, Protocol
	MaxPacketSize uint8
	NumConfigs    uint8
	Reserved      uint8
}

func NewDevQualifier(b []byte) (DevQualifierDescriptor, error) {
	const DQSize = 10
	if len(b) < DQSize {
		return DevQualifierDescriptor{}, errors.New("not enough bytes to create Device Qualifier Descriptor")
	}
	return DevQualifierDescriptor{
		DescHeader: DescHeader{
			Length:     b[0],
			Descriptor: DT(b[1]),
		},
		Version: USBVer(binary.LittleEndian.Uint16(b[2:])),
		DescClasses: DescClasses{
			Class:    USBClass(b[4]),
			SubClass: USBSubClass(b[5]),
			Protocol: USBProtocolDesc(b[6]),
		},
		MaxPacketSize: b[7],
		NumConfigs:    b[8],
		Reserved:      b[9],
	}, nil
}

//@todo: Define the SSEPComp & SSPISOC structs for completeness. ch9.h:670
//@todo: OTG descriptors
//@todo: Interface Assoc Descriptor

type USBVer uint16

func (u USBVer) String() string {
	return fmt.Sprintf("%d.%d", uint16(u>>8), uint16(u&0xff))
}

type USBID uint16

func (id USBID) String() string {
	return fmt.Sprintf("%04x", uint16(id))
}
