// overhead map related functions
package moo

//import "github.com/DrItanium/moo/cseries"

const (
	overheadMapMinimumScale = 1
	overheadMapMaximumScale = 4
	defaultOverheadMapScale = 3
)

type overheadMapMode int16

const (
	// modes
	_rendering_saved_game_preview overheadMapMode = iota
	_rendering_checkpoint_map
	_rendering_game_map
)

type overheadMapData struct {
	mode                  overheadMapMode
	scale                 int16
	origin                WorldPoint2d
	originPolygonIndex    int16
	halfWidth, halfHeight int16
	width, height         int16
	top, left             int16
	drawEverything        bool
}

func initializeOverheadMap() {

}

func renderOverheadMap(data *overheadMapData) {

}

type polygonColor int16

const (
	/* polygon colors */
	_polygon_color polygonColor = iota
	_polygon_platform_color
	_polygon_water_color
	_polygon_lava_color
	_polygon_goo_color
	_polygon_sewage_color
	_polygon_hill_color
)

type lineColor int16

const (
	/* line colors */
	_solid_line_color lineColor = iota
	_elevation_line_color
	_control_panel_line_color
)

type thingColor int16

const (
	/* thing colors */
	_civilian_thing thingColor = iota
	_item_thing
	_monster_thing
	_projectile_thing
	_checkpoint_thing
	NUMBER_OF_THINGS
)

const (
	_rectangle_thing = iota
	_circle_thing
)

const (
	/* render flags */
	_endpoint_on_automap = 0x2000
	_line_on_automap     = 0x4000
	_polygon_on_automap  = 0x8000
)
