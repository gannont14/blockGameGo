package types

import (

)

type ToolItem struct{
  BaseItem
  Durability int 
  ToolType int // could be changed to own type
  Speed int // could also be added to the tooltype type, type type type
}

func (t *ToolItem) Interact(ctx InteractionContext) bool {
  return false
}
