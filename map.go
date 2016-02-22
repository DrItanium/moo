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

	SoundPitch cseries.Fixed
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
	ControlPanelOxygenRefuel ControlPanelSideType = iota
	ControlPanelShieldRefuel
	ControlPanelDoubleShieldRefuel
	ControlPanelTripleShieldRefuel
	ControlPanelLightSwitch
	ControlPanelPlatformSwitch
	ControlPanelTagSwitch
	ControlPanelPatternBuffer
	ControlPanelComputerTerminal

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

	AmbientDelta cseries.Fixed
}

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

type HorizontalSurfaceData struct {
	Height                         WorldDistance
	LightSourceIndex               int16
	Texture                        ShapeDescriptor
	TransferMode, TransferModeData int16

	Origin WorldPoint2d
}

type PolygonData struct {
	Type        int16
	Flags       cseries.Word
	Permutation int16

	VertexCount     int16
	EndpointIndexes [MaximumVerticesPerPolygon]int16
	LineIndexes     [MaximumVerticesPerPolygon]int16

	FloorTexture, CeilingTexture                   ShapeDescriptor
	FloorHeight, CeilingHeight                     WorldDistance
	FloorLightsourceIndex, CeilingLightsourceIndex int16

	Area int32 // in WorldDistance^2 units

	FirstObject int16

	FirstExclusionZoneIndex int16
	LineExclusionZoneCount  int16
	PointExclusionZoneCount int16

	FloorTransferMode   int16
	CeilingTransferMode int16

	AdjacentPolygonIndexes [MaximumVerticesPerPolygon]int16

	FirstNeightborIndex int16
	NeighborCount       int16

	Center WorldPoint2d

	SideIndexes [MaximumVerticesPerPolygon]int16

	FloorOrigin, CeilingOrigin WorldPoint2d

	MediaIndex            int16
	MediaLightsourceIndex int16

	SoundSourceIndexes int16

	AmbientSoundImageIndex, RandomSoundImageIndex int16
}

type GameDifficultyLevel int16

const (
	DifficultyWuss GameDifficultyLevel = iota
	DifficultyEasy
	DifficultyNormal
	DifficultyMajorCarnage
	DifficultyTotalCarnage
	DifficultyCount
)

const MaximumObjectTypes = 64

const (
	// flags for object_frequency_definition
	ReappearsInRandomLocations = 0x0001
)

type ObjectFrequencyDefinition struct {
	Flags cseries.Word

	InitialCount int16 // number that initially appear. can be greater than maximum_count
	MinimumCount int16 // this number of objects will be maintained.
	MaximumCount int16 // can't exceed this, except at the beginning of the level.

	RandomCount  int16        // maximum random occurences of the object
	RandomChance cseries.Word // in (0, 65535]
}

type MissionFlagsDescription int16

const (
	// mission flags
	MissionNone          MissionFlagsDescription = 0x0000
	MissionExtermination                         = 0x0001
	MissionExploration                           = 0x0002
	MissionRetrieval                             = 0x0004
	MissionRepair                                = 0x0008
	MissionRescue                                = 0x0010
)

type EnvironmentFlagsDescription int16

const (
	/* environment flags */
	EnvironmentNormal     EnvironmentFlagsDescription = 0x0000
	EnvironmentVacuum                                 = 0x0001 // prevents certain weapons from working, player uses oxygen
	EnvironmentMagnetic                               = 0x0002 // motion sensor works poorly
	EnvironmentRebellion                              = 0x0004 // makes clients fight pfhor
	EnvironmentLowGravity                             = 0x0008 // low gravity

	EnvironmentNetwork      = 0x2000 // these two pseudo-environments are used to prevent items
	EnvironmentSinglePlayer = 0x4000 // from arriving in the items.c code.
)

type StaticData struct {
	EnvironmentCode int16

	PhysicsModel     int16
	SongIndex        int16
	MissionFlags     MissionFlagsDescription
	EnvironmentFlags EnvironmentFlagsDescription

	BallInPlay      bool
	LevelName       string
	EntryPointFlags int32
}
type GameOptions int16

const ( /* game options.. */
	MultiplayerGame             GameOptions = 0x0001 /* multi or single? */
	AmmoReplenishes                         = 0x0002 /* Does or doesn't */
	WeaponsReplenish                        = 0x0004 /* Weapons replenish? */
	SpecialsReplenish                       = 0x0008 /* Invisibility, Ammo? */
	MonstersReplenish                       = 0x0010 /* Monsters are lazarus.. */
	MotionSensorDoesNotWork                 = 0x0020 /* Motion sensor works */
	OverheadMapIsOmniscient                 = 0x0040 /* Only show teammates on overhead map */
	BurnItemsOnDeath                        = 0x0080 /* When you die, you lose everything but the initial crap.. */
	LiveNetworkStats                        = 0x0100
	GameHasKillLimit                        = 0x0200 /* Game ends when the kill limit is reached. */
	ForceUniqueTeams                        = 0x0400 /* every player must have a unique team */
	DyingIsPenalized                        = 0x0800 /* time penalty for dying */
	SuicideIsPenalized                      = 0x1000 /* time penalty for killing yourselves */
	OverheadMapShowsItems                   = 0x2000
	OverheadMapShowsMonsters                = 0x4000
	OverheadMapShowsProjectiles             = 0x8000
)

const (
	LevelUnfinished = iota
	LevelFinished
	LevelFailed
)

type GameType int16

// game types
const (
	GameOfKillMonsters    GameType = iota // single player & combative use this
	GameOfCooperativePlay                 // multiple players, working together
	GameOfCaptureTheFlag                  // A team game.
	GameOfKingOfTheHill
	GameOfKillManWithBall
	GameOfDefense
	GameOfRugby
	GameOfTag
	NumberOfGameTypes
)

//#define GET_GAME_TYPE() (dynamic_world->game_information.game_type)
//#define GET_GAME_OPTIONS() (dynamic_world->game_information.game_options)
//#define GET_GAME_PARAMETER(x) (dynamic_world->game_information.parameters[(x)])

type GameData struct {
	TimeRemaining     int32
	Type              GameType
	Options           GameOptions
	KillLimit         int16
	InitialRandomSeed int16
	Difficulty        GameDifficultyLevel
	// Parameters [2]int16 // use these later. for now memset to 0
}

type DynamicData struct {
	TickCount int32

	RandomSeed cseries.Word

	// this is stored in the DynamicData so that it is valid across saves
	GameInformation GameData

	PlayerCount         int16
	SpeakingPlayerIndex int16

	// Unused int16
	PlatformCount                                 int16
	EndpointCount                                 int16
	LineCount                                     int16
	SideCount                                     int16
	PolygonCount                                  int16
	LightsourceCount                              int16
	MapIndexCount                                 int16
	AmbientSoundImageCount, RandomSoundImageCount int16

	// statistically unlikely to be valid
	ObjectCount             int16
	MonsterCount            int16
	ProjectileCount         int16
	EffectCount             int16
	LightCount              int16
	DefaultAnnotationCount  int16
	PersonalAnnotationCount int16
	InitialObjectsCount     int16
	GarbageObjectCount      int16

	/* used by move_monsters() to decide who gets to generate paths, etc. */
	LastMonsterIndexToGetTime, LastMonsterIndexToBuildPath int16

	/* variables used by new_monster() to adjust for different difficulty levels */
	NewMonsterManglerCookie, NewMonsterVanishingCookie int16

	/* number of civilians killed by players; periodically decremented */
	CiviliansKilledByPlayers int16

	/* used by the item placement stuff */
	RandomMonstersLeft  [MaximumObjectTypes]int16
	CurrentMonsterCount [MaximumObjectTypes]int16
	RandomItemsLeft     [MaximumObjectTypes]int16
	CurrentItemCount    [MaximumObjectTypes]int16

	CurrentLevelNumber int16 // what level the user is currently exploring.

	CurrentCivilianCausalties, CurrentCivilianCount int16
	TotalCivilianCausalties, TotalCivilianCount     int16

	GameBeacon      WorldPoint2d
	GamePlayerIndex int16
}

var StaticWorld StaticData
var DynamicWorld DynamicData
var Objects []ObjectData

var MapPolygons []PolygonData
var MapSides []SideData
var MapLines []LineData
var MapEndpoints []EndpointData

var AmbientSoundImages []AmbientSoundImageData
var RandomSoundImages []RandomSoundImageData

var MapIndexes []short

var AutomapLines []byte
var AutomapPolygons []byte

var MapAnnotations []MapAnnotation
var SavedObjects []MapObject

var GameIsNetworked bool // true if this is a network game

//#define ADD_LINE_TO_AUTOMAP(i) (automap_lines[(i)>>3] |= (byte) 1<<((i)&0x07))
//#define LINE_IS_IN_AUTOMAP(i) ((automap_lines[(i)>>3]&((byte)1<<((i)&0x07)))?(TRUE):(FALSE))
//
//#define ADD_POLYGON_TO_AUTOMAP(i) (automap_polygons[(i)>>3] |= (byte) 1<<((i)&0x07))
//#define POLYGON_IS_IN_AUTOMAP(i) ((automap_polygons[(i)>>3]&((byte)1<<((i)&0x07)))?(TRUE):(FALSE))
