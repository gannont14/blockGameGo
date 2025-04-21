package types

// constants "blockProject/constants"

var activeSelectedSlot int = -1

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
  slot, sameItem := inv.nextSlot(item, count)

  if slot == -1 { return false }

  // nextSlot already check if they are the same item, so just add count
  inv.Slots[slot].Count += count

  if sameItem == false{
    inv.Slots[slot].Item = item
  }

  return true
}

// func (inv *Inventory) RemoveItem(slotInd int, count int) Item {
//   return 
// }

// find the next best slot in the inventory for the item trying to be added
func (inv *Inventory) nextSlot (item Item, count int) (int, bool) {
  for idx, currItem := range inv.Slots{
    // nothing in the slot
    if currItem.Count == 0 { return idx, false }

    // item of same type in slot, with enough room, could get rid of nested loop
    if currItem.Item.GetID() == item.GetID() {
      // make sure that there is enough room
      if currItem.Count + count <= currItem.Item.GetMaxStackSize(){
        return idx, true
      }
      // otherwise no
    }
  }
  // -1 meaning no good slot, inventory full
  return -1, false
}

// allows drawing function to update which slot the player is currently looking at,
// will later update based on the user input
func UpdateActiveSelectedSlot(slot int){
  activeSelectedSlot = slot
}


