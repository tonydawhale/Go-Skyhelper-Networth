package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"slices"

	"github.com/Tnze/go-mc/nbt"
)
type NBT struct {
	Type int `json:"type"`
	Data string `json:"data"`
}
type NBTItemData struct {
	Count int `json:"Count"`
	Damage int `json:"Damage"`
	ID string `json:"id"`
	Tag struct {
		ExtraAttributes struct {
			DonatedMuseum bool `json:"donated_museum"`
			DungeonItemLevel int `json:"dungeon_item_level"`
			Enchantments map[string]int `json:"enchantments"`
			Gems map[string]int `json:"gems"`
			HotPotatoCount int `json:"hot_potato_count"`
			Id string `json:"id"`
			Modifier string `json:"modifier"`
			OriginTag string `json:"origin_tag"`
			Runes map[string]int `json:"runes"`
			RarityUpgrades bool `json:"rarity_upgrades"`
			Timestamp string `json:"timestamp"`
			Uuid string `json:"uuid"`
		} `json:"ExtraAttributes"`
		HideFlags int `json:"HideFlags"`
		Unbreakable bool `json:"Unbreakable"`
		Display struct {
			Lore []string `json:"Lore"`
			Name string `json:"Name"`
		} `json:"display"`
		SkullOwner struct {
			Id string `json:"Id"`
			Properties struct {
				Textures []struct {
					Name string `json:"Name"`
					Value string `json:"Value"`
				} `json:"textures"`
			} `json:"Properties"`
		}
		Ench []struct {
			Id int `json:"id"`
			Level int `json:"lvl"`
		} `json:"ench"`
	} `json:"tag"`
}
type ProfileData struct {
	Inventory struct {
		InvContents NBT `json:"inv_contents"`
		EnderChestContents NBT `json:"ender_chest_contents"`
		BackpackIcons map[string]NBT `json:"backpack_icons"`
		BagContents struct {
			PotionBag NBT `json:"potion_bag"`
			TalismanBag NBT `json:"talisman_bag"`
			FishingBag NBT `json:"fishing_bag"`
			Quiver NBT `json:"quiver"`
		} `json:"bag_contents"`
		InvArmor NBT `json:"inv_armor"`
		EquipmentContents NBT `json:"equipment_contents"`
		PersonalVaultContents NBT `json:"personal_vault_contents"`
		BackpackContents map[string]NBT `json:"backpack_contents"`
		SackCounts map[string]int `json:"sack_counts"`
		WardrobeContents NBT `json:"wardrobe_contents"`
	}
}

func main() {
	var base64String = "H4sIAAAAAAAAAFVU227iRhgekmxLqKqV9mqvdkdt2galEB+wA7kjQMBsbJZDIKaqIh8GZ8iMjXwImHdor3pVrXobqY+RR9kHqfrbJFv1wof5/+//Dh7bJYQOUYGWEEKFPbRH3cJvBfSqFSR+XCih/djy9tFhj7rkklleBKh/Sqg0vk8YG6x9EhbRnuaiI0tWFMlxnYrcsOyKvGgsKg35TK6ooiPIykIV7TMX5j6GwYqEMSXRISrGZBMnIYly6SJ6NbVYQgp/kXXgaa2+YM1E5sijO/umSbV24OkTc21MzFTf6pKx9DYGXX/QWk3q9PoPc86i+TW712hT1Vqaos8A252zweSC61tT0CVNHHRhtnu91beOOF960rytK/OxFrVo09P8i9SW5iu7Ox2YoLvj6U/cbiO1p4Zgy/2YfMEaq7mk3Lm9aTqf9plzM105fLrT7o1Sd3b9jBsx0huJ0NvuelHmN8s2GQts8P9ahp+mdkvzBrRJrd5IcNrBw5X8H8cVF1c2ny4dfsndlpLMb4YPbnday32MG3zQ1jfGhC2NWZ8a0rCmt4ei0R2m87YpGFyDniMbXU3WJx3JnF1SU9K25swUzaUpG1v3Tl86orE1uLkdwvNxasZkxDUvyL0thnDtCR8WQ9irEvrapdGKWekhOrgKQlKE4hv0/unxrEusEI8dqJ3jp0dXVWtwqR+LdVUpozIAesRi8V3WtE4kQapKcNc4PhGVcg48UepiVSwjWJy1yYL4Edlh63IOBYQk16u1MvoOEOMVIe4z145HeqGpNtQyEgCj+TFhjHrEd56plJr4LCq9qIo/K2qtChbfZetf4Pj86Xc4//pl+ecf2RJyfg+cI+ImDolwfEewa3HLIzgNEhxb9wQvwoCjtwBaU2iHEbZTIHFE4QeoVYHgx6dH9RK+HjwmMb4I/CQ6x7McawehnwU/Fk4h4Ps8oLX2I2zhHRnm1KeBj8kDCdPcydMjkYVsJiJO4LsRTlY4DmCAWxvKE46OcpCVJX52VMVmkIT4xd6aMpY/zDi0HgjLp30Xk82KBS7BoObDntop+gYwxCccvt0sxjuI8fTIrjrdjtFujkzcvja6nYGBe50rvTMpogPD4iQzqc5oRPA4DkL+U4R7hHEIDuXPn/7eneGFet3ZgH4zjkNqJzGJiqjIA5cuKAnRwRrm99G3ycoLLZfcMojPwMFeMftboTczbdy5nWmTXmd0+0U8SaB1dLawJbUmLCpCY2FVarIqVux6jVRcu27XieDIdVUuosOYchLFFl+h1/VTST2VZCyfizX8UQcV9FU732K0j9C/qwQGpiwFAAA="

	var data map[string]interface{}
	if err := decodeNBT(base64String, &data); err != nil {
		panic(err)
	}
}

func decodeNBT(in string, data any) (error) {
	z, _ := base64.StdEncoding.DecodeString(in)
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
	"candy_inventory": "candy_inventory_contents",
}
var bagContainers = []string{
	"fishing_bag",
	"potion_bag",
	"talisman_bag",
}
var sharedContainers = []string{
	"candy_intventory_contents",
}

func consalidateInventories(profile ProfileData, museum map[string]interface{}) {
	var items map[string]interface{}

	// Single Containers 
	for _, container := range singleContainers {
		items[container] = []NBTItemData{}
		var inventory_data interface{}
		if slices.Contains(bagContainers, container) {
			switch container {
			case "potion_bag":
				inventory_data = profile.Inventory.BagContents.PotionBag
			case "talisman_bag":
				inventory_data = profile.Inventory.BagContents.TalismanBag
			case "fishing_bag":
				inventory_data = profile.Inventory.BagContents.FishingBag
			case "quiver":
				inventory_data = profile.Inventory.BagContents.Quiver
			}
		} else {
			switch container {
			case "armor":
				inventory_data = profile.Inventory.InvArmor
			case "equipment":
				inventory_data = profile.Inventory.EquipmentContents
			case "wardrobe":
				inventory_data = profile.Inventory.WardrobeContents
			case "inventory":
				inventory_data = profile.Inventory.InvContents
			case "enderchest":
				inventory_data = profile.Inventory.EnderChestContents
			case "personal_vault":
				inventory_data = profile.Inventory.PersonalVaultContents
			case "backpack":
				inventory_data = profile.Inventory.BackpackContents
			}
		
		}
		var data map[string]interface{}
		if err := decodeNBT(inventory_data.(string), &data); err != nil {
			panic(err)
		}
		items[container] = append(items[container].([]NBTItemData), data["i"].([]NBTItemData)...)
	}
	// Storage
	items["storage"] = []NBTItemData{}
	inv := profile.Inventory
	if inv.BackpackContents != nil && inv.BackpackIcons != nil {
		for _, backpack := range inv.BackpackContents {
			var data map[string]interface{}
			if err := decodeNBT(backpack.Data, &data); err != nil {
				panic(err)
			}
			items["storage"] = append(items["storage"].([]NBTItemData), data["i"].([]NBTItemData)...)
		}
	}
	fmt.Println(items["storage"].([]NBTItemData)[0])
}