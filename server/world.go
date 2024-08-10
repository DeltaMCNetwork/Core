package server

type IWorld interface {
	GetBlock(Vec3) IBlock
	SetBlock(Vec3, IBlock)
	GetChunk(Vec2) IChunk
	GetGenerator() IGenerator
	SetGenerator(IGenerator)
}

type IChunk interface {
	GetBlock(Vec3) IBlock
	SetBlock(Vec3, IBlock)
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

func (chunk *BasicChunk) GetBlock(pos Vec3) IBlock {
	index := int(pos.x + 16*pos.z + 16*16*pos.y)

	if index > 0 && index < len(chunk.blocks) {
		return &BasicBlock{
			id: chunk.blocks[index],
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
