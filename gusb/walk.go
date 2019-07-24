package gusb

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/apex/log"
)

type DevicePath struct {
	Bus     int
	Dev     int
	SysPath string
}

func support(p string) bool {
	_, err := os.Stat(p)
	return !os.IsNotExist(err)
}

func Walk(cb walkCB) ([]DeviceDescriptor, error) {
	// if Linux kernel 2.6.26 +
	// we can get most of the information from sysfs (/sys/bus/usb/devices..)
	// instead of usbfs (/dev/bus/usb...). Usbfs is occasionally slower and wakes
	// up USB devices.
	const (
		SYSFS = "/sys/bus/usb/devices"
		USBFS = "/dev/bus/usb"
	)

	useSys := support(SYSFS)
	useUSB := support(USBFS)

	if !useSys && !useUSB {
		return nil, fmt.Errorf("Not supported. Could not find %s or %s", SYSFS, USBFS)
	}
	if useSys {
		return walker(SYSFS, walkSysFs, cb)
	} else {
		return walker(USBFS, walkUsbFs, cb)
	}
}

type walkCB func(*DeviceDescriptor) error

type walkMethod func(path string, info os.FileInfo) (DeviceDescriptor, error)

func walker(tree string, method walkMethod, cb walkCB) ([]DeviceDescriptor, error) {
	devs := make([]DeviceDescriptor, 0, 20) // randomly preallocate some space. Because I wanted to?

	err := filepath.Walk(tree, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == tree {
			return nil
		}
		d, err := method(path, info)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return err
		}

		if d.Length != 0 {
			// callback if supplied, to modify device and/or stop iteration
			if cb != nil {
				err := cb(&d)
				if err != nil {
					return err
				}
			}
			devs = append(devs, d)
		}
		return nil
	})
	return devs, err
}

// SYSFS structure:
// - usbX
// 		X: ID of host controller on the machine
// - X-A.B.C
//		X: Host Controller ID as above
//		A.B.C: Physical path to port where device is connected
// - X-A.B.C:Y.Z
//		X-A.B.C: as above
//		Y: Active configuration
//		Z: bInterfaceNumber

func walkSysFs(path string, info os.FileInfo) (DeviceDescriptor, error) {
	name := info.Name()
	//ch, _ := utf8.DecodeRuneInString(name)
	if /*!unicode.IsDigit(ch) || name[:3] == "usb" ||*/ strings.Contains(name, ":") {
		return DeviceDescriptor{}, nil
	}
	f, err := os.Open(filepath.Join(path, "descriptors"))
	if err != nil {
		return DeviceDescriptor{}, err
	}
	defer f.Close()
	dsc, err := ParseDescriptor(f)
	if err != nil {
		return dsc, err
	}
	dsc.PathInfo.SysPath = path

	return dsc, nil
}

func walkUsbFs(path string, info os.FileInfo) (DeviceDescriptor, error) {
	f, err := os.Open(path)
	if err != nil {
		return DeviceDescriptor{}, err
	}
	defer f.Close()
	dsc, err := ParseDescriptor(f)
	if err != nil {
		return dsc, err
	}

	dev, err := strconv.Atoi(info.Name())
	if err != nil {
		return dsc, err
	}
	bus, err := strconv.Atoi(filepath.Base(filepath.Dir(path)))
	if err != nil {
		return dsc, err
	}

	dsc.PathInfo.Bus = bus
	dsc.PathInfo.Dev = dev
	return dsc, nil
}

func ParseDescriptor(r io.Reader) (DeviceDescriptor, error) {
	var dev DeviceDescriptor
	var curConf int
	var curIntf int
	var curEp int

	f, err := ioutil.ReadAll(r)
	if err != nil {
		return dev, err
	}

	buf := bytes.NewBuffer(f)
	epNumForInterf := map[int]int{}

	for buf.Len() > 0 {
		if length, err := buf.ReadByte(); err != nil {
			return dev, err
		} else if err := buf.UnreadByte(); err != nil {
			return dev, err
		} else {
			body := make([]byte, length)
			if n, err := buf.Read(body); err != nil {
				return dev, err
			} else if n != int(length) || length < 2 {
				return dev, errors.New("short read")
			} else {
				h := DescHeader{
					Length:     body[0],
					Descriptor: DT(body[1]),
				}
				switch h.Descriptor {
				case DTDevice:
					dev, err = NewDevice(body)
					if err != nil {
						return dev, err
					}
				case DTConfig:
					cfg, err := NewConfig(body)
					if err != nil {
						return dev, err
					}
					curConf = int(cfg.Value - 1) // not zero-based
					dev.Configs[curConf] = cfg
				case DTString:
					//dsc, err := NewString(body) don't know what to do here
				case DTInterface:
					intf, err := NewInterface(body)
					if err != nil {
						return dev, err
					}
					curIntf = int(intf.InterfaceNumber)
					epNumForInterf[curIntf] = 0
					dev.Configs[curConf].Interfaces[curIntf] = intf
				case DTEndpoint:
					ep, err := NewEndpoint(body)
					if err != nil {
						return dev, err
					}
					curEp = epNumForInterf[curIntf]
					epNumForInterf[curIntf]++
					dev.Configs[curConf].Interfaces[curIntf].Endpoints[curEp] = ep
				default:
					log.WithFields(log.Fields{
						"descriptor": h.Descriptor.String(),
						"length":     h.Length,
						"body":       body[2:],
					}).Debug("got unknown descriptor")
					continue
				}
			}
		}
	}
	return dev, nil
}
