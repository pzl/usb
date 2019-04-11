package usb

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/apex/log"
)

var vmap vMap

type vMap map[uint16]_vendor
type _vendor struct {
	name     string
	products map[uint16]string
}

func createIDMap() vMap {
	m := vMap{}
	idbuf := make([]byte, 2)
	lastVID := uint16(0)

	usbids := usbIDs()
	defer usbids.Close()

	scanner := bufio.NewScanner(usbids)
SCANNER:
	for scanner.Scan() {
		l := scanner.Bytes()

		if len(l) < 1 {
			continue
		}
		if l[0] == '#' {
			continue
		}

		switch {
		case bytes.Equal(l[:2], []byte("C ")): //@todo
			//@todo
			break SCANNER
		case bytes.Equal(l[:2], []byte("R ")): //@todo
		case bytes.Equal(l[:2], []byte("L ")): //@todo
		case bytes.Equal(l[:3], []byte("AT ")): //@todo
		case bytes.Equal(l[:3], []byte("VT ")): //@todo
		case bytes.Equal(l[:4], []byte("HID ")): //@todo
		case bytes.Equal(l[:4], []byte("HCC ")): //@todo
		case bytes.Equal(l[:4], []byte("PHY ")): //@todo
		case bytes.Equal(l[:4], []byte("HUT ")): //@todo
		case bytes.Equal(l[:5], []byte("BIAS ")): //@todo
		case l[0] == '\t' && l[1] != '\t':
			if lastVID == 0 {
				continue
			}
			if _, err := hex.Decode(idbuf, l[1:5]); err != nil {
				continue
			}
			m[lastVID].products[binary.BigEndian.Uint16(idbuf)] = strings.TrimSpace(string(l[5:]))
			//single sub-item
		case l[0] == '\t' && l[1] == '\t':
			// sub-sub-item
		default:
			//vendor ID
			if _, err := hex.Decode(idbuf, l[:4]); err != nil {
				log.WithField("line", string(l)).Info("failed parsing line in usb.ids")
				continue
			}
			vid := binary.BigEndian.Uint16(idbuf)
			m[vid] = _vendor{
				name:     strings.TrimSpace(string(l[4:])),
				products: make(map[uint16]string),
			}
			lastVID = vid
		}

	}

	return m
}

func vendorName(id uint16) string {
	if vmap == nil {
		vmap = createIDMap()
	}
	if v, exists := vmap[id]; exists {
		return v.name
	}
	return ""
}

func productName(vid uint16, pid uint16) string {
	if vmap == nil {
		vmap = createIDMap()
	}
	if v, exists := vmap[vid]; exists {
		if p, ex := v.products[pid]; ex {
			return p
		}
	}
	return ""
}

func usbIDs() (r io.ReadCloser) {

	idPaths := []string{"/usr/share/hwdata/usb.ids", "/usr/share/usb.ids", "/usr/share/libosinfo/usb.ids", "/usr/share/kcmusb/usb.ids", "/var/lib/usbutils/usb.ids"}

	for i := range idPaths {
		log.WithField("path", idPaths[i]).Debug("checking for usb.ids at")
		if f, err := os.OpenFile(idPaths[i], os.O_RDONLY, 0644); err != nil {
			continue
		} else {
			log.Debug("success")
			r = f
			break
		}
	}
	if r == nil {
		// must not have found any usable files in the above paths
		// use the version we shipped with
		log.Debug("using built-in usb.ids")
		r = ioutil.NopCloser(shippedUsbIds()) // turns the reader into a ReadCloser where Close is a no-op
	}

	return r
}
