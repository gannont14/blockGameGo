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
      ID: types.GetNewItemID(), // will this work without static?
      Name: "Red Block",
      MaxStackSize: 64,
    },
    Type: types.RedBlock,
  }
}

// blue block ( the red block, but blue! )
func NewBlueBlockItem() *types.BlockItem {
  return &types.BlockItem{
    BaseItem: types.BaseItem{
      ID: types.GetNewItemID(),
      Name: "Blue Block",
      MaxStackSize: 64,
    },
    Type: types.BlueBlock,
  }
}
