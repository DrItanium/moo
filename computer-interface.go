// operations relating to computer interfaces
/*
	computer_interface.h
	Tuesday, August 23, 1994 11:25:40 PM (ajr)
	Thursday, May 25, 1995 5:18:03 PM- rewriting.

	New paradigm:
	Groups each start with one of the following groups:
	 #UNFINISHED, #SUCCESS, #FAILURE

	First is shown the
	#LOGON XXXXX

	Then there are any number of groups with:
	#INFORMATION, #CHECKPOINT, #SOUND, #MOVIE, #TRACK

	And a final:
	#INTERLEVEL TELEPORT, #INTRALEVEL TELEPORT

	Each group ends with:
	#END

	Groupings:
	#logon XXXX- login message (XXXX is shape for login screen)
	#unfinished- unfinished message
	#success- success message
	#failure- failure message
	#information- information
	#briefing XX- briefing, then load XX
	#checkpoint XX- Checkpoint xx (associated with goal)
	#sound XXXX- play sound XXXX
	#movie XXXX- play movie XXXX (from Movie file)
	#track XXXX- play soundtrack XXXX (from Music file)
	#interlevel teleport XXX- go to level XXX
	#intralevel teleport XXX- go to polygon XXX
	#pict XXXX- diplay the pict resource XXXX

	Special embedded keys:
	$B- Bold on
	$b- bold off
	$I- Italic on
	$i- italic off
	$U- underline on
	$u- underline off
	$- anything else is passed through unchanged

	Preprocessed format:
	static:
		long total_length;
		short grouping_count;
		short font_changes_count;
		short total_text_length;
	dynamic:
		struct terminal_groupings groups[grouping_count];
		struct text_face_data[font_changes_count];
		char text;
*/
package moo

//import "fmt"

type StaticPreprocessedTerminalData struct {
	TotalLength      int16
	Flags            int16
	LinesPerPage     int16
	GroupingCount    int16
	FontChangesCount int16
}

type ViewTerminalData struct {
	Top, Left, Bottom, Right int16
	VerticalOffset           int16
}

// this could become a structure or just a slice typedef
var MapTerminalData []byte
var MapTerminalDataLength int32

///* ------------ prototypes */
//void initialize_terminal_manager(void);
//void initialize_player_terminal_info(short player_index);
//void enter_computer_interface(short player_index, short text_number, short completion_flag);
//void _render_computer_interface(struct view_terminal_data *data);
//void update_player_for_terminal_mode(short player_index);
//void update_player_keys_for_terminal(short player_index, long action_flags);
//long build_terminal_action_flags(char *keymap);
//void dirty_terminal_view(short player_index);
//void abort_terminal_mode(short player_index);
//
//boolean player_in_terminal_mode(short player_index);
//
//void *get_terminal_data_for_save_game(void);
//long calculate_terminal_data_length(void);
//
///* This returns the text.. */
//void *get_terminal_information_array(void);
//long calculate_terminal_information_length(void);
//
//#ifdef PREPROCESSING_CODE
//struct static_preprocessed_terminal_data *preprocess_text(char *text, short length);
//struct static_preprocessed_terminal_data *get_indexed_terminal_data(short id);
//void encode_text(struct static_preprocessed_terminal_data *terminal_text);
//void decode_text(struct static_preprocessed_terminal_data *terminal_text);
//void find_all_picts_references_by_terminals(byte *compiled_text, short terminal_count,
//	short *picts, short *picture_count);
//void find_all_checkpoints_references_by_terminals(byte *compiled_text,
//	short terminal_count, short *checkpoints, short *checkpoint_count);
//boolean terminal_has_finished_text_type(short terminal_id, short finished_type);
//#endif

// Macro functions
//#define TERMINAL_IS_DIRTY(term) ((term)->flags & _terminal_is_dirty)
//#define SET_TERMINAL_IS_DIRTY(term, v) ((v)? ((term)->flags |= _terminal_is_dirty) : ((term)->flags &= ~_terminal_is_dirty))
const (
	LabelInset               = 3
	LogDurationBeforeTimeout = 2 * TicksPerSecond
	BorderHeight             = 10
	BorderInset              = 9
	FudgeFactor              = 1
	// terminal-states
	ReadingTerminal = iota
	NoTerminalState
	TerminalStateCount
	// terminal flags
	TerminalIsDirty = 0x01
	// terminal keys
	//enum {
	//	_any_abort_key_mask= _action_trigger_state,
	//	_terminal_up_arrow= _moving_forward,
	//	_terminal_down_arrow= _moving_backward,
	//	_terminal_page_down= _turning_right,
	//	_terminal_page_up= _turning_left,
	//	_terminal_next_state= _left_trigger_state
	//};
	strComputerLabels = 135
	// computer tags
	MarathonName = iota
	ComputerStartingUp
	ComputerManufacturer
	ComputerAddress
	ComputerTerminal
	ScrollingMessage
	AcknowledgementMessage
	DisconnectingMessage
	ConnectionTerminatedMessage
	DateFormat
	// Macros
	MaximumFaceChangesPerTextGrouping = 128
	// text flags?
	TextIsEncodedFlag = 0x0001

	LogonGroup = iota
	UnfinishedGroup
	SuccessGroup
	FailureGroup
	InformationGroup
	EndGroup
	Interlevel_teleportGroup // permutation is level to go to
	Intralevel_teleportGroup // permutation is polygon to go to.
	CheckpointGroup          // permutation is the goal to show
	SoundGroup               // permutation is the sound id to play
	MovieGroup               // permutation is the movie id to play
	TrackGroup               // permutation is the track to play
	PictGroup                // permutation is the pict to display
	LogoffGroup
	CameraGroup //  permutation is the object index
	StaticGroup // permutation is the duration of static.
	TagGroup    // permutation is the tag to activate
	NumberOfGroupTypes
	// flags to indicate text styles for paragraphs
	PlainText     = 0x00
	BoldText      = 0x01
	ItalicText    = 0x02
	UnderlineText = 0x04

	/* terminal grouping flags */
	DrawObjectOnRight = 0x01 // for drawing checkpoints, picts, movies.
	CenterObject      = 0x02
)

type TerminalGroupings struct {
	Flags            int16
	Type             int16
	Permutation      int16
	StartIndex       int16
	Length           int16
	MaximumLineCount int16
}

type TextFaceData struct {
	Index int16
	Face  int16
	Color int16
}

type PlayerTerminalData struct {
	Flags                int16
	Phase                int16
	State                int16
	CurrentGroup         int16
	LevelCompletionState int16
	CurrentLine          int16
	MaximumLine          int16
	TerminalId           int16
	LastActionFlag       int32
}

type TerminalKey struct {
	Keycode    int16
	Offset     int16
	Mask       int16
	ActionFlag int32
}

type FontDimensions struct {
	LinesPerScreen int16
	CharacterWidth int16
}
type playerTerminals []PlayerTerminalData

var PlayerTerminals playerTerminals
var fontData FontDimensions

// TODO: implement this and the corresponding enum
//var terminalKeys = []TerminalKey{
//	{0x7e, 0, 0, _terminal_page_up},    // arrow up
//	{0x7d, 0, 0, _terminal_page_down},  // arrow down
//	{0x74, 0, 0, _terminal_page_up},    // page up
//	{0x79, 0, 0, _terminal_page_down},  // page down
//	{0x30, 0, 0, _terminal_next_state}, // tab
//	{0x4c, 0, 0, _terminal_next_state}, // enter
//	{0x24, 0, 0, _terminal_next_state}, // return
//	{0x31, 0, 0, _terminal_next_state}, // space
//	{0x3a, 0, 0, _terminal_next_state}, // command
//	{0x35, 0, 0, _any_abort_key_mask},  // escape
//}

//func (playerTerminals) Data(index int16) (*PlayerTerminalData, error) {
//	if Debug && (index < 0 || index >= MaximumNumberOfPlayers) {
//		return nil, fmt.Errorf("PlayerTerminals.Data: index (%d) is out of range!", index)
//	} else {
//		return &(PlayerTerminals[index]), nil
//	}
//}

/* ------------ private prototypes */
//static void draw_logon_text(Rect *bounds, struct static_preprocessed_terminal_data *terminal_text,
//	short current_group_index, short logon_shape_id);
//static void draw_computer_text(Rect *bounds,
//	struct static_preprocessed_terminal_data *terminal_text, short current_group_index, short current_line);
//static void _draw_computer_text(char *base_text, short start_index, Rect *bounds,
//	struct static_preprocessed_terminal_data *terminal_text, short current_line);
//static void render_terminal(short player_index, struct view_terminal_data *data);
//static short find_group_type(struct static_preprocessed_terminal_data *data,
//	short group_type);
//static void teleport_to_level(short level_number);
//static void teleport_to_polygon(short player_index, short polygon_index);
//static struct terminal_groupings *get_indexed_grouping(
//	struct static_preprocessed_terminal_data *data, short index);
//static struct text_face_data *get_indexed_font_changes(
//	struct static_preprocessed_terminal_data *data, short index);
//static char *get_text_base(struct static_preprocessed_terminal_data *data);
//static void draw_terminal_borders(struct view_terminal_data *data,
//	struct player_terminal_data *terminal_data, Rect *terminal_frame);
//static void next_terminal_state(short player_index);
//static void next_terminal_group(short player_index, struct static_preprocessed_terminal_data *terminal_text);
//static get_date_string(char *date_string);
//static void present_checkpoint_text(Rect *frame,
//	struct static_preprocessed_terminal_data *terminal_text, short current_group_index,
//	short current_line);
//static boolean find_checkpoint_location(short checkpoint_index, world_point2d *location,
//	short *polygon_index);
//struct static_preprocessed_terminal_data *preprocess_text(char *text, short length);
//static void pre_build_groups(struct terminal_groupings *groups,
//	short *group_count, struct text_face_data *text_faces, short *text_face_count,
//	char *base_text, short *base_length);
//static short matches_group(char *base_text, short length, short index, short possible_group,
//	short *permutation);
//static void	set_text_face(struct text_face_data *text_face);
//static void draw_line(char *base_text, short start_index, short end_index, Rect *bounds,
//	struct static_preprocessed_terminal_data *terminal_text, short *text_face_start_index,
//	short line_number);
//static boolean calculate_line(char *base_text, short width, short start_index,
//	short text_end_index, short *end_index);
//static void handle_reading_terminal_keys(short player_index, long action_flags);
//static void calculate_bounds_for_object(Rect *frame, short flags, Rect *bounds, Rect *source);
//static void display_picture(short picture_id, Rect *frame, short flags);
//static void display_picture_with_text(struct player_terminal_data *terminal_data,
//	Rect *bounds, struct static_preprocessed_terminal_data *terminal_text, short current_lien);
//static short count_total_lines(char *base_text, short width, short start_index, short end_index);
//static void calculate_bounds_for_text_box(Rect *frame, short flags, Rect *bounds);
//static void goto_terminal_group(short player_index, struct static_preprocessed_terminal_data *terminal_text,
//	short new_group_index);
//static boolean previous_terminal_group(short player_index, struct static_preprocessed_terminal_data *terminal_text);
//static void fill_terminal_with_static(Rect *bounds);
//static short calculate_lines_per_page(void);
//
//#ifndef PREPROCESSING_CODE
//static struct static_preprocessed_terminal_data *get_indexed_terminal_data(short id);
//static void encode_text(struct static_preprocessed_terminal_data *terminal_text);
//static void decode_text(struct static_preprocessed_terminal_data *terminal_text);
//#endif
