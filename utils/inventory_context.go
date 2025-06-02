package utils

import (
	constants "blockProject/constants"
	types "blockProject/types"
	"fmt"

	. "github.com/gen2brain/raylib-go/raylib"
)

// this is an absolute mess, don't ask
type InventoryDimensionInformation struct {
	DividerHeight   float64
	MarginSpaceY    float64
	MarginSpaceX    int
	SlotSize        int
	InventoryHeight int
}

/*
		Struct built to handle multiple inventories, such as when
		A player opens a chest, and other interactions
*/

type InventoryContext struct{
	Inventories   []*types.Inventory
	Names         []string
	Info 	        []InventoryDimensionInformation
	Positions     []Vector2
	GlobalHand   *types.ItemStack
}

func NewChestInventoryContext(chestInv, playerInv *types.Inventory) *InventoryContext {
	// calculate what they need
	chestInf := CalculatePlayerInventoryDimensionInformation(chestInv.MaxSlots / constants.PlayerInventoryCols,
		constants.PlayerInventoryCols)
	playerInf :=  CalculatePlayerInventoryDimensionInformation(constants.PlayerInventoryRows,
		constants.PlayerInventoryCols)


	totalHeight := playerInf.InventoryHeight + chestInf.InventoryHeight
	w_h := constants.ScreenWidth / 2
	h_h := constants.ScreenHeight / 2

	tw_h := constants.PlayerInventoryWidth / 2
	th_h := totalHeight / 2

	// draw the chest interface above 
	chestUIPos := Vector2{
		X: float32(w_h - tw_h),
		Y: float32(h_h - th_h),
	}
	playerUIPos := Vector2{
		X: float32(w_h - tw_h),
		Y: float32(h_h - th_h + chestInf.InventoryHeight),
	}

	return &InventoryContext{
		Inventories: []*types.Inventory{chestInv, playerInv},
		Names:       []string{"Chest", "Player"},
		Info:        []InventoryDimensionInformation{chestInf, playerInf},
		Positions:   []Vector2{chestUIPos, playerUIPos},
		GlobalHand:  &types.ItemStack{},
	}
}

func (ic *InventoryContext) Update() {
	// update ticks when the inventories are open
	mousePos := GetMousePosition()
	var clickedSlot       int = -1
	var clickedSlotInvInd int = -1
	var clickedSlotInv    *types.Inventory

	// check which slot was clicked
	for invIndex, inv := range ic.Inventories{
		// positions
		pos := ic.Positions[invIndex]
		// dimensions
		dim := ic.Info[invIndex]

		// see which inventory the mouse is over
		if ic.IsMouseOverInventory(mousePos, pos, dim){
			slot := ic.GetSlotFromMousePos(mousePos, pos, dim, invIndex)
			if slot > -1 {
				clickedSlot       = slot
				clickedSlotInvInd = invIndex
				clickedSlotInv    = inv
			}
		}
	}

	// now worry about the clicks

	if IsMouseButtonPressed(MouseLeftButton) && clickedSlotInv != nil {
		ic.HandleLeftClick(clickedSlotInvInd, clickedSlot, clickedSlotInv)

	}

	if IsMouseButtonPressed(MouseRightButton) && clickedSlotInv != nil {
		// TODO
		// ic.HandleRightClick(clickedSlotInvInd, clickedSlot, clickedSlotInv)
		fmt.Println("Right click")
	}


	// now draw
	ic.draw()
}

func (ic *InventoryContext) draw() {
	for invIndex, inv := range ic.Inventories {
		pos :=        ic.Positions[invIndex]
		dim :=        ic.Info[invIndex]
		hasDivider := ic.Names[invIndex] == "Player"

		DrawInventory(inv, pos, false, hasDivider, dim)
	}
}

func (ic *InventoryContext) HandleLeftClick(invIndex, slot int, inv *types.Inventory) {
	// the selected slot
	activeSlot := inv.Slots[slot]

	
	if ic.GlobalHand == nil { // player has nothign in hand 
		// try to pick up item if there is somethign there
		if activeSlot.Item != nil {
			ic.GlobalHand = &activeSlot
			activeSlot = types.ItemStack{}
		}
		
	} else { // player has something in hand
		// empty slot, place item
		if activeSlot.Item == nil {
			activeSlot.Item = ic.GlobalHand.Item
			ic.GlobalHand = &types.ItemStack{}
		} else if activeSlot.Item.GetId() == ic.GlobalHand.Item.GetId() {
			// same item as in hand, fill as many as possible into the inv slot
			maxItems := activeSlot.Item.GetMaxStackSize()
			diff := maxItems - activeSlot.Count

			toAdd := min(diff, ic.GlobalHand.Count)

			activeSlot.Count += toAdd
			ic.GlobalHand.Count -= toAdd
			
			// check if the hand is now empty
			if ic.GlobalHand.Count <= 0{
				ic.GlobalHand = &types.ItemStack{}
			}
		} else { // player has opposite item in hand, swap
			temp := activeSlot
			activeSlot = *ic.GlobalHand
			ic.GlobalHand = &temp
		}
	}
}

func (ic *InventoryContext) HandleRightClick(clickedSlotInvInd, clickedSlot int, clickedSlotInv *types.Inventory) {

}

/*
-------------------Helpers-----------------
*/



func min(val1, val2 int) int {
	if val1 <= val2 {
		return val1
	}
	return val2
}

func (ic *InventoryContext) GetSlotFromMousePos(mousePos, invPos Vector2,
	dim InventoryDimensionInformation, invIndex int) int {

	   // Convert mouse position to slot index
    relativeX := int(mousePos.X - invPos.X)
    relativeY := int(mousePos.Y - invPos.Y)
    
    // Calculate which slot based on relative position
    col := (relativeX - constants.PlayerInventorySlotMargin) / (dim.SlotSize + constants.PlayerInventorySlotMargin)
    row := (relativeY - constants.PlayerInventorySlotMargin) / (dim.SlotSize + constants.PlayerInventorySlotMargin)
    
    if col < 0 || col >= constants.PlayerInventoryCols {
        return -1
    }
    
    maxRows := ic.Inventories[invIndex].MaxSlots / constants.PlayerInventoryCols
    if row < 0 || row >= maxRows {
        return -1
    }
    
    return row*constants.PlayerInventoryCols + col
}

func (ic *InventoryContext) IsMouseOverInventory (mousePos,
	pos Vector2, dim InventoryDimensionInformation) bool {

	return PointInRectangle(mousePos, pos,
		dim.InventoryHeight, constants.PlayerInventoryWidth)
}

/*
                might be the longest function name ever
*/
func CalculatePlayerInventoryDimensionInformation(rows, cols int) InventoryDimensionInformation{
	// divider between inventory and hotbar items
	dividerHeight := constants.PlayerInventorySlotMargin * 3.0
	// total margin space on y axis (numRows + 1)
	marginSpaceY := float64(rows * constants.PlayerInventorySlotMargin) + dividerHeight
	marginSpaceX := cols * constants.PlayerInventorySlotMargin

	// slot size
	slotSize := (constants.PlayerInventoryWidth - (marginSpaceX)) / constants.PlayerInventoryCols

	// calculate what the height should be
	PlayerInventoryHeight := int(marginSpaceY) + (slotSize * rows)

	return InventoryDimensionInformation{
		DividerHeight: dividerHeight,
		MarginSpaceY: marginSpaceY,
		MarginSpaceX: marginSpaceX,
		SlotSize: slotSize,
		InventoryHeight: PlayerInventoryHeight,
	}
}
