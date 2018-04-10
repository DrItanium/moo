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

sizeof(int16) constant sizeof(world-distance)
sizeof(int32) constant sizeof(fixed)
sizeof(int16) constant sizeof(angle)
sizeof(world-distance) 2* constant sizeof(world-point2d)
sizeof(fixed) 3* constant sizeof(fixed-point3d)
sizeof(world-distance) 3* constant sizeof(world-point3d)
sizeof(world-distance) 2* constant sizeof(world-vector2d)
sizeof(world-distance) 3* constant sizeof(world-vector3d)
sizeof(fixed) 3* constant sizeof(fixed-vector3d)
sizeof(world-point3d) sizeof(short) + ( a )
sizeof(angle) 2* + ( b )
sizeof(world-vector3d) +
constant sizeof(world-location3d)

: field-int16 ( adr field -- v ) 2* ( adr field*2 ) + q@ ;
: field-int32 ( adr field -- v ) 4* + h@ ;
: field-world-distance ( adr field -- v ) field-int16 ;
: field-fixed ( adr field -- v ) field-int32 ;
: field-angle ( adr field -- v ) field-int16 ;

: world-point2d.x ( adr -- x ) 0 field-world-distance ;
: world-point2d.y ( adr -- y ) 1 field-world-distance ;

: fixed-point3d.x ( adr -- x ) 0 field-fixed ;
: fixed-point3d.y ( adr -- y ) 1 field-fixed ;
: fixed-point3d.z ( adr -- z ) 2 field-fixed ;

: world-point3d.x ( adr -- x ) 0 field-world-distance ;
: world-point3d.y ( adr -- y ) 1 field-world-distance ;
: world-point3d.z ( adr -- z ) 2 field-world-distance ;

: world-vector2d.x ( adr -- x ) 0 field-world-distance ;
: world-vector2d.y ( adr -- x ) 1 field-world-distance ;

: world-vector3d.x ( adr -- x ) world-vector2d.x ;
: world-vector3d.y ( adr -- y ) world-vector2d.y ;
: world-vector3d.z ( adr -- z ) 2 field-world-distance ;

: fixed-vector3d.i ( adr -- i ) 0 field-fixed ;
: fixed-vector3d.j ( adr -- j ) 1 field-fixed ;
: fixed-vector3d.k ( adr -- k ) 2 field-fixed ;

: world-location3d.point.x ( adr -- n ) world-point2d.x ;
: world-location3d.point.y ( adr -- n ) world-point2d.y ;
: world-location3d.polygon_index ( adr -- v )
  sizeof(world-point2d) + ( nadr ) q@ ;
: world-location3d.yaw ( adr -- v )
  sizeof(world-point2d)
  sizeof(short) +
  + q@ ;
: world-location3d.pitch ( adr -- v )
  sizeof(world-point2d) sizeof(short) +
  sizeof(angle) +
  + q@ ;

: world-location3d.velocity.i ( adr -- v )
  sizeof(world-point2d) sizeof(short) +
  sizeof(angle) 2* +
  + q@ ;
  
: world-location3d.velocity.j ( adr -- v )
  sizeof(world-point2d) sizeof(short) +
  sizeof(angle) 2* +
  sizeof(world-distance) +
  + q@ ;
: world-location3d.velocity.k ( adr -- v )
  sizeof(world-point2d) sizeof(short) +
  sizeof(angle) 2* +
  sizeof(world-distance) 2* +
  + q@ ;
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

: world-point2d.x- ( p0 p1 -- dx )
  world-point2d.x swap
  world-point2d.x swap - ;
: world-point2d.y- ( p0 p1 -- dy )
  world-point2d.y swap
  world-point2d.y swap - ;

: world-point3d.x- ( p0 p1 -- dx )
  world-point3d.x swap
  world-point3d.x swap - ;
: world-point3d.y- ( p0 p1 -- dy )
  world-point3d.y swap
  world-point3d.y swap - ;
: world-point3d.z- ( p0 p1 -- dy )
  world-point3d.z swap
  world-point3d.z swap - ;

: to-int16 ( n -- k ) short-max min ;

  
: guess-distance2d ( p0 p1 -- n )
  2dup ( p0 p1 p0 p1 )
  world-point2d.x- ( p0 p1 dx )
  -rot ( dx p0 p1 )
  world-point2d.y- ( dx dy )
  abs swap abs swap ( dx dy )
  guess-hypotenuse
  to-int16 ;
: square ( n -- n^2 )
  dup * ;
  
\ Taken from the documentation laid out by the original source code
\
: isqrt ( n -- k ) ;
: distance3d ( p0 p1 -- n )
  2dup ( p0 p1 p0 p1 )
  world-point3d.x- ( p0 p1 dx )
  square ( p0 p1 dx^2 )
  -rot
  2dup ( dx^2 p0 p1 p0 p1 )
  world-point3d.y- ( dx p0 p1 dy )
  square ( dx^2 p0 p1 dy^2 )
  -rot ( dx dy p0 p1 )
  world-point3d.z- ( dx dy dz )
  square
  + + ( combined )
  isqrt ( distance )
  to-int16 ;
  
: distance2d ( p0 p1 -- n )
  2dup ( p0 p1 p0 p1 )
  world-point2d.x- square rot-
  world-point2d.y- square +
  isqrt ;
  
  
    
  \ : rotate-point2d ( point origin theta -- point )
  
