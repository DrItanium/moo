{enum
enum: ADD-CHECKSUM
enum: FLETCHER-CHECKSUM
enum: CRC32-CHECKSUM
enum}

\ there is code in checksum.h for the Checksum structure, we'll come back to
\ this
{struct \ checksum
    field(long): &checksum.bogus1
    field(word): &checksum.type
    field(long): &checksum.value
    field(long): &checksum.bogus2
struct} constant sizeof(checksum)

: checksum.bogus1! ( v adr -- ) &checksum.bogus1 + long! ;
: checksum.bogus2! ( v adr -- ) &checksum.bogus2 + long! ;
: checksum.type@ ( adr -- v ) &checksum.type + word@ ;
: checksum.type! ( v adr -- v ) &checksum.type + word! ;
: checksum.value@ ( adr -- v ) &checksum.value + long@ ;
: checksum.value! ( v adr -- ) &checksum.value + long! ;
: checksum.value.add-checksum@ ( adr -- v ) checksum.value@ (word) ;
: ?checksum.type-is-not-add ( adr -- f ) checksum.type@ add-checksum <> ;
: ?checksum-types<> ( c1 c2 -- f ) 
  checksum.type@ swap
  checksum.type@ <> ;
: =checksum ( check1 check2 -- f ) 
  2dup ( c1 c2 c1 c2 )
  ?checksum-types<> abort" Checksum types do not match!"
  over ?checksum.type-is-not-add abort" invalid checksum type!"
  checksum.value.add-checksum@ swap
  checksum.value.add-checksum@ = ; 


: +=checksum.value ( value check* -- ) 
  dup checksum.value@ ( value check* v2 )
  rot ( check* v2 value )
  + ( check* v3 )
  0xFFFF and ( check v4 ) \ make sure it it only 16-bits wide
  swap ( v4 check* ) checksum.value! ;

: update-add-checksum ( check* src* length -- ) 
  dup ?odd if 1- then 
  sizeof(word) / ( check src length )
  0 swap ( check* src* 0 length )
  do 
  >r \ stash the length onto the parameter stack
  2dup swap ( check* src* index src* index ) 
  word@ ( check* src* index v ) swap ( c* s* v i )
  >r \ stash index onto the return stack 
  ( c* s* v ) 
  rot swap over ( s* c* v c* ) 
  +=checksum.value swap ( c* s* )
  r> \ get the index back from the return stack
  1+ \ increment i
  r> \ put it back for the check
  continue 2drop ;


: new-checksum ( check* type -- ) ;
  2dup ( check* type check* type )
  swap ( check* type type check* )
  checksum.type! ( check* type )
  over swap ( check* check* type )
  add-checksum <> if abort" illegal checksum kind!" then 
  0 swap checksum.value! ( check* )
  dup ( check* check* )
  rand swap ( check* v check* )
  checksum.bogus1!
  rand swap ( v check* )
  checksum.bogus2! ;

: update-checksum ( check* src* length -- ) 
  rot dup >r -rot r> ( check* src* length check* )
  ?checksum.type-is-not-add abort" illegal checksum kind!"
  update-add-checksum ;

: equal-checksums ( check1* check2* -- f ) =checksum ;

;s
