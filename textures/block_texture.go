package textures

import (
	. "github.com/gen2brain/raylib-go/raylib"
)

type BlockTexture struct {
	Atlas *BlockAtlas
	Orig Vector2
	BlockTexturemap BlockTextureMap
}

type BlockTextureMap map[string][4]Vector2 // stores the coords for each of the faces in the atlas

func NewBlockTexture(atlas *BlockAtlas, row, col int) BlockTexture {
	// find out row and col
	orig := atlas.GetTextureAtlasOrig(NewTextureAtlasPosition(row, col))


	// Find the Face coords
	blockTextureMap := atlas.RetrieveBlockTextureMap(orig)

	return BlockTexture{
		Atlas: atlas,
		Orig: orig,
		BlockTexturemap: blockTextureMap,
	}
}
