package moo

import (
	"github.com/DrItanium/moo/cseries"
)

const (
	ProjectileBallDropped = 1000

	NormalWeaponDz = 20

	ChargingWeaponAmmoCount = 4 // this is the amount of ammo taht charging weapons use at one time..

	// weapon classes
	MeleeClass           = iota // normal weapon, no ammunition, both triggers do the same thing
	NormalClass                 // normal weapon, one ammunition type, both triggers do the same thing
	DualFunctionClass           // normal weapon, one ammunition type, trigger does something different
	TwofistedPistolClass        // two can be held at once (different triggers), same ammunition
	MultipurposeClass           // two weapons in one (assault rifle, grenade launcher), two different ammunition types with two separate triggers; secondary ammunition is discrete (i.e. it is never loaded explicitly but appears in the weapon)
)
const (
	// weapon flags
	NoFlags                            = 0x0
	WeaponIsAutomatic                  = 0x01
	WeaponDisappearsAfterUse           = 0x02
	WeaponPlaysInstantShellCasingSound = 0x04
	WeaponOverloads                    = 0x08
	WeaponHasRandomAmmoOnPickup        = 0x10
	PowerupIsTemporary                 = 0x20
	WeaponReloadsInOneHand             = 0x40
	WeaponFiresOutOfPhase              = 0x80
	WeaponFiresUnderMedia              = 0x100
	WeaponTriggersShareAmmo            = 0x200
	WeaponSecondaryHasAngularFlipping  = 0x400
)
const (
	WeaponInHandCollection = 1
	FistIdle               = iota
	FistPunching
	PistolIdle
	PistolFiring
	PistolReloading
	ShotgunIdle
	ShotgunFiring
	ShotgunReloading
	AssaultRifleIdle
	AssaultRifleFiring
	AssaultRifleReloading
	FusionIdle
	FusionFiring
	MissileLauncherIdle
	MissileLauncherFiring
	FlamethrowerIdle
	FlamethrowerTransit
	FlamethrowerFiring
	PistolShellCasing
	AlienWeaponIdle
	AlienWeaponFiring
	SmgIdle
	SmgFiring
	SmgReloading
	SmgShellCasing
)
const (
	// shell casing types
	ShellCasingAssaultRifle = iota
	ShellCasingPistol
	ShellCasingPistolLeft
	ShellCasingPistolRight
	ShellCasingSmg
	NumberOfShellCasingTypes
)

type ShellCasingDefinition struct {
	Collection int16
	Shape      int16

	X0, Y0   cseries.Fixed
	Vx0, Vy0 cseries.Fixed
	Dvx, Dvy cseries.Fixed
}

var ShellCasingDefinitions [NumberOfShellCasingTypes]ShellCasingDefinition

//TODO: add shell casing definitions

type TriggerDefinition struct {
	RoundsPerMagazine int16
	AmmunitionType    int16
	TicksPerRound     int16
	RecoveryTicks     int16
	ChargingTicks     int16
	RecoilMagnitude   WorldDistance
	FiringSound       int16
	ClickSound        int16
	ChargingSound     int16
	ShellCasingSound  int16
	ReloadingSound    int16
	ChargedSound      int16
	ProjectileType    int16
	ThetaError        int16
	Dx, Dz            int16
	ShellCasingType   int16
	BurstCount        int16
}

type WeaponDefinition struct {
	ItemType    int16
	PowerupType int16
	WeaponClass int16
	Flags       int16

	FiringLightIntensity      cseries.Fixed
	FiringIntensityDecayTicks int16

	// weapon will come up to FIXED_ONE when fired; idle_height±bob_amplitude should be in the range [0,FIXED_ONE]
	IdleHeight, BobAmplitude, KickHeight, ReloadHeight cseries.Fixed
	IdleWidth, HorizontalAmplitude                     cseries.Fixed

	// weapon has three basic animations: idle, firing and reloading.
	// sounds and frames are pulled from the shape collection.
	// for automatic weapons the firing animation loops until the trigger is released
	// or the gun is empty and the gun begins rising as soon as the trigger is depressed
	// and is not lowered until the firing animation stops.
	// for single shot weapons the animation loops once;
	// the weapon is raised and lowered as soon as the firing animation terminates
	Collection                             int16
	IdleShape, FiringShape, ReloadingShape int16
	Unused                                 int16
	ChargingShape, ChargedShape            int16

	// How long does it take to ready the weapon?
	// load_rounds_tick is the point which you actually load them.
	ReadyTicks, AwaitReloadTicks, LoadingTicks, FinishLoadingTicks, PowerupTicks int16

	WeaponsByTrigger [NumberOfTriggers]TriggerDefinition
}

var WeaponDefinitions []WeaponDefinition

//TODO: add the weapon definitions here
