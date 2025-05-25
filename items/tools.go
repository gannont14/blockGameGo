package items

import (
  "blockProject/types"
)



// Actual Tools now

/*
    Pickaxes
*/


func NewWoodenPickaxeItem() *types.ToolItem {
  return &types.ToolItem{
    BaseItem: types.BaseItem{
      Id: types.GetNewItemID(),
      Name: "Wooden Pickaxe",
      MaxStackSize: 1,
    },
    Durability: 200,
    ToolType: types.ToolWoodenPickaxe,
    BreakLevel: types.BreakLevelWood,
    Speed: 2,
  }
}


func NewStonePickaxeItem() *types.ToolItem {
  return &types.ToolItem{
    BaseItem: types.BaseItem{
      Id: types.GetNewItemID(),
      Name: "Stone Pickaxe",
      MaxStackSize: 1,
    },
    Durability: 400,
    ToolType: types.ToolStonePickaxe,
    BreakLevel: types.BreakLevelWood,
    Speed: 3,
  }
}


func NewIronPickaxeItem() *types.ToolItem {
  return &types.ToolItem{
    BaseItem: types.BaseItem{
      Id: types.GetNewItemID(),
      Name: "Iron Pickaxe",
      MaxStackSize: 1,
    },
    Durability: 800,
    ToolType: types.ToolIronPickaxe,
    BreakLevel: types.BreakLevelWood,
    Speed: 4,
  }
}


func NewGoldPickaxeItem() *types.ToolItem {
  return &types.ToolItem{
    BaseItem: types.BaseItem{
      Id: types.GetNewItemID(),
      Name: "Gold Pickaxe",
      MaxStackSize: 1,
    },
    Durability: 1600,
    ToolType: types.ToolGoldPickaxe,
    BreakLevel: types.BreakLevelWood,
    Speed: 5,
  }
}

func NewDiamondPickaxeItem() *types.ToolItem {
  return &types.ToolItem{
    BaseItem: types.BaseItem{
      Id: types.GetNewItemID(),
      Name: "Diamond Pickaxe",
      MaxStackSize: 1,
    },
    Durability: 3200,
    ToolType: types.ToolDiamondPickaxe,
    BreakLevel: types.BreakLevelWood,
    Speed: 6,
  }
}
