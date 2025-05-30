package types

import(
	"fmt"
	. "github.com/gen2brain/raylib-go/raylib"
) 

/*
   All interactors for different types of interactable blocks
*/

type Interactable interface {
	OnBlockInteract(ctx InteractionContext) bool
}

/*
---------------------Chests---------------------
*/

type ChestInteractor struct {
	Inventory *Inventory
	IsOpen     bool
}

func NewChestInteractor(invSize int) *ChestInteractor{
	return &ChestInteractor{
		Inventory: NewInventory(invSize, "Chest"),
		IsOpen: false,
	}
}

func (ci *ChestInteractor) OnBlockInteract(ctx InteractionContext) bool {
	player := ctx.Player

	// chest open
	if player.ChestOpen != nil{
		// this shouldn't ever happen? if they right click
		// in a chest it should do the inv stuff
		fmt.Println("Right click in open chest")
		// ci.closeChest(player)
	} else {
		ci.openChest(player)
	}

	return true
}





/*
---------------------Utils---------------------
*/

func (ci *ChestInteractor) openChest(player *Player) {
	EnableCursor()
	player.ChestOpen     = ci.Inventory
	player.CanMoveCamera = false
	player.InventoryOpen = true
}

func (ci *ChestInteractor) closeChest(player *Player) {
	DisableCursor()
	player.ChestOpen     = nil
	player.CanMoveCamera = true
	player.InventoryOpen = false
}
