package main

type IPhone interface {
	setRam(ram int)
	setColor(color string)
	getRam() int
	getColor() string
}

type Phone struct {
	ram   int
	color string
}

func (p *Phone) setRam(ram int) {
	p.ram = ram
}

func (p *Phone) setColor(color string) {
	p.color = color
}

func (p *Phone) getRam() int {
	return p.ram
}

func (p *Phone) getColor() string {
	return p.color
}
