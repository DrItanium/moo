// fading related stuffs
package moo

import (
	"fmt"
	"github.com/DrItanium/moo/cseries"
	"math"
)

const (
	NumberOfGammaLevels = 8
	DefaultGammaLevel   = 2

	// fade types

	StartCinematicFadeIn = iota /* force all colors to black immediately */
	CinematicFadeIn             /* fade in from black */
	LongCinematicFadeIn
	CinematicFadeOut    /* fade out from black */
	EndCinematicFadeOut /* force all colors from black immediately */

	FadeRed             /* bullets and fist */
	FadeBigRed          /* bigger bullets and fists */
	FadeBonus           /* picking up items */
	FadeBright          /* teleporting */
	FadeLongBright      /* nuclear monster detonations */
	FadeYellow          /* explosions */
	FadeBigYellow       /* big explosions */
	FadePurple          /* ? */
	FadeCyan            /* fighter staves and projectiles */
	FadeWhite           /* absorbed */
	FadeBigWhite        /* rocket (probably) absorbed */
	FadeOrange          /* flamethrower */
	FadeLongOrange      /* marathon lava */
	FadeGreen           /* hunter projectile */
	FadeLongGreen       /* alien green goo */
	FadeStatic          /* compiler projectile */
	FadeNegative        /* minor fusion projectile */
	FadeBigNegative     /* major fusion projectile */
	FadeFlickerNegative /* hummer projectile */
	FadeDodgePurple     /* alien weapon */
	FadeBurnCyan        /* armageddon beast electricity */
	FadeDodgeYellow     /* armageddon beast projectile */
	FadeBurnGreen       /* hunter projectile */

	FadeTintGreen  /* under goo */
	FadeTintBlue   /* under water */
	FadeTintOrange /* under lava */
	FadeTintGross  /* under sewage */

	NumberOfFadeTypes

	// Effect types
	EffectUnderWater = iota
	EffectUnderLaval
	EffectUnderSewage
	EffectUnderGoo
	NumberOfFadeEffectTypes

	AdjustedTransparencyDownshift = 8

	MinimumFadeRestartTicks = cseries.MachineTicksPerSecond / 2
	MinimumFadeUpdateTicks  = cseries.MachineTicksPerSecond / 8

	FullScreenFlag         = 0x0001
	RandomTransparencyFlag = 0x0002
)

var FadesRandomSeed = cseries.Word(0x1)

func FadesRandom() cseries.Word {
	if (FadesRandomSeed & 1) != 0 {
		FadesRandomSeed = (FadesRandomSeed >> 1) ^ 0xb400
	} else {
		FadesRandomSeed >>= 1
	}
	return FadesRandomSeed
}

type FadeProcedure func(ColorTable, ColorTable, *RgbColor, cseries.Fixed)

type FadeEffectDefinition struct {
	Type         int16
	Transparency cseries.Fixed
}

type FadeDefinition struct {
	Proc                                   FadeProcedure
	Color                                  RgbColor
	InitialTransparency, FinalTransparency cseries.Fixed // in [0, FIXED_ONE]

	Period   int16
	Flags    cseries.Word
	Priority int16 // higher is higher
}

func (this *FadeDefinition) IsActive() bool {
	return this.Flags&cseries.Word(0x8000) != 0
}

func (this *FadeDefinition) SetActive(result bool) {
	if result {
		this.Flags |= cseries.Word(0x8000)
	} else {
		this.Flags &^= cseries.Word(0x8000)
	}
}

type FadeData struct {
	Flags cseries.Word // [active.1] [unused.15]

	Type       int16
	EffectType int16

	StartTick      int64
	LastUpdateTick int64

	OriginalColorTable ColorTable
	AnimatedColorTable ColorTable
}

func (this *FadeData) IsActive() bool {
	return this.Flags&cseries.Word(0x8000) != 0
}

func (this *FadeData) SetActive(result bool) {
	if result {
		this.Flags |= cseries.Word(0x8000)
	} else {
		this.Flags &^= cseries.Word(0x8000)
	}
}

var Fade *FadeData

var FadeDefinitions = [NumberOfFadeTypes]FadeDefinition{
	{TintColorTable, RgbColor{0, 0, 0}, cseries.FixedOne, cseries.FixedOne, 0, FullScreenFlag, 0},                      /* StartCinematicFadeIn */
	{TintColorTable, RgbColor{0, 0, 0}, cseries.FixedOne, 0, cseries.MachineTicksPerSecond / 2, FullScreenFlag, 0},     /* CinematicFadeIn */
	{TintColorTable, RgbColor{0, 0, 0}, cseries.FixedOne, 0, 3 * cseries.MachineTicksPerSecond / 2, FullScreenFlag, 0}, /* LongCinematicFadeIn */
	{TintColorTable, RgbColor{0, 0, 0}, 0, cseries.FixedOne, cseries.MachineTicksPerSecond / 2, FullScreenFlag, 0},     /* CinematicFadeOut */
	{TintColorTable, RgbColor{0, 0, 0}, 0, 0, 0, FullScreenFlag, 0},                                                    /* EndCinematicFadeOut */

	{TintColorTable, RgbColor{65535, 0, 0}, (3 * cseries.FixedOne) / 4, 0, cseries.MachineTicksPerSecond / 4, 0, 0},                 /* FadeRed */
	{TintColorTable, RgbColor{65535, 0, 0}, cseries.FixedOne, 0, (3 * cseries.MachineTicksPerSecond) / 4, 0, 25},                    /* FadeBigRed */
	{TintColorTable, RgbColor{0, 65535, 0}, cseries.FixedOneHalf, 0, cseries.MachineTicksPerSecond / 4, 0, 0},                       /* FadeBonus */
	{TintColorTable, RgbColor{65535, 65535, 50000}, cseries.FixedOne, 0, cseries.MachineTicksPerSecond / 3, 0, 0},                   /* FadeBright */
	{TintColorTable, RgbColor{65535, 65535, 50000}, cseries.FixedOne, 0, 4 * cseries.MachineTicksPerSecond, 0, 100},                 /* FadeLongBright */
	{TintColorTable, RgbColor{65535, 65535, 0}, cseries.FixedOne, 0, cseries.MachineTicksPerSecond / 2, 0, 50},                      /* FadeYellow */
	{TintColorTable, RgbColor{65535, 65535, 0}, cseries.FixedOne, 0, cseries.MachineTicksPerSecond, 0, 75},                          /* FadeBigYellow */
	{TintColorTable, RgbColor{215 * 256, 107 * 256, 65535}, (3 * cseries.FixedOne) / 4, 0, cseries.MachineTicksPerSecond / 4, 0, 0}, /* FadePurple */
	{TintColorTable, RgbColor{169 * 256, 65535, 224 * 256}, (3 * cseries.FixedOne) / 4, 0, cseries.MachineTicksPerSecond / 2, 0, 0}, /* FadeCyan */
	{TintColorTable, RgbColor{65535, 65535, 65535}, cseries.FixedOneHalf, 0, cseries.MachineTicksPerSecond / 4, 0, 0},               /* FadeWhite */
	{TintColorTable, RgbColor{65535, 65535, 65535}, cseries.FixedOne, 0, cseries.MachineTicksPerSecond / 2, 0, 25},                  /* FadeBigWhite */
	{TintColorTable, RgbColor{65535, 32768, 0}, cseries.FixedOne, 0, cseries.MachineTicksPerSecond / 4, 0, 0},                       /* FadeOrange */
	{TintColorTable, RgbColor{65535, 32768, 0}, cseries.FixedOne / 4, 0, 3 * cseries.MachineTicksPerSecond, 0, 25},                  /* FadeLongOrange */
	{TintColorTable, RgbColor{0, 65535, 0}, 3 * cseries.FixedOne / 4, 0, cseries.MachineTicksPerSecond / 2, 0, 0},                   /* FadeGreen */
	{TintColorTable, RgbColor{65535, 0, 65535}, cseries.FixedOne / 4, 0, 3 * cseries.MachineTicksPerSecond, 0, 25},                  /* FadeLongGreen */

	{RandomizeColorTable, RgbColor{0, 0, 0}, cseries.FixedOne, 0, (3 * cseries.MachineTicksPerSecond) / 8, 0, 0},                 /* FadeStatic */
	{NegateColorTable, RgbColor{65535, 65535, 65535}, cseries.FixedOne, 0, cseries.MachineTicksPerSecond / 2, 0, 0},              /* FadeNegative */
	{NegateColorTable, RgbColor{65535, 65535, 65535}, cseries.FixedOne, 0, (3 * cseries.MachineTicksPerSecond) / 2, 0, 25},       /* FadeBigNegative */
	{NegateColorTable, RgbColor{0, 65535, 0}, cseries.FixedOne, 0, cseries.MachineTicksPerSecond / 2, RandomTransparencyFlag, 0}, /* FadeFlickerNegative */
	{DodgeColorTable, RgbColor{0, 65535, 0}, cseries.FixedOne, 0, (3 * cseries.MachineTicksPerSecond) / 4, 0, 0},                 /* FadeDodgePurple */
	{BurnColorTable, RgbColor{0, 65535, 65535}, cseries.FixedOne, 0, cseries.MachineTicksPerSecond, 0, 0},                        /* FadeBurnCyan */
	{DodgeColorTable, RgbColor{0, 0, 65535}, cseries.FixedOne, 0, (3 * cseries.MachineTicksPerSecond) / 2, 0, 0},                 /* FadeDodgeYellow */
	{BurnColorTable, RgbColor{0, 65535, 0}, cseries.FixedOne, 0, 2 * cseries.MachineTicksPerSecond, 0, 0},                        /* FadeBurnGreen */

	{SoftTintColorTable, RgbColor{137 * 256, 0, 137 * 256}, cseries.FixedOne, 0, 2 * cseries.MachineTicksPerSecond, 0, 0}, /* FadeTintPurple */
	{SoftTintColorTable, RgbColor{0, 0, 65535}, cseries.FixedOne, 0, 2 * cseries.MachineTicksPerSecond, 0, 0},             /* FadeTintBlue */
	{SoftTintColorTable, RgbColor{65535, 16384, 0}, cseries.FixedOne, 0, 2 * cseries.MachineTicksPerSecond, 0, 0},         /* FadeTintOrange */
	{SoftTintColorTable, RgbColor{32768, 65535, 0}, cseries.FixedOne, 0, 2 * cseries.MachineTicksPerSecond, 0, 0},         /* FadeTintGross */
}

var FadeEffectDefinitions = [NumberOfFadeEffectTypes]FadeEffectDefinition{
	{FadeTintBlue, 3 * cseries.FixedOne / 4},   /* EffectUnderWater */
	{FadeTintOrange, 3 * cseries.FixedOne / 4}, /* EffectUnderLava */
	{FadeTintGross, 3 * cseries.FixedOne / 4},  /* EffectUnderSewage */
	{FadeTintGreen, 3 * cseries.FixedOne / 4},  /* EffectUnderGoo */
}

var ActualGammaValues = [NumberOfGammaLevels]float32{
	1.3,
	1.15,
	1.0, // Default
	0.95,
	0.90,
	0.85,
	0.77,
	0.70,
}

func InitializeFades() {
	Fade = new(FadeData)
	Fade.SetActive(false)
	Fade.EffectType = cseries.None
}

func UpdateFades() (bool, error) {
	if Fade.IsActive() {
		definition, err := GetFadeDefinition(Fade.Type)
		if err != nil {
			return false, err
		}
		tickCount, err0 := cseries.MachineTickCount()
		if err0 != nil {
			return false, err0
		}
		update := false
		var transparency cseries.Fixed
		phase := int16(tickCount - Fade.StartTick) // hmmm, this could be a problem
		if phase >= definition.Period {
			transparency = definition.FinalTransparency
			Fade.SetActive(false)
			update = true
		} else {
			if (tickCount - Fade.LastUpdateTick) >= MinimumFadeUpdateTicks {
				transparency = definition.InitialTransparency + (cseries.Fixed(phase)*(definition.FinalTransparency-definition.InitialTransparency))/cseries.Fixed(definition.Period)
				if (definition.Flags & RandomTransparencyFlag) != 0 {
					transparency += cseries.Fixed(FadesRandom()) % (definition.FinalTransparency - transparency)
				}
				update = true
			}
			if update {
				RecalculateAndDisplayColorTable(Fade.Type, transparency, Fade.OriginalColorTable, Fade.AnimatedColorTable)
			}
		}

	}
	return Fade.IsActive(), nil
}

func SetFadeEffect(fade int16) {
	if Fade.EffectType != fade {
		Fade.EffectType = fade
		if !Fade.IsActive() {
			if fade == cseries.None {
				AnimateScreenClut(WorldColorTable, false)
			} else {
				RecalculateAndDisplayColorTable(cseries.None, 0, WorldColorTable, VisibleColorTable)
			}
		}
	}
}
func StartFade(fade int16) error {
	return ExplicitStartFade(fade, WorldColorTable, VisibleColorTable)
}

func ExplicitStartFade(fade int16, originalColorTable, animatedColorTable ColorTable) error {
	definition, err1 := GetFadeDefinition(fade)
	if err1 != nil {
		return err1
	}
	tickCount, err0 := cseries.MachineTickCount()
	if err0 != nil {
		return err0
	}
	doFade := true

	if Fade.IsActive() {
		oldDefinition, err2 := GetFadeDefinition(Fade.Type)
		if err2 != nil {
			return err2
		}
		if oldDefinition.Priority > definition.Priority {
			doFade = false
		}

		if ((tickCount - Fade.StartTick) > MinimumFadeRestartTicks) && Fade.Type == fade {
			doFade = false
		}
	}

	if doFade {
		Fade.SetActive(false)

		RecalculateAndDisplayColorTable(fade, definition.InitialTransparency, originalColorTable, animatedColorTable)
		if definition.Period != 0 {
			Fade.Type = fade
			Fade.LastUpdateTick = tickCount
			Fade.StartTick = tickCount
			Fade.OriginalColorTable = originalColorTable
			Fade.AnimatedColorTable = animatedColorTable
			Fade.SetActive(true)
		}
	}
	return nil
}

func StopFade() error {
	if Fade.IsActive() {
		defn, err := GetFadeDefinition(Fade.Type)
		if err != nil {
			return err
		}
		RecalculateAndDisplayColorTable(Fade.Type, defn.FinalTransparency, Fade.OriginalColorTable, Fade.AnimatedColorTable)
		Fade.SetActive(false)
	}
	return nil
}

func FadeFinished() bool {
	return !Fade.IsActive()
}

func FullFade(fade int16, original ColorTable) error {
	animated := make(ColorTable, 0, 256)
	copy(animated[0:len(original)], original)

	// If draw sprocket support isn't there
	ExplicitStartFade(fade, original, animated)
	for {
		if result, err := UpdateFades(); err != nil {
			return err
		} else if !result {
			break
		}
	}
	return nil
}
func GetFadePeriod(fade int16) (int16, error) {
	if defn, err := GetFadeDefinition(fade); err != nil {
		return 0, err
	} else {
		return defn.Period, nil
	}
}

func GammaCorrectColorTable(uncorrected, corrected ColorTable, gammaLevel int16) error {
	var gamma float32
	if !(gammaLevel >= 0 && gammaLevel < NumberOfGammaLevels) {
		return &cseries.AssertionError{
			Function: "GamaCorrectColorTable",
			Message:  fmt.Sprintf("Gamma level (%d) is not between [0,%d)", gammaLevel, NumberOfGammaLevels),
		}
	}
	gamma = ActualGammaValues[gammaLevel]

	fn := func(value cseries.Word, g float32) cseries.Word {
		return cseries.Word(math.Pow(float64(value/65535.0), float64(g)) * 65535.0)
	}
	for i := 0; i < len(uncorrected); i++ {
		corrected[i].Red = fn(uncorrected[i].Red, gamma)
		corrected[i].Green = fn(uncorrected[i].Green, gamma)
		corrected[i].Blue = fn(uncorrected[i].Blue, gamma)
	}

	return nil
}
func GetFadeDefinition(index int16) (*FadeDefinition, error) {
	if !(index >= 0 && index < NumberOfFadeTypes) {
		return nil, &cseries.AssertionError{
			Function: "GetFadeDefinition",
			Message:  fmt.Sprintf("Index (%d) is not in the valid fade type range of [0,%d)", index, NumberOfFadeTypes),
		}
	}

	return &FadeDefinitions[index], nil
}

func GetFadeEffectDefinition(index int16) (*FadeEffectDefinition, error) {
	if !(index >= 0 && index < NumberOfFadeEffectTypes) {
		return nil, &cseries.AssertionError{
			Function: "GetFadeEffectDefinition",
			Message:  fmt.Sprintf("Index (%d) is not in the valid fade effect type range of [0,%d)", index, NumberOfFadeEffectTypes),
		}
	}

	return &FadeEffectDefinitions[index], nil
}

func RecalculateAndDisplayColorTable(fadeType int16, transparency cseries.Fixed, original, animated ColorTable) error {
	fullScreen := false

	if Fade.EffectType != cseries.None {
		effectDefinition, err0 := GetFadeEffectDefinition(Fade.EffectType)
		if err0 != nil {
			return err0
		}
		definition, err1 := GetFadeDefinition(effectDefinition.Type)
		if err1 != nil {
			return err1
		}

		definition.Proc(original, animated, &definition.Color, effectDefinition.Transparency)
		original = animated
	}

	if fadeType != cseries.None {
		definition, err0 := GetFadeDefinition(fadeType)
		if err0 != nil {
			return err0
		}
		definition.Proc(original, animated, &definition.Color, transparency)
		fullScreen = (definition.Flags & FullScreenFlag) != 0
	}
	AnimateScreenClut(animated, fullScreen)
	return nil
}

func TintColorTable(original, animated ColorTable, color *RgbColor, transparency cseries.Fixed) {
	adjustedTransparency := transparency >> AdjustedTransparencyDownshift
	fn := func(unadjustedValue, colorValue cseries.Word) cseries.Word {
		return cseries.Word(cseries.Fixed(unadjustedValue) + ((cseries.Fixed(colorValue-unadjustedValue) * adjustedTransparency) >> (cseries.FixedFractionalBits - AdjustedTransparencyDownshift)))
	}
	for i := 0; i < len(original); i++ {
		animated[i].Red = fn(original[i].Red, color.Red)
		animated[i].Green = fn(original[i].Green, color.Green)
		animated[i].Blue = fn(original[i].Blue, color.Blue)
	}
}

func RandomizeColorTable(original, animated ColorTable, color *RgbColor, transparency cseries.Fixed) {
	var mask cseries.Word
	adjustedTransparency := transparency.Pin(0, 0xffff)
	// calculate a mask which has all bits including and lower than the high-bit in the transparency set
	for mask = 0; (cseries.Word(adjustedTransparency) &^ mask) != 0; mask = (mask << 1) | 1 {
		// empty loop body
	}
	fn := func(value, mask cseries.Word) cseries.Word {
		return value + (FadesRandom() & mask)
	}
	for i := 0; i < len(original); i++ {
		animated[i].Red = fn(original[i].Red, mask)
		animated[i].Green = fn(original[i].Green, mask)
		animated[i].Blue = fn(original[i].Blue, mask)
	}
}

// unlike pathways, all colors won't pass through 50% gray at the same time
func NegateColorTable(original, animated ColorTable, color *RgbColor, transparency cseries.Fixed) {
	transparency = cseries.FixedOne - transparency

	fn := func(ov, cv cseries.Word) cseries.Word {
		var tmp cseries.Word
		if ov > 0x8000 {
			tmp = (ov ^ cv) + cseries.Word(transparency)
			return tmp.Ceiling(ov)
		} else {
			tmp = (ov ^ cv) - cseries.Word(transparency)
			return tmp.Floor(ov)
		}
	}
	for i := 0; i < len(original); i++ {
		animated[i].Red = fn(original[i].Red, color.Red)
		animated[i].Green = fn(original[i].Green, color.Green)
		animated[i].Blue = fn(original[i].Blue, color.Blue)
	}
}

func DodgeColorTable(original, animated ColorTable, color *RgbColor, transparency cseries.Fixed) {
	customCeiling := func(n, ceiling int32) int32 {
		if n > ceiling {
			return ceiling
		} else {
			return n
		}
	}
	computeComponent := func(unadjusted, color cseries.Word) int32 {
		return int32(0xffff - cseries.Fixed((((color ^ 0xffff) * unadjusted) >> cseries.FixedFractionalBits)) - transparency)
	}
	updateElement := func(adjusted cseries.Word, component int32) cseries.Word {
		return cseries.Word(customCeiling(component, int32(adjusted)))
	}
	for i := 0; i < len(original); i++ {
		var component int32

		component = computeComponent(original[i].Red, color.Red)
		animated[i].Red = updateElement(original[i].Red, component)
		component = computeComponent(original[i].Blue, color.Blue)
		animated[i].Blue = updateElement(original[i].Blue, component)
		component = computeComponent(original[i].Green, color.Green)
		animated[i].Green = updateElement(original[i].Green, component)
	}

	// Using the comma operator instead of ; for assignment...why?
	//component= 0xffff - (((color->red^0xffff)*unadjusted->red)>>FIXED_FRACTIONAL_BITS) - transparency, adjusted->red= CEILING(component, unadjusted->red);
	//component= 0xffff - (((color->green^0xffff)*unadjusted->green)>>FIXED_FRACTIONAL_BITS) - transparency, adjusted->green= CEILING(component, unadjusted->green);
	//component= 0xffff - (((color->blue^0xffff)*unadjusted->blue)>>FIXED_FRACTIONAL_BITS) - transparency, adjusted->blue= CEILING(component, unadjusted->blue);
}

func BurnColorTable(original, animated ColorTable, color *RgbColor, transparency cseries.Fixed) {
	customCeiling := func(n, ceiling int32) int32 {
		if n > ceiling {
			return ceiling
		} else {
			return n
		}
	}
	updateElement := func(adjusted cseries.Word, component int32) cseries.Word {
		return cseries.Word(customCeiling(component, int32(adjusted)))
	}
	computeComponent := func(unadjusted, color cseries.Word, transparency cseries.Fixed) int32 {
		return int32(cseries.Fixed(((color * unadjusted) >> cseries.FixedFractionalBits)) + transparency)
	}
	transparency = cseries.FixedOne - transparency
	for i := 0; i < len(original); i++ {
		r := original[i].Red
		g := original[i].Green
		b := original[i].Blue
		component := computeComponent(r, color.Red, transparency)
		animated[i].Red = updateElement(r, component)
		component = computeComponent(g, color.Green, transparency)
		animated[i].Green = updateElement(g, component)
		component = computeComponent(b, color.Blue, transparency)
		animated[i].Blue = updateElement(b, component)
	}

	//component= ((color->red*unadjusted->red)>>FIXED_FRACTIONAL_BITS) + transparency, adjusted->red= CEILING(component, unadjusted->red);
	//component= ((color->green*unadjusted->green)>>FIXED_FRACTIONAL_BITS) + transparency, adjusted->green= CEILING(component, unadjusted->green);
	//component= ((color->blue*unadjusted->blue)>>FIXED_FRACTIONAL_BITS) + transparency, adjusted->blue= CEILING(component, unadjusted->blue);

}

func SoftTintColorTable(original, animated ColorTable, color *RgbColor, transparency cseries.Fixed) {
	adjustedTransparency := cseries.Word(transparency >> AdjustedTransparencyDownshift)
	fn := func(x, y, z, w cseries.Word) cseries.Word {
		return cseries.Word(x + (((((y * w) >> (cseries.FixedFractionalBits - AdjustedTransparencyDownshift)) - x) * z) >> (cseries.FixedFractionalBits - AdjustedTransparencyDownshift)))
	}
	for i := 0; i < len(original); i++ {
		intensity := cseries.Word(cseries.Max(int64(original[i].Red), int64(original[i].Green)))
		intensity = cseries.Word(cseries.Max(int64(intensity), int64(original[i].Blue)) >> AdjustedTransparencyDownshift)

		animated[i].Red = fn(original[i].Red, color.Red, adjustedTransparency, intensity)
		animated[i].Green = fn(original[i].Green, color.Green, adjustedTransparency, intensity)
		animated[i].Blue = fn(original[i].Blue, color.Blue, adjustedTransparency, intensity)

	}
	//adjusted->red= unadjusted->red + (((((color->red*intensity)>>(FIXED_FRACTIONAL_BITS-ADJUSTED_TRANSPARENCY_DOWNSHIFT))-unadjusted->red)*adjusted_transparency)>>(FIXED_FRACTIONAL_BITS-ADJUSTED_TRANSPARENCY_DOWNSHIFT));
	//adjusted->green= unadjusted->green + (((((color->green*intensity)>>(FIXED_FRACTIONAL_BITS-ADJUSTED_TRANSPARENCY_DOWNSHIFT))-unadjusted->green)*adjusted_transparency)>>(FIXED_FRACTIONAL_BITS-ADJUSTED_TRANSPARENCY_DOWNSHIFT));
	//adjusted->blue= unadjusted->blue + (((((color->blue*intensity)>>(FIXED_FRACTIONAL_BITS-ADJUSTED_TRANSPARENCY_DOWNSHIFT))-unadjusted->blue)*adjusted_transparency)>>(FIXED_FRACTIONAL_BITS-ADJUSTED_TRANSPARENCY_DOWNSHIFT));
}
