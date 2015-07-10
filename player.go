// player related functions
package moo

import (
	"github.com/DrItanium/moo/cseries"
)

const (
	// We aren't in the demo so max players will always be 8
	MaximumNumberOfPlayers = 8

	NumberOfItems = 64

	// physics models
	EditorModel = iota
	EarthGravityModel
	LowGravityModel

	// Player actions; irrelevant if the player is dying or something

	PlayerStationary = iota
	PlayerWalking
	PlayerRunning
	PlayerSliding
	PlayerAirborne
	NumberOfPlayerActions

	// team colors
	VioletTeam = iota
	RedTeam
	TanTeam
	LightBlueTeam
	YellowTeam
	BrownTeam
	BlueTeam
	GreenTeam
	NumberOfTeamColors

	// action flags
	AbsoluteYawBits    = 7
	MaximumAbsoluteYaw = 1 << AbsoluteYawBits

	AbsolutePitchBits    = 5
	MaximumAbsolutePitch = 1 << AbsolutePitchBits

	AbsolutePositionBits    = 7
	MaximumAbsolutePosition = 1 << AbsolutePositionBits

	// action flag bit offsets

	AbsoluteYawModeBit = iota
	TurningLeftBit
	TurningRightBit
	SidestepDontTurnBit
	LookingLeftBit
	LookingRightBit
	AbsoluteYawBit0
	AbsoluteYawBit1

	AbsolutePitchModeBit
	LookingUpBit
	LookingDownBit
	LookingCenterBit
	AbsolutePitchBit0
	AbsolutePitchBit1

	AbsolutePositionModeBit
	MovingForwardBit
	MovingBackwardBit
	RunDontWalkBit
	LookDontTurnBit
	AbsolutePositionBit0
	AbsolutePositionBit1
	AbsolutePositionBit2

	SidesteppingLeftBit
	SidesteppingRightBit
	LeftTriggerStateBit
	RightTriggerStateBit
	ActionTriggerStateBit
	CycleWeaponsForwardBit
	CycleWeaponsBackwardBit
	ToggleMapBit
	MicrophoneButtonBit
	SwimBit

	NumberOfActionFlagBits // should be <=32

	// action flags
	AbsoluteYawMode  = 1 << AbsoluteYawModeBit
	TurningLeft      = 1 << TurningLeftBit
	TurningRight     = 1 << TurningRightBit
	SidestepDontTurn = 1 << SidestepDontTurnBit
	LookingLeft      = 1 << LookingLeftBit
	LookingRight     = 1 << LookingRightBit

	AbsolutePitchMode = 1 << AbsolutePitchModeBit
	LookingUp         = 1 << LookingUpBit
	LookingDown       = 1 << LookingDownBit
	LookingCenter     = 1 << LookingCenterBit
	LookDontTurn      = 1 << LookDontTurnBit

	AbsolutePositionMode = 1 << AbsolutePositionModeBit
	MovingForward        = 1 << MovingForwardBit
	MovingBackward       = 1 << MovingBackwardBit
	RunDontWalk          = 1 << RunDontWalkBit

	SidesteppingLeft     = 1 << SidesteppingLeftBit
	SidesteppingRight    = 1 << SidesteppingRightBit
	LeftTriggerState     = 1 << LeftTriggerStateBit
	RightTriggerState    = 1 << RightTriggerStateBit
	ActionTriggerState   = 1 << ActionTriggerStateBit
	CycleWeaponsForward  = 1 << CycleWeaponsForwardBit
	CycleWeaponsBackward = 1 << CycleWeaponsBackwardBit
	ToggleMap            = 1 << ToggleMapBit
	MicrophoneButton     = 1 << MicrophoneButtonBit
	Swim                 = 1 << SwimBit

	Turning           = TurningLeft | TurningRight
	Looking           = LookingLeft | LookingRight
	Moving            = MovingForward | MovingBackward
	Sidestepping      = SidesteppingLeft | SidesteppingRight
	LookingVertically = LookingUp | LookingDown | LookingCenter

	// player flag bits
	RecenteringBit    = 0x8000
	AboveGroundBit    = 0x4000
	BelowGroundBit    = 0x2000
	FeetBelowMediaBit = 0x1000
	HeadBelowMediaBit = 0x0800
	StepPeriodBit     = 0x0400

	// player flags
	PlayerIsInterlevelTeleportingFlag = 0x0100
	PlayerHasCheatedFlag              = 0x0200
	PlayerIsTeleportingFlag           = 0x0400
	PlayerHasMapOpenFlag              = 0x0800
	PlayerIsTotallyDeadFlag           = 0x1000
	PlayerIsZombieFlag                = 0x2000 // IS THIS USED??
	PlayerIsDeadFlag                  = 0x4000

	// constants from player.c
	ActionQueueBufferDiameter  = 0x100
	ActionQueueBufferIndexMask = 0xff
	KInvisibilityDuration      = 70 * TicksPerSecond
	KInvincibilityDuration     = 50 * TicksPerSecond
	KExtravisionDuration       = 3 * TicksPerMinute
	KInfravisionDuration       = 3 * TicksPerMinute

	MinimumReincarnationDelay = TicksPerSecond
	NormalReincarnationDelay  = 10 * TicksPerSecond
	SuicideReincarnationDelay = 15 * TicksPerSecond

	DeadPlayerHeight = WorldOneFourth

	OxygenWarningLevel     = TicksPerMinute
	OxygenWarningFrequency = TicksPerMinute / 4
	OxygenWarningOffset    = 10 * TicksPerSecond

	LastLevel = 100
)

type PhysicsVariables struct {
	HeadDirection                   cseries.Fixed
	LastDirection                   cseries.Fixed
	Direction                       cseries.Fixed
	Elevation                       cseries.Fixed
	AngularVelocity                 cseries.Fixed
	VerticalAngularVelocity         cseries.Fixed
	Velocity, PerpendicularVelocity cseries.Fixed // in and perpendicular to direction, respectively
	LastPosition, Position          FixedPoint3d
	ActualHeight                    cseries.Fixed

	// used by MaskInAbsolutePositioningInformation (because it is not really absolute) to keep track of where we're going
	AdjustedPitch, AdjustedYaw cseries.Fixed

	ExternalVelocity        FixedVector3d // from impacts; slowly absorbed
	ExternalAngularVelocity cseries.Fixed // from impacts; slowly absorbed

	StepPhase     cseries.Fixed // StepPhase is in [0,1) and is some function of the distance travelled (for bobbing the gun and the viewpoint)
	StepAmplitude cseries.Fixed // step amplitude is in [0,1) and is some function of velocity

	FloorHeight   cseries.Fixed // the height of the floor on the polygon where we ended up last time
	CeilingHeight cseries.Fixed // same as above, but ceiling height
	MediaHeight   cseries.Fixed // media height

	Action          int16        // what the player's legs are doing, basically
	OldFlags, Flags cseries.Word // stuff like recentering
}

type DamageRecord struct {
	Damage int32
	Kills  int16
}

type PlayerData struct {
	Identifier int16
	Flags      int16 // [unused.1] [dead.1] [zombie.1] [totally_dead.1] [map.1] [teleporting.1] [unused.10]

	Color int16
	Team  int16
	Name  string

	// Shadowed from physics_variables structure below and the player's object (read-only)
	Location                   WorldPoint3d
	CameraLocation             WorldPoint3d // beginning of fake world_location3d structure
	CameraPolygonIndex         int16
	Facing, Elevation          Angle
	SupportingPolygonIndex     int16 // what polygon is actually supporting our weight
	LastSupportingPolygonIndex int16

	// suit energy shadows vitality in the player's monster slot
	SuitEnergy, SuitOxygen int16

	MonsterIndex int16 // this player's entry in the monster list
	ObjectIndex  int16 // monster->object_index

	// Reset by initialize_player_weapons
	WeaponIntensityDecay int16 // zero is idle intensity
	WeaponIntensity      cseries.Fixed

	// powerups
	InvisibilityDuration  int16
	InvincibilityDuration int16
	InfravisionDuration   int16
	ExtravisionDuration   int16

	// teleporting
	DelayBeforeTeleport     int16 // this is only valid for interlevel teleports (teleporting_destination is a negative number)
	TeleportingPhase        int16 // NONE means no teleporting, other [0,TELEPORTING_PHASE) */
	TeleportingDestination  int16 // level number or NONE if intralevel transporter
	InterlevelTeleportPhase int16 // This is for the other players when someone else initiates the teleport

	// there is no state information associated with items; each slot is only a count
	Items [NumberOfItems]int16

	// used by the game window code to keep track of the interface state
	InterfaceFlags int16
	InterfaceDecay int16

	Variables PhysicsVariables

	TotalDamageGiven                       DamageRecord
	DamageTaken                            [MaximumNumberOfPlayers]DamageRecord
	MonsterDamageTaken, MonsterDamageGiven DamageRecord

	ReincarnationDelay int16

	ControlPanelSideIndex int16 // NONE, or the side index of a control panel the user is using that requires passage of time

	TicksAtLastSuccessfulSave int32

	NetgameParameters [2]int32

	Unused [256]int16
}

type ActionQueue struct {
	ReadIndex, WriteIndex int16

	Buffer chan int32
}

type PlayerShapeDefinitions struct {
	Collection int16

	DyingHard, DyingSoft int16
	DeadHard, DeadSoft   int16
	Legs                 [NumberOfPlayerActions]int16 // Stationary, walking, running, sliding, airborne
	Torsos               [PlayerTorsoShapeCount]int16 // NONE, ..., double pistols
	ChargingTorsos       [PlayerTorsoShapeCount]int16 // NONE, ..., double pistols
	FiringTorsos         [PlayerTorsoShapeCount]int16 // NONE, ..., double pistols
}

type DamageResponseDefinition struct {
	Type            int16
	DamageThreshold int16

	Fade                           int16
	Sound, DeathSound, DeathAction int16
}

var Players []PlayerData

var LocalPlayerIndex int16
var CurrentPlayerIndex int16
var LocalPlayer *PlayerData
var CurrentPlayer *PlayerData

var ActionQueues *ActionQueue

var playerShapes = PlayerShapeDefinitions{
	Collection:     6,
	DyingHard:      9,
	DyingSoft:      8,
	DeadHard:       11,
	DeadSoft:       10,
	Legs:           [NumberOfPlayerActions]int16{7, 0, 0, 24, 23},                             // legs: stationary, walking, runnning, sliding, airborne
	Torsos:         [PlayerTorsoShapeCount]int16{1, 3, 20, 26, 14, 12, 31, 16, 28, 33, 5, 18}, // idle torsos: fists, magnum, fusion, assault, rocket, flamethrower, alien, shotgun, double pistol, double shotgun, da ball
	ChargingTorsos: [PlayerTorsoShapeCount]int16{1, 3, 21, 26, 14, 12, 31, 16, 28, 33, 5, 18}, // charging torsos: fists, magnum, fusion, assault, rocket, flamethrower, alien, shotgun, double pistol, double shotgun, ball
	FiringTorsos:   [PlayerTorsoShapeCount]int16{2, 4, 22, 27, 15, 13, 32, 17, 28, 34, 6, 19}, // firing torsos: fists, magnum, fusion, assault, rocket, flamethrower, alien, shotgun, double pistol, double shotgun, ball
}

var playerInitialItems = []int16{
	ItemMagnum, // First weapon is the weapon he will use...
	ItemKnife,
	ItemKnife,
	ItemMagnumMagazine,
	ItemMagnumMagazine,
	ItemMagnumMagazine,
}

func NumberOfPlayerInitialItems() int {
	return len(playerInitialItems)
}

var DamageResponseDefinitions = []DamageResponseDefinition{
	DamageResponseDefinition{DamageExplosion, 100, FadeYellow, None, SndHumanScream, MonsterIsDyingHard},
	DamageResponseDefinition{DamageCrushing, None, FadeRed, None, SndHumanWail, MonsterIsDyingHard},
	DamageResponseDefinition{DamageProjectile, None, FadeRed, None, SndHumanScream, None},
	DamageResponseDefinition{DamageShotgunProjectile, None, FadeRed, None, SndHumanScream, None},
	DamageResponseDefinition{DamageElectricalStaff, None, FadeCyan, None, SndHumanScream, None},
	DamageResponseDefinition{DamageHulkSlap, None, FadeCyan, None, SndHumanScream, None},
	DamageResponseDefinition{DamageAbsorbed, 100, FadeWhite, SndAbsorbed, None, None},
	DamageResponseDefinition{DamageTeleporter, 100, FadeWhite, SndAbsorbed, None, None},
	DamageResponseDefinition{DamageFlame, None, FadeOrange, None, SndHumanWail, MonsterIsDyingFlaming},
	DamageResponseDefinition{DamageHoundClaws, None, FadeRed, None, SndHumanScream, None},
	DamageResponseDefinition{DamageCompilerBolt, None, FadeStatic, None, SndHumanScream, None},
	DamageResponseDefinition{DamageAlienProjectile, None, FadeDodgePurple, None, SndHumanWail, MonsterIsDyingFlaming},
	DamageResponseDefinition{DamageHunterBolt, None, FadeBurnGreen, None, SndHumanScream, None},
	DamageResponseDefinition{DamageFusionBolt, 60, FadeNegative, None, SndHumanScream, None},
	DamageResponseDefinition{DamageFist, 40, FadeRed, None, SndHumanScream, None},
	DamageResponseDefinition{DamageYetiClaws, None, FadeBurnCyan, None, SndHumanScream, None},
	DamageResponseDefinition{DamageYetiProjectile, None, FadeDodgeYellow, None, SndHumanScream, None},
	DamageResponseDefinition{DamageDefender, None, FadePurple, None, SndHumanScream, None},
	DamageResponseDefinition{DamageLava, None, FadeLongOrange, None, SndHumanWail, MonsterIsDyingFlaming},
	DamageResponseDefinition{DamageGoo, None, FadeLongGreen, None, SndHumanWail, MonsterIsDyingFlaming},
	DamageResponseDefinition{DamageSuffocation, None, None, None, SndSuffocation, MonsterIsDyingSoft},
	DamageResponseDefinition{DamageEnergyDrain, None, None, None, None, None},
	DamageResponseDefinition{DamageOxygenDrain, None, None, None, None, None},
	DamageResponseDefinition{DamageHummerBolt, None, FadeFlickerNegative, None, SndHumanScream, None},
}

func InitializePlayers() {

}
func ResetPlayerQueues() {

}

func AllocatePlayerMemory() {

}

func SetLocalPlayerIndex(playerIndex int16) {

}

func SetCurrentPlayerIndex(playerIndex int16) {

}

func NewPlayer(team, color, playerIdentifier int16) int16 {
	return 0
}

func DeletePlayer(playerNumber int16) {

}

func RecreatePlayersForNewLevel() {

}

func UpdatePlayers() {

}

/* Assumes �t==1 Tick */

func WalkPlayerList() {

}

func QueueActionFlags(playerIndex int16, actionFlags *int32, count int16) {

}

func DequeueActionFlags(playerIndex int16) int32 {

	return 0
}

func GetActionQueueSize(playerIndex int16) int16 {
	return 0

}

func DamagePlayer(monsterIndex, aggressorIndex, aggressorType int16, damage *DamageDefinition) {

}

func MarkPlayerCollections(loading bool) {

}

func PlayerIdentifierToPlayerIndex(playerId int16) int16 {
	return 0

}

func GetPlayerData(playerId int16) *PlayerData {
	return nil

}

//#Define GetPlayerData(I) (Players+(I))

func MonsterIndexToPlayerIndex(monsterIndex int16) int16 {
	return 0

}

func GetPolygonIndexSupportingPlayer(playerIndex int16) int16 {
	return 0

}

func LegalPlayerPowerup(itemIndex, playerIndex int16) bool {
	return false

}

func ProcessPlayerPowerup(playerIndex, itemIndex int16) {

}

func DeadPlayerMinimumPolygonHeight(polygonIndex int16) WorldDistance {
	return 0
}

func TryAndSubtractPlayerItem(playerIndex, itemType int16) bool {
	return false
}

func InitializePlayerPhysicsVariables(playerIndex int16) {

}

func UpdatePlayerPhysicsVariables(playerIndex int16, actionFlags int32) {

}

func AdjustPlayerForPolygonHeightChange(monsterIndex, polygonIndex int16, newFloorHeight, newCeilingHeight WorldDistance) {

}

func AcceleratePlayer(monsterIndex int16, verticalVelocity WorldDistance, direction Angle, velocity WorldDistance) {

}

func KillPlayerPhysicsVariables(playerIndex int16) {

}

func MaskInAbsolutePositioningInformation(actionFlags int32, yaw, pitch, velocity cseries.Fixed) int32 {
	return 0
}

func GetAbsolutePitchRange(min, max *cseries.Fixed) {

}

func InstantiateAbsolutePositioningInformation(playerIndex int16, facing, elevation cseries.Fixed) {

}

func GetBinocularVisionOrigins(playerIndex int16, left *WorldPoint3d, leftPolygonIndex *int16, leftAngle *Angle, right *WorldPoint3d, rightPolygonIndex *int16, rightAngle *Angle) {

}

func GetPlayerForwardVelocityScale(PlayerIndex int16) cseries.Fixed {
	return cseries.FixedOne

}
