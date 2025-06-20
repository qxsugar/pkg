package kit

type Alice struct {
	err error
}

func NewAlice() *Alice {
	return &Alice{}
}

func New(funList ...func() error) *Alice {
	a := NewAlice()
	for _, f := range funList {
		a.Then(f)
	}
	return a
}

func (a *Alice) Then(next func() error) *Alice {
	if a.err != nil {
		return a
	}
	a.err = next()
	return a
}

func (a *Alice) Error() error {
	return a.err
}
