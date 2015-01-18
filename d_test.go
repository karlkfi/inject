package inject

type ImplD struct {}

func NewD() *ImplD {
	return &ImplD{}
}

func (d *ImplD) String() string {
	return "&ImplD{}"
}
