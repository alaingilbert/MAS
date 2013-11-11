package nbt


type TagNodeList struct {
  _type TagType
  m_Length int32
  m_List []TagNode
}


func (t *TagNodeList) Length() int32 {
  return t.m_Length
}


func (t *TagNodeList) Get(i int) TagNode {
  return t.m_List[i]
}


func (t *TagNodeList) Add(item TagNode, i int) {
  t.m_List[i] = item
}


func (t *TagNodeList) ValueType() TagType {
  return t._type
}
