// flood map algorithms
package moo

import "github.com/DrItanium/moo/cseries"

type FloodModes cseries.Word

const (
	// flood modes
	FloodDepthFirst          FloodModes = iota // unsupported
	FloodBreadthFirst                          // significantly faster than _best_first for large domains
	FloodFlaggedBreadthFirst                   // user data is interpreted as a long * to 4 bytes of flags
	FloodBestFirst
)

//typedef long (*cost_proc_ptr)(short source_polygon_index, short line_index, short destination_polygon_index, void *caller_data);
type CostProc func(sourcePolygonIndex, lineIndex, destinationPolygonIndex int16, callerData interface{}) int32

type FloodMapNodeData struct {
	Flags           cseries.Word
	ParentNodeIndex int16 /* node index of the node we came from to get here; only used for backtracking */
	PolygonIndex    int16 /* index of this polygon */
	Cost            int32 /* the cost to evaluate this entry */
	Depth           int16
	UserFlags       int32
}

var nodeCount int16 = 0
var lastNodeIndexExpanded int16 = 0

var nodes []FloodMapNodeData
var visitedPolygons []int16

func AllocatePathfindingMemory() {

}

func ResetPath() {

}

func NewPath(sourcePoint *WorldPoint2d, soucePolygonIndex int16, destinationPoint *WorldPoint2d, destinationPolygonIndex int16, minimumSeparation WorldDistance, cost CostProc, data interface{}) int16 {
	return 0
}

func MoveAlongPath(pathIndex int16, p *WorldPoint2d) bool {
	return false
}

func DeletePath(pathIndex int16) {

}

func AllocateFloodMapMemory() {

}

/* default cost_proc, NULL, is the area of the destination polygon and is significantly faster
than supplying a user procedure */
func FloodMap(firstPolygonIndex int16, maximumCost int32, costProc CostProc, floodMode FloodModes, callerData interface{}) int16 {
	return 0
}

func ReverseFloodMap() int16 {
	return 0
}

func FloodDepth() int16 {
	return 0
}

func ChooseRandomFloodNode(bias *WorldVector2d) {

}
