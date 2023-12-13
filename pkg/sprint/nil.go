package sprint

// DefaultIfNil returns val if val is not nil, def otherwise
func DefaultIfNil[V comparable](val, def *V) *V {
	if val == nil {
		return def
	}
	return val
}
