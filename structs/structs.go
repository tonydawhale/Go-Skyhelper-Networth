package structs

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
