package nbt

import (
  "encoding/binary"
  "io"
  "log"
  "math"
)

// NewNbtTree ...
func NewNbtTree(r io.Reader) *NbtTree {
  tree := new(NbtTree)
  tree.Stream = r
  tree._root = tree.ReadRoot()
  return tree
}

// NbtTree ...
type NbtTree struct {
  Stream    io.Reader
  _root     *TagNodeCompound
  _rootName string
}

// Root ...
func (n *NbtTree) Root() *TagNodeCompound {
  return n._root
}

// ReadRoot ...
func (n *NbtTree) ReadRoot() *TagNodeCompound {
  tagType := TagType(ReadByte(n.Stream))
  if tagType == TagCompound {
    n._rootName = ReadString(n.Stream)
    return n.ReadValue(tagType).(*TagNodeCompound)
  }
  return new(TagNodeCompound)
}

// ReadValue ...
func (n *NbtTree) ReadValue(tagType TagType) ITagNode {
  //fmt.Println("-> (ReadValue)", tagType)
  switch tagType {
  case TagByte:
    return n.ReadByte()
  case TagCompound:
    return n.ReadCompound()
  case TagList:
    return n.ReadList()
  case TagByteArray:
    return n.ReadByteArray()
  case TagLong:
    return n.ReadLong()
  case TagInt:
    return n.ReadInt()
  case TagIntArray:
    return n.ReadIntArray()
  case TagShort:
    return n.ReadShort()
  case TagFloat:
    return n.ReadFloat()
  case TagDouble:
    return n.ReadDouble()
  case TagString:
    return n.ReadString()
  default:
    log.Println("Unknow TagNode", tagType)
    return TagNodeUnknown{}
  }
}

// ReadFloat ...
func (n *NbtTree) ReadFloat() ITagNode {
  val := NewTagNodeFloat(ReadFloat(n.Stream))
  return val
}

// ReadString ...
func (n *NbtTree) ReadString() ITagNode {
  str := ReadString(n.Stream)
  val := NewTagNodeString(str)
  return val
}

// ReadDouble ...
func (n *NbtTree) ReadDouble() ITagNode {
  val := NewTagNodeDouble(ReadDouble(n.Stream))
  return val
}

// ReadShort ...
func (n *NbtTree) ReadShort() ITagNode {
  val := NewTagNodeShort(ReadShort(n.Stream))
  return val
}

// ReadIntArray ...
func (n *NbtTree) ReadIntArray() ITagNode {
  size := ReadInt(n.Stream)
  data := make([]int32, size)
  for i := int32(0); i < size; i++ {
    tmpInt := ReadInt(n.Stream)
    data[i] = tmpInt
  }
  val := NewTagNodeIntArray(data)
  return val
}

// ReadByte ...
func (n *NbtTree) ReadByte() ITagNode {
  val := NewTagNodeByte(ReadByte(n.Stream))
  return val
}

// ReadInt ...
func (n *NbtTree) ReadInt() ITagNode {
  val := NewTagNodeInt(ReadInt(n.Stream))
  return val
}

// ReadLong ...
func (n *NbtTree) ReadLong() ITagNode {
  long := ReadLong(n.Stream)
  val := NewTagNodeLong(long)
  return val
}

// ReadByteArray ...
func (n *NbtTree) ReadByteArray() ITagNode {
  size := ReadInt(n.Stream)
  if size < 0 {
    log.Fatal("Read Neg")
  }
  data := make([]byte, size)
  n.Stream.Read(data)

  val := NewTagNodeByteArray(data)
  return val
}

// ReadList ...
func (n *NbtTree) ReadList() ITagNode {
  tagID := TagType(ReadByte(n.Stream))
  length := ReadInt(n.Stream)
  list := make([]ITagNode, length)
  val := NewTagNodeList(tagID, length, list)
  if val.ValueType() == TagEnd {
    return NewTagNodeList(TagByte, length, list)
  }
  for i := 0; int32(i) < length; i++ {
    val.Add(n.ReadValue(val.ValueType()), i)
  }
  return val
}

// ReadCompound ...
func (n *NbtTree) ReadCompound() ITagNode {
  val := NewTagNodeCompound(make(map[string]ITagNode))
  for n.ReadTag(val) {
  }
  return val
}

// ReadTag ...
func (n *NbtTree) ReadTag(parent *TagNodeCompound) bool {
  tagType := TagType(ReadByte(n.Stream))
  if tagType != TagEnd {
    name := ReadString(n.Stream)
    value := n.ReadValue(tagType)
    parent.Entries[name] = value
    return true
  }
  return false
}

// ReadByte ...
func ReadByte(r io.Reader) (i byte) {
  b := make([]byte, 1)
  r.Read(b)
  i = b[0]
  return
}

// ReadShort ...
func ReadShort(r io.Reader) (i int16) {
  binary.Read(r, binary.BigEndian, &i)
  return
}

// ReadInt ...
func ReadInt(r io.Reader) (i int32) {
  binary.Read(r, binary.BigEndian, &i)
  return
}

// ReadLong ...
func ReadLong(r io.Reader) (i int64) {
  binary.Read(r, binary.BigEndian, &i)
  return
}

// ReadFloat ...
func ReadFloat(r io.Reader) (i float32) {
  b := make([]byte, 4)
  r.Read(b)
  i = math.Float32frombits(binary.BigEndian.Uint32(b))
  return
}

// ReadDouble ...
func ReadDouble(r io.Reader) (i float64) {
  b := make([]byte, 8)
  r.Read(b)
  i = math.Float64frombits(binary.BigEndian.Uint64(b))
  return
}

// ReadByteArray ...
func ReadByteArray(r io.Reader) (i []byte) {
  i = make([]byte, ReadInt(r))
  r.Read(i)
  return
}

// ReadString ...
func ReadString(r io.Reader) string {
  result := make([]byte, ReadShort(r))
  r.Read(result)
  return string(result)
}

// ReadIntArray ...
func ReadIntArray(r io.Reader) (list []int32) {
  length := int(ReadInt(r))
  for i := 0; i < length; i++ {
    list = append(list, ReadInt(r))
  }
  return
}
