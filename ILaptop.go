package main

type ILaptop interface {
	setRam(ram int)
	setColor(color string)
	getRam() int
	getColor() string
}

type Laptop struct {
	ram   int
	color string
}

func (l *Laptop) setRam(ram int) {
	l.ram = ram
}

func (l *Laptop) setColor(color string) {
	l.color = color
}

func (l *Laptop) getRam() int {
	return l.ram
}

func (l *Laptop) getColor() string {
	return l.color
}
