//Package tempconv performs Celsius to Farenheit temperature computations
package main

import "fmt"

type Celsius float64
type Farenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func cToF(c Celsius) Farenheit { return Farenheit(c*9/5 + 32) }

func fToC(f Farenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func main() {
	fmt.Printf("%g° C = %g° F\n", AbsoluteZeroC, cToF(AbsoluteZeroC))
	fmt.Printf("%g° C = %g° F\n", fToC(cToF(FreezingC)), cToF(FreezingC))
}
