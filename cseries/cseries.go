// Port of the cseries functions from marathon infinity
package cseries

const (
	None                  = -1
	Kilo                  = 1024
	Meg                   = Kilo * Kilo
	Gig                   = Kilo * Meg
	MachineTicksPerSecond = 60
)

type Word uint16

func Signum(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	} else {
		return 0
	}
}

func Abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func Min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func Max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

type Fixed int32

const (
	FixedFractionalBits       = 16
	FixedOne            Fixed = Fixed(1 << FixedFractionalBits)
	FixedOneHalf        Fixed = Fixed(1 << (FixedFractionalBits - 1))
)

func (this Fixed) Float() float64 {
	return (float64(this) / float64(FixedOne))
}

func FloatToFixed(f float64) Fixed {
	return Fixed(f) * FixedOne
}

func IntegerToFixed(s int) Fixed {
	return Fixed(s) << FixedFractionalBits
}

func (this Fixed) IntegralPart() int16 {
	return int16(this >> FixedFractionalBits)
}

func (this Fixed) FractionalPart() int16 {
	return int16(Fixed(int16(this)) & (FixedOne - 1))
}
