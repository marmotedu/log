package main

import (
	"flag"

	"github.com/marmotedu/log"
)

var (
	h bool

	level  int
	format string
)

func main() {
	flag.BoolVar(&h, "h", false, "Print this help.")
	flag.IntVar(&level, "l", 0, "Log level.")
	flag.StringVar(&format, "f", "console", "log output format.")

	flag.Parse()

	if h {
		flag.Usage()
		return
	}

	opts := &log.Options{
		Level:            "info",
		Format:           "console",
		EnableColor:      true,
		EnableCaller:     true,
		OutputPaths:      []string{"test.log", "stdout"},
		ErrorOutputPaths: []string{},
	}

	log.Init(opts)
	defer log.Flush()

	log.Debug("This is a debug message")
	log.Info("This is a info message")
	log.Warn("This is a warn message")
	log.Error("This is a error message")

	log.Debugf("This is a %s message", "debug")
	log.Infof("This is a %s message", "info")
	log.Warnf("This is a %s message", "warn")
	log.Errorf("This is a %s message", "error")

	log.Debug("This is a debug message", log.String("key", "value"))
	log.Info("This is a info message", log.Int32("key2", 10))
	log.Warn("This is a warn message", log.Bool("key3", false))
	log.Error("This is a error message", log.Any("key4", "any"))
}
