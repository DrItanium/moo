// game related sounds
package moo

import "github.com/DrItanium/moo/cseries"

const (
	NumberOfSoundVolumeLevels = 8

	MaximumSoundVolumeBits = 8
	MaximumSoundVolume     = 1 << MaximumSoundVolumeBits
)
const (
	// sound sources
	Eightbit22kSource = iota
	Sixteenbit22kSource

	NumberOfSoundSources
)
const (
	//initialization flags
	StereoFlag           = 0x0001 // play sounds in stereo
	DynamicTrackingFlag  = 0x0002 // tracks sound sources during idleProc
	DopplerShiftFlag     = 0x0004 // Adjusts Sound Pitch During IdleProc
	AmbientSoundFlag     = 0x0008 // Plays And Tracks Ambient Sounds (Valid Iff DynamicTrackingFlag)
	SixteeenBitSoundFlag = 0x0010 // Loads 16bit Audio Instead Of 8bit
	MoreSoundsFlag       = 0x0020 // Loads All Permutations; Only Loads #0 If False
	ExtraMemoryFlag      = 0x0040 // Double Usual Memory

	// SoundObstructedProc() flags
	SoundWasObstructed      = 0x0001 // No Clear Path Between Source And Listener
	SoundWasMediaObstructed = 0x0002 // Source And Listener Are On Different Sides Of The Media
	SoundWasMediaMuffled    = 0x0004 // Source And Listener Both Under The Same Media

	// frequencies
	LowerFrequency  = cseries.FixedOne - cseries.FixedOne/8
	NormalFrequency = cseries.FixedOne
	HigherFrequency = cseries.FixedOne + cseries.FixedOne/8
)

// ---------- Sound Codes
const (
	// Ambient Sound Codes
	AmbientSndWater = iota
	AmbientSndSewage
	AmbientSndLava
	AmbientSndGoo
	AmbientSndUnderMedia
	AmbientSndWind
	AmbientSndWaterfall
	AmbientSndSiren
	AmbientSndFan
	AmbientSndSphtDoor
	AmbientSndSphtPlatform
	AmbientSndHeavySphtDoor
	AmbientSndHeavySphtPlatform
	AmbientSndLightMachinery
	AmbientSndHeavyMachinery
	AmbientSndTransformer
	AmbientSndSparkingTransformer
	AmbientSndMachineBinder
	AmbientSndMachineBookpress
	AmbientSndMachinePuncher
	AmbientSndElectric
	AmbientSndAlarm
	AmbientSndNightWind
	AmbientSndPfhorDoor
	AmbientSndPfhorPlatform
	AmbientSndAlienNoise1
	AmbientSndAlienNoise2
	AmbientSndJjaroNoise

	NumberOfAmbientSoundDefinitions
)
const (
	// Random Sound Codes
	RandomSndWaterDrip = iota
	RandomSndSurfaceExplosion
	RandomSndUndergroundExplosion
	RandomSndOwl
	RandomSndJjaroCreak

	NumberOfRandomSoundDefinitions
)
const (
	// Sound Codes
	SndStartup = iota
	SndTeleportIn
	SndTeleportOut
	SndBodyBeingCrunched
	SndJjaroCreak
	SndAbsorbed

	SndBreathing
	SndOxygenWarning
	SndSuffocation

	SndEnergyRefuel
	SndOxygenRefuel
	SndCantToggleSwitch
	SndSwitchOn
	SndSwitchOff
	SndPuzzleSwitch
	SndChipInsertion
	SndPatternBuffer
	SndDestroyControlPanel

	SndAdjustVolume
	SndGotPowerup
	SndGotItem

	SndBulletRicochet
	SndMetallicRicochet
	SndEmptyGun

	SndSphtDoorOpening
	SndSphtDoorClosing
	SndSphtDoorObstructed

	SndSphtPlatformStarting
	SndSphtPlatformStopping

	SndOwl
	SndSmgFiring
	SndSmgReloading

	SndHeavySphtPlatformStarting
	SndHeavySphtPlatformStopping

	SndFistHitting

	SndMagnumFiring
	SndMagnumReloading

	SndAssaultRifleFiring
	SndGrenadeLauncherFiring
	SndGrenadeExploding
	SndGrenadeFlyby

	SndFusionFiring
	SndFusionExploding
	SndFusionFlyby
	SndFusionCharging

	SndRocketExploding
	SndRocketFlyby
	SndRocketFiring

	SndFlamethrower

	SndBodyFalling
	SndBodyExploding
	SndBulletHittingFlesh

	SndFighterActivate
	SndFighterWail
	SndFighterScream
	SndFighterChatter
	SndFighterAttack
	SndFighterProjectileHit
	SndFighterProjectileFlyby

	SndCompilerAttack
	SndCompilerDeath
	SndCompilerHit
	SndCompilerProjectileFlyby
	SndCompilerProjectileHit

	SndCyborgMoving
	SndCyborgAttack
	SndCyborgHit
	SndCyborgDeath
	SndCyborgProjectileBounce
	SndCyborgProjectileHit
	SndCyborgProjectileFlyby

	SndHummerActivate
	SndHummerStartAttack
	SndHummerAttack
	SndHummerDying
	SndHummerDeath
	SndHummerProjectileHit
	SndHummerProjectileFlyby

	SndHumanWail
	SndHumanScream
	SndHumanHit
	SndHumanChatter
	SndAssimilatedHumanChatter
	SndHumanTrashTalk
	SndHumanApology
	SndHumanActivation
	SndHumanClear
	SndHumanStopShootingMeYouBastard
	SndHumanAreaSecure
	SndKillThePlayer

	SndWater
	SndSewage
	SndLava
	SndGoo
	SndUnderMedia
	SndWind
	SndWaterfall
	SndSiren
	SndFan
	SndSphtDoor
	SndSphtPlatform
	SndJjaroNoise
	SndHeavySphtPlatform
	SndLightMachinery
	SndHeavyMachinery
	SndTransformer
	SndSparkingTransformer

	SndWaterDrip

	SndWalkingInWater
	SndExitWater
	SndEnterWater
	SndSmallWaterSplash
	SndMediumWaterSplash
	SndLargeWaterSplash

	SndWalkingInLava
	SndEnterLava
	SndExitLava
	SndSmallLavaSplash
	SndMediumLavaSplash
	SndLargeLavaSplash

	SndWalkingInSewage
	SndExitSewage
	SndEnterSewage
	SndSmallSewageSplash
	SndMediumSewageSplash
	SndLargeSewageSplash

	SndWalkingInGoo
	SndExitGoo
	SndEnterGoo
	SndSmallGooSplash
	SndMediumGooSplash
	SndLargeGooSplash

	SndMajorFusionFiring
	SndMajorFusionCharged

	SndAssaultRifleReloading
	SndAssaultRifleShellCasings

	SndShotgunFiring
	SndShotgunReloading

	SndBallBounce
	SndYouAreIt
	SndGotBall

	SndComputerInterfaceLogon
	SndComputerInterfaceLogout
	SndComputerInterfacePage

	SndHeavySphtDoor
	SndHeavySphtDoorOpening
	SndHeavySphtDoorClosing
	SndHeavySphtDoorOpen
	SndHeavySphtDoorClosed
	SndHeavySphtDoorObstructed

	SndHunterActivate
	SndHunterAttack
	SndHunterDying
	SndHunterLanding
	SndHunterExploding
	SndHunterProjectileHit
	SndHunterProjectileFlyby

	SndEnforcerActivate
	SndEnforcerAttack
	SndEnforcerProjectileHit
	SndEnforcerProjectileFlyby

	SndYetiMeleeAttack
	SndYetiMeleeAttackHit
	SndYetiProjectileAttack
	SndYetiProjectileSewageAttackHit
	SndYetiProjectileSewageFlyby
	SndYetiProjectileLavaAttackHit
	SndYetiProjectileLavaFlyby
	SndYetiDying

	SndMachineBinder
	SndMachineBookpress
	SndMachinePuncher
	SndElectric
	SndAlarm
	SndNightWind

	SndSurfaceExplosion
	SndUndergroundExplosion

	SndDefenderAttack
	SndDefenderHit
	SndDefenderFlyby
	SndDefenderBeingHit
	SndDefenderExploding

	SndTickChatter
	SndTickFalling
	SndTickFlapping
	SndTickExploding

	SndCeilingLampExploding

	SndPfhorPlatformStarting
	SndPfhorPlatformStopping
	SndPfhorPlatform

	SndPfhorDoorOpening
	SndPfhorDoorClosing
	SndPfhorDoorObstructed
	SndPfhorDoor

	SndPfhorSwitchOff
	SndPfhorSwitchOn

	SndJuggernautFiring
	SndJuggernautWarning
	SndJuggernautExploding
	SndJuggernautPreparingToFire

	SndEnforcerExploding

	SndAlienNoise1
	SndAlienNoise2

	SndFusionHumanWail
	SndFusionHumanScream
	SndFusionHumanHit
	SndFusionHumanChatter
	SndAssimilatedFusionHumanChatter
	SndFusionHumanTrashTalk
	SndFusionHumanApology
	SndFusionHumanActivation
	SndFusionHumanClear
	SndFusionHumanStopShootingMeYouBastard
	SndFusionHumanAreaSecure
	SndFusionKillThePlayer

	NumberOfSoundDefinitions
)
