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

;s \ always last in the file
