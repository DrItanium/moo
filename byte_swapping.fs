\ byte swapping related operations
1 constant byteswap_byte 
-2 constant byteswap_2byte
-4 constant byteswap_4byte
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

;s
