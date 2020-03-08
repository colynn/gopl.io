package lwconv

import "fmt"

// Feet 英尺
type Feet float64

// Metre 米
type Metre float64

// Pound ..
type Pound float64

// Kilo ..
type Kilo float64

func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }
func (m Metre) String() string { return fmt.Sprintf("%gm", m) }

func (p Pound) String() string { return fmt.Sprintf("%glb", p) }
func (k Kilo) String() string  { return fmt.Sprintf("%gkg", k) }

// FToM converts a feet to Metre.
func FToM(f Feet) Metre { return Metre(f / 39.370) }

// MToF converts a metre to feet.
func MToF(m Metre) Feet { return Feet(m * 39.370) }

// PToK converts a feet to Metre.
func PToK(p Pound) Kilo { return Kilo(p / 2.2) }

// KToP converts a metre to feet.
func KToP(k Kilo) Pound { return Pound(k * 2.2) }
