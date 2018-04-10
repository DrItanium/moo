\ a translation of the cseries functions described in the marathon infinity
\ source code


1 constant byteswap_byte 
-2 constant byteswap_2byte
-4 constant byteswap_4byte

\ #define SWAP2(q) (((q)>>8) | (((q)<<8)&0xff00))
: swap2 ( n -- n ) 
	dup ( n n )
	8 u>> swap ( k n )
	8 u<< 0xFF00 and
	 or ;
: swap4 ( n -- n )
	dup dup dup ( n n n n )
	24 u>> swap ( n n l n )
	8 u>> 0xFF00 and ( n n l l2 )
	or swap ( n k n )
	8 u<< 0x00ff00 and ( n k l )
	or swap ( z n ) 
	24 u<< 0xFF000000 and or ;

{enum
enum: ADD-CHECKSUM
enum: FLETCHER-CHECKSUM
enum: CRC32-CHECKSUM
enum}

\ there is code in checksum.h for the Checksum structure, we'll come back to
\ this


\ cseries.h
: true ( -- n ) 1 ;
: none ( -- n ) -1 ;
1024 constant kilo
kilo kilo * constant meg
meg meg * constant gig

60 constant machine-ticks-per-second

: sgn ( x -- n ) 
  dup 0<> 
  if 
      0< if -1 else 1 then 
  else 
      0 
  then ;

\ can we get away with just swapping stack arguments? lets hope so for now
\ abs, min, and max defined in basics.fs

: floor ( n floor -- v ) 2dup < if nip else drop then ;
: ceiling ( n ceiling -- v ) 2dup > if nip else drop then ;
: pin ( n floor ceiling -- v ) \ n < floor ? floor : ceiling ( n , ceiling ) 
  rot ( floor ceiling n )
  tuck ( floor n ceiling n )
  rot ( n ceiling floor n )
  swap ( n ceiling n floor ) 
  2dup  ( n ceiling n floor n floor )
  < if 
  nip ( n ceiling floor )
  -rot ( floor n ceiling )
  2drop ( floor )
  else
  2drop ( n ceiling )
  ceiling ( v ) 
  then ;

: flag ( b -- f ) 
  1 swap ( 1 b )
  u<< ( f ) ;

: test-flag ( f b -- v ) 
  flag ( f m ) 
  and ( v ) ;

\ swap_flag16 not implemented
: set-flag ( f b v -- n )
  if 
  	\ f |= word(Flag(b))
  	flag
  	or 
  else 
  	flag 
  	invert
  	and 
  then ;

\ fixed point coordinate system
16 constant fixed-fractional-bits
1 fixed-fractional-bits u<< constant fixed-one
1 fixed-fractional-bits 1- u<< constant fixed-one-half

\ there are routines for fixed to float and float to fixed, ignoring

: integer-to-fixed ( s -- v ) fixed-fractional-bits << ;
: fixed-integral-part ( s -- v ) fixed-fractional-bits >> ;
: fixed-to-integer ( f -- v ) fixed-integral-part ;
: fixed-to-integer-round ( f -- v ) fixed-one-half + fixed-integral-part ;
: fixed-fractional-part ( f -- v ) fixed-one 1- and ;


\ limits and type values
4294967295 constant *unsigned-long-max*
2147483647 constant *long-max*
-2147483648 constant *long-min*
32 constant *long-bits* 
65535 constant *unsigned-short-max*
32767 constant *short-max*
-32768 constant *short-min*
16 constant *short-bits*
255 constant *unsigned-char-max*
127 constant *char-max*
-128 constant *char-min*
8 constant *char-bits*

\ tons of functions at the bottom of cseries.h which have not been implemented

\ textures.h
256 constant pixel8-maximum-colors
32768 constant pixel16-maximum-colors 
16777216 constant pixel32-maximum-colors
5 constant pixel16-bits
8 constant pixel32-bits

3 constant number-of-color-components
0x1f constant pixel16-maximum-component
0xff constant pixel32-maximum-component 

\ 16-bit color pixels have 5 bits per color channel plus one bit unused?
: red16 ( p -- v ) 10 u>> ;
: green16 ( p -- v ) 
  5 u>> 
  pixel16-maximum-component and ;
: blue16 ( p -- v )
  pixel16-maximum-component and ;
: build-pixel16 ( r g b -- v )
  -rot ( b r g )
   5 u<< swap ( b g! r )
   10 u<< or ( b gr )
   or ;
: 


;s \ must always be last in file
