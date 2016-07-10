package moo

import (
	"fmt"
	"github.com/DrItanium/moo/cseries"
	"math"
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
)
const (
	// shading tables
	ShadingNormal      = iota // to black
	ShadingInfravision        // false color
)
const (
	// macro constants
	NormalFieldOfView      = 80
	ExtravisionFieldOfView = 130
)
const (
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
)
const (
	RenderFlagsBufferSize = 8 * cseries.Kilo
	// from render.c
	MaxPolygonQueueSize             = 256
	MaximumVerticiesPerWorldPolygon = MaximumVerticesPerPolygon + 4
	ExplosionEffectRange            = WorldOne / 12
	ClipIndexBufferSize             = 4096

	MaximumLineClips       = 256
	MaximumEndpointClips   = 64
	MaximumClippingWindows = 256
)

const (

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

type Coordinate struct {
	X int16
	Y int16
}

type Point2d struct {
	X int16
	Y int16
}

type DefinitionHeader struct {
	Tag       int16
	ClipLeft  int16
	ClipRight int16
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

type RenderFlags []cseries.Word

func (this RenderFlags) Test(index, flag int16) bool {
	return (this[index] & cseries.Word(flag)) == 1
}
func (this RenderFlags) Set(index, flag int16) {
	this[index] |= cseries.Word(flag)
}

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

var HasAmbiguousFlags = false
var ExceededMaxNodeAliases = false

// in screen.c/h

type renderState struct {
	Flags                    RenderFlags
	Nodes                    []NodeData
	PolygonQueue             []int16
	PolygonQueueIndex        int16
	SortedNodes              []SortedNodeData
	RenderObjects            []RenderObjectData
	EndpointClips            []EndpointClipData
	LineClips                []LineClipData
	LineClipIndexes          []int16
	ClippingWindows          []ClippingWindowData
	EndpointXCoordinates     []ClippingWindowData
	PolygonIndexToSortedNode []*SortedNodeData
}

var renderGlobals = renderState{}

func AllocateRenderMemory() error {
	if NumberOfRenderFlags > 16 {
		return fmt.Errorf("AllocateRenderMemory: too many render flags!")
	} else if MaximumLinesPerMap > RenderFlagsBufferSize {
		return fmt.Errorf("AllocateRenderMemory: MaximumLinesPerMap > RenderFlagsBufferSize")
	}
	// add more asserts
	renderGlobals.Flags = make([]cseries.Word, RenderFlagsBufferSize)
	// assert
	renderGlobals.Nodes = make([]NodeData, MaximumNodes)
	renderGlobals.PolygonQueue = make([]int16, MaxPolygonQueueSize)
	renderGlobals.PolygonQueueIndex = 0
	renderGlobals.SortedNodes = make([]SortedNodeData, MaximumSortedNodes)
	renderGlobals.RenderObjects = make([]RenderObjectData, MaximumRenderObjects)
	renderGlobals.EndpointClips = make([]EndpointClipData, MaximumEndpointClips)
	renderGlobals.LineClips = make([]LineClipData, MaximumLineClips)
	renderGlobals.LineClipIndexes = make([]int16, MaximumLinesPerMap)
	renderGlobals.ClippingWindows = make([]ClippingWindowData, MaximumClippingWindows)
	renderGlobals.EndpointXCoordinates = make([]ClippingWindowData, MaximumEndpointsPerMap)
	renderGlobals.PolygonIndexToSortedNode = make([]*SortedNodeData, MaximumPolygonsPerMap)
	return nil
}

func (view *ViewData) Initialize() {
	twoPi := 8.0 * math.Atan(1.0)
	halfCone := float64(view.FieldOfView) * float64(twoPi/360.0) / 2.0
	adjustedHalfCone := math.Asin(float64(view.ScreenWidth) * math.Sin(halfCone) / float64(view.StandardScreenWidth))
	var worldToScreen float64

	view.HalfScreenWidth = view.ScreenWidth / 2
	view.HalfScreenHeight = view.ScreenHeight / 2

	// if there's a round-off error in half_cone, we want to make the cone too big (so when we clip lines to the edge of the screen they're actually off the screen, thus +1.0)
	view.HalfCone = Angle(adjustedHalfCone*(float64(NumberOfAngles))/twoPi + 1.0)

	// calculate world_to_screen; we could calculate this with standard_screen_width/2 and the old half_cone and get the same result
	worldToScreen = float64(view.HalfScreenWidth) / math.Tan(adjustedHalfCone)
	tmp0 := int16((worldToScreen / float64(view.HorizontalScale)) + 0.5)
	tmp1 := int16((worldToScreen / float64(view.VerticalScale)) + 0.5)
	view.WorldToScreen.X = tmp0
	view.RealWorldToScreen.X = tmp0
	view.WorldToScreen.Y = tmp1
	view.RealWorldToScreen.Y = tmp1

	// cacluate the vertical cone angle; again, overflow instead of underflow when rounding
	view.HalfVerticalCone = Angle(NumberOfAngles*math.Atan((float64(view.HalfScreenHeight*view.VerticalScale)/worldToScreen))/twoPi + 1.0)

	// calculate left edge vector
	view.UntransformedLeftEdge.I = WorldDistance(view.WorldToScreen.X)
	view.UntransformedLeftEdge.J = WorldDistance(-view.HalfScreenWidth)

	// calculate right edge vector (negative, so it clips in the right direction)
	view.UntransformedRightEdge.I = WorldDistance(-view.WorldToScreen.X)
	view.UntransformedRightEdge.J = WorldDistance(-view.HalfScreenWidth)

	// reset any effects
	view.Effect = cseries.None
}
func (view *ViewData) RenderView(destination *BitmapDefinition) {
	view.UpdateViewData()

	// clear the render flags
	for i := 0; i < len(renderGlobals.Flags); i++ {
		renderGlobals.Flags[i] = 0
	}

	if view.TerminalModeActive {
		// render the computer interface
		view.RenderComputerInterface()
	} else {
		// build the render tree, regardless of map node, so the automap updates while active
		view.BuildRenderTree()
		if view.OverheadMapActive {
			view.RenderOverheadMap()
		} else {
			// do something complicated and difficult to explain
			// sor the render tree (so we have a depth-ordering of polygons) and accumulate clipping information for each polygon
			view.SortRenderTree()

			// build the render object list by looking at the sorted render tree
			view.BuildRenderObjectList()

			// render the object list, back to front, doing clipping on each surface before passing it to the texture-mapping code
			view.RenderTree(destination)

			// render the player's weapons, etc..
			view.RenderViewerSpriteLayer(destination)
		}
	}
}

func (view *ViewData) StartRenderEffect(effect int16) {
	view.Effect = effect
	view.EffectPhase = cseries.None
}
func WrapLow16(n, max int16) int16 {
	if n != 0 {
		return n - 1
	} else {
		return max
	}
}
func WrapHigh16(n, max int16) int16 {
	if n == max {
		return 0
	} else {
		return n + 1
	}
}

func (this *renderState) PushPolygonIndex(polygonIndex int16) {
	if !this.Flags.Test(polygonIndex, PolygonIsVisible) {
		this.PolygonQueue[this.PolygonQueueIndex] = polygonIndex
		this.PolygonQueueIndex++
		this.Flags.Set(polygonIndex, PolygonIsVisible)
	}
}

func (view *ViewData) UpdateViewData() {
	if view.Effect == cseries.None {
		view.WorldToScreen = view.RealWorldToScreen
	} else {
		view.UpdateRenderEffect()
	}

	view.UntransformedLeftEdge.I = WorldDistance(view.WorldToScreen.X)
	view.UntransformedRightEdge.I = WorldDistance(-view.WorldToScreen.X)
	// calculate worldToScreen.Y * tan(pitch)
	view.Dtanpitch = int16(view.WorldToScreen.Y*SineTable[view.Pitch]) / int16(CosineTable[view.Pitch])

	// calculate left cone vector
	theta := NormalizeAngle(view.Yaw - view.HalfCone)
	view.RightEdge.I = WorldDistance(CosineTable[theta])
	view.RightEdge.J = WorldDistance(SineTable[theta])

	// calculate top cone vector (negative to clip the right direction)
	view.TopEdge.I = WorldDistance(-view.WorldToScreen.Y)
	view.TopEdge.J = WorldDistance(-(view.HalfScreenHeight + view.Dtanpitch)) // == k

	// calculate bottom cone vector
	view.BottomEdge.I = WorldDistance(view.WorldToScreen.Y)
	view.BottomEdge.J = WorldDistance(view.HalfScreenHeight + view.Dtanpitch) // == k

	// if we're sitting on one of the endpoints in our origin polygon, move us back slightly (+/- 1) into
	// that polygon. When we split rays we're assuming that we'll never pass through a given vertex in
	// different directions (because if we do the tree becomes a graph) but when we start on a vertex
	// this can happen. This is a destructive modification of the origin
	polygon := GetPolygonData(view.OriginPolygonIndex)
	for i := int16(0); i < polygon.VertexCount; i++ {
		vertex := GetEndpointData(polygon.EndpointIndexes[i]).Vertex
		if vertex.X == view.Origin.X && vertex.Y == view.Origin.Y {
			ccwVertex := GetEndpointData(polygon.EndpointIndexes[WrapLow16(i, polygon.VertexCount-1)]).Vertex
			cwVertex := GetEndpointData(polygon.EndpointIndexes[WrapHigh16(i, polygon.VertexCount-1)]).Vertex
			var insetVector WorldVector2d
			insetVector.I = (ccwVertex.X - vertex.X) + (cwVertex.X - vertex.X)
			insetVector.J = (ccwVertex.Y - vertex.Y) + (cwVertex.Y - vertex.Y)
			view.Origin.X += WorldDistance(cseries.Signum(int64(insetVector.I)))
			view.Origin.Y += WorldDistance(cseries.Signum(int64(insetVector.J)))
			break
		}
	}

	// determine whether we are under or over the media boundary of our polygon; we will see all
	// other media boundaries from this orientation (above or below) or fail to draw them.
	if polygon.MediaIndex == cseries.None {
		view.UnderMediaBoundary = false
	} else {
		media := GetMediaData(polygon.MediaIndex)
		view.UnderMediaBoundary = media.UnderMedia(view.Origin.Z)
		view.UnderMediaIndex = polygon.MediaIndex
	}

}

func (view *ViewData) RenderOverheadMap() {

}

func (view *ViewData) RenderComputerInterface() {

}

func (view *ViewData) BuildRenderTree() {

}

func (view *ViewData) SortRenderTree() {

}
func (view *ViewData) BuildRenderObjectList() {

}

func (view *ViewData) RenderTree(destination *BitmapDefinition) {
}

func (view *ViewData) RenderViewerSpriteLayer(destination *BitmapDefinition) {

}

func (view *ViewData) UpdateRenderEffect() {

}
