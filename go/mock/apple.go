package mock

type Apple interface {
	AppleSize(size int) int
}

type apple struct {
}

func (a *apple) AppleSize(size int) int {
	return size
}

func NewApple() Apple {
	return &apple{}
}
