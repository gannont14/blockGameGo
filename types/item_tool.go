package types

import (

)

type BreakLevel int

const (
  BreakLevelBase BreakLevel = iota // Base break level for all 
  BreakLevelWood
  BreakLevelStone
  BreakLevelIron
  BreakLevelGold
  BreakLevelDiamond
)


type ToolItem struct{
  BaseItem
  Durability int 
  ToolType ToolType // could be changed to own type
  BreakLevel BreakLevel
  Speed int // could also be added to the tooltype type, type type type
}

func (t *ToolItem) Interact(ctx InteractionContext) bool {
  return false
}

func (t *ToolItem) DegradeTool() {
  // take down Durability
  t.Durability--

  // check if it's 0 or below
  if t.Durability <= 0 {
    // nillify the item, going to just not let it break anything for now, could 
    // actually break it later, could also handle enchantments here
    t.BreakLevel = -1 
  }
}


/*
  Types
*/


type ToolType int 

const (
  //  Wooden tools
  ToolWoodenAxe ToolType = iota
  ToolWoodenPickaxe
  ToolWoodenSword
  ToolWoodenShovel
  ToolWoodenHoe
  //  Stone Tools
  ToolStoneAxe
  ToolStonePickaxe
  ToolStoneSword
  ToolStoneShovel
  ToolStoneHoe
  //  Iron Tools
  ToolIronAxe
  ToolIronPickaxe
  ToolIronSword
  ToolIronShovel
  ToolIronHoe
  //  Gold Tools
  ToolGoldAxe
  ToolGoldPickaxe
  ToolGoldSword
  ToolGoldShovel
  ToolGoldHoe
  //  Diamond Tools
  ToolDiamondAxe
  ToolDiamondPickaxe
  ToolDiamondSword
  ToolDiamondShovel
  ToolDiamondHoe
)
