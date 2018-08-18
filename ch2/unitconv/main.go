package main

import (
	"fmt"
	"github.com/kambleajay/gopl/ch2/lenconv"
	"github.com/kambleajay/gopl/ch2/tempconv"
	"github.com/kambleajay/gopl/ch2/weightconv"
	"os"
	"strconv"
)

func printTemperatures(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("Temperatures: %s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
}

func printLengths(t float64) {
	f := lenconv.Foot(t)
	m := lenconv.Meter(t)
	fmt.Printf("Lengths: %s = %s, %s = %s\n", f, lenconv.FToM(f), m, lenconv.MToF(m))
}

func printWeights(t float64) {
	p := weightconv.Pound(t)
	k := weightconv.Kilogram(t)
	fmt.Printf("Weights: %s = %s, %s = %s", p, weightconv.PToKg(p), k, weightconv.KgToP(k))
}

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		printTemperatures(t)
		printLengths(t)
		printWeights(t)
	}
}
