package mock

type Size struct {
	app Apple
}

func (s *Size) GetSize(size int) string {
	a := size
	if s.app.AppleSize(a) == a {
		return "aaaaa"
	}
	return "bbbbb"
}
