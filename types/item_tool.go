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
  ToolType ToolType
  ToolLevel ToolLevel 
  BreakLevel BreakLevel
  Speed float64 // could also be added to the tooltype type, type type type
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
  ToolTypeNone ToolType = iota
  ToolTypeAxe
  ToolTypePickaxe
  ToolTypeSword
  ToolTypeShovel
  ToolTypeHoe
)

type ToolLevel int 

const (
	ToolLevelBase ToolLevel = iota
	ToolLevelWooden
	ToolLevelStone
	ToolLevelIron
	ToolLevelGold
	ToolLevelDiamond
)
