// Package unitconv performs Foot and Rice or Pound and Kilogram conversions.
package unitconv

import "fmt"

type Foot float64     // 英尺
type Meter float64    // 米
type Pound float64    // 磅
type Kilogram float64 // 公斤

func (f Foot) String() string {
	return fmt.Sprintf("%g英尺", f)
}

func (m Meter) String() string {
	return fmt.Sprintf("%g米", m)
}

func (p Pound) String() string {
	return fmt.Sprintf("%g磅", p)
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%g公斤", k)
}
