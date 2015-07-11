// fading related stuffs
package moo

import (
	"fmt"
	"github.com/DrItanium/moo/cseries"
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
	Index        int16
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

	StartTick      int32
	LastUpdateTick int32

	OriginalColorTable ColorTable
	AnimatedColorTable ColorTable
}

var Fade *FadeData

func TintColorTable(original, animated ColorTable, color *RgbColor, transparency cseries.Fixed) {

}

func RandomizeColorTable(original, animated ColorTable, color *RgbColor, transparency cseries.Fixed) {

}

func NegateColorTable(original, animated ColorTable, color *RgbColor, transparency cseries.Fixed) {

}

func DodgeColorTable(original, animated ColorTable, color *RgbColor, transparency cseries.Fixed) {

}

func BurnColorTable(original, animated ColorTable, color *RgbColor, transparency cseries.Fixed) {

}

func SoftTintColorTable(original, animated ColorTable, color *RgbColor, transparency cseries.Fixed) {

}

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

func GetFadeDefinition(index int16) *FadeDefinition {
	return nil
}
func GetFadeEffectDefinition(index int16) *FadeEffectDefinition {
	return nil
}

func RecalculateAndDisplayColorTable(fadeType int16, transparency cseries.Fixed, original, animated ColorTable) {

}
func InitializeFades() {
	Fade = new(FadeData)
	Fade.SetActive(false)
	Fade.FadeEffectType = cseries.None
}

func UpdateFades() bool {
	if Fade.IsActive() {
		definition := GetFadeDefinition(Fade.Type)
		tickCount := MachineTickCount()
		update := false
		var transparency cseries.Fixed
		phase := tickCount - Fade.StartTick
		if phase >= definition.Period {
			transparency = definition.FinalTransparency
			Fade.SetActive(false)
			update = true
		} else {
			if tickCount.LastUpdateTick >= MinimumFadeUpdateTicks {
				transparency = definition.InitialTransparency + (phase*(definition.FinalTransparency-definition.InitialTransparency))/definition.Period
				if (definition.Flags & RandomTransparencyFlag) != 0 {
					transparency += FadesRandom() % (definition.FinalTransparency - transparency)
				}
				update = TRUE
			}
			if update {
				RecalculateAndDisplayColorTable(Fade.Type, transparency, Fade.OriginalColorTable, Fade.AnimatedColorTable)
			}
		}

	}
	return Fade.IsActive()
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
func StartFade(fade int16) {
	ExplicitStartFade(fade, WorldColorTable, VisibleColorTable)
}

func ExplicitStartFade(fade int16, originalColorTable, animatedColorTable ColorTable) {
	definition := GetFadeDefinition(fade)
	tickCount := MachineTickCount()
	doFade := true

	if Fade.IsActive() {
		oldDefinition := GetFadeDefinition(fade.Type)
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

}

func StopFade() {
	if Fade.IsActive() {
		RecalculateAndDisplayColorTable(Fade.Type, GetFadeDefinition(Fade.Type).FinalTransparency, Fade.OriginalColorTable, Fade.AnimatedColorTable)
		Fade.SetActive(false)
	}
}

func FadeFinished() bool {
	return !Fade.IsActive()
}

func FullFade(fade int16, original ColorTable) {
	animated := make(ColorTable, 0, 256)
	copy(animated[0:len(original)], original)

	// If draw sprocket support isn't there
	ExplicitStartFade(fade, original, animated)
	for UpdateFades() {
	}
}
func GetFadePeriod(fade int16) int16 {
	return GetFadeDefinition(fade).Period
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

	for i := 0; i < len(uncorrected); i++ {
		corrected[i].Red = Math.Pow(uncorrected[i].Red/65535.0, gamma) * 65535.0
		corrected[i].Green = Math.Pow(uncorrected[i].Green/65535.0, gamma) * 65535.0
		corrected[i].Blue = Math.Pow(uncorrected[i].Blue/65535.0, gamma) * 65535.0
	}

	return nil
}

func GetFadeEffect() int16 {
	return 0
}
