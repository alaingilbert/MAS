package nbt


type TagNodeIntArray struct {
  _data []int32
}


func (t TagNodeIntArray) Data() []int32 {
  return t._data
}
