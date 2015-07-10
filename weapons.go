package moo

import (
	"github.com/DrItanium/moo/cseries"
)

const (
	WeaponFist = iota
	WeaponPistol
	WeaponPlasmaPistol
	WeaponAssaultRifle
	WeaponMissileLauncher
	WeaponFlamethrower
	WeaponAlienShotgun
	WeaponShotgun
	WeaponBall // or something similar
	WeaponSmg
	MaximumNumberOfWeapons

	WeaponDoubleFistedPistols = MaximumNumberOfWeapons // this is a pseudo-weapon
	WeaponDoubleFistedShotguns
	PlayerTorsoShapeCount

	ShapeWeaponIdle = iota
	ShapeWeaponCharging
	ShapeWeaponFiring

	PrimaryWeapon = iota
	SecondaryWeapon
	NumberOfTriggers

	// weapon display positioning modes
	PositionLow    = iota // position==0 is invisible, position==FIXED_ONE is sticking out from left/bottom
	PositionCenter        // position==0 is off left/bottom, position==FIXED_ONE is off top/right
	PositionHigh          // position==0 is invisible, position==FIXED_ONE is sticking out from right/top (mirrored, whether you like it or not)

	// weapon states

	WeaponIdle            = iota // if weapon_delay is non-zero, the weapon cannot be fired again yet
	WeaponRaising                // weapon is rising to idle position
	WeaponLowering               // weapon is lowering off the screen
	WeaponCharging               // Weapon is charging to fire..
	WeaponCharged                // Ready to fire..
	WeaponFiring                 // in firing animation
	WeaponRecovering             // Weapon is recovering from firing.
	WeaponAwaitingReload         // About to start reload sequence
	WeaponWaitingToLoad          // waiting to actually put bullets in
	WeaponFinishingReload        // finishing the reload

	WeaponLoweringForTwofistedReload    // lowering so the other weapon can reload
	WeaponAwaitingTwofistedReload       // waiting for other to lower..
	WeaponWaitingForTwofistToReload     // we are offscreen, waiting for the other to finish its load
	WeaponSlidingOverToSecondPosition   // pistol is going across when the weapon is present
	WeaponSlidingOverFromSecondPosition // Pistol returning to center of screen..
	WeaponWaitingForOtherIdleToReload   // Pistol awaiting friend's idle..
	NumberOfWeaponStates

	TriggerDown         = 0x0001
	PrimaryWeaponIsUp   = 0x0002
	SecondaryWeaponIsUp = 0x0004
	WantsTwofist        = 0x0008
	FlipState           = 0x0010

	WeaponType = iota
	ShellCasingType
	NumberOfDataTypes

	// For the flags - [11.unused 1.horizontal 1.vertical 3.unused]
	FlipShapeHorizontal = 0x08
	FlipShapeVertical   = 0x10

	PistolSeparationWidth                = cseries.FixedOne / 4
	AutomaticStillFiringDuration         = 4
	FiringBeforeShellCasingSoundIsPlayed = TicksPerSecond / 2
	CostPerChargedWeaponShot             = 4
	AngularVariance                      = 32

	MaximumShellCasings = 4

	// shell casing flags
	ShellCasingIsReversed = 0x0001
)

type WeaponDisplayInformation struct {
	Collection         int16
	LowLevelShapeIndex int16

	VerticalPosition          cseries.Fixed
	HorizontalPosition        cseries.Fixed
	VerticalPositioningMode   int16
	HorizontalPositioningMode int16
	TransferMode              int16
	TransferPhase             cseries.Fixed

	FlipHorizontal bool
	FlipVertical   bool
}

// Called once at startup
func InitializeWeaponManager() {

}

// Initialize weapons for a completely new game.
func InitializePlayerWeaponsForNewGame(playerIndex int16) {

}

// Initialize the given players weapons-> called after creating a player
func InitializePlayerWeapons(playerIndex int16) {

}

func GetWeaponArray() interface{} {
	return nil
}

func CalculateWeaponArrayLength() int32 {
	return 0
}

// While this returns true, keep calling...
func GetWeaponDisplayInformation(count *int16, data *WeaponDisplayInformation) bool {
	return true
}

// When the player runs over an item, check for reloads, etc.
func ProcessNewItemForReloading(playerIndex int16, itemType int16) {

}

// Update the given player's weapons
func UpdatePlayerWeapons(playerIndex int16, actionFlags int32) {

}

// Mark the weapon collections for loading or unloading..
func MarkWeaponCollections(loading bool) {

}

// Called when a player dies to discharge the weapons that they have charged up.
func DischargeChargedWeapons(playerIndex int16) {

}

// Called on entry to a level, and will change weapons if this one doesn't work in the given environment.
func CheckPlayerWeaponsForEnvironmentChange() {

}

// Tell me when one of my projectiles hits, and return the weapon_identifier I passed to new_projectile...
func PlayerHitTarget(playerIndex, weaponIdentifier int16) {

}

// for drawing the player
func GetPlayerWeaponModeAndType(playerIndex int16, shapeWeaponType, shapeMode *int16) {

}

// For the game window to update properly
func GetPlayerDesiredWeapon(playerIndex int16) int16 {
	return 0
}

// This is pinned to the maximum I think I can hold..
func GetPlayerWeaponAmmoCount(playerIndex, whichWeapon, whichTrigger int16) int16 {
	return 0
}

func DebugPrintWeaponStatus() {

}

type TriggerData struct {
	State              int16
	Phase              int16
	RoundsLoaded       int16
	ShotsFired         int16
	ShotsHit           int16
	TicksSinceLastShot int16        // used to play shell casing sound, and to calculate arc for shell casing drawing....
	TicksFiring        int16        // How long have we been firing? (only valid for automatics)
	Sequence           cseries.Word // What step of the animation are we in? (NOT guaranteed to be in sync!)
}

type WeaponData struct {
	WeaponType int16 // Stored here to make life easeir..
	Flags      cseries.Word
	Unused     cseries.Word // non zero-> weapon is powered up
	Triggers   [NumberOfTriggers]TriggerData
}

type ShellCasingData struct {
	Type  int16
	Frame int16

	Flags cseries.Word

	X  cseries.Fixed
	Y  cseries.Fixed
	Vx cseries.Fixed
	Vy cseries.Fixed
}

type PlayerWeaponData struct {
	CurrentWeapon int16
	DesiredWeapon int16
	//Weapons [NumberOfWeapons]WeaponData
}
