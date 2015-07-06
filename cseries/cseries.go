// Port of the cseries functions from marathon infinity
package cseries

const (
	None                  = -1
	Kilo                  = 1024
	Meg                   = Kilo * Kilo
	Gig                   = Kilo * Meg
	MachineTicksPerSecond = 60
)

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
