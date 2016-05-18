package weightconv

import "fmt"

type Kilogram float64
type Pounds float64

func (k Kilogram) String() string { return fmt.Sprintf("%g kilogram", k)}
func (p Pounds) String() string { return fmt.Sprintf("%g pounds", p)}
