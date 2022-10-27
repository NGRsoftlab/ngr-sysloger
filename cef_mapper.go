package sysloger

type FieldInfo struct {
	FullName    string `json:"full_name"`
	ShortName   string `json:"short_name"`
	DataType    string `json:"data_type"`
	Length      int64  `json:"length"`
	Description string `json:"description"`
}

func CheckKey(testKey string) bool {
	_, ok1 := longNamesDictionary[testKey]
	_, ok2 := shortNamesDictionary[testKey]

	return ok1 || ok2
}

func GetShortNameByLong(shortName string) string {
	res, ok := longNamesDictionary[shortName]
	if ok {
		return res.ShortName
	} else {
		return shortName
	}
}

func GetLongNameByShort(longName string) string {
	res, ok := shortNamesDictionary[longName]
	if ok {
		return res.FullName
	} else {
		return longName
	}
}

var shortNamesDictionary = map[string]FieldInfo{
	// common
	"start":       {FullName: "startTime", DataType: "DateTime"},
	"description": {FullName: "description", DataType: "Nullable(String)"},
	"outcome":     {FullName: "outcome", DataType: "Nullable(String)"},
	"app":         {FullName: "applicationProtocol", DataType: "Nullable(String)"},
	"end":         {FullName: "endTime", DataType: "Nullable(DateTime)"},
	"rt":          {FullName: "receiptTime", DataType: "Nullable(DateTime)"},
	"cnt":         {FullName: "baseEventCount", DataType: "Nullable(UInt32)"},
	"externalId":  {FullName: "externalId", DataType: "Nullable(String)"},
	"msg":         {FullName: "message", DataType: "Nullable(String)"},
	"proto":       {FullName: "transportProtocol", DataType: "Nullable(String)"},
	"reason":      {FullName: "reason", DataType: "Nullable(String)"},
	"in":          {FullName: "bytesIn", DataType: "Nullable(UInt32)"},
	"out":         {FullName: "bytesOut", DataType: "Nullable(UInt32)"},
	"rawEvent":    {FullName: "rawEvent", DataType: "Nullable(String)"},
	"eventId":     {FullName: "eventId", DataType: "Nullable(UInt32)"},
	"level":       {FullName: "level", DataType: "Nullable(String)"},

	// source
	"shost":                   {FullName: "sourceHostName", DataType: "Nullable(String)"},
	"smac":                    {FullName: "sourceMacAddress", DataType: "Nullable(String)"},
	"sntdom":                  {FullName: "sourceNtDomain", DataType: "Nullable(String)"},
	"sourceDnsDomain":         {FullName: "sourceDnsDomain", DataType: "Nullable(String)"},
	"sourceServiceName":       {FullName: "sourceServiceName", DataType: "Nullable(String)"},
	"sourceTranslatedAddress": {FullName: "sourceTranslatedAddress", DataType: "Nullable(String)"},
	"sourceTranslatedPort":    {FullName: "sourceTranslatedPort", DataType: "Nullable(UInt32)"},
	"spid":                    {FullName: "sourceProcessId", DataType: "Nullable(UInt32)"},
	"spriv":                   {FullName: "sourceUserPrivileges", DataType: "Nullable(String)"},
	"spt":                     {FullName: "sourcePort", DataType: "Nullable(UInt32)"},
	"src":                     {FullName: "sourceAddress", DataType: "Nullable(String)"},
	"suid":                    {FullName: "sourceUserId", DataType: "Nullable(String)"},
	"suser":                   {FullName: "sourceUserName", DataType: "Nullable(String)"},

	// destination
	"dhost":                        {FullName: "destinationHostName", DataType: "Nullable(String)"},
	"dmac":                         {FullName: "destinationMac", DataType: "Nullable(String)"},
	"dntdom":                       {FullName: "destinationNtDomain", DataType: "Nullable(String)"},
	"dpid":                         {FullName: "destinationProcessId", DataType: "Nullable(UInt32)"},
	"dpriv":                        {FullName: "destinationUserPrivileges", DataType: "Nullable(String)"},
	"dproc":                        {FullName: "destinationProcessName", DataType: "Nullable(String)"},
	"dpt":                          {FullName: "destinationPort", DataType: "Nullable(UInt32)"},
	"dst":                          {FullName: "destinationAddress", DataType: "Nullable(String)"},
	"duid":                         {FullName: "destinationUserId", DataType: "Nullable(String)"},
	"duser":                        {FullName: "destinationUserName", DataType: "Nullable(String)"},
	"destinationDnsDomain":         {FullName: "destinationDnsDomain", DataType: "Nullable(String)"},
	"destinationServiceName":       {FullName: "destinationServiceName", DataType: "Nullable(String)"},
	"destinationTranslatedAddress": {FullName: "destinationTranslatedAddress", DataType: "Nullable(String)"},
	"destinationTranslatedPort":    {FullName: "destinationTranslatedPort", DataType: "Nullable(UInt32)"},

	// device
	"act":                     {FullName: "deviceAction", DataType: "Nullable(String)"},
	"cat":                     {FullName: "deviceCategory", DataType: "Nullable(String)"},
	"dvc":                     {FullName: "deviceAddress", DataType: "Nullable(String)"},
	"dtz":                     {FullName: "deviceTimeZone", DataType: "Nullable(String)"},
	"dvchost":                 {FullName: "deviceHostName", DataType: "Nullable(String)"},
	"dvcpid":                  {FullName: "deviceProcessId", DataType: "Nullable(UInt32)"},
	"devicePayload":           {FullName: "devicePayload", DataType: "Nullable(String)"},
	"deviceDirection":         {FullName: "deviceDirection", DataType: "Nullable(String)"},
	"deviceDnsDomain":         {FullName: "deviceDnsDomain", DataType: "Nullable(String)"},
	"deviceExternalId":        {FullName: "deviceExternalId", DataType: "Nullable(String)"},
	"deviceFacility":          {FullName: "deviceFacility", DataType: "Nullable(String)"},
	"deviceInboundInterface":  {FullName: "deviceInboundInterface", DataType: "Nullable(String)"},
	"deviceMacAddress":        {FullName: "deviceMacAddress", DataType: "Nullable(String)"},
	"deviceNtDomain":          {FullName: "deviceNtDomain", DataType: "Nullable(String)"},
	"deviceOutboundInterface": {FullName: "deviceOutboundInterface", DataType: "Nullable(String)"},
	"deviceProcessName":       {FullName: "deviceProcessName", DataType: "Nullable(String)"},
	"deviceTranslatedAddress": {FullName: "deviceTranslatedAddress", DataType: "Nullable(String)"},

	// request
	"request":                  {FullName: "request", DataType: "Nullable(String)"},
	"requestClientApplication": {FullName: "requestClientApplication", DataType: "Nullable(String)"},
	"requestCookies":           {FullName: "requestCookies", DataType: "Nullable(String)"},
	"requestMethod":            {FullName: "requestMethod", DataType: "Nullable(String)"},
	"requestContext":           {FullName: "requestContext", DataType: "Nullable(String)"},

	// file
	"fileCreateTime":       {FullName: "fileCreateTime", DataType: "Nullable(DateTime)"},
	"fileModificationTime": {FullName: "fileModificationTime", DataType: "Nullable(DateTime)"},
	"fileHash":             {FullName: "fileHash", DataType: "Nullable(String)"},
	"fileId":               {FullName: "fileId", DataType: "Nullable(String)"},
	"filePath":             {FullName: "filePath", DataType: "Nullable(String)"},
	"filePermission":       {FullName: "filePermission", DataType: "Nullable(String)"},
	"fileType":             {FullName: "fileType", DataType: "Nullable(String)"},
	"fname":                {FullName: "fileName", DataType: "Nullable(String)"},
	"fsize":                {FullName: "fileSize", DataType: "Nullable(UInt32)"},

	// old file
	"oldFileCreateTime":       {FullName: "oldFileCreateTime", DataType: "Nullable(DateTime)"},
	"oldFileModificationTime": {FullName: "oldFileModificationTime", DataType: "Nullable(DateTime)"},
	"oldFileHash":             {FullName: "oldFileHash", DataType: "Nullable(String)"},
	"oldFileId":               {FullName: "oldFileId", DataType: "Nullable(String)"},
	"oldFileName":             {FullName: "oldFileName", DataType: "Nullable(String)"},
	"oldFilePath":             {FullName: "oldFilePath", DataType: "Nullable(String)"},
	"oldFilePermission":       {FullName: "oldFilePermission", DataType: "Nullable(String)"},
	"oldFileSize":             {FullName: "oldFileSize", DataType: "Nullable(UInt32)"},
	"oldFileType":             {FullName: "oldFileType", DataType: "Nullable(String)"},

	// custom ...
}

var longNamesDictionary = map[string]FieldInfo{
	// common
	"startTime":           {FullName: "startTime", ShortName: "start", DataType: "DateTime"},
	"description":         {FullName: "description", ShortName: "description", DataType: "Nullable(String)"},
	"outcome":             {FullName: "outcome", ShortName: "outcome", DataType: "Nullable(String)"},
	"applicationProtocol": {FullName: "applicationProtocol", ShortName: "app", DataType: "Nullable(String)"},
	"endTime":             {FullName: "endTime", ShortName: "end", DataType: "Nullable(DateTime)"},
	"receiptTime":         {FullName: "receiptTime", ShortName: "rt", DataType: "Nullable(DateTime)"},
	"baseEventCount":      {FullName: "baseEventCount", ShortName: "cnt", DataType: "Nullable(UInt32)"},
	"externalId":          {FullName: "externalId", ShortName: "externalId", DataType: "Nullable(String)"},
	"message":             {FullName: "message", ShortName: "msg", DataType: "Nullable(String)"},
	"transportProtocol":   {FullName: "transportProtocol", ShortName: "proto", DataType: "Nullable(String)"},
	"reason":              {FullName: "reason", ShortName: "reason", DataType: "Nullable(String)"},
	"bytesIn":             {FullName: "bytesIn", ShortName: "in", DataType: "Nullable(UInt32)"},
	"bytesOut":            {FullName: "bytesOut", ShortName: "out", DataType: "Nullable(UInt32)"},
	"rawEvent":            {FullName: "rawEvent", ShortName: "rawEvent", DataType: "Nullable(String)"},
	"eventId":             {FullName: "eventId", ShortName: "eventId", DataType: "Nullable(UInt32)"},
	"level":               {FullName: "level", ShortName: "level", DataType: "Nullable(String)"},

	// source
	"sourceHostName":          {FullName: "sourceHostName", ShortName: "shost", DataType: "Nullable(String)"},
	"sourceMacAddress":        {FullName: "sourceMacAddress", ShortName: "smac", DataType: "Nullable(String)"},
	"sourceNtDomain":          {FullName: "sourceNtDomain", ShortName: "sntdom", DataType: "Nullable(String)"},
	"sourceDnsDomain":         {FullName: "sourceDnsDomain", ShortName: "sourceDnsDomain", DataType: "Nullable(String)"},
	"sourceServiceName":       {FullName: "sourceServiceName", ShortName: "sourceServiceName", DataType: "Nullable(String)"},
	"sourceTranslatedAddress": {FullName: "sourceTranslatedAddress", ShortName: "sourceTranslatedAddress", DataType: "Nullable(String)"},
	"sourceTranslatedPort":    {FullName: "sourceTranslatedPort", ShortName: "sourceTranslatedPort", DataType: "Nullable(UInt32)"},
	"sourceProcessId":         {FullName: "sourceProcessId", ShortName: "spid", DataType: "Nullable(UInt32)"},
	"sourceUserPrivileges":    {FullName: "sourceUserPrivileges", ShortName: "spriv", DataType: "Nullable(String)"},
	"sourcePort":              {FullName: "sourcePort", ShortName: "spt", DataType: "Nullable(UInt32)"},
	"sourceAddress":           {FullName: "sourceAddress", ShortName: "src", DataType: "Nullable(String)"},
	"sourceUserId":            {FullName: "sourceUserId", ShortName: "suid", DataType: "Nullable(String)"},
	"sourceUserName":          {FullName: "sourceUserName", ShortName: "suser", DataType: "Nullable(String)"},

	// destination
	"destinationHostName":          {FullName: "destinationHostName", ShortName: "dhost", DataType: "Nullable(String)"},
	"destinationMac":               {FullName: "destinationMac", ShortName: "dmac", DataType: "Nullable(String)"},
	"destinationNtDomain":          {FullName: "destinationNtDomain", ShortName: "dntdom", DataType: "Nullable(String)"},
	"destinationProcessId":         {FullName: "destinationProcessId", ShortName: "dpid", DataType: "Nullable(UInt32)"},
	"destinationUserPrivileges":    {FullName: "destinationUserPrivileges", ShortName: "dpriv", DataType: "Nullable(String)"},
	"destinationProcessName":       {FullName: "destinationProcessName", ShortName: "dproc", DataType: "Nullable(String)"},
	"destinationPort":              {FullName: "destinationPort", ShortName: "dpt", DataType: "Nullable(UInt32)"},
	"destinationAddress":           {FullName: "destinationAddress", ShortName: "dst", DataType: "Nullable(String)"},
	"destinationUserId":            {FullName: "destinationUserId", ShortName: "duid", DataType: "Nullable(String)"},
	"destinationUserName":          {FullName: "destinationUserName", ShortName: "duser", DataType: "Nullable(String)"},
	"destinationDnsDomain":         {FullName: "destinationDnsDomain", ShortName: "destinationDnsDomain", DataType: "Nullable(String)"},
	"destinationServiceName":       {FullName: "destinationServiceName", ShortName: "destinationServiceName", DataType: "Nullable(String)"},
	"destinationTranslatedAddress": {FullName: "destinationTranslatedAddress", ShortName: "destinationTranslatedAddress", DataType: "Nullable(String)"},
	"destinationTranslatedPort":    {FullName: "destinationTranslatedPort", ShortName: "destinationTranslatedPort", DataType: "Nullable(UInt32)"},

	// device
	"deviceAction":            {FullName: "deviceAction", ShortName: "act", DataType: "Nullable(String)"},
	"deviceCategory":          {FullName: "deviceCategory", ShortName: "cat", DataType: "Nullable(String)"},
	"deviceAddress":           {FullName: "deviceAddress", ShortName: "dvc", DataType: "Nullable(String)"},
	"deviceTimeZone":          {FullName: "deviceTimeZone", ShortName: "dtz", DataType: "Nullable(String)"},
	"deviceHostName":          {FullName: "deviceHostName", ShortName: "dvchost", DataType: "Nullable(String)"},
	"deviceProcessId":         {FullName: "deviceProcessId", ShortName: "dvcpid", DataType: "Nullable(UInt32)"},
	"devicePayload":           {FullName: "devicePayload", ShortName: "devicePayload", DataType: "Nullable(String)"},
	"deviceDirection":         {FullName: "deviceDirection", ShortName: "deviceDirection", DataType: "Nullable(String)"},
	"deviceDnsDomain":         {FullName: "deviceDnsDomain", ShortName: "deviceDnsDomain", DataType: "Nullable(String)"},
	"deviceExternalId":        {FullName: "deviceExternalId", ShortName: "deviceExternalId", DataType: "Nullable(String)"},
	"deviceFacility":          {FullName: "deviceFacility", ShortName: "deviceFacility", DataType: "Nullable(String)"},
	"deviceInboundInterface":  {FullName: "deviceInboundInterface", ShortName: "deviceInboundInterface", DataType: "Nullable(String)"},
	"deviceMacAddress":        {FullName: "deviceMacAddress", ShortName: "deviceMacAddress", DataType: "Nullable(String)"},
	"deviceNtDomain":          {FullName: "deviceNtDomain", ShortName: "deviceNtDomain", DataType: "Nullable(String)"},
	"deviceOutboundInterface": {FullName: "deviceOutboundInterface", ShortName: "deviceOutboundInterface", DataType: "Nullable(String)"},
	"deviceProcessName":       {FullName: "deviceProcessName", ShortName: "deviceProcessName", DataType: "Nullable(String)"},
	"deviceTranslatedAddress": {FullName: "deviceTranslatedAddress", ShortName: "deviceTranslatedAddress", DataType: "Nullable(String)"},

	// request
	"request":                  {FullName: "request", ShortName: "request", DataType: "Nullable(String)"},
	"requestClientApplication": {FullName: "requestClientApplication", ShortName: "requestClientApplication", DataType: "Nullable(String)"},
	"requestCookies":           {FullName: "requestCookies", ShortName: "requestCookies", DataType: "Nullable(String)"},
	"requestMethod":            {FullName: "requestMethod", ShortName: "requestMethod", DataType: "Nullable(String)"},
	"requestContext":           {FullName: "requestContext", ShortName: "requestContext", DataType: "Nullable(String)"},

	// file
	"fileCreateTime":       {FullName: "fileCreateTime", ShortName: "fileCreateTime", DataType: "Nullable(DateTime)"},
	"fileModificationTime": {FullName: "fileModificationTime", ShortName: "fileModificationTime", DataType: "Nullable(DateTime)"},
	"fileHash":             {FullName: "fileHash", ShortName: "fileHash", DataType: "Nullable(String)"},
	"fileId":               {FullName: "fileId", ShortName: "fileId", DataType: "Nullable(String)"},
	"filePath":             {FullName: "filePath", ShortName: "filePath", DataType: "Nullable(String)"},
	"filePermission":       {FullName: "filePermission", ShortName: "filePermission", DataType: "Nullable(String)"},
	"fileType":             {FullName: "fileType", ShortName: "fileType", DataType: "Nullable(String)"},
	"fileName":             {FullName: "fileName", ShortName: "fname", DataType: "Nullable(String)"},
	"fileSize":             {FullName: "fileSize", ShortName: "fsize", DataType: "Nullable(UInt32)"},

	// old file
	"oldFileCreateTime":       {FullName: "oldFileCreateTime", ShortName: "oldFileCreateTime", DataType: "Nullable(DateTime)"},
	"oldFileModificationTime": {FullName: "oldFileModificationTime", ShortName: "oldFileModificationTime", DataType: "Nullable(DateTime)"},
	"oldFileHash":             {FullName: "oldFileHash", ShortName: "oldFileHash", DataType: "Nullable(String)"},
	"oldFileId":               {FullName: "oldFileId", ShortName: "oldFileId", DataType: "Nullable(String)"},
	"oldFileName":             {FullName: "oldFileName", ShortName: "oldFileName", DataType: "Nullable(String)"},
	"oldFilePath":             {FullName: "oldFilePath", ShortName: "oldFilePath", DataType: "Nullable(String)"},
	"oldFilePermission":       {FullName: "oldFilePermission", ShortName: "oldFilePermission", DataType: "Nullable(String)"},
	"oldFileSize":             {FullName: "oldFileSize", ShortName: "oldFileSize", DataType: "Nullable(UInt32)"},
	"oldFileType":             {FullName: "oldFileType", ShortName: "oldFileType", DataType: "Nullable(String)"},

	// custom ...
}
