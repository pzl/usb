package usb

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/apex/log"
	"github.com/pzl/usb/gusb"
)

// @todo: Class,Subclass,Protocol

const badIndexNumber = "invalid %s value: %d"

var (
	ErrDeviceNotFound = errors.New("Device not found")
)

type ID struct {
	ID             uint16 // ID number, e.g. 0xF00D
	nameFromIdFile string // sourced from usb.ids
	nameFromDevice string // sourced from device string descriptors
}

func (i ID) Name() string {
	if i.nameFromIdFile != "" {
		return i.nameFromIdFile
	} else {
		return i.nameFromDevice
	}
}

type Device struct {
	Bus          int
	Device       int
	Port         int // @todo: keep this up to date with hotplugs, resets?
	Ports        []int
	Vendor       ID
	Product      ID
	Parent       *Device
	Speed        Speed
	Configs      []Configuration
	ActiveConfig *Configuration // can read SYSFSPATH/bConfigurationValue

	dataSource dataBacking
	f          *os.File // USBFS file
	sysPath    string   // SYSFS directory for this device
}

func List() ([]*Device, error) {
	dd, err := gusb.Walk(nil)
	if err != nil {
		return nil, err
	}

	devs := make([]*Device, len(dd))

	for i := range dd {
		devs[i] = toDevice(dd[i])
	}
	return devs, nil
}

func Open(bus int, dev int) (*Device, error) {
	l := log.WithFields(log.Fields{"bus": bus, "dev": dev})
	f, err := os.OpenFile(fmt.Sprintf("/dev/bus/usb/%03d/%03d", bus, dev), os.O_RDWR, 0644)
	if os.IsNotExist(err) {
		return nil, ErrDeviceNotFound
	} else if err != nil {
		l.WithError(err).Error("failed opening file")
		return nil, err
	}

	desc, err := gusb.ParseDescriptor(f)
	if err != nil {
		l.WithError(err).Error("failed parsing descriptor")
		return nil, err
	}
	desc.PathInfo.Bus = bus
	desc.PathInfo.Dev = dev
	d := toDevice(desc)
	d.f = f

	return d, nil
}

func VidPid(vid uint16, pid uint16) (*Device, error) {
	var dev *Device

	gusb.Walk(func(dd *gusb.DeviceDescriptor) error {
		if vid == uint16(dd.Vendor) && pid == uint16(dd.Product) {
			dev = toDevice(*dd)
			return filepath.SkipDir
		}
		return nil
	})
	if dev == nil {
		return nil, ErrDeviceNotFound
	}
	return dev, nil
}

func (d *Device) Open() error {
	if d.f != nil {
		d.f.Close()
	}

	f, err := os.OpenFile(fmt.Sprintf("/dev/bus/usb/%03d/%03d", d.Bus, d.Device), os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	d.f = f
	return nil
}

func (d *Device) Close() error {
	if d.f == nil {
		return errors.New("device not open")
	}

	// @todo release any claimed interfaces
	return d.f.Close()
}

func (d *Device) Interface(i int) (*Interface, error) {
	if d.ActiveConfig == nil {
		//@todo what do we do here?
		log.WithField("interface", i).Error("no active config")
		return nil, errors.New("no active config")
	}
	if i < 0 || i > len(d.ActiveConfig.Interfaces)-1 {
		return nil, fmt.Errorf(badIndexNumber, "interface", i)
	}
	return &d.ActiveConfig.Interfaces[i], nil
}

// Return endpoint by it's Address number.
func (d *Device) Endpoint(num int) (*Endpoint, error) {
	if num < 0 {
		return nil, fmt.Errorf(badIndexNumber, "endpoint", num)
	}
	return nil, nil // @todo, look up endpoint
}

func (d *Device) SetConfiguration(cfg int) error {
	err := d.dataSource.setConfiguration(*d, cfg)
	if err != nil {
		d.ActiveConfig = &d.Configs[cfg-1]
	}
	return err
}
func (d *Device) ClaimInterface(intf int) error { // accept int? or Interface?
	i, err := d.Interface(intf)
	if err != nil {
		return err
	}
	return i.Claim()
}
func (d *Device) ReleaseInterface(intf int) error {
	i, err := d.Interface(intf)
	if err != nil {
		return err
	}
	return i.Release()
}
func (d *Device) Reset() error {
	// https://github.com/libusb/libusb/blob/master/libusb/os/linux_usbfs.c#L1629
	return nil
}
func (d *Device) GetDriver(intf int) (string, error) {
	i, err := d.Interface(intf)
	if err != nil {
		return "", err
	}
	return i.GetDriver()
}

type Configuration struct {
	SelfPowered    bool
	RemoteWakeup   bool
	BatteryPowered bool
	MaxPower       int // in mA
	Value          int
	Interfaces     []Interface

	d *Device
}

type Speed int

const (
	SpeedUnknown Speed = iota
	SpeedLow
	SpeedFull
	SpeedHigh
	SpeedWireless
	SpeedSuper
	SpeedSuperPlus
)

func (s Speed) String() string {
	switch s {
	case SpeedUnknown:
		return "Unknown"
	case SpeedLow:
		return "Low, 1.5 Mbps"
	case SpeedFull:
		return "Full, 12Mbps"
	case SpeedHigh:
		return "High, 480 Mbps"
	case SpeedSuper:
		return "Super, 5 Gbps"
	case SpeedSuperPlus:
		return "Super Plus, 10 Gbps"
	}
	return "invalid"
}
