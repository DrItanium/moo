// item related operations
package moo

import (
//	"fmt"
//"github.com/DrItanium/moo/cseries"
)

const (
	// item types (class)
	Weapon = iota
	Ammunition
	Powerup
	Item
	WeaponPowerup
	Ball

	NumberOfItemTypes
	NetworkStatistics = NumberOfItemTypes // used in game_window.go

	// item types
	ItemKnife = iota
	ItemMagnum
	ItemMagnumMagazine
	ItemPlasmaPistol
	ItemPlasmaMagazine
	ItemAssaultRifle
	ItemAssaultRifleMagazine
	ItemAssaultGrenadeMagazine
	ItemMissileLauncher
	ItemMissileLauncherMagazine
	ItemInvisibilityPowerup
	ItemInvincibilityPowerup
	ItemInfravisionPowerup
	ItemAlienShotgun
	ItemAlienShotgunMagazine
	ItemFlamethrower
	ItemFlamethrowerCanister
	ItemExtravisionPowerup
	ItemOxygenPowerup
	ItemEnergyPowerup
	ItemDoubleEnergyPowerup
	ItemTripleEnergyPowerup
	ItemShotgun
	ItemShotgunMagazine
	ItemSphtDoorKey
	ItemUplinkChip

	BallItemBase
	ItemLightBlueBall = BallItemBase
	ItemRedBall
	ItemVioletBall
	ItemYellowBall
	ItemBrownBall
	ItemOrangeBall
	ItemBlueBall // heh heh
	ItemGreenBall

	ItemSmg
	ItemSmgAmmo

	NumberOfDefinedItems

	Structure_ItemNameList   = 150
	Structure_HeaderNameList = 151

	MaximumArmReach = 3 * WorldOneFourth
)

func NewItem(location *ObjectLocation, itemType int16) (int16, error) {
	//var objectIndex int16
	//definition, err0 := GetItemDefinition(itemType)
	//if err0 != nil {
	//	return 0, err0
	//}
	//addItem := true

	//if len(ItemDefinitions) != NumberOfDefinedItems {
	//	return 0, &cseries.AssertionError{
	//		Function: "NewItem",
	//		Message:  fmt.Sprintf("Number of item definitions (%d) does not equal number of defined items (%d)", len(ItemDefinitions), NumberOfDefinedItems),
	//	}
	//}

	//if DynamicWorld.PlayerCount > 1 {
	//	if (definition.InvalidEnvironments & EnvironmentNetwork) != 0 {
	//		addItem = false
	//	}

	//	if GetItemKind(itemType) == ItemBall && !CurrentGameHasBalls() {
	//		addItem = false
	//	}
	//} else {
	//	if (definition.InvalidEnvironments & EnvironmentSinglePlayer) != 0 {
	//		addItem = false
	//	}
	//}
	//if addItem {
	//	objectIndex, err = NewMapObject(location, definition.BaseShape)
	//	if err != nil {
	//		return 0, err
	//	}
	//	object := GetObjectData(objectIndex)
	//	object.SetOwner(ObjectIsItem)
	//	object.Permutation = itemType
	//	//	if location.Flags &
	//}
	return 0, nil
}

func CalculatePlayerItemArray(playerIndex, itemType int16, items, counts, arrayCount *int16) {

}

func GetHeaderName(buffer string, itemType int16) {

}
func GetItemName(buffer string, itemId int16, plural bool) {

}
func NewItemInRandomLocation(itemType int16) bool {
	return false
}
func CountInventoryLines(playerIndex int16) int16 {
	return 0
}
func SwipeNearbyItems(playerIndex int16) {

}
func MarkItemCollections(loading bool) {

}

func GetItemKind(itemId int16) {

}

func UnretrievedItemsOnMap() bool {
	return false
}

func ItemValidInCurrentEnvironment(itemType int16) bool {
	return false
}

func TriggerNearbyItems(polygonIndex int16) {

}

// returns the color of the ball or NONE if they don't have one
// Returns NONE if this player is not carrying a ball
func FindPlayerBallColor(playerIndex int16) int16 {
	return 0
}

func GetItemShape(itemId int16) int16 {
	return 0
}

func TryAndAddPlayerItem(playerIndex, itemType int16) bool {
	return false
}

func GetItemDefinition(itemType int16) (*ItemDefinition, error) {
	return nil, nil
}

func GetItem(playerIndex, objectIndex int16) bool {
	return false
}

func TestItemRetrieval(polygonIndex1 int16, location1, location2 *WorldPoint3d) bool {
	return false
}
