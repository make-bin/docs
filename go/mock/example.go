package mock

type Size struct {
	app Apple
}

func (s Size) GetSize() {
	s.app.AppleSize()
}
