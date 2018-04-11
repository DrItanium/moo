: *3/4 ( x -- n ) 3 * 4/ ;
10 constant trig-shift
trig-shift 1u<< constant trig-magnitude

9 constant angular-bits
angular-bits 1u<< constant number-of-angles
number-of-angles constant full-circle
number-of-angles 4/ constant quarter-circle
number-of-angles 2/ constant half-circle
number-of-angles *3/4 constant three-quarter-circle
number-of-angles 3 u<< constant eighth-circle
number-of-angles 4 u<< constant sixteenth-circle

10 constant world-fractional-bits
world-fractional-bits 1u<< constant world-one
world-one 2/ constant world-one-half
world-one 4/ constant world-one-fourth
world-one *3/4 constant world-three-fourths

0xfded constant default-random-seed

: integer-to-world ( s -- w ) world-fractional-bits u<< ;
: world-fractional-part ( d -- fx ) world-one 1- and ;
: world-integral-part ( d -- i )  world-fractional-bits u>> ;

fixed-fractional-bits world-fractional-bits - constant fractional-bits-difference

: world-to-fixed ( w -- fx ) fractional-bits-difference u<< ;
: fixed-to-world ( fx -- w ) fractional-bits-difference u>> ;
: normalize-angle ( t -- n ) number-of-angles 1- and ;
: facing4 ( a -- n ) normalize-angle eighth-circle - angular-bits 2- u>> ;
: facing5 ( a -- n )
  normalize-angle full-circle 10 / -
  number-of-angles 5 / 1+ / ;
: facing8 ( a -- n )
  normalize-angle sixteenth-circle -
  angular-bits 3 - u>> ;

: guess-hypotenuse ( x y -- n )
  2dup ( x y x y )
  <= ( x y )
  if swap then
  1 u>> + ;

sizeof(int16) <-> sizeof(world-distance)
sizeof(int16) <-> sizeof(angle)
: field(world-distance): ( a b -- c c ) sizeof(world-distance) field: ;
: field(angle): ( a b -- c c ) sizeof(angle) field: ;
{struct \ world-point2d
    field(world-distance): &world-point2d.x 
    field(world-distance): &world-point2d.y
struct} constant sizeof(world-point2d)


{struct \ world-point3d
    field(world-distance): &world-point3d.x 
    field(world-distance): &world-point3d.y
    field(world-distance): &world-point3d.z
struct} constant sizeof(world-point3d)

{struct \ fixed-point3d 
    field(fixed): &fixed-point3d.x
    field(fixed): &fixed-point3d.y
    field(fixed): &fixed-point3d.z
struct} constant sizeof(fixed-point3d)
: field(world-point2d) ( a b -- c c ) sizeof(world-point2d) field: ;
: field(world-point3d) ( a b -- c c ) sizeof(world-point3d) field: ;
: field(fixed-point3d) ( a b -- c c ) sizeof(fixed-point3d) field: ;

{struct \ world-vector2d
    field(world-distance): &world-vector2d.i 
    field(world-distance): &world-vector2d.j
struct} constant sizeof(world-vector2d)


{struct \ world-vector3d
    field(world-distance): &world-vector3d.i 
    field(world-distance): &world-vector3d.j
    field(world-distance): &world-vector3d.k
struct} constant sizeof(world-vector3d)

{struct \ fixed-vector3d
    field(fixed): &fixed-vector3d.i 
    field(fixed): &fixed-vector3d.j
    field(fixed): &fixed-vector3d.k
struct} constant sizeof(fixed-vector3d)

: field(world-vector2d) ( a b -- c c ) sizeof(world-vector2d) field: ;
: field(world-vector3d) ( a b -- c c ) sizeof(world-vector3d) field: ;
: field(fixed-vector3d) ( a b -- c c ) sizeof(fixed-vector3d) field: ;

{struct \ world-location3d
field(world-point3d): &world-location3d.point
field(short): &world-location3d.polygon-index
field(angle): &world-location3d.yaw
field(angle): &world-location3d.pitch
field(world-vector3d) &world-location3d.velocity
struct} constant sizeof(world-location3d)

: field(world-location3d) ( a b -- c c ) sizeof(world-location3d) field: ;


: world-point2d.x@ ( adr -- x ) &world-point2d.x + @ ;
: world-point2d.y@ ( adr -- y ) &world-point2d.y + @ ;

: world-point3d.x@ ( adr -- x ) &world-point3d.x + @ ;
: world-point3d.y@ ( adr -- y ) &world-point3d.y + @ ;
: world-point3d.z@ ( adr -- z ) &world-point3d.z + @ ;

: fixed-point3d.x@ ( adr -- x ) &fixed-point3d.x + @ ;
: fixed-point3d.y@ ( adr -- y ) &fixed-point3d.y + @ ;
: fixed-point3d.z@ ( adr -- z ) &fixed-point3d.z + @ ;

: world-vector2d.i@ ( adr -- i ) &world-vector2d.i + @ ;
: world-vector2d.j@ ( adr -- j ) &world-vector2d.j + @ ;

: world-vector3d.i@ ( adr -- i ) &world-vector3d.i + @ ;
: world-vector3d.j@ ( adr -- j ) &world-vector3d.j + @ ;
: world-vector3d.k@ ( adr -- k ) &world-vector3d.k + @ ;

: fixed-vector3d.i@ ( adr -- i ) &fixed-vector3d.i + @ ;
: fixed-vector3d.j@ ( adr -- j ) &fixed-vector3d.j + @ ;
: fixed-vector3d.k@ ( adr -- k ) &fixed-vector3d.k + @ ;

: world-location3d.polygon-index@ ( adr -- v ) &world-location3d.polygon-index + @ ;
: world-location3d.yaw@ ( adr -- v ) &world-location3d.yaw + @ ;
: world-location3d.pitch@ ( adr -- v ) &world-location3d.pitch + @ ;
variable random-seed
default-random-seed random-seed !
: set-random-seed ( seed -- )
  dup 0<> if drop default-random-seed then random-seed ! ;
: random ( -- n )
  random-seed @ dup
  1 and
  if
  1 u>> 0xb400 xor
  else
  1 u>>
  then
  dup random-seed ! ;

: world-point2d.x@- ( p0 p1 -- dx )
  world-point2d.x@ swap
  world-point2d.x@ swap - ;
: world-point2d.y@- ( p0 p1 -- dy )
  world-point2d.y@ swap
  world-point2d.y@ swap - ;

: world-point3d.x@- ( p0 p1 -- dx )
  world-point3d.x@ swap
  world-point3d.x@ swap - ;
: world-point3d.y@- ( p0 p1 -- dy )
  world-point3d.y@ swap
  world-point3d.y@ swap - ;
: world-point3d.z@- ( p0 p1 -- dz )
  world-point3d.z@ swap
  world-point3d.z@ swap - ;

: to-int16 ( n -- k ) short-max min ;

: 2abs ( a b -- aa ab ) abs swap abs swap ;
: guess-distance2d ( p0 p1 -- n )
  2dup ( p0 p1 p0 p1 )
  world-point2d.x@- ( p0 p1 dx )
  -rot ( dx p0 p1 )
  world-point2d.y@- ( dx dy )
  2abs ( dx dy )
  guess-hypotenuse
  to-int16 ;
: square ( n -- n^2 )
  dup * ;
: square3 ( a b c -- c^2 a b ) square -rot ;
  
\ Taken from the documentation laid out by the original source code
: isqrt ( n -- k ) ;

: distance3d ( p0 p1 -- n )
  2dup ( p0 p1 p0 p1 )
  world-point3d.x@- ( p0 p1 dx )
  square3
  2dup ( dx^2 p0 p1 p0 p1 )
  world-point3d.y@- ( dx p0 p1 dy )
  square3
  world-point3d.z@- ( dx dy dz )
  square
  + + ( combined )
  isqrt ( distance )
  to-int16 ;
  
: distance2d ( p0 p1 -- n )
  2dup ( p0 p1 p0 p1 )
  world-point2d.x- 
  square3
  world-point2d.y- 
  square +
  isqrt ;
  
  
    
  
;s \ must always be last in file
