// placement related operations
package moo

const (
	NumberOfTicksBetweenRecreation = 15 * TicksPerSecond
	InvisibleRandomPointRetries    = 10
)

var objectPlacementInfo [2 * MaximumObjectTypes]ObjectFrequencyDefinition
var monsterPlacementInfo []ObjectFrequencyDefinition
var itemPlacementInfo []ObjectFrequencyDefinition
