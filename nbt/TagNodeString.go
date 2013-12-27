package nbt

// TagNodeString ...
type TagNodeString struct {
  mData string
}

// NewTagNodeString ...
func NewTagNodeString(data string) *TagNodeString {
  return &TagNodeString{data}
}

// ToString ...
func (t *TagNodeString) ToString() string {
  return t.mData
}
