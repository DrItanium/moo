// collection definition structures
package moo

import (
	"github.com/DrItanium/moo/cseries"
)

const (
	CollectionVersion     = 3
	NumberOfPrivateColors = 3
	// collection types
	UnusedCollection = iota
	WallCollection
	ObjectCollection
	InterfaceCollection
	SceneryCollection

	HighLevelShapeNameLength = 32
	// Low Level Shape Definition Flags
	XMirroredBit        = 0x8000
	YMirroredBit        = 0x4000
	KeypointObscuredBit = 0x2000

	// RgbColorValue flags
	SelfLuminescentColorFlag = 0x80
)

type CollectionDefinition struct {
	Version int16

	Type  int16
	Flags cseries.Word

	ColorCount                      int16
	ClutCount                       int16
	ColorTableOffset                int32
	HighLevelShapeCount             int16
	HighLevelShapeOffsetTableOffset int32
	LowLevelShapeCount              int16
	LowLevelShapeOffsetTableOffset  int32
	PixelsToWorld                   int16
	Size                            int32
	Unused                          [253]int16
}

type HighLevelShapeDefinition struct {
	Type                 int16
	Flags                cseries.Word
	Name                 string
	NumberOfViews        int16
	FramesPerView        int16
	TicksPerFrame        int16
	KeyFrame             int16
	TransferMode         int16
	TransferModePeriod   int16
	FirstFrameSound      int16
	KeyFrameSound        int16
	LastFramesound       int16
	PixelsToWorld        int16
	LoopFrame            int16
	Unused               [14]int16
	LowLevelShapeIndexes [1]int16
}

type LowLevelShapeDefinition struct {
	Flags                 cseries.Word
	MinimumLightIntensity cseries.Fixed
	BitmapIndex           int16
	OriginX               int16
	OriginY               int16
	KeyX                  int16
	KeyY                  int16
	WorldLeft             int16
	WorldRight            int16
	WorldTop              int16
	WorldBottom           int16
	WorldX0               int16
	WorldY0               int16
	Unused                [4]int16
}

type RgbColorValue struct {
	Flags byte
	Value byte
	Red   cseries.Word
	Green cseries.Word
	Blue  cseries.Word
}
