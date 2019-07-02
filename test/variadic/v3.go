package variadic

type V3 struct{}

func NewV3() *V3 {
	return &V3{}
}

func (v3 V3) Something() bool {
	return true
}

func (v3 V3) String() string {
	return "v3"
}
