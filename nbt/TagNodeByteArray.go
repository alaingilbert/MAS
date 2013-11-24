package nbt

// TagNodeByteArray ...
type TagNodeByteArray struct {
  mData []byte
}

// Data ...
func (t *TagNodeByteArray) Data() []byte {
  return t.mData
}
