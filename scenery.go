// scenery related operations

package moo

import (
	"github.com/DrItanium/moo/cseries"
)

const MaximumAnimatedSceneryObjects = 20

const (
	SceneryIsSolid        = 0x0001
	SceneryIsAnimated     = 0x0002
	SceneryCanBeDestroyed = 0x0004
)

type SceneryDefinition struct {
	Flags           cseries.Word
	Shape           ShapeDescriptor
	Radius, Height  WorldDistance
	DestroyedEffect int16
	DestroyedShape  ShapeDescriptor
}

var AnimatedSceneryObjectCount int16
var AnimatedSceneryObjectIndexes []int16

func init() {
	AnimatedSceneryObjectIndexes = make([]int16, MaximumAnimatedSceneryObjects)
}

func NewScenery(location *ObjectLocation, sceneryType int16) int16 {
	var objectIndex int16
	// TODO: body of new_scenery
	return objectIndex
}

func GetSceneryDefinition(sceneryType int16) *SceneryDefinition {
	return nil
}

func AnimateScenery() {
	for i := 0; i < AnimatedSceneryObjectCount; i++ {
		AnimateObject(AnimatedSceneryObjectIndexes[i])
	}
}

func RandomizeSceneryShapes() {
	//	var object ObjectData
	//	var objectIndex int16

	AnimatedSceneryObjectCount = 0

	//TODO: implement randomize_scenery_shapes

}

func GetSceneryDimensions(sceneryType int16) (*WorldDistance, *WorldDistance) {
	definition := GetSceneryDefinition(sceneryType)
	return &(definition.Radius), &(definition.Height)
}

func DamageScenery(objectIndex int16) {
	object := GetObjectData(objectIndex)
	definition := GetSceneryDefinition(object.permutation)
}
