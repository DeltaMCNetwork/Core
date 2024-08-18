package server

import (
	uuid "github.com/satori/go.uuid"
)

type Double = float64

// Last modified: 6/21/2024
type Vec3 struct {
	x Double
	y Double
	z Double
}

func (vec *Vec3) GetX() Double {
	return vec.x
}

func (vec *Vec3) GetY() Double {
	return vec.y
}

func (vec *Vec3) GetZ() Double {
	return vec.z
}

func (vec *Vec3) SetX(x Double) {
	vec.x = x
}

func (vec *Vec3) SetY(y Double) {
	vec.y = y
}

func (vec *Vec3) SetZ(z Double) {
	vec.z = z
}

func CreateVec3(x Double, y Double, z Double) *Vec3 {
	return &Vec3{
		x: x,
		y: y,
		z: z,
	}
}

func (vec *Vec3) Add(vec1 Vec3) {
	vec.x += vec1.x
	vec.y += vec1.y
	vec.z += vec1.z
}

func (vec *Vec3) Sub(vec1 Vec3) {
	vec.x -= vec1.x
	vec.y -= vec1.y
	vec.z -= vec1.z
}

//Last modified: 6/21/2024

type Position struct {
	pos   *Vec3
	yaw   float32
	pitch float32
}

func (pos *Position) GetYaw() float32 {
	return pos.yaw
}

func (pos *Position) GetPitch() float32 {
	return pos.pitch
}

func (pos *Position) SetYaw(f float32) {
	pos.yaw = f
}

func (pos *Position) SetPitch(p float32) {
	pos.pitch = p
}

func (pos *Position) GetVec3() *Vec3 {
	return pos.pos
}

func (pos *Position) SetVec3(vec *Vec3) {
	pos.pos = vec
}

func CreatePosition(x Double, y Double, z Double, yaw float32, pitch float32) *Position {
	return &Position{
		pos:   CreateVec3(x, y, z),
		yaw:   yaw,
		pitch: pitch,
	}
}

func CreatePositionVec(vec *Vec3, yaw float32, pitch float32) *Position {
	return &Position{
		pos:   vec,
		yaw:   yaw,
		pitch: pitch,
	}
}

func CreateEmptyPosition() *Position {
	return &Position{}
}

//Last modified: 6/21/2024

type UUID = uuid.UUID

func CreateUUID() UUID {
	return uuid.NewV4()
}

//Last modified: 6/22/2024

type ServerResponse struct {
}

type Version struct {
}

func CreateServerResponse(server *MinecraftServer) *ServerResponse {
	return &ServerResponse{}
}

//Last modified: 7/16/2024
//Last modified: 7/16/2024

type Vec2 struct {
	x int
	y int
}

func CreateVec2(x int, y int) *Vec2 {
	return &Vec2{
		x: x,
		y: y,
	}
}

func (vec *Vec2) GetX() int {
	return vec.x
}

func (vec *Vec2) SetX(x int) {
	vec.x = x
}

func (vec *Vec2) GetY() int {
	return vec.y
}

func (vec *Vec2) SetY(y int) {
	vec.y = y
}

//Last modified: 8/17/2024

type BlockPos struct {
	X int
	Y int
	Z int
}

func (pos *BlockPos) ToChunkBlockCoords() {
	pos.X %= 16
	pos.Z %= 16
}

func CreateBlockPos(x int, y int, z int) *BlockPos {
	return &BlockPos{
		X: x,
		Y: y,
		Z: z,
	}
}
