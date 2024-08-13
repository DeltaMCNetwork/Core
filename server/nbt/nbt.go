package nbt

type NbtTagType = byte

const (
	TagEnd NbtTagType = iota
	TagByte
	TagShort
	TagInt
	TagLong
	TagFloat
	TagDouble
	TagByteArray
	TagString
	TagList
	TagCompound
	TagIntArray
	TagLongArray
)

type NbtCompound struct {
	name string
	data map[string]any
}

func CreateCompound(name string) *NbtCompound {
	return &NbtCompound{
		data: make(map[string]any, 0),
	}
}

func (compound *NbtCompound) Has(key string) bool {
	return compound.data[key] != nil
}

func (compound *NbtCompound) Set(key string, value any) {
	compound.data[key] = value
}

func (compound *NbtCompound) GetByte(key string) byte {
	return compound.data[key].(byte)
}

func (compound *NbtCompound) GetShort(key string) int16 {
	return compound.data[key].(int16)
}

func (compound *NbtCompound) GetInt(key string) int32 {
	return compound.data[key].(int32)
}

func (compound *NbtCompound) GetLong(key string) int64 {
	return compound.data[key].(int64)
}

func (compound *NbtCompound) GetFloat(key string) float32 {
	return compound.data[key].(float32)
}

func (compound *NbtCompound) GetDouble(key string) float64 {
	return compound.data[key].(float64)
}

func (compound *NbtCompound) GetByteArray(key string) []byte {
	return compound.data[key].([]byte)
}

func (compound *NbtCompound) GetString(key string) string {
	return compound.data[key].(string)
}

func (compound *NbtCompound) GetList(key string) []any {
	return compound.data[key].([]any)
}

func (compound *NbtCompound) GetCompound(key string) NbtCompound {
	return compound.data[key].(NbtCompound)
}

func (compound *NbtCompound) GetIntArray(key string) []int {
	return compound.data[key].([]int)
}

func (compound *NbtCompound) GetLength() int {
	return len(compound.data)
}

// init
func Init() {
	InitWriter()
	InitReader()
}
