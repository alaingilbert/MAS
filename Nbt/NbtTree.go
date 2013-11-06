package NbtTree


import (
  "fmt"
  "io"
)


const (
  TAG_END = iota
  TAG_BYTE
  TAG_SHORT
  TAG_INT
  TAG_LONG
  TAG_FLOAT
  TAG_DOUBLE
  TAG_BYTE_ARRAY
  TAG_STRING
  TAG_LIST
  TAG_COMPOUND
  TAG_INT_ARRAY
)


type NbtTree struct {
  Stream io.Reader
}


func (n *NbtTree) Init(r io.Reader) {
  n.Stream = r
}


func (n *NbtTree) Nothing() {
  fmt.Println("Nothing")
}
