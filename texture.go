// Texture related operations

package moo

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
