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

;s
