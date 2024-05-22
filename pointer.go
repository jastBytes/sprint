package sprint

// ToPointer returns a pointer to v
func ToPointer[V any](v V) *V {
	return &v
}
