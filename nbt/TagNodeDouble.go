package nbt


type TagNodeDouble struct {
  _data float64
}


func (t *TagNodeDouble) Data() float64 {
  return t._data
}
