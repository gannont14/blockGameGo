package types

import (
	"blockProject/textures"
	"blockProject/types"
	"sync"
)

/*
	singleton for shared game state
*/
type GameState struct {
	Player *types.Player
	World  *types.World
	BreakingManager *types.BreakingManager
	ItemAtlas *textures.ItemAtlas
}

var (
	instance *GameState
	once     sync.Once
)

func Get() *GameState {
	once.Do(func() {
		instance = &GameState{}
	})

	return instance
}

func GameStateInit (p *types.Player,
	w *types.World,
	bm *types.BreakingManager,
	ia *textures.ItemAtlas) {
	// state
	s := Get()

	// set vals
	s.Player          = p
	s.World           = w
	s.BreakingManager = bm
	s.ItemAtlas       = ia
}

