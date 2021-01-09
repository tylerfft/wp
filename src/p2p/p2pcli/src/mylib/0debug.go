package mylib

import (
	log "devlog"
	"fmt"
	"os"
	"runtime"
	"strings"
)

const (
	LOG_DEBUG int = iota
	LOG_INFO
	LOG_WARNING
	LOG_ERROR
	LOG_OFF
	LOG_UNKNOWN
)

var ThisPrnLog PrnLogStu

const (
	DEBUGName   = "DEBUG"
	INFOName    = "INFO"
	WARNINGName = "WARNING"
	ERRORName   = "ERROR"
)

const formatLeader string = "→→→"
const formatTag string = "tag="
const formatLongTime string = "+"
const formatLeaderLen int = len(formatLeader)

// sysflag = "smartcom"
type PrnLogStu struct {
	Sysflag string

	LogLevel   int
	DbLogLevel int

	UseFulFmt      bool
	SaveToFileOnly bool
	LogPATHFile    string
	logTimes       uint32

	LimitLenPerLine int
}

func GetPrnLogInstance() (inst *PrnLogStu) {
	inst = &ThisPrnLog
	return
}

func (this *PrnLogStu) SetLogLevel(logLevel, DbLogLevel string) {
	logLevel = strings.TrimSpace(strings.ToUpper(logLevel))
	DbLogLevel = strings.TrimSpace(strings.ToUpper(DbLogLevel))
	switch logLevel {
	case DEBUGName:
		this.LogLevel = LOG_DEBUG
	case INFOName:
		this.LogLevel = LOG_INFO
	case WARNINGName:
		this.LogLevel = LOG_WARNING
	case ERRORName:
		this.LogLevel = LOG_ERROR
	}

	switch DbLogLevel {
	case DEBUGName:
		this.DbLogLevel = LOG_DEBUG
	case INFOName:
		this.DbLogLevel = LOG_INFO
	case WARNINGName:
		this.DbLogLevel = LOG_WARNING
	case ERRORName:
		this.DbLogLevel = LOG_ERROR
	}
}

func (this *PrnLogStu) Debug(v ...interface{}) {
	if this.LogLevel <= LOG_DEBUG {
		this.output(DEBUGName, "", v...)
	}
}

func (this *PrnLogStu) Info(v ...interface{}) {
	if this.LogLevel <= LOG_INFO {
		this.output(INFOName, "", v...)
	}
}

func (this *PrnLogStu) Warning(v ...interface{}) {
	if this.LogLevel <= LOG_WARNING {
		this.output(WARNINGName, "", v...)
	}
}

func (this *PrnLogStu) Error(v ...interface{}) {
	if this.LogLevel <= LOG_ERROR {
		this.output(ERRORName, "", v...)
	}
}

func (this *PrnLogStu) output(info string, format string, v ...interface{}) {

	debug1 := false

	if this.UseFulFmt {
		pc, file, line, _ := runtime.Caller(2)
		short := callerFileNameGet(file)
		f := runtime.FuncForPC(pc)
		fn := f.Name()

		for i := len(fn) - 1; i > 0; i-- {
			if fn[i] == '.' {
				fn = fn[i+1:]
				break
			}
		}

		if debug1 {
			fmt.Println("1", format, formatLeader) //
		}

		if format == "" {
			if debug1 {
				fmt.Println("2") //
			}

			str1 := fmt.Sprintln(v...)
			if len(str1) > this.LimitLenPerLine {
				str1 = str1[:this.LimitLenPerLine] + "..."
			}
			log.Printf("|%v|%v|%v|%v()|%v|%v", info, this.Sysflag, short, fn, line, str1)
		} else if len(format) >= formatLeaderLen && format[:formatLeaderLen] == formatLeader {
			format = format[formatLeaderLen:]
			str1 := format + fmt.Sprintln(v...)

			if debug1 {
				fmt.Println("3", format, str1) //
			}

			if len(str1) > this.LimitLenPerLine {
				str1 = str1[:this.LimitLenPerLine] + "..."
			}
			log.Printf("|%v|%v|%v|%v()|%v|%v", info, this.Sysflag, short, fn, line, str1)
		} else {
			str1 := fmt.Sprintf(format, v...)
			if len(str1) > this.LimitLenPerLine {
				str1 = str1[:this.LimitLenPerLine] + "..."
			}

			if debug1 {
				fmt.Println("4", str1) //
			}

			log.Printf("|%v|%v|%v|%v()|%v|%v", info, this.Sysflag, short, fn, line, str1+"\n")
		}

	} else {
		if format == "" {

			str1 := fmt.Sprintln(v...)
			if len(str1) > this.LimitLenPerLine {
				str1 = str1[:this.LimitLenPerLine] + "..."
			}
			log.Printf("[%s]|%v", info, str1)
		} else {
			str1 := fmt.Sprintf(format, v...)
			if len(str1) > this.LimitLenPerLine {
				str1 = str1[:this.LimitLenPerLine] + "..."
			}
			log.Printf("[%s]|%v", info, str1+"\n")
		}
	}

}

func callerFileNameGet(file_runtime string) string {
	short := file_runtime
	for i := len(file_runtime) - 1; i > 0; i-- {
		if file_runtime[i] == '/' {
			short = file_runtime[i+1:]
			break
		}
	}

	return short
}

func (this *PrnLogStu) Logger() *os.File {
	if !this.SaveToFileOnly {
		return nil
	}
	log.Printf(" 仅打印到文件=%v, 适时刷新到文件、非实时。\n\n", this.LogPATHFile)

	f, err := os.OpenFile(this.LogPATHFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(f)

	return f
}

func (this *PrnLogStu) Init(prnIn *PrnLogStu) {

	prnIn.SaveToFileOnly = len(prnIn.LogPATHFile) > 0

	this.Sysflag = prnIn.Sysflag
	this.LogLevel = prnIn.LogLevel
	this.DbLogLevel = prnIn.DbLogLevel
	this.UseFulFmt = prnIn.UseFulFmt
	this.SaveToFileOnly = prnIn.SaveToFileOnly
	this.LogPATHFile = prnIn.LogPATHFile
	this.logTimes = 0

	if prnIn.LimitLenPerLine <= 0 {
		prnIn.LimitLenPerLine = 5000
	}
	this.LimitLenPerLine = prnIn.LimitLenPerLine

	log.SetFlags(log.Ldd | log.Ltime /* | log.Lmicroseconds*/)

	this.Logger()
}

func init() {

	prnLog := &PrnLogStu{
		Sysflag:         "gosrc",
		LogLevel:        LOG_DEBUG,
		DbLogLevel:      LOG_DEBUG,
		UseFulFmt:       true,
		SaveToFileOnly:  false,
		LogPATHFile:     "",
		LimitLenPerLine: 0,
	}
	ThisPrnLog.Init(prnLog)
}
