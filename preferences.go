// preferences
package moo

/* New preferences junk */
type graphicsPreferencesData struct {
	screenMode            interface{} //screenModeData
	doResolutionSwitching bool
	deviceSpec            interface{}
}

type serial_number_data struct {
	networkOnly           bool
	longSerialNumber      [10]byte
	userName              string
	tokenizedSerialNumber string
}

type network_preferences_data struct {
	allowMicrophone bool
	gameIsUntimed   bool
	typ             int16 // look in network_dialogs.c for _ethernet, etc...
	gameType        int16
	difficultyLevel GameDifficultyLevel
	gameOptions     GameOptions // Penalize suicide, etc... see map.h for constants
	timeLimit       int32
	killLimit       int16
	entryPoint      int16
}

type player_preferences_data struct {
	name              string
	color             int16
	team              int16
	lastTimeRan       uint32
	difficultyLevel   GameDifficultyLevel
	backgroundMusicOn bool
}

type input_preferences_data struct {
	inputDevice int16
	keycodes    []int16 // [NUMBER_OF_KEYS]int16
}

const maximumPatchesPerEnvironment = 32

type environment_preferences_data struct {
	mapFile     interface{}
	physicsFile interface{}
	shapesFile  interface{}
	soundsFile  interface{}

	map_checksum     uint32
	physics_checksum uint32
	shapes_mod_date  uint32
	sounds_mod_date  uint32
	patches          [maximumPatchesPerEnvironment]uint32
}

/* New preferences.. (this sorta defeats the purpose of this system, but not really) */
var graphics_preferences []graphicsPreferencesData
var serial_preferences []serial_number_data
var network_preferences []network_preferences_data
var player_preferences []player_preferences_data
var input_preferences []input_preferences_data

//var sound_preferences []sound_manager_parameters
var environment_preferences []environment_preferences_data

/* --------- functions */
func initializePreferences() {

}

func handlePreferences() {

}

func writePreferences() {

}
