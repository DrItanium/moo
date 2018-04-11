\ a translation of the cseries functions described in the marathon infinity
\ source code



sizeof(byte) constant sizeof(char)
: field(char): ( a b -- c c ) sizeof(char) field: ;
: (int16) ( n -- n ) 0xFFFF and ;
: (int32) ( n -- n ) 0xFFFFFFFF and ;
: (char) ( n -- n ) 0xFF and ;
sizeof(int16) constant sizeof(short)
: field(short): ( a b -- c c ) sizeof(short) field: ;
sizeof(int32) constant sizeof(long)
: field(long): ( a b -- c c ) sizeof(long) field: ;
sizeof(int16) constant sizeof(word)
: field(word): ( a b -- c c ) sizeof(word) field: ;
sizeof(int32) constant sizeof(fixed)
: field(fixed): ( a b -- c c ) sizeof(fixed) field: ;
: long@ ( adr -- v ) int32@ ;
: short@ ( adr -- v ) int16@ ;
: word@ ( adr -- v ) int16@ ;
: fixed@ ( adr -- v ) int32@ ;
: long! ( v adr -- ) int32! ;
: short! ( v adr -- ) int16! ;
: word! ( v adr -- ) int16! ;
: fixed! ( v adr -- ) int32! ;

: (byte) ( n -- v ) (char) ;
: (word) ( n -- v ) (int16) ;
: (long) ( n -- v ) (int32) ;
: (short) ( n -- v ) (int16) ;
: (fixed) ( n -- v ) (int32) ;



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
: top-else-lower ( a b -- a | b ) if nip else drop then ;
: floor ( n floor -- v ) 2dup < top-else-lower ;
: ceiling ( n ceiling -- v ) 2dup > top-else-lower ;
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

: flag ( b -- f ) 1u<< ;

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

{enum
enum: fatal-error
enum: info-error
enum}

\ : alert-user ( type resource-number error-number identifier -- ) ;
\ : get-cstr ( buffer collection-number string-number -- ptr ) ;

\ tons of functions at the bottom of cseries.h which have not been implemented



;s \ must always be last in file

