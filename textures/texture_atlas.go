package textures

import (
	"fmt"

	. "github.com/gen2brain/raylib-go/raylib"
)

type TextureAtlasPosition struct {
	Row int
	Col int
}

func NewTextureAtlasPosition(row, col int) TextureAtlasPosition{
	return TextureAtlasPosition{
		Row: row,
		Col: col,
	}
}

func (tap *TextureAtlasPosition) Val() (int, int){
	return tap.Row, tap.Col
}

type TextureAtlas struct {
	FilePath   string
	NumBlocksX int
	NumBlocksY int
	BlockSize  int
	Texture Texture2D
}

func NewTextureAtlas(filePath string, numBlocksX, numBlocksY, blockSize int) TextureAtlas {
	// create the texture object 
	tex := LoadTexture(filePath)
	fmt.Print("Loaded Texture")
	fmt.Printf("Texture ID: %d\n", tex.ID)
	fmt.Printf("Texture Width: %d\n", tex.Width)
	fmt.Printf("Texture Height: %d\n", tex.Height)
	fmt.Printf("Texture Format: %d\n", tex.Format)

	atlas := TextureAtlas{
		FilePath: filePath,
		NumBlocksX: numBlocksX,
		NumBlocksY: numBlocksY,
		BlockSize: blockSize,
		Texture: tex,
	}

	return atlas
}

func (ta *TextureAtlas) GetTextureAtlasOrig(atlasPos TextureAtlasPosition) Vector2{
	// find how much X a block takes up
	bw := ta.BlockSize * 3 // 3 faces horizontally
	// find how much Y a block takes up
	bh := ta.BlockSize * 2 // 2 faces vertically

	x := float32(bw * atlasPos.Col)
	y := float32(bh * atlasPos.Row)

	// fmt.Printf("Received texture from atlas starting at : [%f, %f]\n", x, y)

	return NewVector2(x, y)
}
/* 
	Gets the coords for each corner of each face for a given origin in the atlas
*/

func (ta *TextureAtlas) RetrieveBlockTextureMap (orig Vector2) BlockTextureMap {
	// takes in row, col to get corner points
	getTexCoords := func(x, y int) [4]Vector2 {
		w := float32(ta.BlockSize)
		h := float32(ta.BlockSize)

		// calculate pixel coordinates
		pixelU := orig.X + float32(x) * w
		pixelV := orig.Y + float32(y) * h

		// convert to normalized UV coordinates => 0-1
		u := pixelU / float32(ta.Texture.Width)
		v := pixelV / float32(ta.Texture.Height)
		uW := w / float32(ta.Texture.Width)
		vH := h / float32(ta.Texture.Height)

		return [4]Vector2{
			{X: u, Y: v},           // Top left
			{X: u + uW, Y: v},      // Top right
			{X: u, Y: v + vH},      // Bottom left
			{X: u + uW, Y: v + vH}, // Bottom right
		}
	}

	// now use this function to get all of the faces
	faceTexCoords := map[string][4]Vector2{
		"top":    getTexCoords(0, 0),
		"bottom": getTexCoords(1, 0),
		"left":   getTexCoords(2, 0),
		"right":  getTexCoords(0, 1),
		"front":  getTexCoords(1, 1),
		"back":   getTexCoords(2, 1),
	}

// // debug the UV coordinates
// 	for face, coords := range faceTexCoords {
// 		fmt.Printf("Face %s: UV coords [(%f,%f), (%f,%f), (%f,%f), (%f,%f)]\n", 
// 			face, coords[0].X, coords[0].Y, coords[1].X, coords[1].Y,
// 			coords[2].X, coords[2].Y, coords[3].X, coords[3].Y)
// 	}

	return faceTexCoords
}
