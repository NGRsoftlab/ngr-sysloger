// Copyright 2020-2024 NGR Softlab
package sysloger

import (
	"time"

	syslog "github.com/RackSec/srslog"
)

/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////

// SendListToSyslog - Send list of msgs to syslog
func SendListToSyslog(params SyslogParams, formatter syslog.Formatter, msgList []string) error {
	writer, err := NewSyslogWriter(params, formatter)
	if err != nil {
		logger.Errorf("ERROR failed to create writer: %s", err.Error())
		return err
	}

	defer func() { _ = writer.Close() }()

	for i, msg := range msgList {
		_, errW := doSuitableMethod(writer, params.Level, msg)
		if errW != nil {
			logger.Warningf("ERROR writing to syslog: %d %v", i, errW)
		}
	}

	return nil
}

// SendSingleSyslogMsg - Send single syslog msg
func SendSingleSyslogMsg(params SyslogParams, formatter syslog.Formatter, msg string) error {
	writer, err := NewSyslogWriter(params, formatter)
	if err != nil {
		logger.Errorf("ERROR failed to create writer: %s", err.Error())
		return err
	}

	defer func() { _ = writer.Close() }()

	_, errW := doSuitableMethod(writer, params.Level, msg)
	if errW != nil {
		logger.Errorf("ERROR writing to syslog: %s", errW.Error())
		return errW
	}

	return nil
}

// SendListToSyslogWithTimeout - Send list of msgs to syslog with timeout
func SendListToSyslogWithTimeout(params SyslogParams, formatter syslog.Formatter, msgList []string, timeout time.Duration) error {
	writer, err := NewSyslogWriterWithTimeout(params, formatter, timeout)
	if err != nil {
		logger.Errorf("ERROR failed to create writer: %s", err.Error())
		return err
	}

	defer func() { _ = writer.Close() }()

	for i, msg := range msgList {
		_, errW := doSuitableMethod(writer, params.Level, msg)
		if errW != nil {
			logger.Warningf("ERROR writing to syslog: %d %v", i, errW)
		}
	}

	return nil
}

// SendSingleSyslogMsgWithTimeout - Send single syslog msg with timeout
func SendSingleSyslogMsgWithTimeout(params SyslogParams, formatter syslog.Formatter, msg string, timeout time.Duration) error {
	writer, err := NewSyslogWriterWithTimeout(params, formatter, timeout)
	if err != nil {
		logger.Errorf("ERROR failed to create writer: %s", err.Error())
		return err
	}

	defer func() { _ = writer.Close() }()

	_, errW := doSuitableMethod(writer, params.Level, msg)
	if errW != nil {
		logger.Errorf("ERROR writing to syslog: %s", errW.Error())
		return errW
	}

	return nil
}
