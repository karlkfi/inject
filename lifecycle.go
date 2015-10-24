package inject

// Initializable describes an object that needs initialization after being created
type Initializable interface {
	Initialize()
}

// Finalizable describes an object that needs finalization before being destroyed
type Finalizable interface {
	Finalize()
}
