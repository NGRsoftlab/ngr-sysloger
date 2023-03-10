// Copyright 2020-2022 NGR Softlab
//
package sysloger

import (
	"fmt"
	"time"

	syslog "github.com/RackSec/srslog"
	log "github.com/sirupsen/logrus"
)

/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////

// SendListToSyslog - Send list of msgs to syslog
func SendListToSyslog(params SyslogParams, formatter syslog.Formatter, msgList []string) error {
	writer, err := NewSyslogWriter(params, formatter)
	if err != nil {
		log.Error("ERROR failed to create writer:", err)
		return err
	}

	defer func() { _ = writer.Close() }()

	log.Info(fmt.Sprintf("need write: %d msgs to syslog", len(msgList)))

	for i, msg := range msgList {
		_, errW := doSuitableMethod(writer, params.Level, msg)
		if errW != nil {
			log.Warning(fmt.Sprintf("ERROR writing to syslog: %d %v", i, errW))
		}
	}

	return nil
}

// SendSingleSyslogMsg - Send single syslog msg
func SendSingleSyslogMsg(params SyslogParams, formatter syslog.Formatter, msg string) error {
	writer, err := NewSyslogWriter(params, formatter)
	if err != nil {
		log.Error("ERROR failed to create writer:", err)
		return err
	}

	defer func() { _ = writer.Close() }()

	log.Info("need write: 1 msg to syslog")

	_, errW := doSuitableMethod(writer, params.Level, msg)
	if errW != nil {
		log.Error("ERROR writing to syslog: ", errW)
		return errW
	}

	return nil
}

// SendListToSyslogWithTimeout - Send list of msgs to syslog with timeout
func SendListToSyslogWithTimeout(params SyslogParams, formatter syslog.Formatter, msgList []string, timeout time.Duration) error {
	writer, err := NewSyslogWriterWithTimeout(params, formatter, timeout)
	if err != nil {
		log.Error("ERROR failed to create writer:", err)
		return err
	}

	defer func() { _ = writer.Close() }()

	log.Info(fmt.Sprintf("need write: %d msgs to syslog", len(msgList)))

	for i, msg := range msgList {
		_, errW := doSuitableMethod(writer, params.Level, msg)
		if errW != nil {
			log.Warning(fmt.Sprintf("ERROR writing to syslog: %d %v", i, errW))
		}
	}

	return nil
}

// SendSingleSyslogMsgWithTimeout - Send single syslog msg with timeout
func SendSingleSyslogMsgWithTimeout(params SyslogParams, formatter syslog.Formatter, msg string, timeout time.Duration) error {
	writer, err := NewSyslogWriterWithTimeout(params, formatter, timeout)
	if err != nil {
		log.Error("ERROR failed to create writer:", err)
		return err
	}

	defer func() { _ = writer.Close() }()

	log.Info("need write: 1 msg to syslog")

	_, errW := doSuitableMethod(writer, params.Level, msg)
	if errW != nil {
		log.Error("ERROR writing to syslog: ", errW)
		return errW
	}

	return nil
}
