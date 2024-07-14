package server

import (
	uuid "github.com/satori/go.uuid"
)

type double float64

// Last modified: 6/21/2024
type Vec3 struct {
	x double
	y double
	z double
}

func (vec *Vec3) GetX() double {
	return vec.x
}

func (vec *Vec3) GetY() double {
	return vec.y
}

func (vec *Vec3) GetZ() double {
	return vec.z
}

func (vec *Vec3) SetX(x double) {
	vec.x = x
}

func (vec *Vec3) SetY(y double) {
	vec.y = y
}

func (vec *Vec3) SetZ(z double) {
	vec.z = z
}

func CreateVec3(x double, y double, z double) *Vec3 {
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

func CreatePosition(x double, y double, z double, yaw float32, pitch float32) *Position {
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

type Message struct {
}

func CreateServerResponse(server MinecraftServer) *ServerResponse {
	return &ServerResponse{}
}
