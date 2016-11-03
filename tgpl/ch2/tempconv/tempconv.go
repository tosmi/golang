package tempconv

import "fmt"

type Celcius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC = -273.15
	FreezingC = 0
	BoilingC = 100
)

func (c Celcius) String() string { return fmt.Sprintf("%g^C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g^F", f) }
func (k Kelvin) String() string { return fmt.Sprintf("%g^K", k) }

