package types

// interaction type
type InteractionType int 
const (
  LEFT_CLICK = iota
  RIGHT_CLICK
)

type InteractionContext struct {
  World             *World
  Player            *Player
  FocusedBlock      *Block
  PotentialBlock    *Block
  InteractionType   InteractionType
}

// generator
func NewInteractionContext(
  world             *World,
  player            *Player,
  focusedBlock      *Block,
  potentialBlock    *Block,
  interactionType   InteractionType,
) InteractionContext {
  return InteractionContext{
    World: world,
    Player: player,
    FocusedBlock: focusedBlock,
    PotentialBlock: potentialBlock,
    InteractionType: interactionType,
  }
}
