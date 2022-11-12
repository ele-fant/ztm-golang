package main

import "fmt"

type Fuel struct {
	Diesel Diesel
}

type Diesel struct {
	Amount  *float64
	Density float64
	Unit    string
}

func main() {
	d := Diesel{}
	if d.Amount != nil {
		fmt.Println(d.Amount)
	} else {
		fmt.Println("No value")
	}
	x := 3.0
	d.Amount = &x

	a := 0.0
	if d.Amount != nil {
		a += *d.Amount
	}
	fmt.Println(a)

	fmt.Println(d.Density)
	fmt.Println()
	e := Diesel{}

	f := Fuel{
		Diesel: e,
	}
	fmt.Println(f)

	g := Fuel{}
	g.Diesel = e
	fmt.Println(g)
}
