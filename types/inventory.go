package types

import (
  // constants "blockProject/constants"
)

type Inventory struct {
  // not sure if this should be a 1d or 2d array but we will go with 1
  Slots []ItemStack
  MaxSlots int
  OwnerType string // could get it's own type if needed later
}

type ItemStack struct {
  Item Item
  Count int
}

func (inv *Inventory) AddItem(item Item, count int) bool {
  return true
}


// func (inv *Inventory) RemoveItem(slotInd int, count int) Item {
//   return 
// }
