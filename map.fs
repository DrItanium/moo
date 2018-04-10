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
kilo constant maximum-polygons-per-map
kilo 4 * constant maximum-sides-per-map
kilo 8 * constant maximum-endpoints-per-map 
kilo 4 * constant maximum-lines-per-map
128 constant maximum-levels-per-map

64 1+ constant level-name-length

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
sizeof(short) 4* sizeof(fixed) + constant sizeof(damage-definition) 

: @damage-definition.type ( adr -- fx ) @q ;
: @damage-definition.flags ( adr -- fx ) sizeof(short) + @q ;
: @damage-definition.base ( adr -- fx ) sizeof(short) 2* + @q ;
: @damage-definition.random ( adr -- fx ) sizeof(short) 3 * + @q ;
: @damage-definition.scale ( adr -- fx ) sizeof(short) 4 * + @h ;
: !damage-definition.type ( value adr -- fx ) !q ;
: !damage-definition.flags ( value adr -- fx ) sizeof(short) + !q ;
: !damage-definition.base ( value adr -- fx ) sizeof(short) 2* + !q ;
: !damage-definition.random ( value adr -- fx ) sizeof(short) 3 * + !q ;
: !damage-definition.scale ( value adr -- fx ) sizeof(short) 4 * + !h ;


384 constant maximum-saved-objects

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

sizeof(short) 4 * sizeof(word-point3d) + sizeof(word) + constant sizeof(map-object)
0 
: @map-object.type ( adr -- t ) @q ;
: @map-object.index ( adr -- t ) [ sizeof(short) + dup ]L + @q ;
: @map-object.facing ( adr -- t ) [ sizeof(short) + dup ]L + @q ;
: @map-object.polygon-index ( adr -- t ) [ sizeof(short) + dup ]L + @q ;
: &map-object.location ( adr -- adr ) [ sizeof(short) + dup ]L + ;
: @map-object.flags ( adr -- t ) [ sizeof(world-point3d) + ]L + @word ;

0 
: !map-object.type ( value adr -- ) !q ;
: !map-object.index ( value adr -- ) [ sizeof(short) + dup ]L + @q ;
: !map-object.facing ( value adr -- ) [ sizeof(short) + dup ]L + @q ;
: !map-object.polygon-index ( value adr -- ) [ sizeof(short) + dup ]L + @q ;
: !map-object.flags ( value adr -- ) [ sizeof(world-point3d) + ]L + @q ;

: ?map-object.invisible ( adr -- f ) @map-object.flags MapObjectIsInvisibleFlag bitset? ;
: ?map-object.platform-sound ( adr -- f ) @map-object.flags MapObjectIsPlatformSoundFlag bitset? ;
: ?map-object.hanging-from-ceiling ( adr -- f ) @map-object.flags MapObjectIsHangingFromCeilingFlag bitset? ;
: ?map-object.blind ( adr -- f ) @map-object.flags MapObjectIsBlindFlag bitset? ;
: ?map-object.deaf ( adr -- f ) @map-object.flags MapObjectIsDeafFlag bitset? ;
: ?map-object.floats ( adr -- f ) @map-object.flags MapObjectFloatsFlag bitset? ;
: ?map-object.network-only ( adr -- f ) @map-object.flags MapObjectIsNetworkOnlyFlag bitset? ;

: ?map-object.monster ( t -- f ) @map-object.type SavedMonster = ;
: ?map-object.object ( t -- f ) @map-object.type SavedObject = ;
: ?map-object.item ( t -- f ) @map-object.type SavedItem = ;
: ?map-object.player ( t -- f ) @map-object.type SavedPlayer = ;
: ?map-object.goal ( t -- f ) @map-object.type SavedGoal = ;
: ?map-object.sound-source ( t -- f ) @map-object.type SavedSoundSource = ;

: @map-object.sound-type ( t -- f ) @map-object.index ;
: @map-object.sound-volume ( t -- f ) @map-object.facing ;
: @map-object.monster-type ( t -- f ) @map-object.index ;
: @map-object.scenery-type ( t -- f ) @map-object.index ;
: @map-object.item-type  ( t -- f ) @map-object.index ;
: @map-object.team-bitfield  ( t -- f ) @map-object.index ;
: @map-object.goal-number ( t -- f ) @map-object.index ;
: !map-object.sound-type ( v t -- f ) !map-object.index ;
: !map-object.sound-volume ( v t -- f ) !map-object.facing ;
: !map-object.monster-type ( v t -- f ) !map-object.index ;
: !map-object.scenery-type ( v t -- f ) !map-object.index ;
: !map-object.item-type  ( v t -- f ) !map-object.index ;
: !map-object.team-bitfield  ( v t -- f ) !map-object.index ;
: !map-object.goal-number ( v t -- f ) !map-object.index ;

: @map-object.monster-activation-bias ( t -- f ) @map-object.flags decode-activation-bias ;
: !map-object.monster-activation-bias ( v t -- ) 
  dup @map-object.flags
  swap encode-activation-bias or ( t v ) 
  swap !map-object.flags ;
  
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

sizeof(short) level-name-length + constant sizeof(entry-point)

: &entry-point.level-number ( adr -- adr ) ;
: &entry-point.level-name ( adr -- adr ) sizeof(short) + ;
: @entry-point.level-number ( adr -- n ) &entry-point.level-number @q ;
: @entry-point.level-name:index ( offset adr -- n ) &entry-point.level-name + @ ;
: !entry-point.level-number ( value adr -- n ) &entry-point.level-number !q ;
: !entry-point.level-name:index ( value offset adr -- n ) &entry-point.level-name + ! ;

32 constant maximum-plyaer-start-name-length

sizeof(short) 3 * maximum-player-start-name-length 1+ + constant sizeof(player-start-data)

: &player-start-data.team ( adr -- adr ) ;
: &player-start-data.identifier ( adr -- adr ) sizeof(short) + ;
: &player-start-data.color ( adr -- adr ) sizeof(short) 2* + ;
: &player-start-data.name ( adr -- adr ) sizeof(short) 3 * + ;
: @player-start-data.team ( adr -- n ) &player-start-data.team @q ;
: @player-start-data.identifier ( adr -- n ) &player-start-data.identifier @q ;
: @player-start-data.color ( adr -- n ) &player-start-data.color @q ;
: @player-start-data.name:index ( offset adr -- n ) &player-start-data.name + @ ;
: !player-start-data.team ( value adr -- n ) &player-start-data.team !q ;
: !player-start-data.identifier ( value adr -- n ) &player-start-data.identifier !q ;
: !player-start-data.color ( value adr -- n ) &player-start-data.color !q ;
: !player-start-data.name:index ( value offset adr -- n ) &player-start-data.name + ! ;

;s \ always last in the file
