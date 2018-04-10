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
;s \ always last in the file
