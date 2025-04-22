package types


// constants "blockProject/constants"

var activeSelectedSlot int = -1

type Inventory struct {
  // not sure if this should be a 1d or 2d array but we will go with 1
  Slots []ItemStack
  Hand ItemStack
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
    Hand: NewItemStack(), // for the player to be able to pick up items
    MaxSlots: MaxSlots,
    OwnerType: OwnerType,
  }
}

func NewItemStack() ItemStack{
  return ItemStack{}
}

func generateEmptyItemSlots(count int) []ItemStack {
  return make([]ItemStack, count)
}

// for things like shift clicking into the inventory
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

func (inv *Inventory) HandleItemClick (itemSlot int) {
  // player didn't click on an actual slot, later check if outside of inv to drop items
  if itemSlot == -1 {
    return
  }

  // case where the player clicks on an empty slot
  if inv.Slots[itemSlot].Item == nil {
    // player has nothing in his hand, do nothing
    if inv.Hand.Item == nil {
      // nothing
      // fmt.Println("Player clicked on empty cell")
      return
    }

    // player has an item in his hand, add to that slot
    inv.Slots[itemSlot] = inv.Hand
    // fmt.Println("Adding item from player's hand to cell")
    // empty out the players hand
    inv.Hand.nillify()

    return
  }

  inv.updateItemSlot(itemSlot)
}

// return what the player should have in their hand after the add
func (inv *Inventory) updateItemSlot(itemSlot int) (bool, ItemStack) {
  slot := inv.Slots[itemSlot]

  if(inv.Hand.Item == nil){
    // fmt.Println("Player picking up item")
    temp := slot
    inv.Slots[itemSlot] = inv.Hand
    inv.Hand = temp
    return true, temp
  }

  // slot has the same item...
  if slot.Item.GetID() == inv.Hand.Item.GetID(){
    // check if there is room to add, or if they can only add so many
    // player is just adding to current stack
    if slot.Count + inv.Hand.Count <= slot.Item.GetMaxStackSize(){
      inv.Slots[itemSlot].Count = slot.Count + inv.Hand.Count
      inv.Hand.nillify()
      return false, inv.Hand
    }

    // player is able to add partially to the stack, but still has items in hand
    excess := (slot.Count + inv.Hand.Count) - inv.Hand.Item.GetMaxStackSize()
    inv.Slots[itemSlot].Count = slot.Item.GetMaxStackSize()
    inv.Hand.Count = excess
    return true, inv.Hand
  }

  // slot has different item, swap the items
  temp := slot
  inv.Slots[itemSlot] = inv.Hand
  inv.Hand = temp
  return true, temp
}


// find the next best slot in the inventory for the item trying to be added
func (inv *Inventory) nextSlot (item Item, count int) (int, bool) {
  for idx, currItem := range inv.Slots{
    // nothing in the slot
    if currItem.Item == nil { return idx, false }

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

func (item *ItemStack) nillify() {
  item.Count = 0
  item.Item = nil
}


