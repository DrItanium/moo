package moo

import "github.com/DrItanium/moo/cseries"

const (
	MaximumScreenWidth  = 640
	MaximumScreenHeight = 480
)

// TODO: finish up this file
const (
	ShadelessBit = 0x8000
)

type RectangleDefinition struct {
	Flags   cseries.Word
	Texture *BitmapDefinition
	// screen coordinates; x0<x1, y0<y1
	X0, Y0 int16
	X1, Y1 int16

	// screen coordinates
	ClipLeft, ClipRight int16
	ClipTop, ClipBottom int16

	Depth WorldDistance // depth at logical center (used to calculate light due to viewer)

	AmbientShade cseries.Fixed // Ambient shading table index; may objects will be self-luminescent, so this may have nothing to do with the polygon the object is sitting in

	ShadingTables interface{} // All of the shading tables, crammed together in memory

	TransferMode, TransferData int16 // _tinted, _textured, and _static are supported; _solid would be silly and landscape would be hard (but might be cool)

	FlipVertical, FlipHorizontal bool // mirrored horizontally and vertically if TRUE */
}
