package server

type IWorld interface {
	GetBlock(BlockPos) *Block
	SetBlock(BlockPos, *Block)
	GetChunk(*Vec2) IChunk
	SetChunk(*Vec2, IChunk)
	GetGenerator() IGenerator
	SetGenerator(IGenerator)
	GetWorldBorder() *WorldBorder
	SetWorldBorder(*WorldBorder)
	GetDimension() DimensionType
}

type IChunk interface {
	GetBlock(BlockPos) *Block
	SetBlock(BlockPos, *Block)
	ToBytes(*Vec2) []byte
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
	CreateChunk(*Vec2, IWorld) IChunk
}

type BasicWorld struct {
	IWorld
	generator IGenerator
	chunks    map[Vec2]IChunk
	border    *WorldBorder
	dimension DimensionType
}

func CreateBasicWorld() IWorld {
	return &BasicWorld{
		generator: CreateBasicGenerator(),
		chunks:    map[Vec2]IChunk{},
	}
}

func (world *BasicWorld) GetBlock(pos BlockPos) *Block {
	chunk := world.GetChunk(CreateVec2(pos.X>>4, pos.Z>>4))
	pos.ToChunkBlockCoords()

	return chunk.GetBlock(pos)
}

func (world *BasicWorld) SetBlock(pos BlockPos, block *Block) {
	chunk := world.GetChunk(CreateVec2(pos.X>>4, pos.Z>>4))
	pos.ToChunkBlockCoords()

	chunk.SetBlock(pos, block)
}

func (world *BasicWorld) GetGenerator() IGenerator {
	return world.generator
}

func (world *BasicWorld) SetGenerator(gen IGenerator) {
	world.generator = gen
}

func (world *BasicWorld) GetWorldBorder() *WorldBorder {
	return world.border
}

func (world *BasicWorld) SetWorldBorder(border *WorldBorder) {
	world.border = border
}

func (world *BasicWorld) GetChunk(pos *Vec2) IChunk {
	chunk, found := world.chunks[*pos]

	if !found {
		chunk = world.generator.CreateChunk(pos, world)
		world.chunks[*pos] = chunk
	}

	return chunk
}

func (world *BasicWorld) SetChunk(pos *Vec2, chunk IChunk) {
	world.chunks[*pos] = chunk
}

func (world *BasicWorld) GetDimension() DimensionType {
	return world.dimension
}

// changing chunks to be a single array with metadata in it cuz im silly
// ref: https://github.com/SharpMC/SharpMC/blob/master/src/SharpMC.Core/Worlds/ChunkColumn.cs#L193
// Optimization: Set the value to condense the metadata and blockid inside of the #IChunk.SetBlock()

type BasicChunk struct {
	IChunk
	data []uint8
}

func (chunk *BasicChunk) GetBlock(pos BlockPos) *Block {
	index := int(pos.X + 16*pos.Z + 16*16*pos.Y)

	if index > 0 && index < CHUNK_MAX_LENGTH {
		val := chunk.data[index]
		block := materials.FromId(int32(val >> 4)).Block()
		//block.metadata = val & 15 // messy bit operator hack

		return block
	}

	return nil
}

func (chunk *BasicChunk) SetBlock(pos BlockPos, block *Block) {
	index := int(pos.X+16*pos.Z+16*16*pos.Y) * 2

	if index > 0 && index < CHUNK_MAX_LENGTH*2 {
		//chunk.data[index*2] = uint16(block.GetBlockId()<<4) | block.metadata
		//chunk.data[index*2] =
	}
}

func (chunk *BasicChunk) ToBytes(pos *Vec2) []byte {
	buf := createBasicBuffer()

	buf.WriteInt(int32(pos.x))
	buf.WriteInt(int32(pos.y))
	buf.WriteBool(true)
	buf.WriteUInt16(0xffff)

	//data len, to be changed according
	buf.WriteVarInt(0)

	return buf.GetBytes()

}

func CreateBasicChunk() IChunk {
	return &BasicChunk{
		data: make([]uint8, CHUNK_MAX_LENGTH*2),
	}
}

type BasicGenerator struct {
	IGenerator
}

func CreateBasicGenerator() *BasicGenerator {
	return &BasicGenerator{}
}

func (gen *BasicGenerator) CreateChunk(vec *Vec2, world IWorld) IChunk {
	return nil
}
