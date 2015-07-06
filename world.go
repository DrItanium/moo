// world related operations
package moo

import (
	"github.com/DrItanium/moo/cseries"
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
