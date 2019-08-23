package bunnlog

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func GetBunnLog(verbosity, flag int) BunnyLog {
	return BunnyLog{Logger: log.New(ioutil.Discard, "", flag), Verbosity: verbosity}
}

func (b *BunnyLog) SetVerbosity(verbosity int) {
	b.Verbosity = verbosity
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

func (b *BunnyLog) Debug(msg interface{}) {
	if b.Verbosity <= VerbosityDEBUG {
		b.Logger.SetPrefix("[DEBUG]:\t")
		b.Logger.Print(msg)
	}
}

func (b *BunnyLog) Debugln(msg interface{}) {
	if b.Verbosity <= VerbosityDEBUG {
		b.Logger.SetPrefix("[DEBUG]:\t")
		b.Logger.Println(msg)
	}
}

func (b *BunnyLog) Debugf(format string, v interface{}) {
	if b.Verbosity <= VerbosityDEBUG {
		b.Logger.SetPrefix("[DEBUG]:\t")
		b.Logger.Printf(format, v)
	}
}

func (b *BunnyLog) Info(msg interface{}) {
	if b.Verbosity <= VerbosityINFO {
		b.Logger.SetPrefix("[INFO]:\t")
		b.Logger.Print(msg)
	}
}

func (b *BunnyLog) Infoln(msg interface{}) {
	if b.Verbosity <= VerbosityINFO {
		b.Logger.SetPrefix("[INFO]:\t")
		b.Logger.Println(msg)
	}
}

func (b *BunnyLog) Infof(format string, v interface{}) {
	if b.Verbosity <= VerbosityINFO {
		b.Logger.SetPrefix("[INFO]:\t")
		b.Logger.Printf(format, v)
	}
}

func (b *BunnyLog) Warn(msg interface{}) {
	if b.Verbosity <= VerbosityWARNING {
		b.Logger.SetPrefix("[WARNING]:\t")
		b.Logger.Print(msg)
	}
}

func (b *BunnyLog) Warnln(msg interface{}) {
	if b.Verbosity <= VerbosityWARNING {
		b.Logger.SetPrefix("[WARNING]:\t")
		b.Logger.Println(msg)
	}
}

func (b *BunnyLog) Warnf(format string, v interface{}) {
	if b.Verbosity <= VerbosityWARNING {
		b.Logger.SetPrefix("[WARNING]:\t")
		b.Logger.Printf(format, v)
	}
}

func (b *BunnyLog) Error(msg interface{}) {
	if b.Verbosity <= VerbosityERROR {
		b.Logger.SetPrefix("[ERROR]:\t")
		b.Logger.Print(msg)
	}
}

func (b *BunnyLog) Errorln(msg interface{}) {
	if b.Verbosity <= VerbosityERROR {
		b.Logger.SetPrefix("[ERROR]:\t")
		b.Logger.Println(msg)
	}
}

func (b *BunnyLog) Errorf(format string, v interface{}) {
	if b.Verbosity <= VerbosityERROR {
		b.Logger.SetPrefix("[ERROR]:\t")
		b.Logger.Printf(format, v)
	}
}

func (b *BunnyLog) Fatal(msg interface{}) {
	if b.Verbosity <= VerbosityERROR {
		b.Logger.SetPrefix("[FATAL]:\t")
		b.Logger.Fatal(msg)
	}
}

func (b *BunnyLog) Fatalln(msg interface{}) {
	if b.Verbosity <= VerbosityERROR {
		b.Logger.SetPrefix("[FATAL]:\t")
		b.Logger.Fatalln(msg)
	}
}

func (b *BunnyLog) Fatalf(format string, v interface{}) {
	if b.Verbosity <= VerbosityERROR {
		b.Logger.SetPrefix("[FATAL]:\t")
		b.Logger.Fatalf(format, v)
	}
}
