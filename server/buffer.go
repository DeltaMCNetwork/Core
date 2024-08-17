package server

import (
	"encoding/binary"
	"math"

	uuid "github.com/satori/go.uuid"
)

type IBuffer interface {
	Read(count int) []byte
	ReadByte() byte
	ReadBool() bool
	ReadUInt8() uint8
	ReadInt8() int8
	ReadUInt16() uint16
	ReadInt16() int16
	ReadInt() int32
	ReadUInt() uint32
	ReadVarInt() int32
	ReadFloat() float32
	ReadDouble() Double
	ReadLong() int64
	ReadULong() uint64
	ReadVarLong() int64
	ReadPosition() *Position
	ReadVec3() *Vec3
	ReadString() string
	ReadByteArray() []byte
	ReadUUID() UUID
	Write([]byte)
	WriteByte(byte)
	WriteBool(bool)
	WriteUInt8(uint8)
	WriteInt8(int8)
	WriteUInt16(uint16)
	WriteInt16(int16)
	WriteInt(int32)
	WriteUInt(uint32)
	WriteVarInt(int32)
	WriteFloat(float32)
	WriteDouble(Double)
	WriteLong(int64)
	WriteULong(uint64)
	WriteVarLong(int64)
	WritePosition(*Position)
	WriteVec3(Vec3)
	WriteString(string)
	WriteByteArray(bytes []byte)
	WriteUUID(UUID)

	SetData([]byte)
	GetBytes() []byte
	GetLength() int
	GetPointer() int
	SetPointer(int)
}

func createBasicBuffer() IBuffer {
	return &BasicBuffer{}
}

type BasicBuffer struct {
	data    []byte
	pointer int
}

func (buffer *BasicBuffer) WriteVarLong(value int64) {
	for {
		temp := value & 0x7F
		value >>= 7

		if value != 0 {
			temp |= 0x80
		}

		buffer.WriteByte(byte(temp))

		if value == 0 {
			break
		}
	}
}

func (buffer *BasicBuffer) Read(count int) []byte {
	index := buffer.pointer + count
	data := buffer.data[buffer.pointer:index]
	buffer.pointer = index

	return data
}

func (buffer *BasicBuffer) ReadByte() byte {
	return buffer.Read(1)[0]
}

func (buffer *BasicBuffer) ReadBool() bool {
	b := buffer.ReadByte()

	return b == 0x01
}

func (buffer *BasicBuffer) ReadUInt8() uint8 {
	return uint8(buffer.Read(1)[0])
}

func (buffer *BasicBuffer) ReadInt8() int8 {
	return int8(buffer.Read(1)[0])
}

func (buffer *BasicBuffer) ReadUInt16() uint16 {
	return binary.BigEndian.Uint16(buffer.Read(2))
}

func (buffer *BasicBuffer) ReadInt16() int16 {
	return int16(binary.BigEndian.Uint16(buffer.Read(2)))
}

func (buffer *BasicBuffer) ReadInt() int32 {
	return int32(binary.BigEndian.Uint32(buffer.Read(4)))
}

func (buffer *BasicBuffer) ReadUInt() uint32 {
	return binary.BigEndian.Uint32(buffer.Read(4))
}

func (buffer *BasicBuffer) ReadVarInt() int32 {
	var num int
	var res int32

	for {
		val := buffer.ReadByte()
		tmp := int32(val)
		res |= (tmp & 0x7F) << uint(num*7)

		if num++; num > 5 {
			panic("Value too big!")
		}

		if tmp&0x80 != 0x80 {
			break
		}
	}

	return res
}

func (buffer *BasicBuffer) ReadFloat() float32 {
	return math.Float32frombits(binary.BigEndian.Uint32(buffer.Read(4)))
}

func (buffer *BasicBuffer) ReadDouble() Double {
	return Double(math.Float64frombits(binary.BigEndian.Uint64(buffer.Read(8))))
}

func (buffer *BasicBuffer) ReadLong() int64 {
	return int64(binary.BigEndian.Uint64(buffer.Read(8)))
}

func (buffer *BasicBuffer) ReadULong() uint64 {
	return binary.BigEndian.Uint64(buffer.Read(4))
}

func (buffer *BasicBuffer) ReadVarLong() int64 {
	var num int
	var res int64

	for {
		val := buffer.ReadByte()
		tmp := int64(val)
		res |= (tmp & 0x7F) << uint(num*7)

		if num++; num > 10 {
			panic("Value too big!")
		}

		if tmp&0x80 != 0x80 {
			break
		}
	}

	return res
}

func (buffer *BasicBuffer) ReadPosition() *Position {
	v := buffer.ReadLong()
	pos := CreateEmptyPosition()
	vec := pos.GetVec3()

	vec.SetX(Double(v >> 38))
	vec.SetY(Double(v >> 26 & 0xFFF))
	vec.SetZ(Double(v << 38 >> 38))

	return pos
}

func (buffer *BasicBuffer) ReadVec3() *Vec3 {
	v := buffer.ReadLong()
	vec := &Vec3{}

	// might do problems idk
	// should be fine

	vec.SetX(Double(v >> 38))
	vec.SetY(Double(v >> 26 & 0xFFF))
	vec.SetZ(Double(v << 38 >> 38))

	return vec
}

func (buffer *BasicBuffer) ReadString() string {
	length := buffer.ReadVarInt()

	return string(buffer.Read(int(length)))
}

func (buffer *BasicBuffer) ReadUUID() UUID {
	mBytes := make([]byte, 8)
	lBytes := make([]byte, 8)

	binary.BigEndian.PutUint64(mBytes, uint64(buffer.ReadLong()))
	binary.BigEndian.PutUint64(lBytes, uint64(buffer.ReadLong()))

	id, _ := uuid.FromBytes(append(mBytes, lBytes...))

	return id
}

func (buffer *BasicBuffer) Write(data []byte) {
	buffer.data = append(buffer.data, data...)
}

func (buffer *BasicBuffer) WriteAll(data ...byte) {
	buffer.data = append(buffer.data, data...)
}

func (buffer *BasicBuffer) WriteByte(data byte) {
	buffer.data = append(buffer.data, data)
}

func (buffer *BasicBuffer) WriteBool(value bool) {
	b := byte(0x00)

	if value {
		b = 0x01
	}

	buffer.WriteByte(b)
}

func (buffer *BasicBuffer) WriteUInt8(value uint8) {
	buffer.WriteByte(value)
}

func (buffer *BasicBuffer) WriteInt8(value int8) {
	buffer.WriteByte(uint8(value))
}

func (buffer *BasicBuffer) WriteUInt16(value uint16) {
	buffer.data = binary.BigEndian.AppendUint16(buffer.data, value)
}

func (buffer *BasicBuffer) WriteInt16(value int16) {
	buffer.WriteUInt16(uint16(value))
}

func (buffer *BasicBuffer) WriteInt(value int32) {
	buffer.WriteUInt(uint32(value))
}

func (buffer *BasicBuffer) WriteUInt(value uint32) {
	buffer.data = binary.BigEndian.AppendUint32(buffer.data, value)
}

func (buffer *BasicBuffer) WriteVarInt(value int32) {
	for {
		temp := value & 0x7F
		value >>= 7

		if value != 0 {
			temp |= 0x80
		}

		buffer.WriteByte(byte(temp))

		if value == 0 {
			break
		}
	}
}

func (buffer *BasicBuffer) WriteFloat(float float32) {
	buffer.WriteUInt(math.Float32bits(float))
}

func (buffer *BasicBuffer) WriteDouble(value Double) {
	buffer.WriteULong(math.Float64bits(float64(value)))
}

func (buffer *BasicBuffer) WriteLong(value int64) {
	buffer.WriteULong(uint64(value))
}

func (buffer *BasicBuffer) WriteULong(value uint64) {
	var data [8]byte
	binary.BigEndian.PutUint64(data[:], value)
	buffer.Write(data[:])
}

func (buffer *BasicBuffer) WritePosition(pos *Position) {
	vec := pos.GetVec3()

	buffer.WriteULong(((uint64(vec.x) & 0x3FFFFFF) << 38) | ((uint64(vec.z) & 0x3FFFFFF) << 12) | (uint64(vec.y) & 0xFFF))
}

func (buffer *BasicBuffer) WriteString(value string) {
	buffer.WriteVarInt(int32(len(value)))

	buffer.Write([]byte(value))
}

func (buffer *BasicBuffer) WriteByteArray(bytes []byte) {
	buffer.WriteVarInt(int32(len(bytes)))

	buffer.Write(bytes)
}

func (buffer *BasicBuffer) WriteUUID(uuid UUID) {
	bytes := uuid.Bytes()

	msb := 0
	lsb := 0

	for i := 0; i < 8; i++ {
		msb = (msb << 0x08) | int(bytes[i]&0xFF)
	}

	for i := 8; i < 16; i++ {
		lsb = (lsb << 0x08) | int(bytes[i]&0xFF)
	}

	buffer.Write(bytes)
}

func (buffer *BasicBuffer) ReadByteArray() []byte {
	return buffer.Read(int(buffer.ReadVarInt()))
}

func (buffer *BasicBuffer) GetBytes() []byte {
	return buffer.data
}

func (buffer *BasicBuffer) GetLength() int {
	return len(buffer.data)
}

func (buffer *BasicBuffer) GetPointer() int {
	return buffer.pointer
}

func (buffer *BasicBuffer) WriteVec3(vec Vec3) {
	buffer.WriteULong(((uint64(vec.x) & 0x3FFFFFF) << 38) | ((uint64(vec.z) & 0x3FFFFFF) << 12) | (uint64(vec.y) & 0xFFF))
}

func (buffer *BasicBuffer) SetData(data []byte) {
	buffer.data = data
}

func (buffer *BasicBuffer) SetPointer(pointer int) {
	buffer.pointer = pointer
}

var _ IBuffer = (*BasicBuffer)(nil)
