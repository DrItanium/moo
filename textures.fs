
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

;s \ must always be last in file
