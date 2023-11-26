package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/Tnze/go-mc/nbt"

	"github.com/tonydawhale/Go-Skyhelper-Networth/constants"
)

func TitleCast(str string) string {
	splitStr := strings.Split(strings.ReplaceAll(strings.ToLower(str), "_", " "), " ")
	for i := 0; i < len(splitStr); i++ {
		if splitStr[i] == "" {
			continue
		}
		splitStr[i] = strings.ToUpper(string(splitStr[i][0])) + splitStr[i][1:]
	}
	return strings.Join(splitStr, " ")
}

func Flatten(nested []interface{}) []interface{} {
	flattened := make([]interface{}, 0)

	for _, i := range nested {
		switch i.(type) {
		case []interface{}:
			flattenedSubArray := Flatten(i.([]interface{}))
			flattened = append(flattened, flattenedSubArray...)
		case interface{}:
			flattened = append(flattened, i)
		}
	}

	return flattened
}

func DecodeNBT(in interface{}, data any, isBytes bool) (error) {
	var z []byte
	if !isBytes {
		z, _ = base64.StdEncoding.DecodeString(in.(string))
	} else {
		z = in.([]byte)
	}
	reader := bytes.NewReader(z)
	gzreader, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	output, err := io.ReadAll(gzreader)
	if err != nil {
		return err
	}
	
	if err := nbt.Unmarshal(output, &data); err != nil {
		return err
	}

	return nil
}

var singleContainers = map[string]string{
	"armor": "inv_armor",
	"equipment": "equipment_contents",
	"wardrobe": "wardrobe_contents",
	"inventory": "inv_contents",
	"enderchest": "ender_chest_contents",
	"accessories": "talisman_bag",
	"personal_vault": "personal_vault_contents",
	"fishing_bag": "fishing_bag",
	"potion_bag": "potion_bag",
	// "candy_inventory": "candy_inventory_contents",
}
var bagContainers = []string{
	"fishing_bag",
	"potion_bag",
	"talisman_bag",
}
var sharedContainers = []string{
	"candy_intventory_contents",
}

func ParseItems(profile map[string]interface{}, museum map[string]interface{}, v2Endpoint bool) (map[string][]interface{}, error) {
	var items map[string][]interface{} = make(map[string][]interface{})
	var inventory = profile["inventory"].(map[string]interface{})

	// Single Containers 
	for _, container := range singleContainers {
		var inventory_data string
		if v2Endpoint {
			if slices.Contains(bagContainers, container) {
				inventory_data = inventory["bag_contents"].(map[string]interface{})[container].(map[string]interface{})["data"].(string)
			} else {
				inventory_data = inventory[container].(map[string]interface{})["data"].(string)
			}
		} else {
			if inventory[container] != nil {
				inventory_data = inventory[container].(map[string]interface{})["data"].(string)
			}
		}
		if inventory_data == "" {
			continue
		}
		var data map[string]interface{}
		if err := DecodeNBT(inventory_data, &data, false); err != nil {
			return nil, err
		}
		items[container] = data["i"].([]interface{})
	}
	// Storage
	items["storage"] = []interface{}{}
	if inventory["backpack_contents"] != nil && inventory["backpack_icons"] != nil {
		for _, backpack := range inventory["backpack_contents"].(map[string]interface{}) {
			var data map[string]interface{}
			if err := DecodeNBT(backpack.(map[string]interface{})["data"].(string), &data, false); err != nil {
				return nil, err
			}
			if len(data["i"].([]interface{})) == 0 { continue }
			items["storage"] = append(items["storage"], data["i"].([]interface{})...)
		}
	}

	// Museum
	items["museum"] = []interface{}{}
	if museum != nil && museum["items"] != nil {
		for _, data := range museum["items"].(map[string]interface{}) {
			if data.(map[string]interface{})["borrowing"].(bool) {
				continue
			}
			if data.(map[string]interface{})["items"].(map[string]interface{})["data"] == nil {
				continue
			}

			var itemsData map[string]interface{}
			if err := DecodeNBT(data.(map[string]interface{})["items"].(map[string]interface{})["data"].(string), &itemsData, false); err != nil {
				return nil, err
			}
			items["museum"] = append(items["museum"], itemsData["i"].([]interface{})...)
		}
		for _, data := range museum["special"].(map[string]interface{}) {
			if data.(map[string]interface{})["items"].(map[string]interface{})["data"] == nil {
				continue
			}

			var itemsData map[string]interface{}
			if err := DecodeNBT(data.(map[string]interface{})["items"].(map[string]interface{})["data"].(string), &itemsData, false); err != nil {
				return nil, err
			}
			items["museum"] = append(items["museum"], itemsData["i"].([]interface{})...)
		}
	}

	if err := PostParseItems(profile, items, v2Endpoint); err != nil {
		return nil, err
	}

	return items,nil
}

func PostParseItems(profile map[string]interface{}, items map[string][]interface{}, v2Endpojnt bool) (error) {
	// Parse Cake Bags
	for _, categoryItems := range items {
		for _, item := range categoryItems {
			if item.(map[string]interface{})["tag"] != nil && item.(map[string]interface{})["tag"].(map[string]interface{})["ExtraAttributes"] != nil {
				if item.(map[string]interface{})["tag"].(map[string]interface{})["ExtraAttributes"].(map[string]interface{})["new_year_cake_bag_data"] != nil { 
					var cakes map[string]interface{}
					if err := DecodeNBT(item.(map[string]interface{})["tag"].(map[string]interface{})["ExtraAttributes"].(map[string]interface{})["new_year_cake_bag_data"].([]byte), &cakes, true); err != nil {
						fmt.Println("Error decoding cake bag data")
						return err
					}

					if item.(map[string]interface{})["tag"].(map[string]interface{})["ExtraAttributes"].(map[string]interface{}) != nil {
						item.(map[string]interface{})["tag"].(map[string]interface{})["ExtraAttributes"].(map[string]interface{})["new_year_cake_bag_years"] = filterCakes(cakes["i"].([]interface{}))
					}
				}
			}
		}
	}

	// Parse Sacks
	items["sacks"] = []interface{}{}
	if profile["sacks_counts"] != nil {
		for id, amount := range profile["sacks_counts"].(map[string]interface{}) {
			if int(amount.(float64)) > 0 {
				items["sacks"] = append(items["sacks"], map[string]interface{}{
					"id": id,
					"amount": int(amount.(float64)),
				})
			}
		}
	} else if profile["inventory"].(map[string]interface{})["sacks_counts"] != nil {
		for id, amount := range profile["inventory"].(map[string]interface{})["sacks_counts"].(map[string]interface{}) {
			if int(amount.(float64)) > 0 {
				items["sacks"] = append(items["sacks"], map[string]interface{}{
					"id": id,
					"amount": int(amount.(float64)),
				})
			}
		}
	}
	// Parse Essence
	items["essence"] = []interface{}{}
	if v2Endpojnt {
		if profile["currencies"].(map[string]interface{})["essence"] != nil {
			for id := range profile["currencies"].(map[string]interface{})["essence"].(map[string]interface{}) {
				items["essence"] = append(items["essence"], map[string]interface{}{
					"id": "essence_" + id,
					"amount": int(profile["currencies"].(map[string]interface{})["essence"].(map[string]interface{})[id].(map[string]interface{})["current"].(float64)),
				})
			}
		}
	} else {
		for id := range profile {
			if strings.HasPrefix(id, "essence_") {
				items["essence"] = append(items["essence"], map[string]interface{}{
					"id": id,
					"amount": int(profile[id].(float64)),
				})
			}
		}
	}
	// Parse Pets
	items["pets"] = []interface{}{}
	var data []interface{}
	if profile["pets"] != nil {
		data = profile["pets"].([]interface{})
	} else if profile["pets_data"].(map[string]interface{})["pets"] != nil {
		data = profile["pets_data"].(map[string]interface{})["pets"].([]interface{})
	}
	for _, pet := range data {
		newPet := pet
		levelData := constants.GetPetLevel(newPet.(map[string]interface{}))
		newPet.(map[string]interface{})["level"] = levelData.Level
		newPet.(map[string]interface{})["xp_max"] = levelData.XpMax
		items["pets"] = append(items["pets"], newPet)
	}
	return nil
}

func filterCakes(cakes []interface{}) []int {
	var filteredCakes []int

	for _, cake := range cakes {
		if cake.(map[string]interface{})["id"] != nil && cake.(map[string]interface{})["tag"].(map[string]interface{})["ExtraAttributes"].(map[string]interface{})["new_years_cake"] != nil {
			filteredCakes = append(filteredCakes, int(cake.(map[string]interface{})["tag"].(map[string]interface{})["ExtraAttributes"].(map[string]interface{})["new_years_cake"].(int32)))
		}
	}
	return filteredCakes
}