package nbt


type TagNodeByteArray struct {
  _data []byte
}


func (t *TagNodeByteArray) Data() []byte {
  return t._data
}
