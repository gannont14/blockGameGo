package types

import (
	"blockProject/constants"
	"fmt"

	. "github.com/gen2brain/raylib-go/raylib"
)

/*
  Couldn't really think of a better way to handle this without have to do some state machine
  type thing, this could change for sure
*/

type BreakingManager struct{
  // item registry
  ItemRegistry *ItemRegistry

  // default config, could map players eventually?
  ActiveBreaking map[*Player]*BreakingInfo

  // default config
  defaultBreakingTime float32
  // .. others for sound, particles

}
/*
  Generator for BreakingManager
*/
func NewBreakingManager(itemRegistry ItemRegistry) *BreakingManager{
  return &BreakingManager{
    ItemRegistry: &itemRegistry,
    ActiveBreaking: make(map[*Player]*BreakingInfo),
    defaultBreakingTime: constants.DefaultBreakTime,
  }
}

type BreakingInfo struct {
  // block to break
  TargetBlock *Block

  // time info
  StartTime float64
  Duration float64
  Progress float32

  // item used, anythign not a tool can be default
  ToolUsed Item

  // ... Audio and particle states
}
/*
  Generator for BreakingManager
*/
func NewBreakingInfo( targetBlock *Block, startTime float64, duration float64, tool Item) *BreakingInfo {
  return &BreakingInfo{
    TargetBlock: targetBlock,
    StartTime: startTime,
    Duration: duration,
    Progress: 0.0,
    ToolUsed: tool,
  }
}


/*
  Core Operations --------------------------------
*/


/*
  Starts breaking the block
*/

func (bm *BreakingManager) StartBreaking(player *Player, block *Block){
  // find start time
  currentTime := GetTime()

  // find tool they are holding
  activeSlot := player.ActiveItemSlot
  var tool Item = nil
  if player.Inventory.Slots[activeSlot].Count > 0{
    tool = player.Inventory.Slots[activeSlot].Item
  }

  // find the duration it's going to take
  duration := bm.calculateBreakTime(block, tool)

  // create info object
  bm.ActiveBreaking[player] = NewBreakingInfo(
    block,
    currentTime,
    duration,
    tool,)
}

func (bm *BreakingManager) UpdateBreaking(p *Player, w *World) BreakingResult {
  // info map of active breaks 
  info, isBreaking := bm.ActiveBreaking[p]
  if !isBreaking {
    return BreakingResultNone
  }

  currentTime := GetTime()
  elapsed := currentTime - info.StartTime

  info.Progress = float32(elapsed / info.Duration)
  fmt.Println("Breaking progress: ", info.Progress, "%")

  // check if break is done
  if info.Progress >= 1.0 {
    // breasking operation
    bm.CompleteBreaking(p)
    return BreakingResultComplete
  }

  // otherwise it is in Progress
  return BreakingResultInProgress
}

func (bm *BreakingManager) StopBreaking(p *Player){
  if info, wasBreaking := bm.ActiveBreaking[p]; wasBreaking{
    // other information about player done breaking 
    fmt.Println("Player stopped breaking at ", info.Progress, "%")
    delete(bm.ActiveBreaking, p)
  }
}

/*
  final call when player is done breaking a block

*/
func (bm *BreakingManager) CompleteBreaking(p *Player){
  info := bm.ActiveBreaking[p]
  // actually break the block
  p.BreakBlock(info, bm)
}


/*
  Helpers
*/

type BreakingResult int 
const(
  BreakingResultNone BreakingResult = iota
  BreakingResultInProgress
  BreakingResultComplete
)

func (bm *BreakingManager) IsPlayerBreaking(p *Player) bool {
  _, isBreaking := bm.ActiveBreaking[p]
  return isBreaking
}

func (bm *BreakingManager)calculateBreakTime(block *Block, tool Item) float64 {
  if tool == nil {
    return constants.DefaultBreakTime
  }
  // just base for now, will change when tools implemented
  // get the blockItem from the itemManager
  blockItem := bm.ItemRegistry.GetBlockByItemType(block.Type)

  // TODO: implement the effeciency based on tools, could implement switch, or more complex
  // get the tool the player is holding
  // tool, isTool := tool.(*ToolItem)


  

  return blockItem.BaseBreakTime
}

func (bm *BreakingManager) GetBreakingTarget(p *Player) *Block {
  if info, exists := bm.ActiveBreaking[p] ; exists {
    return info.TargetBlock
  }
  return nil
}

func (bm *BreakingManager) GetProgress(p *Player) float32 {
  if info, exists := bm.ActiveBreaking[p] ; exists {
    return info.Progress
  }
  return 0.0
}
