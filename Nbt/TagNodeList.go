package Nbt


type TagNodeList struct {
  _type byte
}


func (t *TagNodeList) Add(item TagNode) {
}


func (t *TagNodeList) ValueType() byte {
  return t._type
}
