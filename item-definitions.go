package moo

type ItemDefinition struct {
	ItemKind              int16
	SingularNameId        int16
	PluralNameId          int16
	BaseShape             ShapeDescriptor
	MaximumCountPerPlayer int16
	InvalidEnvironments   int16
}

var ItemDefinitions []ItemDefinition

//TODO: fill in item definitions
