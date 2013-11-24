package logger

import (
  "fmt"
  "os"
  "time"
)

// INFO ...
const INFO int = 1

// DEBUG ...
const DEBUG int = 2

// WARNING ...
const WARNING int = 4

// ERROR ...
const ERROR int = 8

// Logger ...
type Logger struct {
  mLevel int
}

// NewLogger ...
func NewLogger(pLevel int) Logger {
  logger := Logger{}
  logger.mLevel = pLevel
  return logger
}

func (l Logger) printLn(args []interface{}) {
  now := time.Now()
  t := fmt.Sprintf("%04d/%02d/%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
  fmt.Println(t, args)
}

// Debug ...
func (l Logger) Debug(args ...interface{}) {
  if l.mLevel&DEBUG > 0 {
    l.printLn(args)
  }
}

// Info ...
func (l Logger) Info(args ...interface{}) {
  if l.mLevel&INFO > 0 {
    l.printLn(args)
  }
}

// Warning ... 
func (l Logger) Warning(args ...interface{}) {
  if l.mLevel&WARNING > 0 {
    l.printLn(args)
  }
}

func (l Logger) Error(args ...interface{}) {
  if l.mLevel&ERROR > 0 {
    l.printLn(args)
  }
}

// Fatal ...
func (l Logger) Fatal(args ...interface{}) {
  l.printLn(args)
  os.Exit(0)
}
