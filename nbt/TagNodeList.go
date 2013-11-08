package nbt


type TagNodeList struct {
  _type TagType
}


func (t *TagNodeList) Add(item TagNode) {
}


func (t *TagNodeList) ValueType() TagType {
  return t._type
}
