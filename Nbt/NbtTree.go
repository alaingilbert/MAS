package Nbt


import (
  //"fmt"
  "log"
  "io"
  "math"
  "encoding/binary"
)



type NbtTree struct {
  Stream io.Reader
  _root TagNodeCompound
  _rootName string
}

func (n *NbtTree) Root() TagNodeCompound {
  return n._root
}

func (n *NbtTree) Init(r io.Reader) {
  n.Stream = r
  n._root = n.ReadRoot()
}

func (n *NbtTree) ReadRoot() TagNodeCompound {
  tagType := TagType(ReadByte(n.Stream))
  if tagType == TAG_COMPOUND {
    n._rootName = ReadString(n.Stream)
    return n.ReadValue(tagType).(TagNodeCompound)
  }
  return TagNodeCompound{}
}

func (n *NbtTree) ReadValue(tagType TagType) TagNode {
  switch tagType {
  case TAG_BYTE:
    return n.ReadByte()
  case TAG_COMPOUND:
    return n.ReadCompound()
  case TAG_LIST:
    return n.ReadList()
  case TAG_BYTE_ARRAY:
    return n.ReadByteArray()
  case TAG_LONG:
    return n.ReadLong()
  case TAG_INT:
    return n.ReadInt()
  case TAG_INT_ARRAY:
    return n.ReadIntArray()
  case TAG_SHORT:
    return n.ReadShort()
  case TAG_FLOAT:
    return n.ReadFloat()
  case TAG_DOUBLE:
    return n.ReadDouble()
  case TAG_STRING:
    return n.ReadString()
  default:
    log.Println("Unknow TagNode", tagType)
    return TagNodeUnknown{}
  }
}

func (n *NbtTree) ReadFloat() TagNode {
  val := TagNodeFloat{ReadFloat(n.Stream)}
  return val
}

func (n *NbtTree) ReadString() TagNode {
  str := ReadString(n.Stream)
  val := TagNodeString{str}
  return val
}

func (n *NbtTree) ReadDouble() TagNode {
  val := TagNodeDouble{ReadDouble(n.Stream)}
  return val
}

func (n *NbtTree) ReadShort() TagNode {
  val := TagNodeShort{ReadShort(n.Stream)}
  return val
}

func (n *NbtTree) ReadIntArray() TagNode {
  size := ReadInt(n.Stream)
  data := make([]int32, size)
  for i := int32(0); i < size; i++ {
    tmpInt := ReadInt(n.Stream)
    data[i] = tmpInt
  }
  val := TagNodeIntArray{data}
  return val
}

func (n *NbtTree) ReadByte() TagNode {
  val := TagNodeByte{ReadByte(n.Stream)}
  return val
}

func (n *NbtTree) ReadInt() TagNode {
  val := TagNodeInt{ReadInt(n.Stream)}
  return val
}

func (n *NbtTree) ReadLong() TagNode {
  long := ReadLong(n.Stream)
  val := TagNodeLong{long}
  return val
}

func (n *NbtTree) ReadByteArray() TagNode {
  size := ReadInt(n.Stream)
  if size < 0 {
    log.Fatal("Read Neg")
  }
  data := make([]byte, size)
  n.Stream.Read(data)

  val := TagNodeByteArray{data}
  return val
}

func (n *NbtTree) ReadList() TagNode {
  tagId := TagType(ReadByte(n.Stream))
  val := TagNodeList{tagId}
  length := ReadInt(n.Stream)
  if val.ValueType() == TAG_END {
    return TagNodeList{TAG_BYTE}
  }
  for i := 0; int32(i) < length; i++ {
    val.Add(n.ReadValue(val.ValueType()))
  }
  return val
}

func (n *NbtTree) ReadCompound() TagNode {
  val := TagNodeCompound{make(map[string]TagNode)}
  for n.ReadTag(val) {}
  return val
}

func (n *NbtTree) ReadTag(parent TagNodeCompound) bool {
  tagType := TagType(ReadByte(n.Stream))
  if tagType != TAG_END {
    name := ReadString(n.Stream)
    value := n.ReadValue(tagType)
    //fmt.Println(name, value)
    parent.Entries[name] = value
    return true
  }
  return false
}

func ReadByte(r io.Reader) (i byte) {
  b := make([]byte, 1)
  r.Read(b)
  i = b[0]
  return
}

func ReadShort(r io.Reader) (i int16) {
  binary.Read(r, binary.BigEndian, &i)
  return
}
 
func ReadInt(r io.Reader) (i int32) {
  binary.Read(r, binary.BigEndian, &i)
  return
}
 
func ReadLong(r io.Reader) (i int64) {
  binary.Read(r, binary.BigEndian, &i)
  return
}
 
func ReadFloat(r io.Reader) (i float32) {
  b := make([]byte, 4)
  r.Read(b)
  i = math.Float32frombits(binary.BigEndian.Uint32(b))
  return
}
 
func ReadDouble(r io.Reader) (i float64) {
  b := make([]byte, 8)
  r.Read(b)
  i = math.Float64frombits(binary.BigEndian.Uint64(b))
  return
}
 
func ReadByteArray(r io.Reader) (i []byte) {
  i = make([]byte, ReadInt(r))
  r.Read(i)
  return
}
 
func ReadString(r io.Reader) string {
  result := make([]byte, ReadShort(r))
  r.Read(result)
  return string(result)
}
 
func ReadIntArray(r io.Reader) (list []int32) {
  length := int(ReadInt(r))
  for i := 0; i < length; i++ {
    list = append(list, ReadInt(r))
  }
  return
}
