
package tempconv

import (
	"fmt"
)

type Celsius	float64
type Fahrenheit	float64
type Kelvin		float64

const (
	AbsoluteZeroC	Celsius = -273.15
	FreezingC		Celsius = 0
	BoilingC		Celsius = 100

	AbsoluteZeroK	Kelvin = 0
	FreezingK		Kelvin = -Kelvin(AbsoluteZeroC * 1)
	BoilingK		Kelvin = Kelvin(BoilingC * 1) + FreezingK
)

func (c Celsius) String() string	{ return fmt.Sprintf("%gC", c) }
func (f Fahrenheit) String() string	{ return fmt.Sprintf("%gF", f) }
func (k Kelvin) String() string		{ return fmt.Sprintf("%gK", k) }

