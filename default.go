package sprint

// DefaultIfNil returns val if val is not nil, def otherwise
func DefaultIfNil[V any](val, def *V) *V {
	if val == nil {
		return def
	}
	return val
}

// DefaultIfEmpty returns v if not empty otherwise d
func DefaultIfEmpty(v string, d string) string {
	if v == "" {
		return d
	}
	return v
}
