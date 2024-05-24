// Copyright 2020-2024 NGR Softlab
package sysloger

import (
	"testing"
	"time"
)

func TestSendSingleSyslogMsg(t *testing.T) {
	header := CefHeader{
		Version:            0,
		DeviceVendor:       "Test",
		DeviceProduct:      "TestProd",
		DeviceVersion:      "1.0",
		DeviceEventClassId: "testing",
		Name:               "TEST",
		Severity:           "Low",
	}

	testMap := map[string]interface{}{
		"src":                      "HOOOST",
		"requestClientApplication": "Test-cli",
	}

	// for CEFFormatter
	testData, err := MakeCefString(header, testMap, false, true, false)
	if err != nil {
		t.Fatal(err)
	}

	// upd test (for no response control)
	err = SendSingleSyslogMsg(SyslogParams{
		Level:    5,
		Host:     "127.0.0.1",
		Port:     555,
		Protocol: "udp",
		Priority: 0,
		Tag:      "test",
		//NeedTls:  false,
		//TlsConf:  &tl,
	},
		CEFFormatter,
		testData)
	if err != nil {
		t.Fatal("Bad TestSendSingleSyslogMsg: ", err)
	}
}

func TestSendSingleSyslogMsgWithTimeout(t *testing.T) {
	header := CefHeader{
		Version:            0,
		DeviceVendor:       "Test",
		DeviceProduct:      "TestProd",
		DeviceVersion:      "1.0",
		DeviceEventClassId: "testing",
		Name:               "TEST",
		Severity:           "Low",
	}

	testMap := map[string]interface{}{
		"src":                      "HOOOST",
		"requestClientApplication": "Test-cli",
	}

	// for CEFFormatter
	testData, err := MakeCefString(header, testMap, false, true, false)
	if err != nil {
		t.Fatal(err)
	}

	// upd test (for no response control)
	err = SendSingleSyslogMsgWithTimeout(SyslogParams{
		Level:    5,
		Host:     "127.0.0.1",
		Port:     555,
		Protocol: "udp",
		Priority: 0,
		Tag:      "test",
		//NeedTls:  false,
		//TlsConf:  &tl,
	},
		CEFFormatter,
		testData,
		2*time.Second)
	if err != nil {
		t.Fatal("Bad TestSendSingleSyslogMsg: ", err)
	}
}

func TestSendListToSyslog(t *testing.T) {
	header := CefHeader{
		Version:            0,
		DeviceVendor:       "Test",
		DeviceProduct:      "TestProd",
		DeviceVersion:      "1.0",
		DeviceEventClassId: "testing",
		Name:               "TEST",
		Severity:           "Low",
	}

	testMap := []map[string]interface{}{
		{
			"src":                      "HOOOST",
			"requestClientApplication": "Test-cli",
		},
		{
			"src":                      "BOOOST",
			"requestClientApplication": "Best-cli",
		},
	}

	testData := make([]string, 0)
	for _, el := range testMap {
		// for CEFFormatter
		testDataStr, err := MakeCefString(header, el, false, true, false)
		if err != nil {
			t.Fatal(err)
		}

		testData = append(testData, testDataStr)
	}

	// upd test (for no response control)
	err := SendListToSyslog(SyslogParams{
		Level:    3,
		Host:     "127.0.0.1",
		Port:     555,
		Protocol: "udp",
		Priority: 0,
		Tag:      "test",
		//NeedTls:  false,
		//TlsConf:  &tl,
	},
		CEFFormatter,
		testData)
	if err != nil {
		t.Fatal("Bad TestSendListToSyslog: ", err)
	}
}

func TestSendListToSyslogWithTimeout(t *testing.T) {
	header := CefHeader{
		Version:            0,
		DeviceVendor:       "Test",
		DeviceProduct:      "TestProd",
		DeviceVersion:      "1.0",
		DeviceEventClassId: "testing",
		Name:               "TEST",
		Severity:           "Low",
	}

	testMap := []map[string]interface{}{
		{
			"src":                      "HOOOST",
			"requestClientApplication": "Test-cli",
		},
		{
			"src":                      "BOOOST",
			"requestClientApplication": "Best-cli",
		},
	}

	testData := make([]string, 0)
	for _, el := range testMap {
		// for CEFFormatter
		testDataStr, err := MakeCefString(header, el, false, true, false)
		if err != nil {
			t.Fatal(err)
		}

		testData = append(testData, testDataStr)
	}

	// upd test (for no response control)
	err := SendListToSyslogWithTimeout(SyslogParams{
		Level:    3,
		Host:     "127.0.0.1",
		Port:     555,
		Protocol: "udp",
		Priority: 0,
		Tag:      "test",
		//NeedTls:  false,
		//TlsConf:  &tl,
	},
		CEFFormatter,
		testData,
		1*time.Minute)
	if err != nil {
		t.Fatal("Bad TestSendListToSyslog: ", err)
	}
}
