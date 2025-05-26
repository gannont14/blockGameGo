package types

import (
	// "blockProject/items"
	"fmt"
)

/*
   Wanted to add a simple way to register items, and look them up efficiently,
   this can be expanded upon by introducing other ways to look up an item,
   such as by name
*/

type ItemRegistry struct{
  items map[int]Item // items stored by their ID
}

func NewItemRegistry() *ItemRegistry{
  return &ItemRegistry{
    items: make(map[int]Item),
  }
}

func (r *ItemRegistry) RegisterItem (item Item) {
  id := item.GetId()

  // make sure that the item isn't already registed
  if existing, exists := r.items[id]; exists {
    panic(fmt.Sprintf("Duplicate item ID %d: %s and %s",
      id, existing.GetName(), item.GetName()))
  }

  r.items[id] = item
}

func (r *ItemRegistry) GetItemByName(itemName string) (Item, bool) {
	for _, item := range r.items {
		if item.GetName() == itemName { 
			return item, true
		}
	}

	return nil, false
}

func (r *ItemRegistry) GetItemByID(id int) (Item, bool) {
  item, exists := r.items[id]
  if !exists {
    return nil, false
  }

  return item, true
}

func (r *ItemRegistry) GetBlockByItemType(BlockType BlockType) *BlockItem {
  for _, item := range r.items {
    if blockItem, ok := item.(*BlockItem); ok{
      if blockItem.Type == BlockType {
        return blockItem
      }
    }
  }

  return nil
}

func (r *ItemRegistry) GetAllItems() []Item {
  res := make([]Item, 0, len(r.items))

  for _, it := range r.items {
    res = append(res, it)
  }

  return res
}
