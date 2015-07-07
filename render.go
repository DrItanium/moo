package moo

import "github.com/DrItanium/moo/cseries"

const (
	MinimumObjectDistance           int16 = int16(WorldOne / 20)
	MinimumVerticesPerScreenPolygon int16 = 3
	MaximumVerticesPerScreenPolygon int16 = 16
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

type ViewData struct {
	FieldOfView                    int16 /* width of the view cone, in degrees (!) */
	StandardScreenWidth            int16 /* this is *not* the width of the projected image (see initialize_view_data() in RENDER.C */
	ScreenWidth, ScreenHeight      int16 /* dimensions of the projected image */
	HorizontalScale, VerticalScale int16

	HalfScreenWidth, HalfScreenHeight int16
	WorldToScreenX, WorldToScreenY    int16
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

	Effect, EffectPhase                    int16
	RealWorldToScreenX, RealWorldToScreenY int16

	OverheadMapActive bool
	OverheadMapScale  int16

	UnderMediaBoundary bool
	UnderMediaIndex    int16

	TerminalModeActive bool
}

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

	PolygonIsVisible           = 1 << PolygonIsVisibleBit
	EndpointHasBeenVisited     = 1 << EndpointHasBeenVisitedBit
	EndpointIsVisible          = 1 << EndpointIsVisibleBit
	SideIsVisible              = 1 << SideIsVisibleBit
	LineHasClipData            = 1 << LineHasClipDataBit
	EndpointHasClipData        = 1 << EndpointHasClipDataBit
	EndpointHasBeenTransformed = 1 << EndpointHasBeenTransformedBit
)

var RenderFlags []cseries.Word

func AllocateRenderMemory() {

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
