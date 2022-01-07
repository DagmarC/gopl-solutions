package unitsconv

import "fmt"

type Meters float64
type Feets float64

const FeetConst = 3.2808

func (m Meters) Converts() float64 {
	return float64(MToF(m))
}

func (ft Feets) Converts() float64 {
	return float64(FToM(ft))
}

func MToF(m Meters) Feets {
	return Feets(m * FeetConst)
}

func FToM(f Feets) Meters {
	return Meters(f / FeetConst)
}

func (m Meters) String() string { return fmt.Sprintf("%gm", m) }
func (f Feets) String() string  { return fmt.Sprintf("%gfeets", f) }
