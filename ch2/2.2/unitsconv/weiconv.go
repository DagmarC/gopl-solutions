package unitsconv

import "fmt"

type Kilograms float64
type Pounds float64

const poundConst = 2.2046

func KgToP(kg Kilograms) Pounds {
	return Pounds(kg * poundConst)
}

func PToKg(p Pounds) Kilograms {
	return Kilograms(p / poundConst)
}

func (kg Kilograms) Converts() float64 {
	return float64(KgToP(kg))
}

func (p Pounds) Converts() float64 {
	return float64(PToKg(p))
}

func (kg Kilograms) String() string { return fmt.Sprintf("%gkg", kg) }
func (p Pounds) String() string     { return fmt.Sprintf("%glbs", p) }
