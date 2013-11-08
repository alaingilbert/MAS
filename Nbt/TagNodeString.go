package Nbt


type TagNodeString struct {
  _data string
}


func (t *TagNodeString) ToString() string {
  return t._data
}
