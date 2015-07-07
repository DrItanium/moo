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

func RenderOverheadMap(view *ViewData) {

}

func RenderComputerInterface(view *ViewData) {

}

var HasAmbiguousFlags = false
var ExceededMaxNodeAliases = false
