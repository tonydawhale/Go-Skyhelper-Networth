package structs

type CalculatorOptions struct {
	V2Endpoint bool
	Cache bool
	OnlyNetworth bool
	Prices map[string]int
	ReturnItemData bool
	MuseumData map[string]interface{}
}

type NetworthItem struct {
	Name        string        `json:"name"`
	Id          string        `json:"id"`
	Price       float64       `json:"price"`
	Calculation []interface{} `json:"calculation"`
	Count       int           `json:"count"`
	Soulbound   bool          `json:"soulbound"`
}

type Item struct {
	Material     string `json:"material"`
	Durability   int    `json:"durability"`
	Name         string `json:"name"`
	Category     string `json:"category"`
	Tier         string `json:"tier"`
	NpcSellPrice int    `json:"npc_sell_price"`
	Id           string `json:"id"`
	Amount       int    `json:"amount"`
}

type Pet struct {
}

type NBTItemData struct {
	Count int `json:"Count" nbt:"Count"`
	Damage int `json:"Damage" nbt:"Damage"`
	ID string `json:"id" nbt:"id"`
	Tag struct {
		ExtraAttributes struct {
			DonatedMuseum bool `json:"donated_museum" nbt:"donated_museum"`
			DungeonItemLevel int `json:"dungeon_item_level" nbt:"dungeon_item_level"`
			Enchantments map[string]int `json:"enchantments" nbt:"enchantments"`
			Gems map[string]int `json:"gems" nbt:"gems"`
			HotPotatoCount int `json:"hot_potato_count" nbt:"hot_potato_count"`
			Id string `json:"id" nbt:"id"`
			Modifier string `json:"modifier" nbt:"modifier"`
			OriginTag string `json:"origin_tag" nbt:"origin_tag"`
			Runes map[string]int `json:"runes" nbt:"runes"`
			RarityUpgrades bool `json:"rarity_upgrades" nbt:"rarity_upgrades"`
			Timestamp string `json:"timestamp" nbt:"timestamp"`
			Uuid string `json:"uuid" nbt:"uuid"`
		} `json:"ExtraAttributes" nbt:"ExtraAttributes"`
		HideFlags int `json:"HideFlags" nbt:"HideFlags"`
		Unbreakable bool `json:"Unbreakable" nbt:"Unbreakable"`
		Display struct {
			Lore []string `json:"Lore" nbt:"Lore"`
			Name string `json:"Name" nbt:"Name"`
		} `json:"display" nbt:"display"`
		SkullOwner struct {
			Id string `json:"Id" nbt:"Id"`
			Properties struct {
				Textures []struct {
					Name string `json:"Name" nbt:"Name"`
					Value string `json:"Value" nbt:"Value"`
				} `json:"textures" nbt:"textures"`
			} `json:"Properties" nbt:"Properties"`
		}
		Ench []struct {
			Id int `json:"id" nbt:"id"`
			Level int `json:"lvl" nbt:"lvl"`
		} `json:"ench" nbt:"ench"`
	} `json:"tag" nbt:"tag"`
}

type NBT struct {
	Type int `json:"type"`
	Data string `json:"data"`
}

type ProfileData struct {
	Members map[string]UserProfile `json:"members"`
}

type UserProfile struct {
	Inventory Inventory `json:"inventory"`
}

type Inventory struct {
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