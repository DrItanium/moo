// interface related code
package moo

import "github.com/DrItanium/moo/cseries"

// local game state objects
type InterfaceGameState struct {
	State                          int16
	Flags                          int16
	User                           int16
	Phase                          int32
	LastTicksOnIdle                int32
	CurrentScreen                  int16
	SuppressBackgroundTasks        bool
	CurrentNetgameAllowsMicrophone bool
	MainMenuDisplayCount           int16
}

type ScreenData struct {
	ScreenBase  int16
	ScreenCount int16
	Duration    int32
}

const TotalShapeCollections = 128

type ShapeInformationData struct {
	XMirrored                                    bool
	YMirrored                                    bool
	KeypointObscured                             bool
	MinimumLightIntensity                        cseries.Fixed
	WorldLeft, WorldRight, WorldTop, WorldBottom int16
	WorldX0, WorldY0                             int16
}

type ShapeAnimationData struct {
	NumberOfViews                                  int16 // must be 1, 2, 5, or 8
	FramesPerView, TicksPerFrame                   int16
	KeyFrame                                       int16
	TransferMode                                   int16
	TransferModePeriod                             int16 // in ticks
	FirstFrameSound, KeyFrameSound, LastFrameSound int16
	PixelsToWorld                                  int16
	LoopFrame                                      int16

	// Number of views * frames per view indexes of low-level shapes follow
	LowLevelShapeIndexes [1]int16
}

var NoFrameRateLimit bool

const (
	// controllers
	SinglePlayer = iota
	NetworkPlayer
	Demo
	Replay
	ReplayFromFile
	NumbeOfPseudoPlayers
)

const (
	// interface states
	DisplayIntroScreens = iota
	DisplayMainMenu
	DisplayChapterHeading
	DisplayPrologue
	DisplayEpilogue
	DisplayCredits
	DisplayIntroScreensForDemo
	DisplayQuitScreens
	NumberOfScreens
	GameInProgress = NumberOfScreens
	QuitGame
	CloseGame
	SwitchDemo
	RevertGame
	ChangeLevel
	BeginDisplayOfEpilogue
	DisplayingNetworkGameDialogs
	NumberOfGameStates
)
