package items

import (
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
      Id: types.GetNewItemID(), // will this work without static?
      Name: "Air Block",
      MaxStackSize: -1,
    },
    Type: types.Air,
    PrefToolType: types.ToolTypePickaxe,
    BaseBreakTime: -1,
    BlockHardness: types.BaseLevel,
    Color: Blank,
  }
}

// redBlock
func NewRedBlockItem() *types.BlockItem {
  return &types.BlockItem{
    BaseItem: types.BaseItem{
      Id: types.GetNewItemID(), // will this work without static?
      Name: "Red Block",
      MaxStackSize: 64,
    },
    Type: types.RedBlock,
    PrefToolType: types.ToolTypePickaxe,
    BaseBreakTime: 3.0,
    BlockHardness: types.BaseLevel,
    Color: Red,
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
	}
}
