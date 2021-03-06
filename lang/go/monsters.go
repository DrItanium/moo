// monster related stuffs, a monster is another creature in the simulation

package moo

const (
	PassOneZoneBorder         = 0x0001
	PassedZoneBorder          = 0x0002
	ActivateInvisibleMonsters = 0x0004 // sound or teleport trigger
	ActivateDeafMonsters      = 0x0008 // i.e., trigger
	PassSolidLines            = 0x0010 // i.e. not a sound (trigger)
	UseActivationBiases       = 0x0020 // inactive monsters follow their editor instructions (trigger)
	ActivationCannotBeAvoided = 0x0040 // cannot be suppressed because of recent activation (trigger)

)

type ActivationBias int16

const (
	// activation biases are only used when the monster is activated by a trigger
	// activation biases (set in editor)
	ActivateOnPlayer ActivationBias = iota
	ActivateOnNearestHostile
	ActivateOnGoal
	ActivateRandomly
)
const MaximumMonstersPerMap = 220

type MonsterType int16

const (
	// player monsters are never active
	// monster types
	MonsterMarine MonsterType = iota
	MonsterTickEnergy
	MonsterTickOxygen
	MonsterTickKamakazi
	MonsterCompilerMinor
	MonsterCompilerMajor
	MonsterCompilerMinorInvisible
	MonsterCompilerMajorInvisible
	MonsterFighterMinor
	MonsterFighterMajor
	MonsterFighterMinorProjectile
	MonsterFighterMajorProjectile
	CivilianCrew
	CivilianScience
	CivilianSecurity
	CivilianAssimilated
	MonsterHummerMinor     // slow hummer
	MonsterHummerMajor     // fast hummer
	MonsterHummerBigMinor  // big hummer
	MonsterHummerBigMajor  // angry hummer
	MonsterHummerPossessed // hummer from durandal
	MonsterCyborgMinor
	MonsterCyborgMajor
	MonsterCyborgFlameMinor
	MonsterCyborgFlameMajor
	MonsterEnforcerMinor
	MonsterEnforcerMajor
	MonsterHunterMinor
	MonsterHunterMajor
	MonsterTrooperMinor
	MonsterTrooperMajor
	MonsterMotherOfAllCyborgs
	MonsterMotherOfAllHunters
	MonsterSewageYeti
	MonsterWaterYeti
	MonsterLavaYeti
	MonsterDefenderMinor
	MonsterDefenderMajor
	MonsterJuggernautMinor
	MonsterJuggernautMajor
	MonsterTinyFighter
	MonsterTinyBob
	MonsterTinyYeti
	VacuumCivilianCrew
	VacuumCivilianScience
	VacuumCivilianSecurity
	VacuumCivilianAssimilated
	NumberOfMonsterTypes
)

type MonsterAction int16

const (
	// monster actions
	MonsterIsStationary MonsterAction = iota
	MonsterIsWaitingToAttackAgain
	MonsterIsMoving
	MonsterIsAttackingClose // melee
	MonsterIsAttackingFar   // ranged
	MonsterIsBeingHit
	MonsterIsDyingHard
	MonsterIsDyingSoft
	MonsterIsDyingFlaming
	MonsterIsTeleporting // transparent
	MonsterIsTeleportingIn
	MonsterIsTeleportingOut
	NumberOfMonsterActions
)

type MonsterMode int16

const (

	// monster modes
	MonsterLocked MonsterMode = iota
	MonsterLosingLock
	MonsterLostLock
	MonsterUnlocked
	MonsterRunning
	NumberOfMonsterModes
)

type MonsterFlags int16

const (
	// monster flags
	MonsterWasPromoted                 MonsterFlags = 0x1
	MonsterWasDemoted                               = 0x2
	MonsterHasNeverBeenActivated                    = 0x4
	MonsterIsBlind                                  = 0x8
	MonsterIsDeaf                                   = 0x10
	MonsterTeleportsOutWhenDeactivated              = 0x20
)
