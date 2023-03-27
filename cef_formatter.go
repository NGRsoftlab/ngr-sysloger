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

// CefHeader CEF header parameters
type CefHeader struct {
	Version                                    int
	DeviceVendor, DeviceProduct, DeviceVersion string
	DeviceEventClassId                         string
	Name                                       string
	Severity                                   string
}

// CEFFormatter custom Formatter for CEF (for github.com/RackSec/srslog lib)
func CEFFormatter(p syslog.Priority, hostname, tag, content string) string {
	timestamp := time.Now().Format(time.Stamp)
	msg := fmt.Sprintf("%s %s %s", timestamp, hostname, content)
	return msg
}

// MakeCefString making CEF string from custom header params and content map
func MakeCefString(header CefHeader, contentMap map[string]interface{}, keysAreLong, useDefault, useCustom bool) (string, error) {
	stringMap := make(map[string]string)

	for key, value := range contentMap {
		valueToSend := ""

		switch value.(type) {
		case float32:
			if value == float32(int32(value.(float32))) {
				valueToSend = fmt.Sprintf("%d", int32(value.(float32)))
			} else {
				valueToSend = fmt.Sprintf("%f", value)
			}
		case float64:
			if value == float64(int64(value.(float64))) {
				valueToSend = fmt.Sprintf("%d", int64(value.(float64)))
			} else {
				valueToSend = fmt.Sprintf("%f", value)
			}
		default:
			valueToSend = fmt.Sprintf("%v", value)
		}

		// get real name from cef maps
		if !useDefault {
			okKey := CheckKey(key)

			if !okKey {
				// if we need some additional fields
				if useCustom {
					stringMap[key] = valueToSend
				}
			} else {
				if keysAreLong {
					stringMap[GetShortNameByLong(key)] = valueToSend
				} else {
					stringMap[GetLongNameByShort(key)] = valueToSend
				}
			}
		} else {
			stringMap[key] = valueToSend
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
