package server

type IWorld interface {
	GetBlock(Vec3) *Block
	SetBlock(Vec3, *Block)
	GetChunk(Vec2) IChunk
	GetGenerator() IGenerator
	SetGenerator(IGenerator)
}

type IChunk interface {
	GetBlock(Vec3) *Block
	SetBlock(Vec3, *Block)
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
