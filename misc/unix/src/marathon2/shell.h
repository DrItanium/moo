#ifndef __SHELL_H__
#define __SHELL_H__
/*
SHELL.H
Saturday, August 22, 1992 2:18:48 PM

Saturday, January 2, 1993 10:22:46 PM
	thank god c doesn�t choke on incomplete structure references.
*/

/* ---------- constants */

#define MAXIMUM_COLORS ((short)256)

enum /* window reference numbers */
{
	refSCREEN_WINDOW= 1000
};

enum /* dialog reference numbers */
{
	refPREFERENCES_DIALOG= 8000,
	refCONFIGURE_KEYBOARD_DIALOG,
	refNETWORK_SETUP_DIALOG,
	refNETWORK_GATHER_DIALOG,
	refNETWORK_JOIN_DIALOG,
	refNETWORK_CARNAGE_DIALOG,
	
	LAST_DIALOG_REFCON= refNETWORK_CARNAGE_DIALOG,
	FIRST_DIALOG_REFCON= refPREFERENCES_DIALOG
};

#define sndCHANGED_VOLUME_SOUND 2000

/* ---------- resources */

enum {
	strPROMPTS= 131,
	_save_game_prompt= 0,
	_save_replay_prompt,
	_select_replay_prompt,
	_default_prompt
};

/* ---------- structures */

struct screen_mode_data
{
	short size;
	short acceleration;
	
	boolean high_resolution;
	boolean texture_floor, texture_ceiling;
	boolean draw_every_other_line;
	
	short bit_depth;  // currently 8 or 16
	short gamma_level;
	
	short unused[6];	// two shorts removed for preferences use
};

#define NUMBER_OF_KEYS 21
#define NUMBER_UNUSED_KEYS 10

#define PREFERENCES_VERSION 17
#define PREFERENCES_CREATOR '26.�'
#define PREFERENCES_TYPE 'pref'
#define PREFERENCES_NAME_LENGTH 32

enum // input devices
{
	_keyboard_or_game_pad,
	_mouse_yaw_pitch,
	_mouse_yaw_velocity,
	_cybermaxx_input, 						// only put "_input" here because it was defined elsewhere.
	_input_sprocket_yaw_pitch,
};

struct system_information_data
{
	boolean has_seven;
	boolean has_apple_events;
	boolean appletalk_is_available;
	boolean machine_is_68k;
	boolean machine_is_68040;
	boolean machine_is_ppc;
	boolean machine_has_network_memory;
#ifdef SUPPORT_INPUT_SPROCKET
	boolean has_input_sprocket;
#endif
#ifdef SUPPORT_DRAW_SPROCKET
	boolean has_draw_sprocket;
#endif
};

/* ---------- globals */

extern struct system_information_data *system_information;
#ifdef SUPPORT_INPUT_SPROCKET
extern boolean use_input_sprocket;

#if defined(envppc) && defined(__INPUTSPROCKET__)
extern ISpElementReference *input_sprocket_elements;
#endif
#endif
#ifdef SUPPORT_DRAW_SPROCKET
extern boolean use_draw_sprocket;
#endif

/* ---------- prototypes/SHELL.C */

void handle_game_key(EventRecord *event, short key);

/* ---------- prototypes/SHAPES.C */

void initialize_shape_handler(void);
PixMapHandle get_shape_pixmap(short shape, boolean force_copy);

void open_shapes_file(FSSpec *spec);

/* ---------- prototypes/SCREEN_DRAWING.C */

void _get_player_color(short color_index, RGBColor *color);
void _get_interface_color(short color_index, RGBColor *color);

/* ---------- protoypes/INTERFACE_MACINTOSH.C */
boolean try_for_event(boolean *use_waitnext);
void process_game_key(EventRecord *event, short key);
void update_game_window(WindowPtr window, EventRecord *event);
boolean has_cheat_modifiers(EventRecord *event);

/* ---------- prototypes/PREFERENCES.C */
void load_environment_from_preferences(void);

#endif