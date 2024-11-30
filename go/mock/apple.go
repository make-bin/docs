package mock

type Apple interface {
	AppleSize() int
}

type apple struct {
}

func (a *apple) AppleSize() int {
	return 10
}

func NewApple() Apple {
	return &apple{}
}
