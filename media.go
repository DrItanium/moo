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
	_media_snd_feet_entering MediaSoundType = iota
	_media_snd_feet_leaving
	_media_snd_head_entering
	_media_snd_head_leaving
	_media_snd_splashing
	_media_snd_ambient_over
	_media_snd_ambient_under
	_media_snd_platform_entering
	_media_snd_platform_leaving

	NUMBER_OF_MEDIA_SOUNDS
)

type mediaData struct {
	_type                       MediaType
	mediaSoundObstructedByFloor bool

	/* this light is not used as a real light; instead, the intensity of this light is used to
	determine the height of the media: height= low + (high-low)*intensity ... this sounds
	gross, but it makes media heights as flexible as light intensities; clearly discontinuous
	light functions (e.g., strobes) should not be used */
	lightIndex int16

	/* this is the maximum external velocity due to current; acceleration is 1/32nd of this */
	currentDirection      Angle
	currentMagnitude      WorldDistance
	high, low             WorldDistance
	origin                WorldPoint2d
	height                WorldDistance
	minimumLightIntensity cseries.Fixed
	texture               ShapeDescriptor
	transferMode          int16
}

var medias []mediaData

func updateMedias() {

}

func getMediaDetonationEffect(mediaIndex int16, dType MediaDetonationEffectType, detonationEffects *int16) {

}

func getMediaSound(mediaIndex, t int16) int16 {
	return 0
}

func getMediaSubmergedFadeEffect(mediaIndex int16) int16 {
	return 0
}

func getMediaDamage(mediaIndex int16, scale cseries.Fixed) *DamageDefinition {
	return nil
}

func mediaInEnvironment(mediaType, environment int16) bool {
	return false
}
