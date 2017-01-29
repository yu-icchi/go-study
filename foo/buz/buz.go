package buz

type Printer interface {
	P()
}

func Call(p Printer) {
	p.P()
}
