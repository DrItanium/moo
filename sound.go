// sound related functions
package moo

import "github.com/DrItanium/moo/cseries"

const MaximumPermutationsPerSound = 5

const (
	// sound behaviors
	SoundIsQuiet = iota
	SoundIsNormal
	SoundIsLoud
	NumberOfSoundBehaviorDefinitions
)

const (
	// flags
	SoundCannotBeRestarted       = 0x0001
	SoundDoesNotSelfAbort        = 0x0002
	SoundResistsPitchChanges     = 0x0004 // 0.5 external pitch changes
	SoundCannotChangePitch       = 0x0008 // no external pitch changes
	SoundCannotBeObstructed      = 0x0010 // ignore obstructions
	SoundCannotBeMediaObstructed = 0x0020 // ignore media obstructions
	SoundIsAmbient               = 0x0040 // will not be loaded unless Ambient_sound_flag is asserted
)
const (
	// sound chances
	TenPercent     = 32768 * 9 / 10
	TwentyPercent  = 32768 * 8 / 10
	ThirtyPercent  = 32768 * 7 / 10
	FourtyPercent  = 32768 * 6 / 10
	FiftyPercent   = 32768 * 5 / 10
	SixtyPercent   = 32768 * 4 / 10
	SeventyPercent = 32768 * 3 / 10
	EightyPercent  = 32768 * 2 / 10
	NintyPercent   = 32768 * 1 / 10
	Always         = 0

	SoundFileVersion = 1
	SoundFileTag     = "snd2"
)

type AmbientSoundDefinition struct {
	SoundIndex int16
}

type RandomSoundDefinition struct {
	SoundIndex int16
}

type SoundFileHeader struct {
	Version int32
	Tag     int32

	SourceCount int16 // usually 2 (8-bit, 16-bit)
	SoundCount  int16

	Unused [124]int16

	// immediately followed by PermutationCount * SoundCount SoundDefinition structures
}

type SoundDefinition struct { // 64 bytes
	SoundCode int16

	BehaviorIndex int16
	Flags         cseries.Word

	Chance cseries.Word // play sound if AbsRandom()) >= chance

	// if low_pitch==0, use FIXED_ONE; if high_pitch==0 use low pitch; else choose in [low_pitch,high_pitch]
	LowPitch, HighPitch cseries.Fixed

	// filled in later
	Permutations                           int16
	PermutationsPlayed                     cseries.Word
	GroupOffset, SingleLength, TotalLength int32                              // magic numbers necessary to load sounds
	SoundOffsets                           [MaximumPermutationsPerSound]int32 // zero-based from group offset

	LastPlayed uint32 // machine ticks

	Handle int32 // (machine-specific pointer type) zero if not loaded

	unused [2]int16
}

type DepthCurveDefinition struct {
	MaximumVolume, MaximumVolumeDistance int16
	MinimumVolume, MinimumVolumeDistance int16
}

type SoundBehaviorDefinition struct {
	ObstructedCurve, UnobstructedCurve DepthCurveDefinition
}
