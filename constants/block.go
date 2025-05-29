package constants


const (
  // block idmension
  BlockWidth  = 2  //X
  BlockDepth  = 2  //Y
  BlockHeight = 2 //Z

  // references to which face of block is being referred to
  BlockHitFaceNegX = 0
  BlockHitFacePosX = 1
  BlockHitFaceNegY = 2
  BlockHitFacePosY = 3
  BlockHitFaceNegZ = 4
  BlockHitFacePosZ = 5

  // Block breaking information
  DefaultBreakTime = 3.0

	// item related consts, could be own file
	DefaultSharpness = 2.0

	// block specific
	SingleChestInventorySize = 27
)
