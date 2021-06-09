# ngr-sysloger
Syslog sending wrapper (with custom CEF formatter)

# import
```import "github.com/NGRsoftlab/ngr-sysloger"```

# example (see _test files)
```
	
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
	testData, err := MakeCefString(header, testMap)
	if err != nil {
		log.Fatal(err)
	}

	err = SendSingleSyslogMsg(SyslogParams{
		Level:    5,
		Host:     "127.0.0.1",
		Port:     "555",
		Protocol: "tcp",
		Priority: 0,
		Tag:      "test",
		//NeedTls:  false,
		//TlsConf:  &tl,
	},
		CEFFormatter,
		testData)
	if err != nil {
		log.Fatal("Bad TestSendSingleSyslogMsg: ", err)
	}
```
