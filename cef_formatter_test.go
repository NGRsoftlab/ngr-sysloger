// Copyright 2020-2024 NGR Softlab
package sysloger

import (
	"fmt"
	"testing"

	syslog "github.com/RackSec/srslog"
)

// ///////////////////////////////////////////////
func TestFormatter(t *testing.T) {
	f := make(map[string]interface{})

	f["src"] = "127.0.0.1"
	f["num"] = 6
	f["deviceAction"] = "uuuuu"
	f["requestClientApplication"] = "Go-http-client/1.1"

	stringEvent, err := MakeCefString(CefHeader{
		Version:            0,
		DeviceVendor:       "Cool Vendor",
		DeviceProduct:      "Cool Product",
		DeviceVersion:      "1.0",
		DeviceEventClassId: "FLAKY_EVENT",
		Name:               "Something flaky happened.",
		Severity:           "Medium",
	}, f, true, false, true)
	if err != nil {
		t.Fatal("Need to handle this.")
	}

	fmt.Println(CEFFormatter(syslog.LOG_INFO, "hostname", "", stringEvent))
}
