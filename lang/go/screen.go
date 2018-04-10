// screen related functions
package moo

const (
	_full_screen = iota
	_100_percent
	_75_percent
	_50_percent
)

const (
	/* hardware acceleration codes */
	_no_acceleration = iota
	_valkyrie_acceleration
)

var WorldColorTable, VisibleColorTable, InterfaceColorTable ColorTable

func ChangeScreenClut(table ColorTable) {

}
func AnimateScreenClut(table ColorTable, fullScreen bool) {

}

func ChangeInterfaceClut(table ColorTable) {

}

func StartTeleportingEffect(out bool) {

}

func StartExtravisionEffect(out bool) {

}

func RenderScreen(ticksElapsed int16) {

}

func BuildDirectColorTable(table ColorTable, bitDepth int16) {

}

func ToggleOverheadMapDisplayStatus() {

}

func ZoomOverheadMapOut() {

}

func ZoomOverheadMapIn() {

}

func EnterScreen() {

}

func ExitScreen() {

}

func DarkenWorldWindow() {

}

func ValidateWorldWindow() {

}

func ChangeGammaLevel(gammaLevel int16) {

}

func AssertWorldColorTable(worldColorTable, interfaceColorTable ColorTable) {

}

const (
	INTRO_SCREEN_BASE       = 1000
	MAIN_MENU_BASE          = 1100
	PROLOGUE_SCREEN_BASE    = 1200
	EPILOGUE_SCREEN_BASE    = 1300
	CREDIT_SCREEN_BASE      = 1400
	CHAPTER_SCREEN_BASE     = 1500
	COMPUTER_INTERFACE_BASE = 1600
	INTERFACE_PANEL_BASE    = 1700
	FINAL_SCREEN_BASE       = 1800
)

/* rectangle id's */
const (
	/* game window rectangles */
	_player_name_rect = iota
	_oxygen_rect
	_shield_rect
	_motion_sensor_rect
	_microphone_rect
	_inventory_rect
	_weapon_display_rect

	/* interface rectangles */
	_new_game_button_rect
	_load_game_button_rect
	_gather_button_rect
	_join_button_rect
	_prefs_button_rect
	_replay_last_button_rect
	_save_last_button_rect
	_replace_saved_button_rect
	_credits_button_rect
	_quit_button_rect
	_center_button_rect
	NUMBER_OF_INTERFACE_RECTANGLES
)

const (
	/* Colors for drawing.. */
	_energy_weapon_full_color = iota
	_energy_weapon_empty_color
	_black_color
	_inventory_text_color
	_inventory_header_background_color
	_inventory_background_color
	PLAYER_COLOR_BASE_INDEX
)
const (
	_white_color = 14
	_invalid_weapon_color
	_computer_border_background_text_color
	_computer_border_text_color
	_computer_interface_text_color
	_computer_interface_color_purple
	_computer_interface_color_red
	_computer_interface_color_pink
	_computer_interface_color_aqua
	_computer_interface_color_yellow
	_computer_interface_color_brown
	_computer_interface_color_blue
)

const (
	/* justification flags for _draw_screen_text */
	_no_flags          = 0
	_center_horizontal = 0x01
	_center_vertical   = 0x02
	_right_justified   = 0x04
	_top_justified     = 0x08
	_bottom_justified  = 0x10
	_wrap_text         = 0x20
)

const (
	/* Fonts for the interface et al.. */
	_interface_font = iota
	_weapon_name_font
	_player_name_font
	_interface_item_count_font
	_computer_interface_font
	_computer_interface_title_font
	_net_stats_font
	NUMBER_OF_INTERFACE_FONTS
)

/* Structure for portable rectangles.  notice it is exactly same as Rect */
type screenRectangle struct {
	top, left     int16
	bottom, right int16
}
