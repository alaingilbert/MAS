package nbt

import (
  "encoding/binary"
  "io"
  "log"
  "mas/logger"
  "math"
)

// NewNbtTree ...
func NewNbtTree() NbtTree {
  tree := NbtTree{}
  tree.mLogger = logger.NewLogger(logger.INFO)
  return tree
}

// NbtTree ...
type NbtTree struct {
  Stream    io.Reader
  _root     TagNodeCompound
  _rootName string
  mLogger   logger.Logger
}

// Root ...
func (n *NbtTree) Root() TagNodeCompound {
  n.mLogger.Debug("Root")
  return n._root
}

// Init ...
func (n *NbtTree) Init(r io.Reader) {
  n.mLogger.Debug("Init")
  n.Stream = r
  n._root = n.ReadRoot()
}

// ReadRoot ...
func (n *NbtTree) ReadRoot() TagNodeCompound {
  n.mLogger.Debug("ReadRoot")
  tagType := TagType(ReadByte(n.Stream))
  if tagType == TagCompound {
    n._rootName = ReadString(n.Stream)
    return n.ReadValue(tagType).(TagNodeCompound)
  }
  return TagNodeCompound{}
}

// ReadValue ...
func (n *NbtTree) ReadValue(tagType TagType) TagNode {
  n.mLogger.Debug("ReadValue")
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
func (n *NbtTree) ReadFloat() TagNode {
  n.mLogger.Debug("ReadFloat")
  val := TagNodeFloat{ReadFloat(n.Stream)}
  return val
}

// ReadString ...
func (n *NbtTree) ReadString() TagNode {
  n.mLogger.Debug("ReadString")
  str := ReadString(n.Stream)
  val := TagNodeString{str}
  return val
}

// ReadDouble ...
func (n *NbtTree) ReadDouble() TagNode {
  n.mLogger.Debug("ReadDouble")
  val := TagNodeDouble{ReadDouble(n.Stream)}
  return val
}

// ReadShort ...
func (n *NbtTree) ReadShort() TagNode {
  n.mLogger.Debug("ReadShort")
  val := TagNodeShort{ReadShort(n.Stream)}
  return val
}

// ReadIntArray ...
func (n *NbtTree) ReadIntArray() TagNode {
  n.mLogger.Debug("ReadIntArray")
  size := ReadInt(n.Stream)
  data := make([]int32, size)
  for i := int32(0); i < size; i++ {
    tmpInt := ReadInt(n.Stream)
    data[i] = tmpInt
  }
  val := TagNodeIntArray{data}
  return val
}

// ReadByte ...
func (n *NbtTree) ReadByte() TagNode {
  n.mLogger.Debug("ReadByte")
  val := TagNodeByte{ReadByte(n.Stream)}
  return val
}

// ReadInt ...
func (n *NbtTree) ReadInt() TagNode {
  n.mLogger.Debug("ReadInt")
  val := TagNodeInt{ReadInt(n.Stream)}
  return val
}

// ReadLong ...
func (n *NbtTree) ReadLong() TagNode {
  n.mLogger.Debug("ReadLong")
  long := ReadLong(n.Stream)
  val := TagNodeLong{long}
  return val
}

// ReadByteArray ...
func (n *NbtTree) ReadByteArray() TagNode {
  n.mLogger.Debug("ReadByteArray")
  size := ReadInt(n.Stream)
  if size < 0 {
    log.Fatal("Read Neg")
  }
  data := make([]byte, size)
  n.Stream.Read(data)

  val := TagNodeByteArray{data}
  return val
}

// ReadList ...
func (n *NbtTree) ReadList() TagNode {
  n.mLogger.Debug("ReadList")
  tagID := TagType(ReadByte(n.Stream))
  length := ReadInt(n.Stream)
  list := make([]TagNode, length)
  val := TagNodeList{tagID, length, list}
  if val.ValueType() == TagEnd {
    return TagNodeList{TagByte, length, list}
  }
  for i := 0; int32(i) < length; i++ {
    val.Add(n.ReadValue(val.ValueType()), i)
  }
  return val
}

// ReadCompound ...
func (n *NbtTree) ReadCompound() TagNode {
  n.mLogger.Debug("ReadCompound")
  val := TagNodeCompound{make(map[string]TagNode)}
  for n.ReadTag(val) {
  }
  return val
}

// ReadTag ...
func (n *NbtTree) ReadTag(parent TagNodeCompound) bool {
  n.mLogger.Debug("ReadTag")
  tagType := TagType(ReadByte(n.Stream))
  //fmt.Println("-> (ReadTag)", tagType)
  if tagType != TagEnd {
    name := ReadString(n.Stream)
    value := n.ReadValue(tagType)
    //fmt.Println(name, value)
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
