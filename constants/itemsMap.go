package constants

import (
	"encoding/json"
	"io"
	"os"
)

type ItemInfo struct {
	Material     string `json:"material"`
	Durability   int    `json:"durability"`
	Name         string `json:"name"`
	Category     string `json:"category"`
	Tier         string `json:"tier"`
	NpcSellPrice int    `json:"npc_sell_price"`
	Id           string `json:"id"`
}

var ItemsMap = make(map[string]ItemInfo)
var mapPath = "constants/items.json"

func init() {
	file, err := os.Open(mapPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	val, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	var temp []ItemInfo
	json.Unmarshal(val, &temp)
	for _, item := range temp {
		ItemsMap[item.Id] = item
	}
}

func GetHypixelItemInformationFromId(id string) ItemInfo {
	return ItemsMap[id]
}
