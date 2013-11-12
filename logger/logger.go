package logger


import (
  "fmt"
  "os"
  "time"
)


const INFO int = 1
const DEBUG int = 2
const WARNING int = 4
const ERROR int = 8


type Logger struct {
  m_Level int
}


func NewLogger(p_Level int) Logger {
  logger := Logger{}
  logger.m_Level = p_Level
  return logger
}


func (l Logger) printLn(args []interface{}) {
  now := time.Now()
  t := fmt.Sprintf("%04d/%02d/%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
  fmt.Println(t, args)
}


func (l Logger) Debug(args ...interface{}) {
  if l.m_Level & DEBUG > 0 {
    l.printLn(args)
  }
}


func (l Logger) Info(args ...interface{}) {
  if l.m_Level & INFO > 0 {
    l.printLn(args)
  }
}

func (l Logger) Warning(args ...interface{}) {
  if l.m_Level & WARNING > 0 {
    l.printLn(args)
  }
}


func (l Logger) Error(args ...interface{}) {
  if l.m_Level & ERROR > 0 {
    l.printLn(args)
  }
}


func (l Logger) Fatal(args ...interface{}) {
  l.printLn(args)
  os.Exit(0)
}
