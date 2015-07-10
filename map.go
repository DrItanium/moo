// map related things
package moo

import (
	"github.com/DrItanium/moo/cseries"
)

const (
	TicksPerSecond = 30
	TicksPerMinute = 60 * TicksPerSecond

	MapIndexBufferSize              = 8192
	MinimumSeparationFromWall       = WorldOne / 4
	MinimumSeparationFromProjectile = ((3 * WorldOne) / 4)

	TeleportingMidpoint = TicksPerSecond / 2
	TeleportingDuration = 2 * TeleportingMidpoint

	MaximumPolygonsPerMap  = cseries.Kilo
	MaximumSidesPerMap     = 4 * cseries.Kilo
	MaximumEndpointsPerMap = 8 * cseries.Kilo
	MaximumLinesPerMap     = 4 * cseries.Kilo
	MaximumLevelsPerMap    = 128

	LevelNameLength = 64 + 1

	MaximumVerticesPerPolygon = 8
)

type SideTextureDefinition struct {
	X0, Y0  WorldDistance
	Texture ShapeDescriptor
}

type ObjectLocation struct {
	P            WorldPoint3d
	PolygonIndex int16

	Yaw, Pitch Angle
	Flags      cseries.Word
}
