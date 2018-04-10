// wad related functions
package moo

const (
	PreEntryPointWadfileVersion = 0
	WadfileHasDirectoryEntry    = 1
	WadfileSupportsOverlays     = 2
	CurrentWadfileVersion       = WadfileSupportsOverlays

	InfintyWadfileVersion = CurrentWadfileVersion + 2

	MaximumDirectoryEntriesPerFile = 64
	MaximumWadfileNameLength       = 64
	MaximumUnionWadfiles           = 16
	MaximumOpenWadfiles            = 3
)

type WadDataType uint32

type WadHeader struct {
	Version                              int16
	DataVersion                          int16
	Filename                             string
	Checksum                             uint32
	DirectoryOffset                      int32
	WadCount                             int16
	ApplicationSpecificDirectoryDataSize int16
	EntryHeaderSize                      int16
	ParentChecksum                       uint32
	unused                               [20]int16 // TODO: remove this, originally for alignment purposes in the old code
}

type oldDirectoryEntry struct {
	OffsetToStart int32 // from start of file
	Length        int32 // of total level
}
type DirectoryEntry struct {
	OffsetToStart int32 // from start of file
	Length        int32 // of total level
	Index         int16 // For inplace modification of the wadfile!
}

type oldEntryHeader struct {
	Tag        WadDataType
	NextOffset int32 // From current file location-> ie directory_entry.offset_to_start+next_offset
	Length     int32 // of entry

	// element size?

	// data follows
}

type EntryHeader struct {
	Tag        WadDataType
	NextOffset int32 // from current file location-> ie directory_entry.offset_to_start+next_offset
	Length     int32 // of entry
	Offset     int32 // offset for inplace expansion of data

	// element size?

	// data follows
}

// memory data structures
type TagData struct {
	Tag    WadDataType // what type of data is this?
	Data   []byte      // offset into the wad
	Length int32       // length of the data
	Offset int32       // offset for patches
}

type WadData struct {
	TagCount     int16 // tag count
	Padding      int16
	ReadOnlyData []byte    // if this is non NULL, we are read-only
	TagData      []TagData // Tag data array
}
