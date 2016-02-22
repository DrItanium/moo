//import of shape_descriptors.h
package moo

import (
	"github.com/DrItanium/moo/cseries"
)

type ShapeDescriptor int16 /* [clut.3] [collection.5] [shape.8] */

const (
	DescriptorShapeBits        = 8
	DescriptorCollectionBits   = 5
	DescriptorClutBits         = 3
	MaximumCollections         = 1 << DescriptorCollectionBits
	MaximumShapesPerCollection = 1 << DescriptorShapeBits
	MaximumClutsPerCollection  = 1 << DescriptorClutBits
)

const (
	// collection numbers
	CollectionInterface      = iota // 0
	CollectionWeaponsInHand         // 1
	CollectionJuggernaut            // 2
	CollectionTick                  // 3
	CollectionRocket                // 4
	CollectionHunter                // 5
	CollectionPlayer                // 6
	CollectionItems                 // 7
	CollectionTrooper               // 8
	CollectionFighter               // 9
	CollectionDefender              // 10
	CollectionYeti                  // 11
	CollectionCivilian              // 12
	CollectionVacuumCivilian        // 13
	CollectionEnforcer              // 14
	CollectionHummer                // 15
	CollectionCompiler              // 16
	CollectionWalls1                // 17
	CollectionWalls2                // 18
	CollectionWalls3                // 19
	CollectionWalls4                // 20
	CollectionWalls5                // 21
	CollectionScenery1              // 22
	CollectionScenery2              // 23
	CollectionScenery3              // 24
	CollectionScenery4              // 25 pathways
	CollectionScenery5              // 26 alien
	CollectionLandscape1            // 27 day
	CollectionLandscape2            // 28 night
	CollectionLandscape3            // 29 moon
	CollectionLandscape4            // 30
	CollectionCyborg                // 31

	NumberOfCollections
)

func (this ShapeDescriptor) Shape() cseries.Word {
	return cseries.Word(this) & cseries.Word(MaximumShapesPerCollection-1)
}

//#define GET_DESCRIPTOR_COLLECTION(d) (((d)>>DESCRIPTOR_SHAPE_BITS)&(word)((1<<(DESCRIPTOR_COLLECTION_BITS+DESCRIPTOR_CLUT_BITS))-1))
func (this ShapeDescriptor) Collection() cseries.Word {
	return cseries.Word(this>>DescriptorShapeBits) & cseries.Word((1<<(DescriptorCollectionBits+DescriptorClutBits))-1)
}

//#define GET_COLLECTION_CLUT(collection) (((collection)>>DESCRIPTOR_COLLECTION_BITS)&(word)(MAXIMUM_CLUTS_PER_COLLECTION-1))
func (this ShapeDescriptor) Clut() cseries.Word {
	return cseries.Word((this >> DescriptorCollectionBits)) & cseries.Word(MaximumClutsPerCollection-1)
}

//TODO: the original code is relying on compile time macros (which have no type checking to do many operations). This is not possible in golang so other steps have to be taken. I just don't want to do them right now
//#define BUILD_DESCRIPTOR(collection,shape) (((collection)<<DESCRIPTOR_SHAPE_BITS)|(shape))
//func BuildDescriptor(collection int16, shape byte) ShapeDescriptor {
//	return ShapeDescriptor((ShapeDescriptor(collection) << DescriptorShapeBits) | ShapeDescriptor(shape))
//}

//#define BUILD_COLLECTION(collection,clut) ((collection)|(word)((clut)<<DESCRIPTOR_COLLECTION_BITS))
//func BuildCollection(collection, clut) {
//
//}
//#define GET_COLLECTION(collection) ((collection)&(MAXIMUM_COLLECTIONS-1))
