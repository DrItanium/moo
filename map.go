// map related things
package moo

import (
	"github.com/DrItanium/moo/cseries"
)

const (
	TicksPerSecond = 30
	TicksPerMinute = 60 * TicksPerSecond

	MapIndexBufferSize              = 8192
	MinimumSeparationFromWall       = WorldOne / 4
	MinimumSeparationFromProjectile = ((3 * WorldOne) / 4)

	TeleportingMidpoint = TicksPerSecond / 2
	TeleportingDuration = 2 * TeleportingMidpoint

	MaximumPolygonsPerMap  = cseries.Kilo
	MaximumSidesPerMap     = 4 * cseries.Kilo
	MaximumEndpointsPerMap = 8 * cseries.Kilo
	MaximumLinesPerMap     = 4 * cseries.Kilo
	MaximumLevelsPerMap    = 128

	LevelNameLength = 64 + 1
)
const (
	// damage types
	DamageExplosion = iota
	DamageElectricalStaff
	DamageProjectile
	DamageAbsorbed
	DamageFlame
	DamageHoundClaws
	DamageAlienProjectile
	DamageHulkSlap
	DamageCompilerBolt
	DamageFusionBolt
	DamageHunterBolt
	DamageFist
	DamageTeleporter
	DamageDefender
	DamageYetiClaws
	DamageYetiProjectile
	DamageCrushing
	DamageLava
	DamageSuffocation
	DamageGoo
	DamageEnergyDrain
	DamageOxygenDrain
	DamageHummerBolt
	DamageShotgunProjectile
)
const (
	// damage flags
	AlienDamage = 0x1

	MaximumSavedObjects = 384
)
const (
	// map object types
	SavedMonster     = iota // index is monster type
	SavedObject             // index is scenery type
	SavedItem               // index is item type
	SavedPlayer             // index is team bitfield
	SavedGoal               // index is goal number
	SavedSoundSource        // index is source type, facing is sound volume
)
const (
	MapObjectIsInvisible = 0x0001

	MaximumVerticesPerPolygon = 8

	MaximumObjectTypes = 64
)

type DamageDefinition struct {
	Type   int16
	Flags  int16
	Base   int16
	Random int16
	Scale  cseries.Fixed
}
type SideTextureDefinition struct {
	X0, Y0  WorldDistance
	Texture ShapeDescriptor
}

type ObjectLocation struct {
	P            WorldPoint3d
	PolygonIndex int16

	Yaw, Pitch Angle
	Flags      cseries.Word
}
type StaticData struct {
	EnviromentCode int16

	PhysicsModel     int16
	SongIndex        int16
	MissionFlags     int16
	EnvironmentFlags int16

	BallInPlay bool // true if there's a ball in play
	unused1    bool
	unused     [3]int16

	LevelName       string
	EntryPointFlags int32
}
type DynamicData struct {
	// ticks since the beginning of the game
	TickCount int32

	// the real seed is static in WORLD.C; must call set_random_seed()
	RandomSeed cseries.Word

	// this is stored in the dynamic_data so that it is valid across saves
	//GameInformation GameData

	PlayerCount         int16
	SpeakingPlayerIndex int16

	unused                                        int16
	PlatformCount                                 int16
	EndpointCount                                 int16
	LineCount                                     int16
	SideCount                                     int16
	PolygonCount                                  int16
	LightsourceCount                              int16
	MapIndexCount                                 int16
	AmbientSoundImageCount, RandomSoundImageCount int16

	//statistically unlikely to be valid

	ObjectCount     int16
	MonsterCount    int16
	ProjectileCount int16
	EffectCount     int16
	LightCount      int16

	DefaultAnnotationCount  int16
	PersonalAnnotationCount int16

	InitialObjectsCount int16

	GarbageObjectCount int16

	// used by MoveMonsters to decide who gets to generate paths, etc.
	LastMonsterIndexToGetTime, LastMonsterIndexToBuildPath int16

	// variables used by NewMonster to adjust for different difficulty levels
	NewMonsterManglerCookie, NewMonsterVanishingCookie int16

	// number of civilians killed by players; periodically decremented
	CiviliansKilledByPlayers int16

	// Used by the item placement stuff
	RandomMonstersLeft  [MaximumObjectTypes]int16
	CurrentMonsterCount [MaximumObjectTypes]int16
	RandomItemsLeft     [MaximumObjectTypes]int16
	CurrentItemCount    [MaximumObjectTypes]int16

	CurrentLevelNumber int16 // what level the user is currently exploring

	CurrentCivilianCausalties, CurrentCivilianCount int16
	TotalCivilianCausalties, TotalCivilianCount     int16

	GameBeacon      WorldPoint2d
	GamePlayerIndex int16
}

var StaticWorld *StaticData
var DynamicWorld *DynamicData

type ObjectData struct {
	Location WorldPoint3d
	Polygon  int16

	Facing Angle

	Shape ShapeDescriptor

	Sequence, Flags cseries.Word

	TransferMode, TransferPeriod int16

	TransferPhase   int16
	Permutation     int16
	NextObject      int16
	ParasiticObject int16

	SoundPitch Fixed
}

type EndpointData struct {
	Solid       bool
	Transparent bool
	Elevation   bool

	HighestAdjacentFloorHeight, LowestAdjacentCeilingHeight WorldDistance
	Vertex                                                  WorldPoint2d
	Transformed                                             WorldPoint2d

	SupportingPolygonIndex int16
}

type LineData struct {
	EndpointIndexes [2]int16

	Solid                                       bool
	Transparent                                 bool
	Landscape                                   bool
	Elevation                                   bool
	VariableElevation                           bool
	LineHasTransparentSide                      bool
	Length                                      WorldDistance
	HighestAdjacentFloor, LowestAdjacentCeiling WorldDistance

	ClockwisePolygonSideIndex, CounterClockwisePolygonSideIndex int16

	ClockwisePolygonOwner, CounterClockwisePolygonOwner int16
}

type SideExclusionZone struct {
	E0, E1, E2, E3 WorldPoint2d
}

type ControlPanelSideType int16

const (
	OxygenRefuel ControlPanelSideType = iota
	ShieldRefuel
	DoubleShieldRefuel
	TripleShieldRefuel
	LightSwitch
	PlatformSwitch
	TagSwitch
	PatternBuffer
	ComputerTerminal

	NumberOfControlPanels
)

type SideData struct {
	Type int16

	ControlPanelStatus        bool
	ControlPanel              bool
	RepairSwitch              bool
	DestructiveSwitch         bool
	LightedSwitch             bool
	CanBeDestroyed            bool
	CanOnlyBeHitByProjectiles bool

	EditorDirty bool // this is probably not relevant

	Primary     SideTextureDefinition
	Secondary   SideTextureDefinition
	Transparent SideTextureDefinition

	ExclusionZone SideExclusionZone

	ControlPanelType        ControlPanelSideType
	ControlPanelPermutation int16

	PrimaryTransferMode     int16
	SecondaryTransferMode   int16
	TransparentTransferMode int16

	PolygonIndex, LineIndex int16

	PrimaryLightsourceIndex     int16
	SecondaryLightsourceIndex   int16
	TransparentLightsourceIndex int16

	AmbientDelta Fixed
}

const (
	MaximumVerticesPerPolygon = 8
)

type PolygonType int16

const (
	NormalPolygon PolygonType = iota
	ItemImpassablePolygon
	MonsterImpassiblePolygon
	HillPolygon
	BasePolygon
	PlatformPolygon
	LightOnTriggerPolygon
	PlatformOnTriggerPolygon
	LightOffTriggerPolygon
	PlatformOffTriggerPolygon
	TeleporterPolygon
	ZoneBorderPolygon
	GoalPolygon
	VisibleMonsterTriggerPolygon
	InvisibleMonsterTriggerPolygon
	DualMonsterTriggerPolygon
	ItemTriggerPolygon
	MustBeExploredPolygon
	AutomaticExitPolygon
)
