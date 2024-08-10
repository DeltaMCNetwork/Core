package server

import (
	"encoding/json"
	"os"
)

type Item struct {
	Id           uint16
	Metadata     int8
	MaxStackSize int8
	Count        int8
}

type Material struct {
	RegistryName  string
	DisplayName   string
	Id            int32
	BlockId       int32
	MaxDurability int16
	MaxStackSize  int8
	IsBlock       bool
	IsArmor       bool
}

type MaterialRegistry struct {
	Materials map[int32]*Material
	StringMap map[string]int32
}

func CreateMaterialRegistry() *MaterialRegistry {
	return &MaterialRegistry{
		Materials: make(map[int32]*Material),
		StringMap: map[string]int32{},
	}
}

func (mat *MaterialRegistry) Load(path string) {
	data, err := os.ReadFile("resources/item.json")

	if err != nil {
		Fatal("Failed to load resources! %s", err.Error())
	}

	var materials []map[string]interface{}

	err = json.Unmarshal(data, &materials)

	if err != nil {
		Fatal("Failed to load resources! %s", err.Error())
	}

	Debug("Data: " + string(data))

	for i := range materials {
		registerMaterial(mat, materials[i])
	}
}

func registerMaterial(matReg *MaterialRegistry, values map[string]interface{}) {
	material := &Material{}

	material.MaxStackSize = int8(values["stackSize"].(float64))
	material.DisplayName = values["displayName"].(string)
	material.MaxDurability = int16(values["maxDurability"].(float64))
	material.Id = int32(values["id"].(float64))
	material.RegistryName = values["name"].(string)

}
