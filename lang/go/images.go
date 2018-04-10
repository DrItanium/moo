// image manager
package moo

const (
	ImagesFileDelta16   = 1000
	ImagesFileDelta32   = 2000
	ScenarioFileDelta16 = 10000
	ScenarioFileDelta32 = 20000
)

var interfaceBitDepth int16
var screenWindow interface{}

var imagesFileHandle int16
var scenarioFileHandle int16

//TODO: implement image manager

func InitializeImagesManager() {

}

func ImagesPictureExists(baseResource int16) bool {
	return false
}

func ScenarioPictureExists(baseResource int16) bool {
	return false
}

func CalculatePictureClut(pictResourceNumber int16) *ColorTable {
	return nil
}

func Build8BitSystemColorTable() *ColorTable {
	return nil
}

func SetScenarioImagesFile(file interface{}) {

}

func DrawFullScreenPictResourceFromImages(pictReousrceNumber int16) {

}

func DrawFullScreenPictResourceFromScenario(pictResourceNumber int16) {

}

func ScrollFullScreenPictResourceFromScenario(pictResourceNumber int16, textBlock bool) {

}
