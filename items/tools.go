package items

import (
	"blockProject/constants"
	"blockProject/types"
	"blockProject/textures"
)

// Actual Tools now

/*
   Shovels
*/

func NewWoodenShovelItem() *types.ToolItem {
	return &types.ToolItem{
		BaseItem: types.BaseItem{
			Id:           types.GetNewItemID(),
			Name:         "Wooden Shovel",
			MaxStackSize: 1,
			AtlasPosition: textures.NewTextureAtlasPosition(0, 0),
		},
		Durability: 200,
		MaxDurability: 200,
		ToolType:   types.ToolTypeShovel,
		ToolLevel:  types.ToolLevelWooden,
		Speed:      2,
		Sharpness: constants.DefaultSharpness,
	}
}

func NewStoneShovelItem() *types.ToolItem {
	return &types.ToolItem{
		BaseItem: types.BaseItem{
			Id:           types.GetNewItemID(),
			Name:         "Stone Shovel",
			MaxStackSize: 1,
			AtlasPosition: textures.NewTextureAtlasPosition(0, 0),
		},
		Durability: 400,
		MaxDurability: 400,
		ToolType:   types.ToolTypeShovel,
		ToolLevel:  types.ToolLevelStone,
		Speed:      3,
		Sharpness: constants.DefaultSharpness,
	}
}

func NewIronShovelItem() *types.ToolItem {
	return &types.ToolItem{
		BaseItem: types.BaseItem{
			Id:           types.GetNewItemID(),
			Name:         "Iron Shovel",
			MaxStackSize: 1,
			AtlasPosition: textures.NewTextureAtlasPosition(0, 0),
		},
		Durability: 800,
		MaxDurability: 800,
		ToolType:   types.ToolTypeShovel,
		ToolLevel:  types.ToolLevelIron,
		Speed:      4,
		Sharpness: constants.DefaultSharpness,
	}
}

func NewGoldShovelItem() *types.ToolItem {
	return &types.ToolItem{
		BaseItem: types.BaseItem{
			Id:           types.GetNewItemID(),
			Name:         "Gold Shovel",
			MaxStackSize: 1,
			AtlasPosition: textures.NewTextureAtlasPosition(0, 0),
		},
		Durability: 200,
		MaxDurability: 200,
		ToolType:   types.ToolTypeShovel,
		ToolLevel:  types.ToolLevelGold,
		Speed:      5,
		Sharpness: constants.DefaultSharpness,
	}
}

func NewDiamondShovelItem() *types.ToolItem {
	return &types.ToolItem{
		BaseItem: types.BaseItem{
			Id:           types.GetNewItemID(),
			Name:         "Diamond Shovel",
			MaxStackSize: 1,
			AtlasPosition: textures.NewTextureAtlasPosition(0, 0),
		},
		Durability: 3200,
		MaxDurability: 3200,
		ToolType:   types.ToolTypeShovel,
		ToolLevel:  types.ToolLevelDiamond,
		Speed:      6,
		Sharpness: constants.DefaultSharpness,
	}
}

/*
   Axes
*/

func NewWoodenAxeItem() *types.ToolItem {
	return &types.ToolItem{
		BaseItem: types.BaseItem{
			Id:           types.GetNewItemID(),
			Name:         "Wooden Axe",
			MaxStackSize: 1,
			AtlasPosition: textures.NewTextureAtlasPosition(0, 0),
		},
		Durability: 200,
		MaxDurability: 200,
		ToolType:   types.ToolTypeAxe,
		ToolLevel:  types.ToolLevelWooden,
		Speed:      2,
		Sharpness: 5.0,
	}
}

func NewStoneAxeItem() *types.ToolItem {
	return &types.ToolItem{
		BaseItem: types.BaseItem{
			Id:           types.GetNewItemID(),
			Name:         "Stone Axe",
			MaxStackSize: 1,
			AtlasPosition: textures.NewTextureAtlasPosition(0, 0),
		},
		Durability: 400,
		MaxDurability: 400,
		ToolType:   types.ToolTypeAxe,
		ToolLevel:  types.ToolLevelStone,
		Speed:      3,
		Sharpness: 5.0,
	}
}

func NewIronAxeItem() *types.ToolItem {
	return &types.ToolItem{
		BaseItem: types.BaseItem{
			Id:           types.GetNewItemID(),
			Name:         "Iron Axe",
			MaxStackSize: 1,
			AtlasPosition: textures.NewTextureAtlasPosition(0, 0),
		},
		Durability: 800,
		MaxDurability: 800,
		ToolType:   types.ToolTypeAxe,
		ToolLevel:  types.ToolLevelIron,
		Speed:      4,
		Sharpness: 5.0,
	}
}

func NewGoldAxeItem() *types.ToolItem {
	return &types.ToolItem{
		BaseItem: types.BaseItem{
			Id:           types.GetNewItemID(),
			Name:         "Gold Axe",
			MaxStackSize: 1,
			AtlasPosition: textures.NewTextureAtlasPosition(0, 0),
		},
		Durability: 200,
		MaxDurability: 200,
		ToolType:   types.ToolTypeAxe,
		ToolLevel:  types.ToolLevelGold,
		Speed:      5,
		Sharpness: 5.0,
	}
}

func NewDiamondAxeItem() *types.ToolItem {
	return &types.ToolItem{
		BaseItem: types.BaseItem{
			Id:           types.GetNewItemID(),
			Name:         "Diamond Axe",
			MaxStackSize: 1,
			AtlasPosition: textures.NewTextureAtlasPosition(0, 0),
		},
		Durability: 3200,
		MaxDurability: 3200,
		ToolType:   types.ToolTypeAxe,
		ToolLevel:  types.ToolLevelDiamond,
		Speed:      6,
		Sharpness: 5.0,
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
			AtlasPosition: textures.NewTextureAtlasPosition(0, 0),
		},
		Durability: 200,
		MaxDurability: 200,
		ToolType:   types.ToolTypePickaxe,
		ToolLevel:  types.ToolLevelWooden,
		Speed:      2,
		Sharpness: constants.DefaultBreakTime,
	}
}

func NewStonePickaxeItem() *types.ToolItem {
	return &types.ToolItem{
		BaseItem: types.BaseItem{
			Id:           types.GetNewItemID(),
			Name:         "Stone Pickaxe",
			MaxStackSize: 1,
			AtlasPosition: textures.NewTextureAtlasPosition(0, 0),
		},
		Durability: 400,
		MaxDurability: 400,
		ToolType:   types.ToolTypePickaxe,
		ToolLevel:  types.ToolLevelStone,
		Speed:      3,
		Sharpness: constants.DefaultBreakTime,
	}
}

func NewIronPickaxeItem() *types.ToolItem {
	return &types.ToolItem{
		BaseItem: types.BaseItem{
			Id:           types.GetNewItemID(),
			Name:         "Iron Pickaxe",
			MaxStackSize: 1,
			AtlasPosition: textures.NewTextureAtlasPosition(0, 0),
		},
		Durability: 800,
		MaxDurability: 800,
		ToolType:   types.ToolTypePickaxe,
		ToolLevel:  types.ToolLevelIron,
		Speed:      4,
		Sharpness: constants.DefaultBreakTime,
	}
}

func NewGoldPickaxeItem() *types.ToolItem {
	return &types.ToolItem{
		BaseItem: types.BaseItem{
			Id:           types.GetNewItemID(),
			Name:         "Gold Pickaxe",
			MaxStackSize: 1,
			AtlasPosition: textures.NewTextureAtlasPosition(0, 0),
		},
		Durability: 10,
		MaxDurability: 10,
		ToolType:   types.ToolTypePickaxe,
		ToolLevel:  types.ToolLevelGold,
		Speed:      5,
		Sharpness: constants.DefaultBreakTime,
	}
}

func NewDiamondPickaxeItem() *types.ToolItem {
	return &types.ToolItem{
		BaseItem: types.BaseItem{
			Id:           types.GetNewItemID(),
			Name:         "Diamond Pickaxe",
			MaxStackSize: 1,
			AtlasPosition: textures.NewTextureAtlasPosition(0, 0),
		},
		Durability: 3200,
		MaxDurability: 3200,
		ToolType:   types.ToolTypePickaxe,
		ToolLevel:  types.ToolLevelDiamond,
		Speed:      6,
		Sharpness: constants.DefaultBreakTime,
	}
}
