package server

type IWorld interface {
	GetBlock(Vec3) *Block
	SetBlock(Vec3, *Block)
	GetChunk(Vec2) IChunk
	GetGenerator() IGenerator
	SetGenerator(IGenerator)
	GetWorldBoarder() *WorldBorder
	SetWorldBoarder(*WorldBorder)
	GetDimension() DimensionType
}

type IChunk interface {
	GetBlock(Vec3) *Block
	SetBlock(Vec3, *Block)
}

type WorldBorder struct {
	sendOnJoin bool

	originX                Double
	originZ                Double
	oldRadius              Double
	newRadius              Double // current size
	speed                  int64
	portalTeleportBoundary int32
	warningTime            int32
	warningBlocks          int32
}

func CreateWorldBorder(sendOnJoin bool, originX Double, originZ Double, oldRadius Double, newRadius Double, speed int64, portalTeleportBoundary int32, warningTime int32, warningBlocks int32) *WorldBorder {
	return &WorldBorder{sendOnJoin: sendOnJoin,
		originX:                originX,
		originZ:                originZ,
		oldRadius:              oldRadius,
		newRadius:              newRadius,
		speed:                  speed,
		portalTeleportBoundary: portalTeleportBoundary,
		warningTime:            warningTime,
		warningBlocks:          warningBlocks,
	}
}

// SetSize The param players is the players to update the boarder for. Leave nil to send the packet to everyone in this world. Changes will only save if it's sending to all players. Vice versa, changes will not affect the fields if the update is sent to selected players.
func (w *WorldBorder) SetSize(size Double, players []byte) {
	if players == nil {
		w.oldRadius = 0
		w.newRadius = size
		w.speed = 0
	}

}

// GetSize Returns the latest radius.
func (w *WorldBorder) GetSize() Double {
	return w.newRadius
}

// Lerp The param players is the players to update the boarder for. Leave nil to send the packet to everyone in this world. Changes will only save if it's sending to all players. Vice versa, changes will not affect the fields if the update is sent to selected players.
func (w *WorldBorder) Lerp(before Double, after Double, speed int64, players []byte) {
	if players == nil {
		w.oldRadius = before
		w.newRadius = after
		w.speed = speed
	}
}

// SetOrigin The param players is the players to update the boarder for. Leave nil to send the packet to everyone in this world. Changes will only save if it's sending to all players. Vice versa, changes will not affect the fields if the update is sent to selected players.
func (w *WorldBorder) SetOrigin(x Double, z Double, players []byte) {
	if players == nil {
		w.originX = x
		w.originZ = z
	}
}

// SetWarningTime The param players is the players to update the boarder for. Leave nil to send the packet to everyone in this world. Changes will only save if it's sending to all players. Vice versa, changes will not affect the fields if the update is sent to selected players.
func (w *WorldBorder) SetWarningTime(time int32, players []byte) {
	if players == nil {
		w.warningTime = time
	}
}

// SetWarningBlocks The param players is the players to update the boarder for. Leave nil to send the packet to everyone in this world. Changes will only save if it's sending to all players. Vice versa, changes will not affect the fields if the update is sent to selected players.
func (w *WorldBorder) SetWarningBlocks(blocks int32, players []byte) {
	if players == nil {
		w.warningBlocks = blocks
	}
}

// SetSendOnJoin Sets whether the chunk data should be sent on player join or not.
func (w *WorldBorder) SetSendOnJoin(shouldSend bool) {
	if w.sendOnJoin == shouldSend {
		// same thing
		return
	}
}

type IGenerator interface {
	CreateChunk(Vec2) IChunk
}

type BasicWorld struct {
	IWorld
}

func CreateBasicWorld() IWorld {
	return &BasicWorld{}
}

type BasicChunk struct {
	IChunk
	blocks   []uint16
	metadata []uint16
}

func (chunk *BasicChunk) GetBlock(pos Vec3) *Block {
	index := int(pos.x + 16*pos.z + 16*16*pos.y)

	if index > 0 && index < len(chunk.blocks) {
		return &Block{
			material: materials.FromId(int32(chunk.blocks[index])),
		}
	}

	return nil
}

func CreateBasicChunk() IChunk {
	return &BasicChunk{
		blocks:   make([]uint16, 16*16*256),
		metadata: make([]uint16, 16*16*256),
	}
}
