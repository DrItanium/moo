// Texture related operations

package moo

import (
	"github.com/DrItanium/moo/cseries"
)

type Pixel8 byte
type Pixel16 uint16
type Pixel32 uint32

type Pixel interface {
	Red() byte
	Green() byte
	Blue() byte
}

const (
	Pixel8MaxColors  = 256
	Pixel16MaxColors = 32768
	Pixel32MaxColors = 16777216

	Pixel16Bits = 5
	Pixel32Bits = 8

	NumberOfColorComponents = 3
	Pixel16MaximumComponent = 0x1f
	Pixel32MaximumComponent = 0xff
)

func (this Pixel16) Red() byte {
	return byte(this >> 10)
}
func (this Pixel16) Green() byte {
	return byte((this >> 5)) & Pixel16MaximumComponent

}

func (this Pixel16) Blue() byte {
	return byte(this) & Pixel16MaximumComponent
}

func NewPixel16(r, g, b byte) Pixel16 {
	return (Pixel16(r) << 10) | (Pixel16(g) << 5) | Pixel16(b)
}

//TODO: this function
//#define RGBCOLOR_TO_PIXEL16(r,g,b) (((pixel16)((r)>>1)&(pixel16)0x7c00)|((pixel16)((g)>>6)&(pixel16)0x03e0)|((pixel16)((b)>>11)&(pixel16)0x1f))
//func RgbColorToPixel16(r, g, b byte) Pixel16 {
// ((pixel16)((r)>>1)&(pixel16)0x7c00)
// ((pixel16)((g)>>6)&(pixel16)0x03e0)
// ((pixel16)((b)>>11)&(pixel16)0x1f)
//	rComponent := Pixel16((Pixel16(r)>>1) & Pixel16(0x7c00))
//	gComponent := Pixel16((Pixel16(g)>>6) & Pixel16(0x03e0))
//	bComponent := Pixel16((Pixel16(b)>>11) & Pixel
//}

func (this Pixel32) Red() byte {
	return byte(this >> 16)
}

func (this Pixel32) Green() byte {
	return byte(this>>8) & Pixel32MaximumComponent
}

func (this Pixel32) Blue() byte {
	return byte(this) & Pixel32MaximumComponent
}

func NewPixel32(r, g, b byte) Pixel32 {
	return (Pixel32(r) << 16) | (Pixel32(g) << 8) | Pixel32(b)
}

type RgbColor struct {
	Red   cseries.Word
	Green cseries.Word
	Blue  cseries.Word
}
type ColorTable []RgbColor // use make to simulate a capacity and length

// bitmap flags
const (
	ColumnOrderBit = 0x8000
	TransparentBit = 0x4000
)

type BitmapDefinition struct {
	Width, Height int16
	BytesPerRow   int16
	Flags         uint16
	BitDepth      int16
	// unused field
	// unused [8]int16
	RowAddresses [1]*Pixel8
}

// this function is wierd but before I explain it here is the c-code
//pixel8 *calculate_bitmap_origin(
//	struct bitmap_definition *bitmap)
//{
//	pixel8 *origin;
//
//	origin= (pixel8 *) (((byte *)bitmap) + sizeof(struct bitmap_definition));
//	if (bitmap->flags&_COLUMN_ORDER_BIT)
//	{
//		origin+= bitmap->width*sizeof(pixel8 *);
//	}
//	else
//	{
//		origin+= bitmap->height*sizeof(pixel8 *);
//	}
//
//	return origin;
//}
// So it looks like the objective is to compute the start of the given bitmap.
// Initially the origin is address of bitmap.RowAddresses in memory.
// If the ColumnOrderBit is set for bitmap then add the width of the bitmap times the sizeof a pixel8 pointer to origin
// If the ColumnOrderBit is not set for the bitmap then add the height of the bitmap times the sizeof a pixel8 pointer to origin
// return origin
//
// So following the bitmap structure itself the bitmap associated with the given header is found (encoding yay!)
// The origin is defined as the pixel furthest away from the start of the bitmap data
// There is also a comment found in textures.h attached to calculate_bitmap_origin which says:
// /* assumes pixel data follows bitmap_definition structure immediately */
// This proves my assumption of how this is laid out
//
//func (this *BitmapDefinition) CalculateBitmapOrigin() *Pixel8 {
//	var origin *Pixel8
//	origin =
//}
func (this *BitmapDefinition) ColumnOrder() bool {
	return (this.Flags & ColumnOrderBit) != 0
}
func (this *BitmapDefinition) Transparent() bool {
	return (this.Flags & TransparentBit) != 0
}

var Marathon1 = marathon1_game{}
var Marathon2 = marathon2_game{}

type marathon1_game struct{}
type marathon2_game struct{}

func (marathon1_game) IsMarathon() bool {
	return true
}

func (marathon2_game) IsMarathon() bool {
	return false
}

type Game interface {
	IsMarathon1() bool
}

// comment from original code: "must initialize bytes_per_row, height, and row_address[0]"
func (this *BitmapDefinition) PrecalculateBitmapRowAddresses(game Game) {
	var rows int16
	if this.ColumnOrder() {
		rows = this.Width
	} else {
		rows = this.Height
	}
	//rowAddress := this.RowAddresses[0]
	//table := this.RowAddresses
	bytesPerRow := this.BytesPerRow
	if bytesPerRow != cseries.None {
		for row := int16(0); row < rows; row++ {
			//*table++ = row_address; // WHY??? This is ambiguous
			// should be
			// table++;
			// *table = row_address;
		}
	} else {
		if game.IsMarathon1() {
			// put code here from textures.c
		} else {
			// put code here from textures.c
		}
	}
}

/*
void map_bytes(byte *buffer, byte *table, long size);
void remap_bitmap(struct bitmap_definition *bitmap,	pixel8 *table);

void erase_bitmap(struct bitmap_definition *bitmap, long pel);
*/

func MapBytes(buffer, table *byte, size int32) {

}

func (this *BitmapDefinition) RemapBitmap(table *Pixel8) {

}

func (this *BitmapDefinition) EraseBitmap(pel int32) {

}
