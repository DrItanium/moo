package moo

import (
	"github.com/DrItanium/moo/cseries"
)

type EffectFlags cseries.Word

const (
	// flags
	EndWhenAnimationLoops         EffectFlags = 0x0001
	EndWhenTransferAnimationLoops EffectFlags = 0x0002
	SoundOnly                     EffectFlags = 0x0004 // play the animation's initial sound and noting else
	MakeTwinVisible               EffectFlags = 0x0008
	MediaEffect                   EffectFlags = 0x0010
)

type EffectDefinition struct {
	Collection, Shape int16
	SoundPitch        cseries.Fixed
	Flags             EffectFlags
	Delay, DelaySound int16
}

//TODO: the effect definitions themselves
