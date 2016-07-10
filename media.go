// media
package moo

import "github.com/DrItanium/moo/cseries"

type MediaType int16

const MaximumMediasPerMap = 16
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
	Type  MediaType
	Flags GenericFlags

	/* this light is not used as a real light; instead, the intensity of this light is used to
	determine the height of the media: height= low + (high-low)*intensity ... this sounds
	gross, but it makes media heights as flexible as light intensities; clearly discontinuous
	light functions (e.g., strobes) should not be used */
	LightIndex int16

	/* this is the maximum external velocity due to current; acceleration is 1/32nd of this */
	CurrentDirection      Angle
	CurrentMagnitude      WorldDistance
	High                  WorldDistance
	Low                   WorldDistance
	Origin                WorldPoint2d
	Height                WorldDistance
	MinimumLightIntensity cseries.Fixed
	Texture               ShapeDescriptor
	TransferMode          int16
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

var Medias [MaximumMediasPerMap]MediaData
var MediaDefinitions [NumberOfMediaTypes]MediaDefinition

type mediaIndex int16

func UpdateMedias() {
	for ind, mediaPos := mediaIndex(0), 0; ind < MaximumMediasPerMap; ind, mediaPos = ind+1, mediaPos+1 {
		media := &Medias[mediaPos]
		if media.Flags.SlotIsUsed() {
			ind.UpdateOneMedia(false)

			media.Origin.X = WorldDistance(media.Origin.X + ((WorldDistance(CosineTable[media.CurrentDirection]) * media.CurrentMagnitude) >> TrigShift)).FractionalPart()
			media.Origin.Y = WorldDistance(media.Origin.Y + ((WorldDistance(SineTable[media.CurrentDirection]) * media.CurrentMagnitude) >> TrigShift)).FractionalPart()
		}
	}
}
func (this mediaIndex) GetMediaDetonationEffect(dType MediaDetonationEffectType, detonationEffect *int16) {
	definition := GetMediaDefinition(this.GetMediaData().Type)

	if dType != MediaDetonationEffectType(cseries.None) {
		// assert (dType >= 0 && dType < NumberOfMediaDetonationTypes
		if definition.DetonationEffects[dType] != int16(cseries.None) {
			*detonationEffect = definition.DetonationEffects[dType]
		}
	}
}
func (this mediaIndex) GetMediaSound(t int16) int16 {
	// assert t >= 0 && t < NumberOfMediaSounds
	return GetMediaDefinition(this.GetMediaData().Type).Sounds[t]
}

func (this mediaIndex) GetMediaDamage(scale cseries.Fixed) *DamageDefinition {
	definition := GetMediaDefinition(this.GetMediaData().Type)
	damage := &definition.Damage

	damage.Scale = scale
	if damage.Type == cseries.None || ((DynamicWorld.TickCount & int32(definition.DamageFrequency)) != 0) {
		return nil
	} else {
		return damage
	}
}

func (this mediaIndex) GetMediaSubmergedFadeEffect() int16 {
	return GetMediaDefinition(this.GetMediaData().Type).SubmergedFadeEffect
}

func MediaInEnvironment(mediaType, environment int16) bool {
	//return CollectionInEnvironment(GetMediaDefinition(mediaType).Collection, environment)
	return true
}

func (this *MediaData) UnderMedia(z WorldDistance) bool {
	return z <= this.Height
}

func (index mediaIndex) GetMediaData() *MediaData {
	return &Medias[index]
}

func (this *MediaData) CalculateMediaHeight() int16 {
	return this.Low + cseries.Fixed((this.Height-this.Low)*GetLightIntensity(this.LightIndex)).IntegralPart()
}

func NewMedia(initializer *MediaData) mediaIndex {
	var ind mediaIndex
	var slotPos int16
	for ind, slotPos = 0, 0; ind < MaximumMediasPerMap; ind, slotPos = ind+1, slotPos+1 {

		if Medias[slotPos].Flags.SlotIsFree() {
			Medias[slotPos] = *initializer
			Medias[slotPos].Flags.MarkSlotAsUsed()

			Medias[slotPos].Origin.X = 0
			Medias[slotPos].Origin.Y = 0
			ind.UpdateOneMedia(true)
			break
		}
	}
	if ind == MaximumMediasPerMap {
		ind = mediaIndex(cseries.None)
	}
	return mediaIndex

}

func GetMediaDefinition(t MediaType) *MediaDefinition {
	return &MediaDefinitions[t]
}

func (this mediaIndex) UpdateOneMedia(forceUpdate bool) {
	media := this.GetMediaData()
	def := GetMediaDefinition(media.Type)

	// update height
	media.Height = media.Low + cseries.Fixed((media.High-media.Low)*GetLightIntensity(media.LightIndex)).IntegralPart()

	// update texture
	media.Texture = BuildDescriptor(def.Collection, def.Shape)
	media.TransferMode = def.TransferMode

}
