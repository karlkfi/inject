package variadic

type V1 struct{}

func NewV1() *V1 {
	return &V1{}
}

func (v1 V1) Something() bool {
	return true
}

func (v1 V1) String() string {
	return "v1"
}
