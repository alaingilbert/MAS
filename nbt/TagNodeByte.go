package nbt

// TagNodeByte ...
type TagNodeByte struct {
  mData byte
}

// NewTagNodeByte ...
func NewTagNodeByte(data byte) *TagNodeByte {
  return &TagNodeByte{data}
}
