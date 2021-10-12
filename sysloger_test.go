// Copyright 2020 NGR Softlab
//
package sysloger

import (
	"testing"
)

const CertPath = ""

/////////////////////////////////////////////////
func TestSendSingleSyslogMsg(t *testing.T) {
	//t.Parallel()

	//GlobalCaCert, err := ioutil.ReadFile(CertPath)
	//if err != nil {
	//	t.Fatal("Bad TestMakeSyslogWithFormatter1: ", err)
	//}
	//
	//CaCertPool := x509.NewCertPool()
	//CaCertPool.AppendCertsFromPEM(GlobalCaCert)
	//
	//tl := tls.Config{
	//	RootCAs:            CaCertPool,
	//	InsecureSkipVerify: true,
	//}

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

	err = SendSingleSyslogMsg(SyslogParams{
		Level:    5,
		Host:     "127.0.0.1",
		Port:     555,
		Protocol: "tcp",
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

/////////////////////////////////////////////////
func TestSendListToSyslog(t *testing.T) {
	//t.Parallel()

	//GlobalCaCert, err := ioutil.ReadFile(CertPath)
	//if err != nil {
	//	t.Fatal("Bad TestMakeSyslogWithFormatter1: ", err)
	//}
	//
	//CaCertPool := x509.NewCertPool()
	//CaCertPool.AppendCertsFromPEM(GlobalCaCert)
	//
	//tl := tls.Config{
	//	RootCAs:            CaCertPool,
	//	InsecureSkipVerify: true,
	//}

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

	err := SendListToSyslog(SyslogParams{
		Level:    3,
		Host:     "127.0.0.1",
		Port:     555,
		Protocol: "tcp",
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
