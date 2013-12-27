package nbt

// TagNodeDouble ...
type TagNodeDouble struct {
  Data float64
}

// NewTagNodeDouble ...
func NewTagNodeDouble(data float64) *TagNodeDouble {
  tagNode := &TagNodeDouble{data}
  return tagNode
}
