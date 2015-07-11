// fading related stuffs
package moo

import (
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

type FadeProcedure func(*ColorTable, *ColorTable, *RgbColor, cseries.Fixed)

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

func (this *FadeDefinition) FadeIsActive() bool {
	return this.Flags&cseries.Word(0x8000) != 0
}

func (this *FadeDefinition) SetFadeIsActive(result bool) {
	if result {
		this.Flags |= cseries.Word(0x8000)
	} else {
		this.Flags &^= cseries.Word(0x8000)
	}
}

type FadeData struct {
	Flags cseries.Word // [active.1] [unused.15]

	Type           int16
	FadeEffectType int16

	StartTick      int32
	LastUpdateTick int32

	OriginalColorTable *ColorTable
	AnimatedColorTable *ColorTable
}

var Fade *FadeData

func TintColorTable(original, animated *ColorTable, color *RgbColor, transparency cseries.Fixed) {

}

func RandomizeColorTable(original, animated *ColorTable, color *RgbColor, transparency cseries.Fixed) {

}

func NegateColorTable(original, animated *ColorTable, color *RgbColor, transparency cseries.Fixed) {

}

func DodgeColorTable(original, animated *ColorTable, color *RgbColor, transparency cseries.Fixed) {

}

func BurnColorTable(original, animated *ColorTable, color *RgbColor, transparency cseries.Fixed) {

}

func SoftTintColorTable(original, animated *ColorTable, color *RgbColor, transparency cseries.Fixed) {

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

func InitializeFades() {

}

func UpdateFades() bool {
	return false
}

func StartFade(fade int16) {

}

func StopFade() {

}

func FadeFinished() bool {
	return false
}

func SetFadeEffect(fade int16) {

}

func GetFadeEffect() int16 {
	return 0
}

func GammaCorrectColorTable(uncorrectedColorTable, correctedColorTable *ColorTable, gammaLevel int16) {

}

func ExplicitStartFade(fade int16, originalColorTable, animatedColorTable *ColorTable) {

}

func FullFade(fade int16, originalColorTable *ColorTable) {

}
