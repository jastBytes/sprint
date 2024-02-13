package sprint

// PanicOnError panics if the given error is not nil
// DO NOT USE WITH DEFER
func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

// PanicOnErrorFunc panics if the given function execution error is not nil
func PanicOnErrorFunc(f func() error) {
	if err := f(); err != nil {
		panic(err)
	}
}
