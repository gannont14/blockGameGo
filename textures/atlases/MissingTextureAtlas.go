package textures

import (
	. "github.com/gen2brain/raylib-go/raylib"
)

func GenerateMissingTextureAtlas(texSize int32) *RenderTexture2D {

	renderTex := LoadRenderTexture(texSize, texSize)
	defer UnloadRenderTexture(renderTex)

	// Start drawing to the texture
	BeginTextureMode(renderTex)
	ClearBackground(Black)

	// Draw the checker pattern (2x2)
	squareSize := texSize / 2
	DrawRectangle(0, 0, squareSize, squareSize, Magenta)
	DrawRectangle(squareSize, squareSize, squareSize, squareSize, Magenta)
	DrawRectangle(squareSize, 0, squareSize, squareSize, Purple)
	DrawRectangle(0, squareSize, squareSize, squareSize, Purple)

	EndTextureMode()


	return &renderTex
}
