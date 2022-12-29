package main

import (
	"fmt"
	"strings"
)

type EnergyUnit string

// Sets the string values of the respective constants.
const (
	EnergyUnit_GAL   EnergyUnit = "GAL"
	EnergyUnit_FT3   EnergyUnit = "FT3"
	EnergyUnit_KWH   EnergyUnit = "KWH"
	EnergyUnit_LITER EnergyUnit = "LITER"
)

func main() {
	var f int
	fmt.Println(float64(f))

	t := "gal"
	s := "kwh"
	fmt.Println(EnergyUnit(strings.ToUpper(t)))
	fmt.Printf("%s\n", t)

	switch s {
	case "kwh":
		//fmt.Println("did not fall through")
	case "gal":
		fmt.Println("fell through")
	}
	u, v := "test", EnergyUnit("gogo")
	fmt.Println(EnergyUnit(strings.ToUpper(u)), v)
}
