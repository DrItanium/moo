package moo

import "github.com/DrItanium/moo/cseries"

const maximum_projectiles_per_map = 32

const ( /* projectile types */
	_projectile_rocket = iota
	_projectile_grenade
	_projectile_pistol_bullet
	_projectile_rifle_bullet
	_projectile_shotgun_bullet
	_projectile_staff
	_projectile_staff_bolt
	_projectile_flamethrower_burst
	_projectile_compiler_bolt_minor
	_projectile_compiler_bolt_major
	_projectile_alien_weapon
	_projectile_fusion_bolt_minor
	_projectile_fusion_bolt_major
	_projectile_hunter
	_projectile_fist
	_projectile_unused
	_projectile_armageddon_electricity
	_projectile_juggernaut_rocket
	_projectile_trooper_bullet
	_projectile_trooper_grenade
	_projectile_minor_defender
	_projectile_major_defender
	_projectile_juggernaut_missile
	_projectile_minor_energy_drain
	_projectile_major_energy_drain
	_projectile_oxygen_drain
	_projectile_minor_hummer
	_projectile_major_hummer
	_projectile_durandal_hummer
	_projectile_minor_cyborg_ball
	_projectile_major_cyborg_ball
	_projectile_ball
	_projectile_minor_fusion_dispersal
	_projectile_major_fusion_dispersal
	_projectile_overloaded_fusion_dispersal
	_projectile_yeti
	_projectile_sewage_yeti
	_projectile_lava_yeti
	_projectile_smg_bullet
	NUMBER_OF_PROJECTILE_TYPES
)

//#define PROJECTILE_HAS_MADE_A_FLYBY(p) ((p)->flags&(word)0x4000)
//#define SET_PROJECTILE_FLYBY_STATUS(p,v) ((v)?((p)->flags|=(word)0x4000):((p)->flags&=(word)~0x4000))
//
///* only used for persistent projectiles */
//#define PROJECTILE_HAS_CAUSED_DAMAGE(p) ((p)->flags&(word)0x2000)
//#define SET_PROJECTILE_DAMAGE_STATUS(p,v) ((v)?((p)->flags|=(word)0x2000):((p)->flags&=(word)~0x2000))
//
//#define PROJECTILE_HAS_CROSSED_MEDIA_BOUNDARY(p) ((p)->flags&(word)0x1000)
//#define SET_PROJECTILE_CROSSED_MEDIA_BOUNDARY_STATUS(p,v) ((v)?((p)->flags|=(word)0x1000):((p)->flags&=(word)~0x1000))

/* uses SLOT_IS_USED(), SLOT_IS_FREE(), MARK_SLOT_AS_FREE(), MARK_SLOT_AS_USED() macros (0x8000 bit) */

type projectile_data struct {
	typ int16

	object_index int16

	target_index int16 /* for guided projectiles, the current target index */

	elevation Angle /* facing is stored in the projectileÕs object */

	owner_index int16        /* ownerless if NONE */
	owner_type  int16        /* identical to the monster type which fired this projectile (valid even if owner==NONE) */
	flags       cseries.Word /* [slot_used.1] [played_flyby_sound.1] [has_caused_damage.1] [unused.13] */

	/* some projectiles leave n contrails effects every m ticks */
	ticks_since_last_contrail, contrail_count int16

	distance_travelled WorldDistance

	gravity WorldDistance /* velocity due to gravity for projectiles affected by it */

	damage_scale cseries.Fixed

	permutation int16 /* item type if we create one */

}

const (
	/* projectile flags */
	_guided                           = 0x0001
	_stop_when_animation_loops        = 0x0002
	_persistent                       = 0x0004 /* does stops doing damage and stops moving against a target but doesn't vanish */
	_alien_projectile                 = 0x0008 /* does less damage and moves slower on lower levels */
	_affected_by_gravity              = 0x0010
	_no_horizontal_error              = 0x0020
	_no_vertical_error                = 0x0040
	_can_toggle_control_panels        = 0x0080
	_positive_vertical_error          = 0x0100
	_melee_projectile                 = 0x0200 /* can use a monsterÕs custom melee detonation */
	_persistent_and_virulent          = 0x0400 /* keeps moving and doing damage after a successful hit */
	_usually_pass_transparent_side    = 0x0800
	_sometimes_pass_transparent_side  = 0x1000
	_doubly_affected_by_gravity       = 0x2000
	_rebounds_from_floor              = 0x4000  /* unless v.z<kvzMIN */
	_penetrates_media                 = 0x8000  /* huh uh huh ... i said penetrate */
	_becomes_item_on_detonation       = 0x10000 /* item type in .permutation field of projectile */
	_bleeding_projectile              = 0x20000 /* can use a monsterÕs custom bleeding detonation */
	_horizontal_wander                = 0x40000 /* random horizontal error perpendicular to direction of movement */
	_vertical_wander                  = 0x80000 /* random vertical movement perpendicular to direction of movement */
	_affected_by_half_gravity         = 0x100000
	_projectile_passes_media_boundary = 0x200000
)

/* ---------- structures */

type projectile_definition struct {
	collection, shape                                           int16 /* collection can be NONE (invisible) */
	detonation_effect, media_detonation_effect                  int16
	contrail_effect, ticks_between_contrails, maximum_contrails int16 /* maximum of NONE is infinite */
	media_projectile_promotion                                  int16

	radius         WorldDistance /* can be zero and will still hit */
	area_of_effect WorldDistance /* one target if ==0 */
	damage         DamageDefinition

	flags uint32

	speed         WorldDistance
	maximum_range WorldDistance

	sound_pitch                cseries.Fixed
	flyby_sound, rebound_sound int16
}

const (
	GRAVITATIONAL_ACCELERATION = 1 // per tick

	WANDER_MAGNITUDE = WorldOne / TicksPerSecond

	MINIMUM_REBOUND_VELOCITY = GRAVITATIONAL_ACCELERATION * TicksPerSecond / 3
)

const (
	/* translate_projectile() flags */
	_flyby_of_current_player  = 0x0001
	_projectile_hit           = 0x0002
	_projectile_hit_monster   = 0x0004 // monster_index in *obstruction_index
	_projectile_hit_floor     = 0x0008 // polygon_index in *obstruction_index
	_projectile_hit_media     = 0x0010 // polygon_index in *obstruction_index
	_projectile_hit_landscape = 0x0020
	_projectile_hit_scenery   = 0x0040
)
const (
	/* things the projectile can hit in detonate_projectile() */
	_hit_nothing = iota
	_hit_floor
	_hit_media
	_hit_ceiling
	_hit_wall
	_hit_monster
	_hit_scenery
)

const maximumProjectileElevation = QuarterCircle / 2
