package server

type Block struct {
	material *Material
	state    int8
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

func (b *Block) GetState() int8 {
	return b.state
}

func (b *Block) SetState(state int8) {
	b.state = state
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
