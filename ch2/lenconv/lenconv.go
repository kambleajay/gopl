// Package lenconv converts the lengths in Feet and Meters
package lenconv

import "fmt"

type Foot float64
type Meter float64

func (f Foot) String() string { return fmt.Sprintf("%g foot", f) }

func (m Meter) String() string { return fmt.Sprintf("%g meter", m) }

func FToM(f Foot) Meter { return Meter(f * 0.3048) }

func MToF(m Meter) Foot { return Foot(m * 3.28084) }
