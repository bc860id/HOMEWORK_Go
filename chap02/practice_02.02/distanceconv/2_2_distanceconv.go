
package distanceconv

import (
	"fmt"
)

type Inch		float64
type Millimeter	float64

const (
	OneInchToMillimeters	Millimeter = 25.4
	OneMillimeterToInches	Inch = Inch(1 / OneInchToMillimeters)
)

func (inch Inch) String() string		{ return fmt.Sprintf("%ginch", inch) }
func (mm Millimeter) String() string	{ return fmt.Sprintf("%gmm", mm) }

func InchToMm(inch Inch) Millimeter		{ return Millimeter(inch) * OneInchToMillimeters }
func MmToInch(mm Millimeter) Inch		{ return Inch(mm) * OneMillimeterToInches }

