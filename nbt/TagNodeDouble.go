package nbt

// TagNodeDouble ...
type TagNodeDouble struct {
  mData float64
}

// Data ...
func (t *TagNodeDouble) Data() float64 {
  return t.mData
}
