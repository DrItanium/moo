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

type EffectTypes cseries.Word

const ( /* effect types */
	EffectRocketExplosion EffectTypes = iota
	EffectRocketContrail
	EffectGrenadeExplosion
	EffectGrenadeContrail
	EffectBulletRicochet
	EffectAlienWeaponRicochet
	EffectFlamethrowerBurst
	EffectFighterBloodSplash
	EffectPlayerBloodSplash
	EffectCivilianBloodSplash
	EffectAssimilatedCivilianBloodSplash
	EffectEnforcerBloodSplash
	EffectCompilerBoltMinorDetonation
	EffectCompilerBoltMajorDetonation
	EffectCompilerBoltMajorContrail
	EffectFighterProjectileDetonation
	EffectFighterMeleeDetonation
	EffectHunterProjectileDetonation
	EffectHunterSpark
	EffectMinorFusionDetonation
	EffectMajorFusionDetonation
	EffectMajorFusionContrail
	EffectFistDetonation
	EffectMinorDefenderDetonation
	EffectMajorDefenderDetonation
	EffectDefenderSpark
	EffectTrooperBloodSplash
	EffectWaterLampBreaking
	EffectLavaLampBreaking
	EffectSewageLampBreaking
	EffectAlienLampBreaking
	EffectMetallicClang
	EffectTeleportObjectIn
	EffectTeleportObjectOut
	EffectSmallWaterSplash
	EffectMediumWaterSplash
	EffectLargeWaterSplash
	EffectLargeWaterEmergence
	EffectSmallLavaSplash
	EffectMediumLavaSplash
	EffectLargeLavaSplash
	EffectLargeLavaEmergence
	EffectSmallSewageSplash
	EffectMediumSewageSplash
	EffectLargeSewageSplash
	EffectLargeSewageEmergence
	EffectSmallGooSplash
	EffectMediumGooSplash
	EffectLargeGooSplash
	EffectLargeGooEmergence
	EffectMinorHummerProjectileDetonation
	EffectMajorHummerProjectileDetonation
	EffectDurandalHummerProjectileDetonation
	EffectHummerSpark
	EffectCyborgProjectileDetonation
	EffectCyborgBloodSplash
	EffectMinorFusionDispersal
	EffectMajorFusionDispersal
	EffectOverloadedFusionDispersal
	EffectSewageYetiBloodSplash
	EffectSewageYetiProjectileDetonation
	EffectWaterYetiBloodSplash
	EffectLavaYetiBloodSplash
	EffectLavaYetiProjectileDetonation
	EffectYetiMeleeDetonation
	EffectJuggernautSpark
	EffectJuggernautMissileContrail
	EffectSmallJjaroSplash
	EffectMediumJjaroSplash
	EffectLargeJjaroSplash
	EffectLargeJjaroEmergence
	EffectVacuumCivilianBloodSplash
	EffectAssimilatedVacuumCivilianBloodSplash
	NumberOfEffectTypes
)

type EffectData struct {
	Type        EffectTypes
	ObjectIndex int16

	Flags cseries.Word // slotUsed

	Data  int16 // used for special effects (effects)
	Delay int16 // the effect is invisible and inactive for this many ticks
}

var Effects []EffectData

func NewEffect(origin *WorldPoint3d, polygonIndex int16, t EffectTypes, facing Angle) int16 {
	return 0
}

func UpdateEffects() {

}

func RemoveAllNonPersistentEffects() {

}

func RemoveEffect(effectIndex int16) {

}

func MarkEffectCollections(t EffectTypes, loading bool) {

}

func TeleportObjectIn(objectIndex int16) {

}

func TeleportObjectOut(objectIndex int16) {

}
