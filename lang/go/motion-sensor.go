// motion sensor imp

package moo

const (
	MaximumMotionSensorEntities = 12
	NumberOfPreviousLocations   = 6
	MotionSensorUpdateFrequency = 5
	MotionSensorRescanFrequency = 15
	MotionSensorRange           = 8 * WorldOne
	MotionScaleScale            = 7
	FlickerFrequency            = 0x0F

	SlotIsBeingRemovedBit = 0x4000
)

//#define ObjectIsVisibleToMotionSensor(o) TRUE
