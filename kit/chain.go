package kit

// Alice implements a chain-of-responsibility pattern for error handling.
// It executes functions sequentially and stops at the first error.
type Alice struct {
	err error
}

// NewAlice creates a new Alice instance for chaining operations.
func NewAlice() *Alice {
	return &Alice{}
}

// New creates and executes a chain of functions, stopping at the first error.
func New(funList ...func() error) *Alice {
	a := NewAlice()
	for _, f := range funList {
		a.Then(f)
	}
	return a
}

// Then adds a function to the chain. If a previous function failed,
// this function will be skipped.
func (a *Alice) Then(next func() error) *Alice {
	if a.err != nil {
		return a
	}
	a.err = next()
	return a
}

// Error returns the first error encountered in the chain, or nil if all succeeded.
func (a *Alice) Error() error {
	return a.err
}
