#!/bin/bash
function base-convert-newlines {
	tr '\r' '\n' < $1 > $2
}

function convert-asm {
	base-convert-newlines source/marathon2/$1 ../unix/src/marathon2/$1
}


convert-asm quadruple.s 
convert-asm scottish_textures.s
convert-asm network_listener.a
convert-asm scottish_textures16.a
convert-asm scottish_textures.a
convert-asm screen.a
