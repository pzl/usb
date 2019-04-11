package usb

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/apex/log"
	"github.com/pzl/usb/gusb"
)

// @todo: usbfs fallback?
type backingSysfs struct{}

func (b backingSysfs) getDevNum(d Device) (int, error) {
	return readAsInt(filepath.Join(d.sysPath, "devnum"))
}
func (b backingSysfs) getVendorName(d Device) (string, error) {
	data, err := ioutil.ReadFile(filepath.Join(d.sysPath, "manufacturer"))
	return strings.TrimSpace(string(data)), err
}
func (b backingSysfs) getProductName(d Device) (string, error) {
	data, err := ioutil.ReadFile(filepath.Join(d.sysPath, "product"))
	return strings.TrimSpace(string(data)), err
}
func (b backingSysfs) getPort(d Device) (int, error) {
	if has := strings.LastIndexAny(d.sysPath, ".-"); has != -1 {
		// look for last port, separated by .  or top-level hub port after the -
		if port, err := strconv.Atoi(d.sysPath[has+1:]); err == nil {
			return port, nil
		} else {
			return 0, err
		}
	}

	if strings.HasPrefix(filepath.Base(d.sysPath), "usb") {
		return 0, nil // is a hub. is not on a port, usually
	}
	return 0, fmt.Errorf("unable to find port number in path: %s", d.sysPath)
}
func (b backingSysfs) getActiveConfig(d Device) (int, error) {
	cfg, err := readAsInt(filepath.Join(d.sysPath, "bConfigurationValue"))
	return cfg, err
}
func (b backingSysfs) getSpeed(d Device) (Speed, error) {
	speed, err := readAsFloat(filepath.Join(d.sysPath, "speed"))
	return toSpeedSysfs(speed), err
}

func (b backingSysfs) getDriver(d Device, intf int) (string, error) {
	driver := filepath.Join(fmt.Sprintf("%s:%d.%d", d.sysPath, d.ActiveConfig.Value, intf), "driver")
	if drv, err := os.Readlink(driver); err == nil {
		return filepath.Base(drv), nil
	} else {
		log.WithField("path", driver).WithError(err).Error("could not use sysfs to get driver")
		return "", err
	}
}

func (b backingSysfs) setConfiguration(d Device, cfg int) error {
	//	write to sysfs_path/bConfigurationValue
	return ErrNotImplemented
}

/* Not universal funcs */

func (b backingSysfs) getBusNum(d Device) (int, error) {
	return readAsInt(filepath.Join(d.sysPath, "busnum"))
}

func (b backingSysfs) getParent(d Device) (*Device, error) {
	if has := strings.LastIndexAny(d.sysPath, ".-"); has != -1 {
		parent := d.sysPath[:has]
		if !strings.ContainsRune(parent, '-') {
			parent = filepath.Join(filepath.Dir(parent), fmt.Sprintf("usb%s", filepath.Base(parent)))
		}

		if f, err := os.OpenFile(filepath.Join(parent, "descriptors"), os.O_RDONLY, 0644); err == nil {
			defer f.Close()
			if pdesc, err := gusb.ParseDescriptor(f); err == nil {
				pdesc.PathInfo.SysPath = parent
				return toDevice(pdesc), nil
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return nil, nil
}

/*  helpers  */

// in Mbps apparently
func toSpeedSysfs(speed float64) Speed {
	s := int(speed)
	switch s {
	case 1: // truncated from 1.5, which is low speed
		return SpeedLow
	case 12:
		return SpeedFull
	case 480:
		return SpeedHigh
	case 5000:
		return SpeedSuper
	case 10000:
		return SpeedSuperPlus
	}
	return SpeedUnknown
}
