package blocks

import (
	"blockProject/constants"
	"blockProject/textures" // textures
	"blockProject/types"    // types

	. "github.com/gen2brain/raylib-go/raylib" // for the colors
)


func NewChestBlockItem() *types.BlockItem {
  return &types.BlockItem{
    BaseItem: types.BaseItem{
      Id: types.GetNewItemID(), 
      Name: "Chest",
      MaxStackSize: 64,
			AtlasPosition: textures.NewTextureAtlasPosition(0, 2),
    },
    Type: types.ChestBlock,
    PrefToolType: types.ToolTypePickaxe,
    BaseBreakTime: -1,
    BlockHardness: types.BaseLevel,
    Color: Blank,
		AtlasPosition: textures.NewTextureAtlasPosition(0, 0),
		Interactable: types.NewChestInteractor(constants.SingleChestInventorySize),
  }
}
