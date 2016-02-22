// lightsource ops
package moo

import "github.com/DrItanium/moo/cseries"

const MaximumLightsPerMap = 64

const (
	// default light types
	NormalLight = iota
	StrobeLight
	MediaLight
	NumberOfLightTypes
)

const (
	// states
	LightBecomingActive = iota
	LightPrimaryActive
	LightSecondaryActive
	LightBecomingInactive
	LightPrimaryInactive
	LightSecondaryInactive
)

const (
	// lighting functions
	ConstantLightingFunction = iota // maintain final intensity for period
	LinearLightingFunction          // linear transition between initial and final intensity over period
	SmoothLightingFunction          // sine transition between initial and final intensity over period
	FlickerLightingFunction         // intensity in [smooth_intensity(t), final_intensity]
	NumberOfLightingFunctions
)

/* as intensities, transition functions are given the primary periods of the active and inactive
state, plus the intensity at the time of transition */
type LightingFunctionSpecification struct {
	Function                  int16
	Period, DeltaPeriod       int16
	Intensity, DeltaIntensity cseries.Fixed
}

type StaticLightData struct {
	Type                 int16
	IsInitiallyActive    bool
	HasSlavedIntensities bool
	IsStateless          bool
	Phase                int16 // initializer, so lights may start out-of-phase with each other

	PrimaryActive, SecondaryActive, BecomingActive       LightingFunctionSpecification
	PrimaryInactive, SecondaryInactive, BecomingInactive LightingFunctionSpecification

	Tag int16
}

type LightData struct {
	Flags cseries.Word
	State int16

	// result of lighting function
	Intensity cseries.Fixed

	// data recalculated each function changed; passed to lighting_function each update

	Phase, Period                    int16
	InitialIntensity, FinalIntensity cseries.Fixed

	StaticData StaticLightData
}

var Lights []LightData

func NewLight(data *StaticLightData) int16 {
	return 0
}

func GetDefaultsForLightType(lType int16) *StaticLightData {
	return nil
}

func UpdateLights() {

}

func GetLightStatus(lightIndex int16) bool {
	return false
}

func SetLightStatus(lightIndex int16, active bool) bool {
	return false
}

func SetTaggedLightStatuses(tag int16, newStatus bool) bool {
	return false
}

func GetLightIntensity(lightIndex int16) cseries.Fixed {
	return 0
}
