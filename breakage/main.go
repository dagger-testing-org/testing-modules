package main

type Breakage struct{}

func (m *Breakage) Hello() string {
	return "hello"
}
func (m *Breakage) Bye() string {
	return "bye"
}
