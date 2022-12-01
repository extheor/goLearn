package unitconv

// 1英尺=12英寸；1英寸=2.54cm ；1英尺=12 英寸=0.3048 米 =30.48厘米。

// FToM converts a Foot to Meter.
func FToM(f Foot) Meter {
	return Meter(f * 0.3048)
}

// MToF converts Meter to Foot.
func MToF(m Meter) Foot {
	return Foot(m / 0.3048)
}

// RToK converts Pound to Kilogram.
func PToK(p Pound) Kilogram {
	return Kilogram(p * 0.45359237)
}

// KToP converts Kilogram to Pound.
func KToP(k Kilogram) Pound {
	return Pound(k / 0.45359237)
}
