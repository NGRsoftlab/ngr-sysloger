// Copyright 2020 NGR Softlab
//
package sysloger

import (
	"crypto/tls"
	"errors"
	"fmt"

	syslog "github.com/RackSec/srslog"
	log "github.com/sirupsen/logrus"
)

// SyslogParams syslog dial params
type SyslogParams struct {
	Level    int    `json:"level"`    // syslog level info/error/fatal
	Host     string `json:"host"`     // host to send
	Port     int    `json:"port"`     // port to send
	Protocol string `json:"protocol"` // tcp\udp

	Priority syslog.Priority
	Tag      string `json:"tag"` // syslog tag

	NeedTls bool
	TlsConf *tls.Config
}

// NewSyslogWriter create new syslog writer with params and custom formatter
func NewSyslogWriter(params SyslogParams, formatter syslog.Formatter) (*syslog.Writer, error) {
	var sysLogger *syslog.Writer
	var err error

	if params.NeedTls {
		if params.TlsConf == nil {
			log.Error("ERROR nil TlsConf")
			return nil, errors.New("nil TlsConf")
		}

		sysLogger, err = syslog.DialWithTLSConfig(params.Protocol,
			fmt.Sprintf("%s:%d", params.Host, params.Port), params.Priority, params.Tag, params.TlsConf)
	} else {
		sysLogger, err = syslog.Dial(params.Protocol,
			fmt.Sprintf("%s:%d", params.Host, params.Port), params.Priority, params.Tag)
	}

	if err != nil {
		log.Error("ERROR failed to dial syslog:", err)
		return nil, err
	}

	sysLogger.SetFormatter(formatter)

	return sysLogger, nil
}
