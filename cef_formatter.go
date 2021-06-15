// Copyright 2020 NGR Softlab
//
package sysloger

import (
	"fmt"
	"time"

	syslog "github.com/RackSec/srslog"
	"github.com/pcktdmp/cef/cefevent"
	log "github.com/sirupsen/logrus"
)

// CEF header parameters
type CefHeader struct {
	Version                                    int
	DeviceVendor, DeviceProduct, DeviceVersion string
	DeviceEventClassId                         string
	Name                                       string
	Severity                                   string
}

// Custom Formatter for CEF (for github.com/RackSec/srslog lib)
func CEFFormatter(p syslog.Priority, hostname, tag, content string) string {
	timestamp := time.Now().Format(time.Stamp)
	msg := fmt.Sprintf("%s %s %s", timestamp, hostname, content)
	return msg
}

// Making CEF string from custom header params and content map
func MakeCefString(header CefHeader, contentMap map[string]interface{}, keysAreLong, useDefault, useCustom bool) (string, error) {
	stringMap := make(map[string]string)

	for key, value := range contentMap {
		// get real name from cef maps
		if !useDefault {
			if keysAreLong {
				stringMap[GetShortNameByLong(key)] = fmt.Sprintf("%v", value)
			} else {
				stringMap[GetLongNameByShort(key)] = fmt.Sprintf("%v", value)
			}
		} else {
			// if we need some additional fields
			if useCustom {
				stringMap[key] = fmt.Sprintf("%v", value)
			}
		}
	}

	event := cefevent.CefEvent{
		Version:            header.Version,
		DeviceVendor:       header.DeviceVendor,
		DeviceProduct:      header.DeviceProduct,
		DeviceVersion:      header.DeviceVersion,
		DeviceEventClassId: header.DeviceEventClassId,
		Name:               header.Name,
		Severity:           header.Severity,
		Extensions:         stringMap,
	}

	cef, err := event.Generate()
	if err != nil {
		log.Warning("bad cef generation: ", err)
		return "", err
	}

	return cef, nil
}
