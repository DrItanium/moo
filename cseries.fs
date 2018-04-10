\ a translation of the cseries functions described in the marathon infinity
\ source code


1 constant byteswap_byte 
-2 constant byteswap_2byte
-4 constant byteswap_4byte

sizeof(byte) constant sizeof(char)
sizeof(byte) 2 * constant sizeof(int16)
sizeof(byte) 4 * constant sizeof(int32)
: (int16) ( n -- n ) 0xFFFF and ;
: (int32) ( n -- n ) 0xFFFFFFFF and ;
: (char) ( n -- n ) 0xFF and ;
sizeof(int16) constant sizeof(short)
sizeof(int32) constant sizeof(long)
sizeof(int16) constant sizeof(word)
sizeof(int32) constant sizeof(fixed)

: @word ( field adr -- n ) swap sizeof(word) * + @q ;

\ #define SWAP2(q) (((q)>>8) | (((q)<<8)&0xff00))
: swap2 ( q -- n ) 
	dup ( q q )
    8 u<< 0xFF00 and swap 
    8 u>> or ;
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

sizeof(long) 3 * sizeof(word) + constant sizeof(checksum) 

: @checksum.bogus1 ( adr -- v ) @h ;
: @checksum.checksum-type ( adr -- v ) sizeof(long) + @q ;
: @checksum.value ( adr -- v ) sizeof(word) sizeof(long) + + @h ;
: @checksum.bogus2 ( adr -- v ) sizeof(long) 2* sizeof(word) + + @h ;
: @checksum.value.add-checksum ( adr -- v ) @checksum.value (int16) ;
: !checksum.bogus1 ( value adr -- ) !h ;
: !checksum.checksum-type ( value adr -- v ) sizeof(long) + !q ;
: !checksum.value ( value adr -- ) sizeof(long) sizeof(word) + + !h ;
: !checksum.bogus2 ( value adr -- ) sizeof(long) 2* sizeof(word) + + !h ;
: ?checksum-types<> ( c1 c2 -- f ) 
  @checksum.checksum-type swap
  @checksum.checksum-type <> ;
: =checksum ( check1 check2 -- f ) 
  2dup ( c1 c2 c1 c2 )
  ?checksum-types<> abort" Checksum types do not match!"
  over @checksum.checksum-type add-checksum <> abort" invalid checksum type!"
  @checksum.value.add-checksum swap
  @checksum.value.add-checksum = ; 


: +=checksum.value ( value check* -- ) 
  dup @checksum.value ( value check* v2 )
  rot ( check* v2 value )
  + ( check* v3 )
  0xFFFF and ( check v4 ) \ make sure it it only 16-bits wide
  swap ( v4 check* ) !checksum.value ;

: update-add-checksum ( check* src* length -- ) 
  dup ?odd if 1- then 
  sizeof(word) / ( check src length )
  0 swap ( check* src* 0 length )
  do 
  >r \ stash the length onto the parameter stack
  2dup swap ( check* src* index src* index ) 
  @word ( check* src* index v ) swap ( c* s* v i )
  >r \ stash index onto the return stack 
  ( c* s* v ) 
  rot swap over ( s* c* v c* ) +=checksum.value swap ( c* s* )
  r> \ get the index back from the return stack
  1+ \ increment i
  r> \ put it back for the check
  continue 2drop ;


: new-checksum ( check* type -- ) ;
  2dup ( check* type check* type )
  swap ( check* type type check* )
  !checksum.checksum-type ( check* type )
  over swap ( check* check* type )
  add-checksum <> if abort" illegal checksum kind!" then 
  0 swap !checksum.value ( check* )
  dup ( check* check* )
  rand swap ( check* v check* )
  !checksum.bogus1 
  rand swap ( v check* )
  !checksum.bogus2 ;
: update-checksum ( check* src* length -- ) 
  rot dup >r -rot r> ( check* src* length check* )
  @checksum.checksum-type add-checksum <> 
  if abort" illegal checksum kind!" then
  update-add-checksum ;
: equal-checksums ( check1* check2* -- f ) =checksum ;

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

\ textures.h
sizeof(char) constant sizeof(pixel8)
sizeof(int16) constant sizeof(pixel16)
sizeof(int32) constant sizeof(pixel32)
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

: rgbcolor-to-pixel16 ( r g b -- n ) 
  11 u>> 0x1f and ( r g bmod )
  swap ( r bmod g )
  6 u>> 0x03e0 and ( r bmod gmod )
  or swap ( comb r )
  1 u>> 0x7c00 and or ;

: extract-pixel32 ( p -- n ) pixel32-maximum-component and ;
: red32 ( p -- n ) 16 u>> ;
: green32 ( p -- n ) 8 u>> extract-pixel32 ;
: blue32 ( p -- n ) extract pixel32 ;
: build-pixel32 ( r g b -- n ) 
  swap ( r b g )
  8 u<< ( r b gmod )
  or ( r comb )
  swap 
  16 u<< or ;
: rgbcolor-to-pixel32 ( r g b -- n ) 
  8 u>> 0x000000FF and swap ( r bmod g )
  0x0000FF00 and or swap ( comb r ) 
  8 u<< 0x00FF0000 and or ;

sizeof(word) 3* constant sizeof(rgb-color)

: @rgb-color.r ( ptr -- r ) @q ;
: @rgb-color.g ( ptr -- g ) sizeof(word) + @q ;
: @rgb-color.b ( ptr -- b ) sizeof(word) 2* + @q ;

: @rgb-color.rgb ( ptr -- r g b )
  dup dup ( ptr ptr ptr )
  @rgb-color.r rot ( ptr r ptr )
  @rgb-color.g rot ( r g ptr )
  @rgb-color.b ;

sizeof(short) sizeof(rgb-color) 256 * + constant sizeof(color-table)

: @color-table.color-count ( adr -- v ) @q ;
: @color-table.colors-start ( adr -- adr ) sizeof(short) + ;
: @color-table.colors-at ( adr field -- color-rgb& ) sizeof(rgb-color) * swap color-table.colors-start + ;
0x8000 constant column-order-bit
0x4000 constant transparent-bit



\ rle.h\c
\ run length encoding operations
: rle-get-destination-size ( compressed -- n ) @h ;
: uncompress-bytes ( *compressed *raw -- )
  ;
\ : compress-bytes ( raw raw-size compressed max-compressed-size -- n )
  
;s \ must always be last in file
