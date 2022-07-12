package isp

import "fmt"

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type Document struct{}

func (d *Document) Print() {
	fmt.Println(d)
}
