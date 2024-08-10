package server

import (
	"encoding/json"
	"fmt"
	"os"
)

type Item struct {
	material *Material
	metadata int8
	count    int16
}

func CreateItem(mat *Material, count int16) *Item {
	return &Item{
		material: mat,
		count:    count,
	}
}

func (item *Item) GetMetadata() int8 {
	return item.metadata
}

func (item *Item) SetMetadata(value int8) {
	item.metadata = value
}

func (item *Item) GetMaterial() *Material {
	return item.material
}

func (item *Item) SetMaterial(mat *Material) {
	item.material = mat
}

func (item *Item) SetCount(count int16) {
	item.count = count
}

func (item *Item) GetCount() int16 {
	return item.count
}

func (item *Item) Block() *Block {
	return item.material.Block()
}

type Material struct {
	Name          string
	RegistryName  string      //
	DisplayName   string      //
	Id            int32       //
	BlockId       int32       //
	MaxDurability int16       //
	Metadata      int16       //
	MaxStackSize  int8        //
	Hardness      int8        //
	IsBlock       bool        //
	IsArmor       bool        //
	Drops         []*DropData //
}

func (m *Material) Block() *Block {
	if m.IsBlock {
		return CreateBlock(m)
	}

	return nil
}

func (m *Material) Item(count int16) *Item {
	return CreateItem(m, count)
}

type DropData struct {
	Drop     int32 //
	MinDrop  int16 //
	MaxDrop  int16 //
	Metadata int16 //
}

var materials = CreateMaterialRegistry()

type MaterialRegistry struct {
	Materials map[string]*Material
	StringMap map[int32]string
}

func CreateMaterialRegistry() *MaterialRegistry {
	return &MaterialRegistry{
		Materials: make(map[string]*Material),
		StringMap: map[int32]string{},
	}
}

func (mat *MaterialRegistry) From(id string) *Material {
	return mat.Materials[id]
}

func (mat *MaterialRegistry) FromId(id int32) *Material {
	regName := mat.StringMap[id]
	return mat.Materials[regName]
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

	//Debug("Data: " + string(data))

	for i := range materials {
		registerMaterial(mat, materials[i])
	}

	// read blocks to make sure which material is a block

	data, err = os.ReadFile("resources/block.json")

	if err != nil {
		Fatal("Failed to load resources! %s", err.Error())
	}

	var blocks []map[string]interface{}

	err = json.Unmarshal(data, &blocks)

	if err != nil {
		Fatal("Failed to load resources! %s", err.Error())
	}

	for i := range blocks {
		registerBlock(mat, blocks[i])
	}
}

func registerBlock(matReg *MaterialRegistry, values map[string]interface{}) {
	materialId := values["name"].(string)

	material := matReg.Materials[materialId]

	if material == nil {
		return
	}

	blockId := int32(values["id"].(float64))

	dropData := make([]*DropData, 0)

	dropDataJson := values["drops"].([]interface{})

	for i := range dropDataJson {
		jsonData := dropDataJson[i].(map[string]interface{})

		dData := &DropData{}

		if fmt.Sprintf("%T", jsonData["drop"]) == "float64" {
			dData.Drop = int32(jsonData["drop"].(float64))
			dData.Metadata = 0
		} else {
			innerMap := jsonData["drop"].(map[string]interface{})

			dData.Drop = int32(innerMap["id"].(float64))
			dData.Metadata = int16(innerMap["metadata"].(float64))
		}
		dData.MinDrop = 1
		dData.MaxDrop = 1

		if jsonData["minCount"] != nil {
			dData.MinDrop = int16(jsonData["minCount"].(float64))
			dData.MaxDrop = dData.MinDrop
		}

		if jsonData["maxCount"] != nil {
			dData.MaxDrop = int16(jsonData["maxCount"].(float64))
		}

		dropData = append(dropData, dData)
	}

	if len(dropData) != 0 {
		material.Drops = dropData
	}

	material.Hardness = 0

	if values["hardness"] != nil {
		material.Hardness = int8(values["hardness"].(float64))
	}

	material.IsBlock = true
	material.BlockId = blockId

}

func registerMaterial(matReg *MaterialRegistry, values map[string]interface{}) {
	material := &Material{}

	material.MaxStackSize = int8(values["stackSize"].(float64))
	material.DisplayName = values["displayName"].(string)
	//material.MaxDurability = int16(values["maxDurability"].(float64))
	maxDurability := values["maxDurability"]
	if maxDurability != nil {
		material.MaxDurability = int16(maxDurability.(float64))
	} else {
		material.MaxDurability = 0
	}
	material.Id = int32(values["id"].(float64))
	material.RegistryName = values["name"].(string)
	material.Name = material.RegistryName
	material.Metadata = 0

	enchantCategories := values["enchantCategories"]

	if enchantCategories != nil && fmt.Sprintf("%T", enchantCategories) == "[]string" {
		encArr := enchantCategories.([]string)
		if encArr[0] == "armor" {
			material.IsArmor = true
		}
	}

	material.IsBlock = false

	variety := values["variations"]

	if variety != nil {
		variations := variety.([]interface{})

		for i := 1; i < len(variations); i++ {
			variation := variations[i].(map[string]interface{})
			variMaterial := &Material{}
			variMaterial.MaxStackSize = material.MaxStackSize
			variMaterial.DisplayName = variation["displayName"].(string)
			variMaterial.MaxDurability = material.MaxDurability
			variMaterial.Metadata = int16(variation["metadata"].(float64))
			variMaterial.Id = material.Id
			variMaterial.RegistryName = fmt.Sprintf("%s:%d", material.RegistryName, variMaterial.Metadata)
			variMaterial.Name = material.RegistryName

			matReg.Materials[variMaterial.RegistryName] = variMaterial

			defer Debug("Registered %s:%s:%d", variMaterial.DisplayName, variMaterial.RegistryName, variMaterial.Id)
		}
	}

	matReg.Materials[material.RegistryName] = material
	matReg.StringMap[material.Id] = material.RegistryName

	Debug("Registered %s:%s:%d", material.DisplayName, material.RegistryName, material.Id)
}
