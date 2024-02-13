package sprint

// DefaultIfNil returns val if val is not nil, def otherwise
func DefaultIfNil[V comparable](val, def *V) *V {
	if val == nil {
		return def
	}
	return val
}

// ToPointer returns a pointer to v
func ToPointer[V comparable](v V) *V {
	return &v
}
