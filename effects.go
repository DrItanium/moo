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
	effect_rocket_explosion EffectTypes = iota
	effect_rocket_contrail
	effect_grenade_explosion
	effect_grenade_contrail
	effect_bullet_ricochet
	effect_alien_weapon_ricochet
	effect_flamethrower_burst
	effect_fighter_blood_splash
	effect_player_blood_splash
	effect_civilian_blood_splash
	effect_assimilated_civilian_blood_splash
	effect_enforcer_blood_splash
	effect_compiler_bolt_minor_detonation
	effect_compiler_bolt_major_detonation
	effect_compiler_bolt_major_contrail
	effect_fighter_projectile_detonation
	effect_fighter_melee_detonation
	effect_hunter_projectile_detonation
	effect_hunter_spark
	effect_minor_fusion_detonation
	effect_major_fusion_detonation
	effect_major_fusion_contrail
	effect_fist_detonation
	effect_minor_defender_detonation
	effect_major_defender_detonation
	effect_defender_spark
	effect_trooper_blood_splash
	effect_water_lamp_breaking
	effect_lava_lamp_breaking
	effect_sewage_lamp_breaking
	effect_alien_lamp_breaking
	effect_metallic_clang
	effect_teleport_object_in
	effect_teleport_object_out
	effect_small_water_splash
	effect_medium_water_splash
	effect_large_water_splash
	effect_large_water_emergence
	effect_small_lava_splash
	effect_medium_lava_splash
	effect_large_lava_splash
	effect_large_lava_emergence
	effect_small_sewage_splash
	effect_medium_sewage_splash
	effect_large_sewage_splash
	effect_large_sewage_emergence
	effect_small_goo_splash
	effect_medium_goo_splash
	effect_large_goo_splash
	effect_large_goo_emergence
	effect_minor_hummer_projectile_detonation
	effect_major_hummer_projectile_detonation
	effect_durandal_hummer_projectile_detonation
	effect_hummer_spark
	effect_cyborg_projectile_detonation
	effect_cyborg_blood_splash
	effect_minor_fusion_dispersal
	effect_major_fusion_dispersal
	effect_overloaded_fusion_dispersal
	effect_sewage_yeti_blood_splash
	effect_sewage_yeti_projectile_detonation
	effect_water_yeti_blood_splash
	effect_lava_yeti_blood_splash
	effect_lava_yeti_projectile_detonation
	effect_yeti_melee_detonation
	effect_juggernaut_spark
	effect_juggernaut_missile_contrail
	effect_small_jjaro_splash
	effect_medium_jjaro_splash
	effect_large_jjaro_splash
	effect_large_jjaro_emergence
	effect_vacuum_civilian_blood_splash
	effect_assimilated_vacuum_civilian_blood_splash
	numberOfEffectTypes
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
