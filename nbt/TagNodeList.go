package nbt

// TagNodeList ...
type TagNodeList struct {
  mType   TagType
  mLength int32
  mList   []TagNode
}

// Length ...
func (t *TagNodeList) Length() int32 {
  return t.mLength
}

// Get ...
func (t *TagNodeList) Get(i int) TagNode {
  return t.mList[i]
}

// Add ...
func (t *TagNodeList) Add(item TagNode, i int) {
  t.mList[i] = item
}

// ValueType ...
func (t *TagNodeList) ValueType() TagType {
  return t.mType
}
