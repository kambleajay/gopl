// Package weightconv converts weights in Pounds and Kilograms
package weightconv

import "fmt"

type Pound float64
type Kilogram float64

func (p Pound) String() string { return fmt.Sprintf("%g pound", p) }

func (k Kilogram) String() string { return fmt.Sprintf("%g kilograms", k) }

func PToKg(p Pound) Kilogram { return Kilogram(p * 0.453592) }

func KgToP(k Kilogram) Pound { return Pound(k * 2.20462) }
