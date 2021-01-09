package mylib

func init() {
	PrnLog = GetPrnLogInstance()
	prnIn := &PrnLogStu{
		Sysflag:         "p2p",
		LogLevel:        LOG_DEBUG,
		DbLogLevel:      LOG_DEBUG,
		UseFulFmt:       true,
		SaveToFileOnly:  false,
		LogPATHFile:     "",
		LimitLenPerLine: 0,
	}
	PrnLog.Init(prnIn)

}

var PrnLog *PrnLogStu
