package main

import "fmt"

const ebulicaoK = 373.0

type EscalasTermometricas struct {
	Kelvin  float64
	Celsius float64
}

func main() {
	conversor := EscalasTermometricas{
		Kelvin:  ebulicaoK,
		Celsius: 0,
	}

	conversor.Celsius = conversor.Kelvin - 273.0

	fmt.Printf("Temperatura de ebulição da água em Kelvins = %g | Temperatura de ebulição da água em Celsius = %g.", conversor.Kelvin, conversor.Celsius)
}
