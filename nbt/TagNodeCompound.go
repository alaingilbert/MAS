package nbt

// TagNodeCompound ...
type TagNodeCompound struct {
  Entries map[string]ITagNode
}

// NewTagNodeCompound ...
func NewTagNodeCompound(pEntries map[string]ITagNode) *TagNodeCompound {
  return &TagNodeCompound{pEntries}
}
