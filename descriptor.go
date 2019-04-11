package usb

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/apex/log"
	"github.com/pzl/usb/gusb"
)

/* ---------- Descriptors to library-native objects ---------- */

func toDevice(dd gusb.DeviceDescriptor) *Device {
	var err error
	vid := uint16(dd.Vendor)
	pid := uint16(dd.Product)

	d := &Device{
		Bus:     dd.PathInfo.Bus,
		Device:  dd.PathInfo.Dev,
		sysPath: dd.PathInfo.SysPath,
		Vendor:  ID{ID: vid, nameFromIdFile: vendorName(vid)},
		Product: ID{ID: pid, nameFromIdFile: productName(vid, pid)},
		Configs: make([]Configuration, dd.NumConfigs),
	}
	for _, c := range dd.Configs {
		d.Configs[c.Value-1] = toConfig(c, d)
	}
	// walk sysfs path to find matching device, and set d.sysPath
	if d.sysPath == "" {
		d.sysPath = getSysfsFromBusDev(d.Bus, d.Device)
	}

	if d.sysPath != "" {
		d.dataSource = backingSysfs{} //@todo: fall back to usbfs for failures?
	} else {
		d.dataSource = backingUsbfs{}
	}

	if d.Device <= 0 {
		if dev, err := d.dataSource.getDevNum(*d); err != nil {
			log.WithError(err).Error("could not get device number")
		} else {
			d.Device = dev
		}
	}

	if d.Bus <= 0 {
		if sysfs, ok := d.dataSource.(backingSysfs); ok {
			d.Bus, err = sysfs.getBusNum(*d)
			if err != nil {
				log.WithError(err).Error("problem getting bus number")
			}
		}
	}

	d.Vendor.nameFromDevice, err = d.dataSource.getVendorName(*d)
	if err != nil {
		log.WithError(err).Error("problem fetching manufacturer name")
	}
	d.Product.nameFromDevice, err = d.dataSource.getProductName(*d)
	if err != nil {
		log.WithError(err).Error("problem fetching product name")
	}
	d.Port, err = d.dataSource.getPort(*d)
	if err != nil {
		log.WithError(err).Error("problem fetching device port number")
	}
	cfg, err := d.dataSource.getActiveConfig(*d)
	if err != nil {
		log.WithError(err).Error("problem fetching active config")
		cfg = 1 // assume it's the first one ?
	}
	d.ActiveConfig = &d.Configs[cfg-1]
	d.Speed, err = d.dataSource.getSpeed(*d)
	if err != nil {
		log.WithError(err).Error("problem fetching device speed")
		d.Speed = SpeedUnknown
	}

	// things we can only get if we are using sysfs
	if sysfs, ok := d.dataSource.(backingSysfs); ok {
		d.Parent, err = sysfs.getParent(*d)
		if err != nil {
			log.WithError(err).Error("problem determining device parent")
		}
	} else {
		log.Info("sysfs not available, not able to determine device hub parents")
	}
	d.Ports = getPorts(*d)

	return d
}

func toConfig(c gusb.ConfigDescriptor, d *Device) Configuration {
	cfg := Configuration{
		SelfPowered:  c.SelfPowered,
		RemoteWakeup: c.RemoteWakeup,
		MaxPower:     int(c.MaxPower * 2),
		Value:        int(c.Value),
		Interfaces:   make([]Interface, c.NumInterfaces),
		d:            d,
	}
	for _, intf := range c.Interfaces {
		cfg.Interfaces[intf.InterfaceNumber] = toInterface(intf, d)
	}

	return cfg
}

func toInterface(i gusb.InterfaceDescriptor, d *Device) Interface {
	intf := Interface{
		ID:        int(i.InterfaceNumber),
		Alternate: 0, //@todo?
		Endpoints: make([]Endpoint, i.NumEndpoints),
		d:         d,
	}

	for idx, ep := range i.Endpoints {
		intf.Endpoints[idx] = toEndpoint(ep, &intf)
	}

	return intf
}

func toEndpoint(e gusb.EndpointDescriptor, i *Interface) Endpoint {
	ep := Endpoint{
		Address:          int(e.Address),
		TransferType:     int(e.TransferType),
		MaxPacketSize:    int(e.MaxPacketSize),
		MaxISOPacketSize: int(e.MaxPacketSize), //@todo: what
		i:                i,
	}

	return ep
}

/* ---------------- helpers -------------------------- */

func readAsInt(fname string) (int, error) {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(string(data[:len(data)-1]))
}
func readAsFloat(fname string) (float64, error) {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return -1, err
	}
	return strconv.ParseFloat(string(data[:len(data)-1]), 64)
}

func getSysfsFromBusDev(bus int, dev int) string {
	syspath := ""
	filepath.Walk("/sys/bus/usb/devices/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() { // entries are symlinks, and don't catch on this. This is just for the parent dir itself
			return nil
		}
		dv, err := readAsInt(filepath.Join(path, "devnum"))
		if err != nil {
			return nil
		}
		bs, err := readAsInt(filepath.Join(path, "busnum"))
		if err != nil {
			return nil
		}
		if bus == bs && dev == dv {
			syspath = path
			return errors.New("done")
		}
		return nil
	})
	return syspath
}

func getPorts(d Device) []int {
	const MAX_PORTS = 7 // according to USB 3.0 spec, max depth limit
	ports := make([]int, 0, MAX_PORTS)
	for dev := &d; dev != nil; dev = dev.Parent {
		if dev.Port != 0 {
			ports = append(ports, dev.Port)
		}
	}
	//reverse
	for i := len(ports)/2 - 1; i >= 0; i-- {
		swap := len(ports) - 1 - i
		ports[i], ports[swap] = ports[swap], ports[i]
	}
	return ports
}
