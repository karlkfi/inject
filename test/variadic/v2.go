package variadic

type V2 struct {}

func NewV2() V2 {
	return V2{}
}

func (v2 V2) Something() bool {
	return true
}

func (v2 V2) String() string {
	return "v2"
}
