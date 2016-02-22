// song playing
package moo

import "github.com/DrItanium/moo/cseries"

type songIndex int16

const (
	introductionSong songIndex = iota
	numberOfSongs
)

func initializeMusicHandler(songFile interface{}) bool {
	return false
}

func (this songIndex) queueSong() {

}

func musicIdleProc() {

}

func stopMusic() {

}

func pauseMusic(pause bool) {

}

func musicPlaying() bool {
	return false
}

func freeMusicChannel() {

}

func fadeOutMusic(duration int16) {

}

type musicStates int16

const (
	_no_song_playing musicStates = iota
	_playing_introduction
	_playing_chorus
	_playing_trailer
	_delaying_for_loop
	_music_fading
	NUMBER_OF_MUSIC_STATES
)

type musicData struct {
	initialized          bool
	songCompleted        bool
	songPaused           bool
	phase                int16
	fadeDuration         int16
	playCount            int16
	songIndex            int16
	nextSongIndex        int16
	songFileRefnum       int16
	fadeIntervalDuration int16
	fadeIntervalTicks    int16
	ticksAtLastUpdate    int32
	soundBuffer          []byte
	soundBufferSize      int32
	channel              interface{} // SndChannelPtr
	completionProc       interface{} // FilePlayCompletionUPP
}

const kDefaultSoundBufferSize = 500 * cseries.Kilo

var musicState []musicData

func shutdownMusicHandler() {

}

//static pascal void file_play_completion_routine(SndChannelPtr channel);
func allocateMusicChannel() {

}

func getSoundVolume() int16 {
	return 0
}
