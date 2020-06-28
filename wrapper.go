package bunnlog

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

func GetBunnLog(printStacks bool, verbosity, flag int) BunnyLog {
	// Manually set stackOnErr for less confusion
	return BunnyLog{Logger: log.New(ioutil.Discard, "", flag), Verbosity: verbosity, StackOnErr: false, StackOnFatal: printStacks}
}

func (b *BunnyLog) SetVerbosity(verbosity int) {
	b.Verbosity = verbosity
}

func (b *BunnyLog) SetPrintStacksOnError(value bool) {
	b.StackOnErr = value
}

func (b *BunnyLog) SetPrintStacksOnFatal(value bool) {
	b.StackOnFatal = value
}

func (b *BunnyLog) SetOutputFile(file *os.File) {
	b.Logger.SetOutput(file)
}

func (b *BunnyLog) SetOutputTerminal() {
	b.Logger.SetOutput(os.Stdout)
}

func (b *BunnyLog) SetOutputFileAndTerminal(file *os.File) {
	wrt := io.MultiWriter(os.Stdout, file)
	b.Logger.SetOutput(wrt)
}

func (b *BunnyLog) Debug(msg ...interface{}) {
	if b.Verbosity <= VerbosityDEBUG && msg != nil {
		b.Logger.SetPrefix("[DEBUG]:\t")
		b.Logger.Print(msg...)
	}
}

func (b *BunnyLog) Debugln(msg ...interface{}) {
	if b.Verbosity <= VerbosityDEBUG && msg != nil {
		b.Logger.SetPrefix("[DEBUG]:\t")
		b.Logger.Println(msg...)
	}
}

func (b *BunnyLog) Debugf(format string, v ...interface{}) {
	if b.Verbosity <= VerbosityDEBUG {
		b.Logger.SetPrefix("[DEBUG]:\t")
		b.Logger.Printf(format, v...)
	}
}

func (b *BunnyLog) Info(msg ...interface{}) {
	if b.Verbosity <= VerbosityINFO && msg != nil {
		b.Logger.SetPrefix("[INFO]:\t")
		b.Logger.Print(msg...)
	}
}

func (b *BunnyLog) Infoln(msg ...interface{}) {
	if b.Verbosity <= VerbosityINFO && msg != nil {
		b.Logger.SetPrefix("[INFO]:\t")
		b.Logger.Println(msg...)
	}
}

func (b *BunnyLog) Infof(format string, v ...interface{}) {
	if b.Verbosity <= VerbosityINFO {
		b.Logger.SetPrefix("[INFO]:\t")
		b.Logger.Printf(format, v...)
	}
}

func (b *BunnyLog) Warn(msg ...interface{}) {
	if b.Verbosity <= VerbosityWARNING && msg != nil {
		b.Logger.SetPrefix("[WARNING]:\t")
		b.Logger.Print(msg...)
	}
}

func (b *BunnyLog) Warnln(msg ...interface{}) {
	if b.Verbosity <= VerbosityWARNING && msg != nil {
		b.Logger.SetPrefix("[WARNING]:\t")
		b.Logger.Println(msg...)
	}
}

func (b *BunnyLog) Warnf(format string, v ...interface{}) {
	if b.Verbosity <= VerbosityWARNING {
		b.Logger.SetPrefix("[WARNING]:\t")
		b.Logger.Printf(format, v...)
	}
}

func (b *BunnyLog) Error(msg ...interface{}) {
	if b.Verbosity <= VerbosityERROR && msg != nil {
		b.Logger.SetPrefix("[ERROR]:\t")
		if b.StackOnErr {
			b.Logger.Print(b.StackTrace())
		}
		b.Logger.Print(msg...)
	}
}

func (b *BunnyLog) Errorln(msg ...interface{}) {
	if b.Verbosity <= VerbosityERROR && msg != nil {
		b.Logger.SetPrefix("[ERROR]:\t")
		if b.StackOnErr {
			b.Logger.Println(b.StackTrace())
		}
		b.Logger.Println(msg...)
	}
}

func (b *BunnyLog) Errorf(format string, v ...interface{}) {
	if b.Verbosity <= VerbosityERROR {
		b.Logger.SetPrefix("[ERROR]:\t")
		if b.StackOnErr {
			b.Logger.Print(b.StackTrace())
		}
		b.Logger.Printf(format, v...)
	}
}

func (b *BunnyLog) Fatal(msg ...interface{}) {
	if b.Verbosity <= VerbosityERROR && msg != nil {
		b.Logger.SetPrefix("[FATAL]:\t")
		if b.StackOnFatal {
			b.Logger.Print(b.StackTrace())
		}
		b.Logger.Fatal(msg...)
	}
}

func (b *BunnyLog) Fatalln(msg ...interface{}) {
	if b.Verbosity <= VerbosityERROR && msg != nil {
		b.Logger.SetPrefix("[FATAL]:\t")
		if b.StackOnFatal {
			b.Logger.Println(b.StackTrace())
		}
		b.Logger.Fatalln(msg...)
	}
}

func (b *BunnyLog) Fatalf(format string, v ...interface{}) {
	if b.Verbosity <= VerbosityERROR {
		b.Logger.SetPrefix("[FATAL]:\t")
		if b.StackOnFatal {
			b.Logger.Print(b.StackTrace())
		}
		b.Logger.Fatalf(format, v...)
	}
}

func (b *BunnyLog) StackTrace() string {
	buf := make([]byte, 2048)
	n := runtime.Stack(buf, false)
	return string(buf[:n])
}
