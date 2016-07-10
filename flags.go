// common flags class
package moo

type GenericFlags int32

func (this *GenericFlags) MarkSlotAsFree() {
	*this &= GenericFlags(^0x8000)
}
func (this *GenericFlags) MarkSlotAsUsed() {
	*this |= GenericFlags(0x8000)
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
