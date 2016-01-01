package main

import (
	"os"
	"syscall"
	"time"

	log "github.com/cihub/seelog"
)

func openOutLog(filename string) *os.File {
	// Move existing out file to a dated file if it exists
	if _, err := os.Stat(filename); err == nil {
		if err = os.Rename(filename, filename+"."+time.Now().Format("2006-01-02_15:04:05")); err != nil {
			log.Criticalf("Cannot move old out file: %v", err)
			os.Exit(1)
		}
	}

	// Redirect stdout and stderr to out file
	logFile, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0644)
	syscall.Dup2(int(logFile.Fd()), 1)
	syscall.Dup2(int(logFile.Fd()), 2)
	return logFile
}

func NewLogger(cfgfile string) {
	logger, err := log.LoggerFromConfigAsFile(cfgfile)
	if err != nil {
		log.Criticalf("Cannot start logger: %v", err)
		os.Exit(1)
	}
	log.ReplaceLogger(logger)
}

func testLog() {
	logdir := "log"
	logconfig := "config/logging.cfg"
	openOutLog(logdir + "/burrow.out")
	NewLogger(logconfig)

	log.Debug("debug log info")
	log.Info("info log info")
	log.Warn("warn log info")
	log.Error("error log info")
	log.Criticalf("criticalf log info")
}
