package nbt

// TagNodeString ...
type TagNodeString struct {
  mData string
}

// ToString ...
func (t *TagNodeString) ToString() string {
  return t.mData
}
