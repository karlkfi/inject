package test

type alpha struct {
	name string
}

func (a alpha) Name() string {
	return a.name
}

type beta struct {
	name string
}

func (b beta) Name() string {
	return b.name
}

type omega interface {
	Name() string
}

// similar type that doesn't satisfy the omega interface
type gamma struct {
	name string
}
