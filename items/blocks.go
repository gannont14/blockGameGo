package items

import (
  "blockProject/types"
)

/*
   Factories for all of the items that need to be registered to the
   registry
*/

// WHY ARE INTERFACES NOT EXPLICIT

// redBlock
func NewRedBlockItem() *types.BlockItem {
  return &types.BlockItem{
    BaseItem: types.BaseItem{
      Id: types.GetNewItemID(), // will this work without static?
      Name: "Red Block",
      MaxStackSize: 64,
    },
    Type: types.RedBlock,
    BaseBreakTime: 3.0,
    BlockHardness: types.BaseLevel,
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
    BaseBreakTime: 3.0,
    BlockHardness: types.BaseLevel,
  }
}
