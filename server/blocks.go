package server

type IBlock interface {
}

type BasicBlock struct {
	IBlock
	id uint16
}

type BlockRegistry struct {
}
