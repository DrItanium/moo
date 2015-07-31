package moo

import (
	"fmt"
	"github.com/DrItanium/moo/cseries"
)

const (
	MinimumObjectDistance           int16 = int16(WorldOne / 20)
	MinimumVerticesPerScreenPolygon int16 = 3
	MaximumVerticesPerScreenPolygon int16 = 16
)
const (
	// render effects
	RenderEffectFoldIn = iota
	RenderEffectFoldOut
	RenderEffectGoingFisheye
	RenderEffectLeavingFisheye
	RenderEffectExplosion
	// shading tables
	ShadingNormal      = iota // to black
	ShadingInfravision        // false color
	// macro constants
	NormalFieldOfView      = 80
	ExtravisionFieldOfView = 130
	// render flags
	PolygonIsVisibleBit           = iota /* some part of this polygon is horizontally in the view cone */
	EndpointHasBeenVisitedBit            /* we've already tried to cast a ray out at this endpoint */
	EndpointIsVisibleBit                 /* this endpoint is horizontally in the view cone */
	SideIsVisibleBit                     /* this side was crossed while building the tree and should be drawn */
	LineHasClipDataBit                   /* this line has a valid clip entry */
	EndpointHasClipDataBit               /* this endpoint has a valid clip entry */
	EndpointHasBeenTransformedBit        /* this endpoint has been transformed into screen-space */
	NumberOfRenderFlags                  /* should be <=16 */
)
const (
	PolygonIsVisible           = 1 << PolygonIsVisibleBit
	EndpointHasBeenVisited     = 1 << EndpointHasBeenVisitedBit
	EndpointIsVisible          = 1 << EndpointIsVisibleBit
	SideIsVisible              = 1 << SideIsVisibleBit
	LineHasClipData            = 1 << LineHasClipDataBit
	EndpointHasClipData        = 1 << EndpointHasClipDataBit
	EndpointHasBeenTransformed = 1 << EndpointHasBeenTransformedBit

	RenderFlagsBufferSize = 8 * cseries.Kilo
	// from render.c
	PolygonQueueSize                = 256
	MaximumVerticiesPerWorldPolygon = MaximumVerticesPerPolygon + 4
	ExplosionEffectRange            = WorldOne / 12
	ClipIndexBufferSize             = 4096

	// clip data flags
	ClipLeftFlag  = 0x0001
	ClipRightFlag = 0x0002
	ClipUpFlag    = 0x0003
	ClipDownFlag  = 0x0004
)
const (
	// left and right sides of screen
	Index_LeftSideOfScreen = iota
	Index_RightSideOfScreen
	NumberOfInitialEndPointClips
)
const (
	// top and bottom sides of screen
	Index_TopAndBottomOfScreen = iota
	NumberOfInitialLineClips
)
const (
	MaximumNodes                    = 512
	MaximumClippingEndpointsPerNode = 4
	MaximumClippingLinesPerNode     = MaximumVerticesPerPolygon - 2

	MaximumSortedNodes = 128

	MaximumRenderObjects = 72
)

type Point2d struct {
	X int16
	Y int16
}

type DefinitionHeader struct {
	Tag       int16
	ClipLeft  int16
	ClipRight int16
}

type Coordinate struct {
	X int16
	Y int16
}
type ViewData struct {
	FieldOfView                    int16 /* width of the view cone, in degrees (!) */
	StandardScreenWidth            int16 /* this is *not* the width of the projected image (see initialize_view_data() in RENDER.C */
	ScreenWidth, ScreenHeight      int16 /* dimensions of the projected image */
	HorizontalScale, VerticalScale int16

	HalfScreenWidth, HalfScreenHeight int16
	WorldToScreen                     Coordinate
	Dtanpitch                         int16 /* world_to_screen*tan(pitch) */
	HalfCone                          Angle /* often ==field_of_view/2 (when screen_width==standard_screen_width) */
	HalfVerticalCone                  Angle

	UntransformedLeftEdge, UntransformedRightEdge WorldVector2d
	LeftEdge, RightEdge, TopEdge, BottomEdge      WorldVector2d

	TicksElapsed          int16
	TickCount             int32 /* for effects and transfer modes */
	OriginPolygonIndex    int16
	Yaw, Pitch, Roll      Angle
	Origin                WorldPoint3d
	MaximumDepthIntensity cseries.Fixed /* in fixed units */

	ShadingMode int16

	Effect, EffectPhase int16
	RealWorldToScreen   Coordinate

	OverheadMapActive bool
	OverheadMapScale  int16

	UnderMediaBoundary bool
	UnderMediaIndex    int16

	TerminalModeActive bool
}

type renderFlags []cseries.Word

func (this renderFlags) Test(index, flag int16) bool {
	return (this[index] & cseries.Word(flag)) == 1
}
func (this renderFlags) Set(index, flag int16) {
	this[index] |= cseries.Word(flag)
}

var RenderFlags renderFlags

type FlaggedWorldPoint2d struct {
	WorldPoint2d
	Flags cseries.Word
}

type FlaggedWorldPoint3d struct {
	WorldPoint3d
	Flags cseries.Word
}

/* it's not worth putting this into the side_data structure, although the transfer mode should
be in the side_texture_definition structure */
type VerticalSurfaceData struct {
	LightSourceIndex int16
	AmbientDelta     cseries.Fixed /* a delta to the lightsourceÕs intensity, then pinned to [0,FIXED_ONE] */

	Length       WorldDistance
	H0, H1, Hmax WorldDistance /* h0<h1; hmax<=h1 and is the height where this wall side meets the ceiling */
	P0, P1       WorldPoint2d  /* will transform into left, right points on the screen (respectively) */

	TextureDefinition *SideTextureDefinition
	TransferMode      int16
}

type EndpointClipData struct {
	Flags  cseries.Word
	X      int16
	Vector WorldVector2d
}

type LineClipData struct {
	Flags  cseries.Word
	X0, X1 int16 /* clipping bounds */

	TopVector, BottomVector WorldVector2d /* viewer-space */
	TopY, BottomY           int16         /* screen-space */
}

type ClippingWindowData struct {
	Left, Right, Top, Bottom WorldVector2d /* j is really k for top and bottom */
	X0, X1, Y0, Y1           int16
	NextWindow               *ClippingWindowData
}

func AllocateRenderMemory() error {
	if NumberOfRenderFlags > 16 {
		return fmt.Errorf("AllocateRenderMemory: too many render flags!")
	}
	return nil
}

type NodeData struct {
	Flags                 cseries.Word
	PolygonIndex          int16
	ClippingEndpointCount int16
	ClippingEndpoints     [MaximumClippingLinesPerNode]int16
	ClippingLineCount     int16
	ClippingLines         [MaximumClippingLinesPerNode]int16
	Parent                *NodeData
	Reference             **NodeData
	Siblings              *NodeData
	Children              []NodeData
}

func NewNodeData(polygonIndex int16, flags cseries.Word, parent *NodeData, reference **NodeData) *NodeData {
	return &NodeData{
		Flags:                 flags,
		PolygonIndex:          polygonIndex,
		ClippingEndpointCount: 0,
		ClippingLineCount:     0,
		Parent:                parent,
		Reference:             reference,
		Siblings:              nil,
	}
}

type SortedNodeData struct {
	PolygonIndex    int16
	InteriorObjects []RenderObjectData
	ExteriorObjects []RenderObjectData
	ClippingWindows []ClippingWindowData
}

type RenderObjectData struct {
	Node            *SortedNodeData
	ClippingWindows []ClippingWindowData
	Next            *RenderObjectData
	Rectangle       RectangleDefinition
	Ymedia          int16
}

func InitializeViewData(view *ViewData) {

}

func RenderView(view *ViewData, destination *BitmapDefinition) {

}

func StartRenderEffect(view *ViewData, effect int16) {

}

// in screen.c/h
func RenderOverheadMap(view *ViewData) {

}

func RenderComputerInterface(view *ViewData) {

}

var HasAmbiguousFlags = false
var ExceededMaxNodeAliases = false
