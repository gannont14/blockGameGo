package types

import "fmt"

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
	fmt.Println("Openign chest")
	return true
}

