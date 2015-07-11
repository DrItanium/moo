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

func (n Word) Floor(floor Word) Word {
	if n < floor {
		return floor
	} else {
		return n
	}
}
func (n Word) Ceiling(ceiling Word) Word {
	if n > ceiling {
		return ceiling
	} else {
		return n
	}
}
func (n Word) Pin(floor, ceiling Word) Word {
	if n < floor {
		return floor
	} else {
		return n.Ceiling(ceiling)
	}
}

func Pin(n, floor, ceiling int64) int64 {
	if n < floor {
		return floor
	} else {
		return Ceiling(n, ceiling)
	}
}
func Floor(n, floor int64) int64 {
	if n < floor {
		return floor
	} else {
		return n
	}
}
func Ceiling(n, ceiling int64) int64 {
	if n > ceiling {
		return ceiling
	} else {
		return n
	}
}

func Signum(x int64) int64 {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	} else {
		return 0
	}
}

func Abs(x int64) int64 {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func Min(x, y int64) int64 {
	if x > y {
		return y
	} else {
		return x
	}
}

func Max(x, y int64) int64 {
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

type Handle interface{}

// constants defined in the cseries lib for emulation purposes
const (
	UnsignedLongMax = 4294967295
	LongMax         = 2147483647
	LongMin         = -2147483648
	LongBits        = 32

	UnsignedShortMax = 65535
	ShortMax         = 32767
	ShortMin         = -32768
	ShortBits        = 16

	UnsignedCharMax = 255
	CharMax         = 127
	CharMin         = -128
	CharBits        = 8
)

// Globals...sigh
var Temporary [256]byte

// error kinds
const (
	FatalError = iota
	InfoError
)

func AlertUser(alertType, resourceNumber, errorNumber, identifier int16) {

}

var Debug bool

func ToggleDebugStats() bool {
	return Debug
}

func InitializeDebugger(forceDebuggerOn bool) {

}
