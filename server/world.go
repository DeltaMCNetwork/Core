package server

type IWorld interface {
	GetBlock(Vec3)
	SetBlock(Vec3)
}
