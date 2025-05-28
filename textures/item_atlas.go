package textures

import (
	. "github.com/gen2brain/raylib-go/raylib"
	"fmt"
)

/*
		atlas for the item models
		This could for sure be changed to fit some sort of interface like the block one
		so that i dont have two interfaces (interfaci?), but oh well
*/

type ItemAtlas struct {
	// config settings
	FilePath string
	NumItemsX int  // num of items the the atlas spans in the X
	NumItemsY int  // num of items the the atlas spans in the X
	ItemSize int   // size of each item
	Texture Texture2D
}

func NewItemAtlas(filePath string, nIX, nIY, iSize int) ItemAtlas {
	// create the texture object 
	tex := LoadTexture(filePath)
	// debug print
	fmt.Print("Loaded Texture")
	fmt.Printf("Texture ID: %d\n", tex.ID)
	fmt.Printf("Texture Width: %d\n", tex.Width)
	fmt.Printf("Texture Height: %d\n", tex.Height)
	fmt.Printf("Texture Format: %d\n", tex.Format)

	return ItemAtlas{
		FilePath: filePath,
		NumItemsX: nIX,
		NumItemsY: nIY,
		ItemSize: iSize,
		Texture: tex,
	}
}

func (ia *ItemAtlas) GetRectanglePos(tap TextureAtlasPosition) Rectangle {
	w := float32(ia.ItemSize)
	h := float32(ia.ItemSize)

	// holy type casting
	row, col := tap.Val()
	x := (float32(col) * w)
	y := (float32(row) * h)

	// fmt.Printf("Item rec  is at [%f %f], with h,w of [%f, %f]\n",
	// 	x, y, h, w)

	return Rectangle{
		X: x,
		Y: y,
		Width: float32(ia.ItemSize),
		Height: float32(ia.ItemSize),
	}
}
