package server

type Block struct {
	material *Material
	metadata uint16
}

func CreateBlock(mat *Material) *Block {
	return &Block{
		material: mat,
	}
}

func (b *Block) GetMaterial() *Material {
	return b.material
}

func (b *Block) SetMaterial(mat *Material) {
	b.material = mat
}

func (b *Block) GetMetadata() uint16 {
	return b.metadata
}

func (b *Block) SetMetadata(state uint16) {
	b.metadata = state
}

func (b *Block) GetBlockId() int32 {
	return b.material.BlockId
}

func (b *Block) GetHardness() int8 {
	return b.material.Hardness
}

func (b *Block) GetDrops() []*DropData {
	return b.material.Drops
}

func (b *Block) GetAsItem(count int8) *Item {
	return b.material.Item(count)
}
