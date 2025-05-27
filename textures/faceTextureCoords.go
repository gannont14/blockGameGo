package textures

import (
	. "github.com/gen2brain/raylib-go/raylib"
)
/*
	Converts the relative coordinates for the portion of the texture into the 
	actual coordinates of the texture 

	i.e. if you want texture for face with is in top left, calling 
	getTextureCoords(0, 0) => the coords for each corner of the texture

*/

func getTextureCoords(x, y int, tw, th float32) [4]Vector2 {
	w := tw / 3  // Width of each faces texture
	h := th / 2
	u := float32(x) * w
	v := float32(y) * h
	return [4]Vector2 {
		{X: u, Y: v},         // top left
		{X: u + w, Y: v},     // top right
		{X: u, Y: v + h},     // bottom left
		{X: u + w, Y: v + h}, // bottom right
	}
}


