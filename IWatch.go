package main

type IWatch interface {
	setRam(ram int)
	setColor(color string)
	getRam() int
	getColor() string
}

type Watch struct {
	ram   int
	color string
}

func (w *Watch) setRam(ram int) {
	w.ram = ram
}

func (w *Watch) setColor(color string) {
	w.color = color
}

func (w *Watch) getRam() int {
	return w.ram
}

func (w *Watch) getColor() string {
	return w.color
}
