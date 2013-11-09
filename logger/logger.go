package logger


import (
  //"fmt"
)


type Logger struct {
  myvar int
}


func NewLogger() Logger {
  logger := Logger{}
  return logger
}


func (l Logger) Debug(msg string) {
  //fmt.Println(msg)
}
