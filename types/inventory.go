package types

// constants "blockProject/constants"

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

func NewInventory(MaxSlots int, OwnerType string) *Inventory {
  return &Inventory{
    Slots: generateEmptyItemSlots(MaxSlots),
    MaxSlots: MaxSlots,
    OwnerType: OwnerType,
  }
}

func generateEmptyItemSlots(count int) []ItemStack {
  return make([]ItemStack, count)
}

func (inv *Inventory) AddItem(item Item, count int) bool {
  return true
}


// func (inv *Inventory) RemoveItem(slotInd int, count int) Item {
//   return 
// }
