// common flags class
package moo

import (
	"github.com/DrItanium/moo/cseries"
)

type GenericFlags int16

func (this *GenericFlags) MarkSlotAsFree() {
	*flags &= ^0x8000
}
func (this *GenericFlags) MarkSlotAsUsed() {
	*flags |= 0x8000
}

func (this GenericFlags) SlotIsUsed() bool {
	return (this & 0x8000) != 0
}

func (this GenericFlags) SlotIsFree() bool {
	return !this.SlotIsUsed()
}

func (this GenericFlags) ObjectWasRendered() bool {
	return (this & 0x4000) != 0
}

func (this *GenericFlags) SetObjectRenderedFlag() {
	*this |= 0x4000
}

func (this *GenericFlags) ClearObjectRenderedFlag() {
	*this &= ^0x4000
}
