\ depends on world.fs

\ framerate information, yes this game is tied to framerate so 60fps causes
\ issues as the game will run twice as fast!
30 constant ticks-per-second
ticks-per-second 60 * constant ticks-per-minute

8192 constant map-index-buffer-size
world-one 4 / constant minimum-separation-from-wall
world-one *3/4 constant minimum-separation-from-projectile

ticks-per-second 2/ constant teleporting-midpoint
teleporting-midpoint 2* constant teleporting-duration

\ map limits
kilo constant max-polygons-per-map
kilo 4 * constant max-sides-per-map
kilo 8 * constant max-endpoints-per-map 
kilo 4 * constant max-lines-per-map
128 constant max-levels-per-map

64 1+ constant level-name-len

\ shape descriptor is included here... not sure why

{enum \ damage types
enum: DamageExplosion
enum: DamageElectricalStaff
enum: DamageProjectile
enum: DamageAbsorbed
enum: DamageFlame
enum: DamageHoundClaws
enum: DamageAlienProjectile
enum: DamageHulkSlap
enum: DamageCompilerBolt
enum: DamageFusionBolt
enum: DamageHunterBolt
enum: DamageFist
enum: DamageTeleporter
enum: DamageDefender
enum: DamageYetiClaws
enum: DamageYetiProjectile
enum: DamageCrushing
enum: DamageLava
enum: DamageSuffocation
enum: DamageGoo
enum: DamageEnergyDrain
enum: DamageOxygenDrain
enum: DamageHummerBolt
enum: DamageShotgunProjectile
enum}

\ damage flags
0x1 constant FlagAlienDamage \ will be decreased at lower difficulty levels

\ damage-definition struct
{struct
 field(int16): &damage-definition.type 
 field(int16): &damage-definition.flags
 field(int16): &damage-definition.base
 field(int16): &damage-definition.random
 field(int32): &damage-definition.scale
struct} constant sizeof(damage-definition)

384 constant max-saved-objects

{enum \ map object types
enum: SavedMonster \ .index is monster type 
enum: SavedObject \ scenery type
enum: SavedItem \ item type
enum: SavedPlayer \ team bitfield
enum: SavedGoal \ goal number
enum: SavedSoundSource \ source type with .facing being sound volume
enum}


\ map object flags
0x0001 constant MapObjectIsInvisibleFlag \ initially invisible 
0x0001 constant MapObjectIsPlatformSoundFlag 
0x0002 constant MapObjectIsHangingFromCeilingFlag \ used for calculating absolute .z coordinate
0x0004 constant MapObjectIsBlindFlag \ monster cannot activate by sight
0x0008 constant MapObjectIsDeafFlag \ monster cannot activate by sound
0x0010 constant MapObjectFloatsFlag \ used by sound sources caused by media 
0x0020 constant MapObjectIsNetworkOnlyFlag \ for items only
\ top four bits is activation bias for monsters
: bitset? ( value flag -- f ) and 0<> ;

: decode-activation-bias ( f -- n ) 12 u>> ;
: encode-activation-bias ( f -- n ) 12 u<< ;

{struct \ map object
 field(int16): &map-object.type 
 field(int16): &map-object.index
 field(int16): &map-object.facing
 field(int16): &map-object.polygon-index
 field(int16): &map-object.location
 sizeof(world-point3d) field: &map-object.location 
 field(int16): &map-object.flags
struct} constant sizeof(map-object)
\ sizeof(short) 4 * sizeof(word-point3d) + sizeof(word) + constant sizeof(map-object)
: map-object.flags@ ( adr -- f ) &map-object.flags + @ ;
: map-object.type@ ( adr -- f ) &map-object.type + @ ;
: map-object.flags! ( value adr -- f ) &map-object.flags + ! ;
: ?map-object.invisible ( adr -- f ) map-object.flags@ MapObjectIsInvisibleFlag bitset? ;
: ?map-object.platform-sound ( adr -- f ) map-object.flags@ MapObjectIsPlatformSoundFlag bitset? ;
: ?map-object.hanging-from-ceiling ( adr -- f ) map-object.flags@ MapObjectIsHangingFromCeilingFlag bitset? ;
: ?map-object.blind ( adr -- f ) map-object.flags@ MapObjectIsBlindFlag bitset? ;
: ?map-object.deaf ( adr -- f ) map-object.flags@ MapObjectIsDeafFlag bitset? ;
: ?map-object.floats ( adr -- f ) map-object.flags@ MapObjectFloatsFlag bitset? ;
: ?map-object.network-only ( adr -- f ) map-object.flags@ MapObjectIsNetworkOnlyFlag bitset? ;

: ?map-object.monster ( t -- f ) map-object.type@ SavedMonster = ;
: ?map-object.object ( t -- f ) map-object.type@ SavedObject = ;
: ?map-object.item ( t -- f ) map-object.type@ SavedItem = ;
: ?map-object.player ( t -- f ) map-object.type@ SavedPlayer = ;
: ?map-object.goal ( t -- f ) map-object.type@ SavedGoal = ;
: ?map-object.sound-source ( t -- f ) map-object.type@ SavedSoundSource = ;
: mo-index<-> ( -- ) &map-object.index <-> ;

mo-index<-> &map-object.sound-type
&map-object.facing <-> &map-object.sound-volume
mo-index<-> &map-object.scenery-type
mo-index<-> &map-object.monster-type
mo-index<-> &map-object.item-type
mo-index<-> &map-object.team-bitfield
mo-index<-> &map-object.goal-number
&map-object.flags <-> &map-object.monster-activation-bias

: map-object.monster-activation-bias@ ( adr -- f ) 
  &map-object.monster-activation-bias + @ 
  decode-activation-bias ;
: map-object.monster-activation-bias! ( v adr -- ) 
  &map-object.monster-activation-bias + dup @ 
  swap encode-activation-bias or swap ! ;
  
sizeof(world-point2d) constant sizeof(saved-map-pt)

\ line-data <-> saved-line
\ side-data <-> saved-side
\ polygon-data <-> saved-poly
\ map-annotation <-> saved-annotation
\ map-object <-> saved-object
\ static-data <-> saved-map-data

\ entry point types
0x01 constant EntryPointSinglePlayer
0x02 constant EntryPointMultiplayerCooperative
0x04 constant EntryPointMultiplayerCarnage
0x08 constant EntryPointCaptureTheFlag
0x10 constant EntryPointKingOfHill
0x20 constant EntryPointDefense
0x40 constant EntryPointRugby

{struct
 field(short): &entry-point.level-number
 level-name-len field: &entry-point.name 
struct} constant sizeof(entry-point)

32 constant max-player-start-name-len 

{struct \ player-start-data
 field(short): &player-start-data.team
 field(short): &player-start-data.identifier
 field(short): &player-start-data.color
 max-player-start-name-len 1+ field: &player-start-data.name 
struct} constant sizeof(player-start-data)

{struct \ directory-data
 field(short)           &directory-data.mission-flags
 field(short)           &directory-data.environment-flags
 field(long)            &directory-data.entry-point-flags
 level-name-len field:  &directory-data.level-name 
struct} constant sizeof(directory-data)

20 constant max-annotations-per-map
64 constant max-annotation-text-len
{struct \ map-annotation
    field(short): &map-annotation.type
    field(world-point2d): &map-annotation.location
    field(short): &map-annotation.polygon-index
    max-annotation-text-len field: text
struct} constant sizeof(map-annotation)

: get-next-map-annotation ( *count -- *ma ) ;

64 constant max-ambient-sound-images-per-map

{struct \ ambient-sound-image-data
\ non directional ambient component
  field(word): &ambient-sound-image-data.flags
  field(short): &ambient-sound-image-data.sound_index
  field(short): &ambient-sound-image-data.volume
  5 sizeof(short) * field: &ambient-sound-image-data.unused
struct} constant sizeof(ambient-sound-image-data)

64 constant max-random-sound-images-per-map

0x0001 constant SoundImageIsNonDirectionalFlag \ ignore direction

{struct \ random-sound-image-data
 field(word):             &random-sound-image-data.flags
 field(short):            &random-sound-image-data.sound-index
 field(short):            &random-sound-image-data.volume 
 field(short):            &random-sound-image-data.delta-volume 
 field(short):            &random-sound-image-data.period
 field(short):            &random-sound-image-data.delta-period
 field(angle):            &random-sound-image-data.direction
 field(angle):            &random-sound-image-data.delta-direction
 field(fixed):            &random-sound-image-data.pitch
 field(fixed):            &random-sound-image-data.delta-pitch
 field(short):            &random-sound-image-data.phase \ only used at runtime; initialize to none
 3 sizeof(short) * field: &random-sound-image-data.unused
struct} constant sizeof(random-sound-image-data)
\ the map contains a series of slots in which map objects are placed
384 constant max-objects-per-map

: ?flag-marked ( adr mask -- b ) swap word@ and (word) ;
: ?slot-used ( adr -- b ) 0x8000 ?flag-marked ;
: ?slot-free ( adr -- b ) ?slot-used invert (word) ;
: set-flag ( adr flag -- ) (word) over word@ or swap word! ;
: clear-flag ( adr flag -- ) invert (word) over word@ and swap word! ;
: mark-slot-as-free ( adr -- ) 0x8000 clear-flag ;
: mark-slot-as-used ( adr -- ) 0x8000 set-flag ;
: ?object-was-rendered ( adr -- f ) 0x4000 ?flag-marked ;
: mark-object-rendered ( adr -- ) 0x4000 set-flag ;
: mark-object-not-rendered ( adr -- ) 0x4000 clear-flag ;
\ map.h: line 238
;s \ always last in the file
