package nbt

// TagNodeIntArray ...
type TagNodeIntArray struct {
  mData []int32
}

// Data ...
func (t TagNodeIntArray) Data() []int32 {
  return t.mData
}
