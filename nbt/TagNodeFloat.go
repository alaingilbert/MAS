package nbt

// TagNodeFloat ...
type TagNodeFloat struct {
  mData float32
}

// NewTagNodeFloat ...
func NewTagNodeFloat(data float32) *TagNodeFloat {
  return &TagNodeFloat{data}
}
