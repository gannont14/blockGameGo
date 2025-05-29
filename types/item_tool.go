package types



type ToolItem struct{
  BaseItem
  Durability int 
	MaxDurability int
  ToolType ToolType
  ToolLevel ToolLevel 
  Speed float64 // could also be added to the tooltype type, type type type
	Sharpness float64
}

func (t *ToolItem) Clone() Item {
	clone := *t
	return &clone
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
    t.ToolLevel = -1 
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
