\ the game divides the different shapes (art assets) into sets of collections with each
\ collection being part of a color lookup table kind
sizeof(short) constant sizeof(shape-descriptor) \ [clut.3] [collection.5] [shape.8]
: (shape-descriptor) ( v -- n ) (int16) ;

8 constant descriptor-shape-bits
5 constant descriptor-collection-bits
3 constant descriptor-clut-bits

\ the game divides 
descriptor-collection-bits 1u<< constant maximum-collections
descriptor-shape-bits 1u<< constant maximum-shapes-per-collection
descriptor-clut-bits 1u<< constant maximum-cluts-per-collection

{enum \ collection numbers
enum: CollectionInterface
enum: CollectionWeaponsInHand
enum: CollectionJuggernaut
enum: CollectionTick
enum: CollectionRocket
enum: CollectionHunter
enum: CollectionPlayer
enum: CollectionItems
enum: CollectionTrooper
enum: CollectionFighter
enum: CollectionDefender
enum: CollectionYeti
enum: CollectionCivilian
enum: CollectionVacuumCivilian
enum: CollectionEnforcer
enum: CollectionHummer
enum: CollectionCompiler
enum: CollectionWalls1
enum: CollectionWalls2
enum: CollectionWalls3
enum: CollectionWalls4
enum: CollectionWalls5
enum: CollectionScenery1
enum: CollectionScenery2
enum: CollectionScenery3
enum: CollectionScenery4
enum: CollectionScenery5
enum: CollectionLandscape1
enum: CollectionLandscape2
enum: CollectionLandscape3
enum: CollectionLandscape4
enum: CollectionLandscape5
enum: CollectionCyborg
enum: NumberOfCollections
enum}


: get-descriptor-shape ( d -- v ) maximum-shapes-per-collection 1- and ;
: get-descriptor-collection ( d -- v ) 
  descriptor-shape-bits u>> 
  descriptor-collection-bits 
  descriptor-clut-bits + 
  1u<< 1- and ;
: build-descriptor ( collection shape -- n ) 
  swap ( shape collection )
  descriptor-shape-bits u<< or (shape-descriptor) ;
: build-collection ( collection clut -- n )
  descriptor-collection-bits u<< or (word) ;
: get-collection-clut ( collection -- n ) 
  descriptor-collection-bits u>> 
  maximum-cluts-per-collection 1-
  and (word) ;
: get-collection ( collection -- n ) maximum-collections 1- and ;

;s \ always last in the file
