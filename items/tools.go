package items

import (
	"blockProject/types"
)

// Actual Tools now

/*
   Axes
*/

func NewWoodenAxeItem() *types.ToolItem {
	return &types.ToolItem{
		BaseItem: types.BaseItem{
			Id:           types.GetNewItemID(),
			Name:         "Wooden Axe",
			MaxStackSize: 1,
		},
		Durability: 200,
		ToolType:   types.ToolTypeAxe,
		ToolLevel:  types.ToolLevelWooden,
		Speed:      2,
	}
}

func NewStoneAxeItem() *types.ToolItem {
	return &types.ToolItem{
		BaseItem: types.BaseItem{
			Id:           types.GetNewItemID(),
			Name:         "Stone Axe",
			MaxStackSize: 1,
		},
		Durability: 400,
		ToolType:   types.ToolTypeAxe,
		ToolLevel:  types.ToolLevelStone,
		Speed:      3,
	}
}

func NewIronAxeItem() *types.ToolItem {
	return &types.ToolItem{
		BaseItem: types.BaseItem{
			Id:           types.GetNewItemID(),
			Name:         "Iron Axe",
			MaxStackSize: 1,
		},
		Durability: 800,
		ToolType:   types.ToolTypeAxe,
		ToolLevel:  types.ToolLevelIron,
		Speed:      4,
	}
}

func NewGoldAxeItem() *types.ToolItem {
	return &types.ToolItem{
		BaseItem: types.BaseItem{
			Id:           types.GetNewItemID(),
			Name:         "Gold Axe",
			MaxStackSize: 1,
		},
		Durability: 200,
		ToolType:   types.ToolTypeAxe,
		ToolLevel:  types.ToolLevelGold,
		Speed:      5,
	}
}

func NewDiamondAxeItem() *types.ToolItem {
	return &types.ToolItem{
		BaseItem: types.BaseItem{
			Id:           types.GetNewItemID(),
			Name:         "Diamond Axe",
			MaxStackSize: 1,
		},
		Durability: 3200,
		ToolType:   types.ToolTypeAxe,
		ToolLevel:  types.ToolLevelDiamond,
		Speed:      6,
	}
}

/*
   Pickaxes
*/

func NewWoodenPickaxeItem() *types.ToolItem {
	return &types.ToolItem{
		BaseItem: types.BaseItem{
			Id:           types.GetNewItemID(),
			Name:         "Wooden Pickaxe",
			MaxStackSize: 1,
		},
		Durability: 200,
		ToolType:   types.ToolTypePickaxe,
		ToolLevel:  types.ToolLevelWooden,
		Speed:      2,
	}
}

func NewStonePickaxeItem() *types.ToolItem {
	return &types.ToolItem{
		BaseItem: types.BaseItem{
			Id:           types.GetNewItemID(),
			Name:         "Stone Pickaxe",
			MaxStackSize: 1,
		},
		Durability: 400,
		ToolType:   types.ToolTypePickaxe,
		ToolLevel:  types.ToolLevelStone,
		Speed:      3,
	}
}

func NewIronPickaxeItem() *types.ToolItem {
	return &types.ToolItem{
		BaseItem: types.BaseItem{
			Id:           types.GetNewItemID(),
			Name:         "Iron Pickaxe",
			MaxStackSize: 1,
		},
		Durability: 800,
		ToolType:   types.ToolTypePickaxe,
		ToolLevel:  types.ToolLevelIron,
		Speed:      4,
	}
}

func NewGoldPickaxeItem() *types.ToolItem {
	return &types.ToolItem{
		BaseItem: types.BaseItem{
			Id:           types.GetNewItemID(),
			Name:         "Gold Pickaxe",
			MaxStackSize: 1,
		},
		Durability: 200,
		ToolType:   types.ToolTypePickaxe,
		ToolLevel:  types.ToolLevelGold,
		Speed:      5,
	}
}

func NewDiamondPickaxeItem() *types.ToolItem {
	return &types.ToolItem{
		BaseItem: types.BaseItem{
			Id:           types.GetNewItemID(),
			Name:         "Diamond Pickaxe",
			MaxStackSize: 1,
		},
		Durability: 3200,
		ToolType:   types.ToolTypePickaxe,
		ToolLevel:  types.ToolLevelDiamond,
		Speed:      6,
	}
}
