// code relating to ingame devices
package moo

import "github.com/DrItanium/moo/cseries"

const (
	OxygenRechargeFrequency = 0
	EnergyRechargeFrequency = 0

	MaximumActivationRange         = 3 * WorldOne
	MaximumPlatformActivationRange = 3 * WorldOne
	MaximumControlActivationRange  = WorldOne + WorldOneHalf

	ObjectRadius = 50

	MinimumResaveTicks = 2 * TicksPerSecond
)

const (
	TargetIsPlatform = iota
	TargetIsControlPanel
)

// control panel sounds
const (
	ActivatingSound = iota
	DeactivatingSound
	UnusableSound
	NumberOfControlPanelSounds // always last
)

type ControlPanelDefinition struct {
	PanelClass int16
	Flags      cseries.Word

	Collection                 int16
	ActiveShape, InactiveShape int16

	Sounds         [NumberOfControlPanelSounds]int16
	SoundFrequency cseries.Fixed

	Item int16
}

//TODO: migrate this to json or something
var ControlPanelDefinitions []ControlPanelDefinition

// TODO: methods
