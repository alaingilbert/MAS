package nbt

// TagNodeInt ...
type TagNodeInt struct {
  mData int32
}

// NewTagNodeInt ...
func NewTagNodeInt(data int32) *TagNodeInt {
  return &TagNodeInt{data}
}
