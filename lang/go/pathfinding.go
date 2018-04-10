// pathfinding algorithms
package moo

const (
	maximumPaths           = 20
	maximumPointsPerPath   = 63
	pathValidationAreaSize = 64 * 1024
)

type pathDefinition struct {
	currentStep int16
	stepCount   int16
	points      [maximumPointsPerPath]WorldPoint2d
}

var paths []pathDefinition
