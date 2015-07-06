// world related operations
package moo

import (
	"fmt"
	"github.com/DrItanium/moo/cseries"
	"math"
)

const (
	TrigShift          = 10
	TrigMagnitude      = 1 << TrigShift
	AngularBits        = 9
	NumberOfAngles     = 1 << AngularBits
	FullCircle         = NumberOfAngles
	QuarterCircle      = NumberOfAngles / 4
	HalfCircle         = NumberOfAngles / 2
	ThreeQuarterCircle = (NumberOfAngles * 3) / 4
	EighthCircle       = NumberOfAngles / 8
	SixteenthCircle    = NumberOfAngles / 16

	WorldFractionalBits               = 10
	WorldOne            WorldDistance = WorldDistance(1 << WorldFractionalBits)
	WorldOneHalf        WorldDistance = WorldDistance(WorldOne / 2)
	WorldOneFourth      WorldDistance = WorldDistance(WorldOne / 4)
	WorldThreeFourths   WorldDistance = WorldDistance((WorldOne * 3) / 4)

	DefaultRandomSeed cseries.Word = cseries.Word(0xfded)
)

type Angle int16
type WorldDistance int16

func IntegerToWorld(s int16) WorldDistance {
	return WorldDistance(s) << WorldFractionalBits
}
func (this WorldDistance) FractionalPart() WorldDistance {
	return this & WorldDistance(WorldOne-1)
}

func (this WorldDistance) IntegralPart() WorldDistance {
	return this >> WorldFractionalBits
}

func (this WorldDistance) ToFixed() cseries.Fixed {
	return cseries.Fixed(this) << (cseries.FixedFractionalBits - WorldFractionalBits)
}

func FixedToWorld(f cseries.Fixed) WorldDistance {
	return WorldDistance(f >> (cseries.FixedFractionalBits - WorldFractionalBits))
}

func NormalizeAngle(a Angle) Angle {
	if -360 < a && a < 720 {
		return a & Angle(NumberOfAngles-1)
	} else {
		theta := a
		for theta < 0 {
			theta += NumberOfAngles
		}
		for theta >= NumberOfAngles {
			theta -= NumberOfAngles
		}
		return theta
	}
}

func (this Angle) Facing4() Angle {
	return NormalizeAngle(this-EighthCircle) >> (AngularBits - 2)
}

func (this Angle) Facing5() Angle {
	return NormalizeAngle(this-FullCircle/10) / ((NumberOfAngles / 5) + 1)
}

func (this Angle) Facing8() Angle {
	return NormalizeAngle(this-SixteenthCircle) >> (AngularBits - 3)
}

func GuessHypotenuse(x, y int64) int64 {
	if x > y {
		return x + (y >> 1)
	} else {
		return y + (x >> 1)
	}
}

type WorldPoint2d struct {
	X WorldDistance
	Y WorldDistance
}

type FunctionError struct {
	Function string
	Message  string
}

func (this FunctionError) Error() string {
	return fmt.Sprintf("%s: %s", this.Function, this.Message)
}
func checkTheta(theta Angle) error {
	if !(theta >= 0 && theta < NumberOfAngles) {
		return &FunctionError{
			Function: "WorldPoint2d.Translate",
			Message:  "Theta is less than zero or theta >= Number of angles!",
		}
	} else {
		return nil
	}
}
func checkCosineTable() error {
	if CosineTable[0] != TrigMagnitude {
		return &FunctionError{
			Function: "WorldPoint2d.Translate",
			Message:  fmt.Sprintf("CosineTable[0] != TrigMagnitude, Actual Value is: %d", CosineTable[0]),
		}
	} else {
		return nil
	}
}
func (this *WorldPoint2d) Translate(distance WorldDistance, theta Angle) error {
	if err := checkTheta(theta); err != nil {
		return err
	} else if err := checkCosineTable(); err != nil {
		return err
	} else {
		this.X += (distance * WorldDistance(CosineTable[theta])) >> TrigShift
		this.Y += (distance * WorldDistance(SineTable[theta])) >> TrigShift
		return nil
	}
}
func (this *WorldPoint2d) Rotate(origin *WorldPoint2d, theta Angle) error {
	var temp WorldPoint2d
	if err := checkTheta(theta); err != nil {
		return err
	} else if err := checkCosineTable(); err != nil {
		return err
	} else {
		temp.X = this.X - origin.X
		temp.Y = this.Y - origin.Y
		this.X = ((temp.X * WorldDistance(CosineTable[theta])) >> TrigShift) + ((temp.Y * WorldDistance(SineTable[theta])) >> TrigShift) + origin.X
		this.Y = ((temp.Y * WorldDistance(CosineTable[theta])) >> TrigShift) - ((temp.X * WorldDistance(SineTable[theta])) >> TrigShift) + origin.Y
		return nil
	}
}

func (this *WorldPoint2d) Transform(origin *WorldPoint2d, theta Angle) error {
	var temp WorldPoint2d
	if err := checkTheta(theta); err != nil {
		return err
	} else if err := checkCosineTable(); err != nil {
		return err
	} else {
		temp.X = this.X - origin.X
		temp.Y = this.Y - origin.Y

		this.X = ((temp.X * WorldDistance(CosineTable[theta])) >> TrigShift) + ((temp.Y * WorldDistance(SineTable[theta])) >> TrigShift)
		this.Y = ((temp.Y * WorldDistance(CosineTable[theta])) >> TrigShift) - ((temp.X * WorldDistance(SineTable[theta])) >> TrigShift)
		return nil
	}
}

func GuessDistance(p0, p1 *WorldPoint2d) WorldDistance {
	dx := int64(p0.X) - int64(p1.X)
	dy := int64(p0.Y) - int64(p1.Y)
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}
	if distance := GuessHypotenuse(dx, dy); distance > 32767 {
		return WorldDistance(32767)
	} else {
		return WorldDistance(distance)
	}
}

type WorldPoint3d struct {
	X WorldDistance
	Y WorldDistance
	Z WorldDistance
}

type FixedPoint3d struct {
	X cseries.Fixed
	Y cseries.Fixed
	Z cseries.Fixed
}

type WorldVector2d struct {
	I WorldDistance
	J WorldDistance
}

type WorldVector3d struct {
	I WorldDistance
	J WorldDistance
	K WorldDistance
}

type FixedVector3d struct {
	I cseries.Fixed
	J cseries.Fixed
	K cseries.Fixed
}

type WorldLocation3d struct {
	Point        WorldPoint3d
	PolygonIndex int16
	Yaw          Angle
	Pitch        Angle
	Velocity     WorldVector3d
}

// definitions for a consine and sine table
var CosineTable []int16
var SineTable []int16
var TangentTable []int32

func BuildTrigTables() {
	SineTable = make([]int16, NumberOfAngles)
	CosineTable = make([]int16, NumberOfAngles)
	TangentTable = make([]int32, NumberOfAngles)
	twoPi := 8.0 * math.Atan(1.0)
	for i := 0; i < NumberOfAngles; i++ {
		theta := twoPi * (float64(i) / float64(NumberOfAngles))
		CosineTable[i] = int16(float64(TrigMagnitude)*math.Cos(theta) + 0.5)
		SineTable[i] = int16(float64(TrigMagnitude)*math.Sin(theta) + 0.5)

		switch i {
		case 0:
			SineTable[i] = 0
			CosineTable[i] = TrigMagnitude
		case QuarterCircle:
			SineTable[i] = TrigMagnitude
			CosineTable[i] = 0
		case HalfCircle:
			SineTable[i] = 0
			CosineTable[i] = -TrigMagnitude
		case ThreeQuarterCircle:
			SineTable[i] = -TrigMagnitude
			CosineTable[i] = 0
		}

		if CosineTable[i] != 0 {
			TangentTable[i] = (TrigMagnitude * int32(SineTable[i])) / int32(CosineTable[i])
		} else {
			TangentTable[i] = -2147483648
		}
	}
}

func Arctangent(x, y WorldDistance) Angle {
	if x != 0 {
		if tangent := (TrigMagnitude * y) / x; tangent != 0 {
			var theta Angle
			if y > 0 {
				theta = 1
			} else {
				theta = HalfCircle + 1
			}
			if theta < 0 {
				theta += QuarterCircle
			}

			lastDifference := int32(tangent) - TangentTable[theta-1]
			for searchArc := QuarterCircle - 1; searchArc != 0; searchArc-- {
				newDifference := int32(tangent) - TangentTable[theta]
				if (lastDifference <= 0 && newDifference >= 0) || (lastDifference >= 0 && newDifference <= 0) {
					if cseries.Abs(int(lastDifference)) < cseries.Abs(int(newDifference)) {
						return theta - 1
					} else {
						return theta
					}
				}
				lastDifference = newDifference
				theta++
			}
			if theta == NumberOfAngles {
				return 0
			} else {
				return theta
			}
		} else {
			if x < 0 {
				return HalfCircle
			} else {
				return 0
			}
		}
	} else {
		if y < 0 {
			return ThreeQuarterCircle
		} else {
			return QuarterCircle
		}
	}
}

// Taken from the original code, I leave it here to describe what is going on
/*
 * It requires more space to describe this implementation of the manual
 * square root algorithm than it did to code it.  The basic idea is that
 * the square root is computed one bit at a time from the high end.  Because
 * the original number is 32 bits (unsigned), the root cannot exceed 16 bits
 * in length, so we start with the 0x8000 bit.
 *
 * Let "x" be the value whose root we desire, "t" be the square root
 * that we desire, and "s" be a bitmask.  A simple way to compute
 * the root is to set "s" to 0x8000, and loop doing the following:
 *
 *      t = 0;
 *      s = 0x8000;
 *      do {
 *              if ((t + s) * (t + s) <= x)
 *                      t += s;
 *              s >>= 1;
 *      while (s != 0);
 *
 * The primary disadvantage to this approach is the multiplication.  To
 * eliminate this, we begin simplying.  First, we observe that
 *
 *      (t + s) * (t + s) == (t * t) + (2 * t * s) + (s * s)
 *
 * Therefore, if we redefine "x" to be the original argument minus the
 * current value of (t * t), we can determine if we should add "s" to
 * the root if
 *
 *      (2 * t * s) + (s * s) <= x
 *
 * If we define a new temporary "nr", we can express this as
 *
 *      t = 0;
 *      s = 0x8000;
 *      do {
 *              nr = (2 * t * s) + (s * s);
 *              if (nr <= x) {
 *                      x -= nr;
 *                      t += s;
 *              }
 *              s >>= 1;
 *      while (s != 0);
 *
 * We can improve the performance of this by noting that "s" is always a
 * power of two, so multiplication by "s" is just a shift.  Also, because
 * "s" changes in a predictable manner (shifted right after each iteration)
 * we can precompute (0x8000 * t) and (0x8000 * 0x8000) and then adjust
 * them by shifting after each step.  First, we let "m" hold the value
 * (s * s) and adjust it after each step by shifting right twice.  We
 * also introduce "r" to hold (2 * t * s) and adjust it after each step
 * by shifting right once.  When we update "t" we must also update "r",
 * and we do so by noting that (2 * (old_t + s) * s) is the same as
 * (2 * old_t * s) + (2 * s * s).  Noting that (s * s) is "m" and that
 * (r + 2 * m) == ((r + m) + m) == (nr + m):
 *
 *      t = 0;
 *      s = 0x8000;
 *      m = 0x40000000;
 *      r = 0;
 *      do {
 *              nr = r + m;
 *              if (nr <= x) {
 *                      x -= nr;
 *                      t += s;
 *                      r = nr + m;
 *              }
 *              s >>= 1;
 *              r >>= 1;
 *              m >>= 2;
 *      } while (s != 0);
 *
 * Finally, we note that, if we were using fractional arithmetic, after
 * 16 iterations "s" would be a binary 0.5, so the value of "r" when
 * the loop terminates is (2 * t * 0.5) or "t".  Because the values in
 * "t" and "r" are identical after the loop terminates, and because we
 * do not otherwise use "t"  explicitly within the loop, we can omit it.
 * When we do so, there is no need for "s" except to terminate the loop,
 * but we observe that "m" will become zero at the same time as "s",
 * so we can use it instead.
 *
 * The result we have at this point is the floor of the square root.  If
 * we want to round to the nearest integer, we need to consider whether
 * the remainder in "x" is greater than or equal to the difference
 * between ((r + 0.5) * (r + 0.5)) and (r * r).  Noting that the former
 * quantity is (r * r + r + 0.25), we want to check if the remainder is
 * greater than or equal to (r + 0.25).  Because we are dealing with
 * integers, we can't have equality, so we round up if "x" is strictly
 * greater than "r":
 *
 *      if (x > r)
 *              r++;
 */
func ISqrt(x uint32) int32 {
	r := uint32(0)
	m := uint32(0x40000000)
	// golang doesn't have do while so I'm emulating it
	for {
		nr := r + m
		if nr <= x {
			x -= nr
			r = nr + m
		}
		r >>= 1
		m >>= 2
		if m != 0 {
			break
		}
	}
	if x > r {
		r++ // was r += 1 the original code
	}
	return int32(r)
}

// make sure that the random function is as close as possible to the original
var randomSeed = cseries.Word(0x1)
var localRandomSeed = cseries.Word(0x1)

func GetRandomSeed() cseries.Word {
	return randomSeed
}

func SetRandomSeed(seed cseries.Word) {
	if seed != 0 {
		randomSeed = seed
	} else {
		randomSeed = DefaultRandomSeed
	}
}

func Random() cseries.Word {
	seed := randomSeed
	if (seed & 1) != 0 {
		seed = (seed >> 1) ^ 0xb400
	} else {
		seed >>= 1
	}

	randomSeed = seed
	return seed
}

func LocalRandom() cseries.Word {
	seed := localRandomSeed

	if (seed & 1) != 0 {
		seed = (seed >> 1) ^ 0xb400
	} else {
		seed >>= 1
	}

	localRandomSeed = seed
	return seed
}
