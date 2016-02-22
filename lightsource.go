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

const (
	// static flags
	LightIsInitiallyActive = iota
	LightHasSlavedIntensities
	LightIsStateless
	NumberOfStaticLightFlags // <=16
)

func init() {
	if NumberOfStaticLightFlags > 16 {
		panic("too many static light flag states defined!")
	}
}
