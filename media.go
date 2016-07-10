// media
package moo

import "github.com/DrItanium/moo/cseries"

type MediaType int16

const (
	// media types
	MediaWater MediaType = iota
	MediaLava
	MediaGoo
	MediaSewage
	MediaJjaro
	NumberOfMediaTypes
)

type MediaDetonationEffectType int16

const (
	// media detonation types
	SmallMediaDetonationEffect MediaDetonationEffectType = iota
	MediumMediaDetonationEffect
	LargeMediaDetonationEffect
	LargeMediaEmergenceEffect
	NumberOfMediaDetonationTypes
)

type MediaSoundType int16

const ( /* media sounds */
	MediaSoundFeetEntering MediaSoundType = iota
	MediaSoundFeetLeaving
	MediaSoundHeadEntering
	MediaSoundHeadLeaving
	MediaSoundSplashing
	MediaSoundAmbientOver
	MediaSoundAmbientUnder
	MediaSoundPlatformEntering
	MediaSoundPlatformLeaving
	NumberOfMediaSounds
)

type MediaData struct {
	Type                        MediaType
	MediaSoundObstructedByFloor bool

	/* this light is not used as a real light; instead, the intensity of this light is used to
	determine the height of the media: height= low + (high-low)*intensity ... this sounds
	gross, but it makes media heights as flexible as light intensities; clearly discontinuous
	light functions (e.g., strobes) should not be used */
	LightIndex int16

	/* this is the maximum external velocity due to current; acceleration is 1/32nd of this */
	CurrentDirection      Angle
	CurrentMagnitude      WorldDistance
	High, low             WorldDistance
	Origin                WorldPoint2d
	Height                WorldDistance
	MinimumLightIntensity cseries.Fixed
	Texture               ShapeDescriptor
	TransferMode          int16
}

var Medias []MediaData

func UpdateMedias() {

}

func GetMediaDetonationEffect(mediaIndex int16, dType MediaDetonationEffectType, detonationEffects *int16) {

}

func GetMediaSound(mediaIndex, t int16) int16 {
	return 0
}

func GetMediaSubmergedFadeEffect(mediaIndex int16) int16 {
	return 0
}

func GetMediaDamage(mediaIndex int16, scale cseries.Fixed) *DamageDefinition {
	return nil
}

func MediaInEnvironment(mediaType, environment int16) bool {
	return false
}

func (this *MediaData) UnderMedia(z WorldDistance) bool {
	return z <= this.Height
}

func GetMediaData(index int16) *MediaData {
	return nil
}

type MediaDefinition struct {
	Collection     int16
	Shape          int16
	ShapeCount     int16
	ShapeFrequency int16
	TransferMode   int16

	DamageFrequency int16 // mask&ticks
	Damage          DamageDefinition

	DetonationEffects [NumberOfMediaDetonationTypes]int16
	Sounds            [NumberOfMediaSounds]int16

	SubmergedFadeEffect int16
}
