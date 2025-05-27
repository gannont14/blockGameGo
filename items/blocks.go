package items

import (
	"blockProject/textures"
	"blockProject/types"

	. "github.com/gen2brain/raylib-go/raylib"
)

/*
   Factories for all of the items that need to be registered to the
   registry
*/

// WHY ARE INTERFACES NOT EXPLICIT

// redBlock
func NewAirBlockItem() *types.BlockItem {
  return &types.BlockItem{
    BaseItem: types.BaseItem{
      Id: types.GetNewItemID(), 
      Name: "Air Block",
      MaxStackSize: -1,
    },
    Type: types.Air,
    PrefToolType: types.ToolTypePickaxe,
    BaseBreakTime: -1,
    BlockHardness: types.BaseLevel,
    Color: Blank,
		AtlasPosition: textures.TextureAtlasPosition{},
  }
}

// redBlock
func NewRedBlockItem() *types.BlockItem {
  return &types.BlockItem{
    BaseItem: types.BaseItem{
      Id: types.GetNewItemID(), 
      Name: "Red Block",
      MaxStackSize: 64,
    },
    Type: types.RedBlock,
    PrefToolType: types.ToolTypePickaxe,
    BaseBreakTime: 3.0,
    BlockHardness: types.BaseLevel,
    Color: Red,
		AtlasPosition: textures.NewTextureAtlasPosition(0, 0),
  }
}

// blue block ( the red block, but blue! )
func NewBlueBlockItem() *types.BlockItem {
  return &types.BlockItem{
    BaseItem: types.BaseItem{
      Id: types.GetNewItemID(),
      Name: "Blue Block",
      MaxStackSize: 64,
    },
    Type: types.BlueBlock,
    PrefToolType: types.ToolTypeNone,
    BaseBreakTime: 3.0,
    BlockHardness: types.BaseLevel,
		Color: Blue,
		AtlasPosition: textures.NewTextureAtlasPosition(0, 1),
  }
}

func NewGreenBlockItem() *types.BlockItem {
	return &types.BlockItem{
		BaseItem: types.BaseItem{
			Id: types.GetNewItemID(),
			Name: "Green Block",
			MaxStackSize: 64,
		},
		Type: types.GreenBlock,
		PrefToolType: types.ToolTypeShovel,
		BaseBreakTime: 3.0,
		BlockHardness: types.BaseLevel,
		Color: Green,
		AtlasPosition: textures.NewTextureAtlasPosition(0, 2),
	}
}


func NewBrownBlockItem() *types.BlockItem {
	return &types.BlockItem{
		BaseItem: types.BaseItem{
			Id: types.GetNewItemID(),
			Name: "Brown Block",
			MaxStackSize: 64,
		},
		Type: types.BrownBlock,
		PrefToolType: types.ToolTypeAxe,
		BaseBreakTime: 3.0,
		BlockHardness: types.BaseLevel,
		Color: Brown,
		AtlasPosition: textures.NewTextureAtlasPosition(0, 3),
	}
}


func NewBlackBlockItem() *types.BlockItem {
	return &types.BlockItem{
		BaseItem: types.BaseItem{
			Id: types.GetNewItemID(),
			Name: "Black Block",
			MaxStackSize: 64,
		},
		Type: types.BlackBlock,
		PrefToolType: types.ToolTypePickaxe,
		BaseBreakTime: 3.0,
		BlockHardness: types.GoldLevel,
		Color: Black,
		AtlasPosition: textures.NewTextureAtlasPosition(1, 0),
	}
}
