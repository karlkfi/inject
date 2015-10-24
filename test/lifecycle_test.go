package test

type initme struct {
	initialized bool
}

func (i *initme) Initialize() {
	i.initialized = true
}

type finalme struct {
	finalized bool
}

func (f *finalme) Finalize() {
	f.finalized = true
}

type lifecycleme struct {
	initialized bool
	finalized bool
}

func (l *lifecycleme) Initialize() {
	l.initialized = true
}

func (l *lifecycleme) Finalize() {
	l.finalized = true
}
