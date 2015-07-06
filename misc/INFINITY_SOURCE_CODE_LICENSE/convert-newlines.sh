#!/bin/bash
function base-convert-newlines {
	tr '\r' '\n' < $1 > $2
}
function convert-newlines {
	base-convert-newlines source/marathon2/$1 unix/src/marathon2/$1
}
function convert-newlines2 {
	mkdir -p unix/src/cseries-interfaces/
	base-convert-newlines source/CSeriesInterfaces/$1 unix/src/cseries-interfaces/$1
}
function convert-newlines3 {
	mkdir -p unix/src/cseries-libraries/
	base-convert-newlines source/CSeriesLibraries/$1 unix/src/cseries-libraries/$1
}
function convert-newlines4 {
	mkdir -p unix/src/cseries/
	base-convert-newlines source/cseries/$1 unix/src/cseries/$1
}
function convert-newlines5 {
	mkdir -p unix/src/cseries.lib/
	base-convert-newlines source/cseries.lib/$1 unix/src/cseries.lib/$1
}
convert-newlines collection_definition.h
convert-newlines computer_interface.c
convert-newlines computer_interface.h
convert-newlines crc.c
convert-newlines crc.h
convert-newlines devices.c
convert-newlines editor.h
convert-newlines effect_definitions.h
convert-newlines effects.c
convert-newlines effects.h
convert-newlines environment.h
convert-newlines export_definitions.c
convert-newlines extensions.h
convert-newlines fades.c
convert-newlines fades.h
convert-newlines files_macintosh.c
convert-newlines find_files.c
convert-newlines find_files.h
convert-newlines flood_map.c
convert-newlines flood_map.h
convert-newlines game_dialogs.c
convert-newlines game_errors.c
convert-newlines game_errors.h
convert-newlines game_sound.c
convert-newlines game_sound.h
convert-newlines game_wad.c
convert-newlines game_wad.h
convert-newlines game_window.c
convert-newlines game_window.h
convert-newlines game_window_macintosh.c
convert-newlines images.c
convert-newlines images.h
convert-newlines import_definitions.c
convert-newlines input_sprocket_needs.h
convert-newlines interface.c
convert-newlines interface.h
convert-newlines interface_macintosh.c
convert-newlines interface_menus.h
convert-newlines item_definitions.h
convert-newlines items.c
convert-newlines items.h
convert-newlines keyboard_dialog.c
convert-newlines key_definitions.h
convert-newlines lightsource.c
convert-newlines lightsource.h
convert-newlines low_level_textures.c
convert-newlines macintosh_input.h
convert-newlines macintosh_network.h
convert-newlines map_accessors.c
convert-newlines map.c
convert-newlines map_constructors.c
convert-newlines map.h
convert-newlines marathon2.c
convert-newlines media.c
convert-newlines media_definitions.h
convert-newlines media.h
convert-newlines monster_definitions.h
convert-newlines monsters.c
convert-newlines monsters.h
convert-newlines motion_sensor.c
convert-newlines motion_sensor.h
convert-newlines mouse.c
convert-newlines mouse.h
convert-newlines music.c
convert-newlines music.h
convert-newlines network_adsp.c
convert-newlines network.c
convert-newlines network_ddp.c
convert-newlines network_dialogs.c
convert-newlines network_games.c
convert-newlines network_games.h
convert-newlines network.h
convert-newlines network_lookup.c
convert-newlines network_microphone.c
convert-newlines network_modem.c
convert-newlines network_modem.h
convert-newlines network_modem_protocol.c
convert-newlines network_modem_protocol.h
convert-newlines network_names.c
convert-newlines network_sound.h
convert-newlines network_speaker.c
convert-newlines network_stream.c
convert-newlines network_stream.h
convert-newlines overhead_map.c
convert-newlines overhead_map.h
convert-newlines overhead_map_macintosh.c
convert-newlines pathfinding.c
convert-newlines physics.c
convert-newlines physics_models.h
convert-newlines physics_patches.c
convert-newlines placement.c
convert-newlines platform_definitions.h
convert-newlines platforms.c
convert-newlines platforms.h
convert-newlines player.c
convert-newlines player.h
convert-newlines portable_files.h
convert-newlines preferences.c
convert-newlines preferences.h
convert-newlines preprocess_map_mac.c
convert-newlines progress.c
convert-newlines progress.h
convert-newlines projectile_definitions.h
convert-newlines projectiles.c
convert-newlines projectiles.h
convert-newlines render.c
convert-newlines render.h
convert-newlines scenery.c
convert-newlines scenery_definitions.h
convert-newlines scenery.h
convert-newlines scottish_textures.c
convert-newlines scottish_textures.h
convert-newlines screen.c
convert-newlines screen_definitions.h
convert-newlines screen_drawing.c
convert-newlines screen_drawing.h
convert-newlines screen.h
convert-newlines serial_numbers.c
convert-newlines shape_definitions.h
convert-newlines shape_descriptors.h
convert-newlines shapes.c
convert-newlines shapes_macintosh.c
convert-newlines shell.c
convert-newlines shell.h
convert-newlines song_definitions.h
convert-newlines sound_definitions.h
convert-newlines sound_macintosh.c
convert-newlines tags.h
convert-newlines textures.c
convert-newlines textures.h
convert-newlines valkyrie.c
convert-newlines valkyrie.h
convert-newlines vbl.c
convert-newlines vbl_definitions.h
convert-newlines vbl.h
convert-newlines vbl_macintosh.c
convert-newlines wad.c
convert-newlines wad.h
convert-newlines wad_macintosh.c
convert-newlines wad_prefs.c
convert-newlines wad_prefs.h
convert-newlines wad_prefs_macintosh.c
convert-newlines weapon_definitions.h
convert-newlines weapons.c
convert-newlines weapons.h
convert-newlines world.c
convert-newlines world.h


convert-newlines2 byte_swapping.h
convert-newlines2 checksum.h
convert-newlines2 cseries.h
convert-newlines2 InputSprocket.h
convert-newlines2 macintosh_cseries.h
convert-newlines2 macintosh_interfaces.c
convert-newlines2 my32bqd.h
convert-newlines2 mytm.h
convert-newlines2 preferences.h
convert-newlines2 proximity_strcmp.h
convert-newlines2 rle.h



convert-newlines3 DrawSprocketDebugLib
convert-newlines3 DrawSprocket.h
convert-newlines3 InputSprocket.h
convert-newlines3 InputSprocketStubLib
convert-newlines3 macintosh_interfaces881.d
convert-newlines3 macintosh_interfaces.d
convert-newlines3 SoundSprocket.h


convert-newlines4 DrawSprocket.h
convert-newlines4 InputSprocket.h
convert-newlines4 SoundSprocket.h


convert-newlines5 AfterDarkGestalt.h
convert-newlines5 beta.h
convert-newlines5 buildprogram
convert-newlines5 byte_swapping.c
convert-newlines5 byte_swapping.h
convert-newlines5 checksum.c
convert-newlines5 checksum.h
convert-newlines5 cseries.h
convert-newlines5 device_dialog.c
convert-newlines5 devices.c
convert-newlines5 dialogs.c
convert-newlines5 final.h
convert-newlines5 InputSprocket.h
convert-newlines5 macintosh_cseries.h
convert-newlines5 macintosh_interfaces.c
convert-newlines5 macintosh_interfaces.d.make
convert-newlines5 macintosh_interfaces.ppc
convert-newlines5 macintosh_utilities.c
convert-newlines5 makefile
convert-newlines5 my32bqd.c
convert-newlines5 my32bqd.h
convert-newlines5 mytm.c
convert-newlines5 mytm.h
convert-newlines5 preferences.c
convert-newlines5 preferences.h
convert-newlines5 proximity_strcmp.c
convert-newlines5 proximity_strcmp.h
convert-newlines5 RequestVideo.c
convert-newlines5 RequestVideo.h
convert-newlines5 rle.c
convert-newlines5 rle.h
convert-newlines5 serial_numbers.makeout
convert-newlines5 textures.h
convert-newlines5 Touch
